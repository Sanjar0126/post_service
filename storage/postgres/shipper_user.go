package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type ShipperUserRepo struct {
	db *sqlx.DB
}

// NewShipperUserRepo ...
func NewShipperUserRepo(db *sqlx.DB) repo.ShipperUserStorageI {
	return &ShipperUserRepo{
		db: db,
	}
}

func (shipUsrRep *ShipperUserRepo) Create(shipperUser *pb.ShipperUser) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		shipper_users
		(
			id,
			name,
			username,
			password,
			phone,
			shipper_id,
			user_role_id
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`
	_, err = shipUsrRep.db.Query(
		insertNew,
		id.String(),
		shipperUser.GetName(),
		shipperUser.GetUsername(),
		shipperUser.GetPassword(),
		shipperUser.GetPhone(),
		shipperUser.GetShipperId(),
		shipperUser.GetUserRoleId(),
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (shipUsrRep *ShipperUserRepo) Get(id string) (*pb.ShipperUser, error) {
	var (
		createdAt   time.Time
		updatedAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
		shipperUser pb.ShipperUser
		column      string
	)

	_, err := uuid.Parse(id)

	if err != nil {
		column = " username "
	} else {
		column = " id "
	}

	row := shipUsrRep.db.QueryRow(`
		SELECT  id,
				name,
				username,
				password,
				phone,

				is_blocked,

				shipper_id,
				user_role_id,

				created_at,
				updated_at
		FROM shipper_users
		WHERE `+column+`=$1 and
		deleted_at=0`, id,
	)

	err = row.Scan(
		&shipperUser.Id,
		&shipperUser.Name,
		&shipperUser.Username,
		&shipperUser.Password,
		&shipperUser.Phone,
		&shipperUser.IsBlocked,

		&shipperUser.ShipperId,
		&shipperUser.UserRoleId,

		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	shipperUser.CreatedAt = createdAt.Format(layoutDate)
	shipperUser.UpdatedAt = updatedAt.Format(layoutDate)

	if err != nil {
		return nil, err
	}

	return &shipperUser, nil
}

func (shipUsrRep *ShipperUserRepo) GetAll(page, limit uint64, shipperID, userRoleID, search string) ([]*pb.ShipperUser, uint64, error) {
	var (
		count        uint64
		createdAt    time.Time
		updatedAt    time.Time
		layoutDate   string = "2006-01-02 15:04:05"
		shipperUsers []*pb.ShipperUser
		filter       string
	)

	offset := (page - 1) * limit
	params := map[string]interface{}{
		"shipper_id":   shipperID,
		"user_role_id": userRoleID,
		"limit":        limit,
		"offset":       offset,
		"search":       search,
	}

	if shipperID != "" {
		filter += ` AND shipper_id=:shipper_id  `
	}

	if userRoleID != "" {
		filter += ` AND user_role_id=:user_role_id  `
	}

	if search != "" {
		filter += ` AND (name ILIKE '%' || :search || '%' OR
			username ILIKE '%' || :search || '%' OR
			phone ILIKE '%' || :search || '%')  `
	}

	query := `
		SELECT  id,
				name,
				username,
				password,
				phone,
				is_blocked,
				shipper_id,
				user_role_id,
				created_at,
				updated_at
		FROM shipper_users
		WHERE deleted_at=0
		` + filter + `
		ORDER BY created_at DESC LIMIT :limit OFFSET :offset`

	rows, err := shipUsrRep.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var shipperUser pb.ShipperUser

		err = rows.Scan(
			&shipperUser.Id,
			&shipperUser.Name,
			&shipperUser.Username,
			&shipperUser.Password,
			&shipperUser.Phone,
			&shipperUser.IsBlocked,

			&shipperUser.ShipperId,
			&shipperUser.UserRoleId,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		shipperUser.CreatedAt = createdAt.Format(layoutDate)
		shipperUser.UpdatedAt = updatedAt.Format(layoutDate)

		shipperUsers = append(shipperUsers, &shipperUser)
	}

	countQuery := `SELECT count(1) FROM shipper_users
		WHERE deleted_at=0 ` + filter

	row, err := shipUsrRep.db.NamedQuery(countQuery, params)
	if err != nil {
		return nil, 0, err
	}

	if row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return nil, 0, err
		}
		row.Close()
	}

	return shipperUsers, count, nil
}

func (shipUsrRep *ShipperUserRepo) Update(shipperUser *pb.ShipperUser) error {
	updateQuery :=
		`UPDATE shipper_users
		 SET
			name=$1,
			username=$2,
			phone=$3,
			user_role_id=$4,
			updated_at=CURRENT_TIMESTAMP,
			is_blocked=$5
		WHERE id=$6`

	result, err := shipUsrRep.db.Exec(
		updateQuery,
		shipperUser.GetName(),
		shipperUser.GetUsername(),
		shipperUser.GetPhone(),
		shipperUser.GetUserRoleId(),
		shipperUser.GetIsBlocked(),
		shipperUser.GetId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (shipUsrRep *ShipperUserRepo) Delete(id, shipperId string) error {
	result, err := shipUsrRep.db.Exec(`UPDATE shipper_users SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int WHERE id=$1 AND deleted_at=0 `, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (shipUsrRep *ShipperUserRepo) GetByUsername(username string) (*pb.ShipperUser, error) {
	var (
		createdAt   time.Time
		updatedAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
		shipperUser pb.ShipperUser
	)

	row := shipUsrRep.db.QueryRow(`
		SELECT id,
			name,
			username,
			password,
			phone,
			is_blocked,
			shipper_id,
			user_role_id,
			created_at,
			updated_at
		FROM shipper_users 
		WHERE username=$1
		AND deleted_at=0`, username)

	err := row.Scan(
		&shipperUser.Id,
		&shipperUser.Name,
		&shipperUser.Username,
		&shipperUser.Password,
		&shipperUser.Phone,
		&shipperUser.IsBlocked,
		&shipperUser.ShipperId,
		&shipperUser.UserRoleId,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	shipperUser.CreatedAt = createdAt.Format(layoutDate)
	shipperUser.UpdatedAt = updatedAt.Format(layoutDate)

	if err != nil {
		return nil, err
	}

	return &shipperUser, nil
}

func (shipUsrRep *ShipperUserRepo) ChangePassword(id, password string) error {
	result, err := shipUsrRep.db.Exec(`update shipper_users set password = $1 where id=$2`, password, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
