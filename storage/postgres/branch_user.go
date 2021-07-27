package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/pkg/etc"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type BranchUserRepo struct {
	db *sqlx.DB
}

// NewBranchUserRepo ...
func NewBranchUserRepo(db *sqlx.DB) repo.BranchUserStorageI {
	return &BranchUserRepo{
		db: db,
	}
}

func (branchUsrRep *BranchUserRepo) Create(branchUser *pb.BranchUser) (string, error) {
	var fcmToken sql.NullString = etc.NullString(branchUser.FcmToken)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		branch_users
		(
			id,
			name,
			phone,
			shipper_id,
			branch_id,
			user_role_id,
			fcm_token
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7)`
	_, err = branchUsrRep.db.Exec(
		insertNew,
		id.String(),
		branchUser.GetName(),
		branchUser.GetPhone(),
		branchUser.GetShipperId(),
		branchUser.GetBranchId(),
		branchUser.GetUserRoleId(),
		fcmToken,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (branchUsrRep *BranchUserRepo) Get(id string) (*pb.BranchUser, error) {
	var (
		createdAt  time.Time
		updatedAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		branchUser pb.BranchUser
		fcmToken   sql.NullString
		platformID sql.NullString
	)

	row := branchUsrRep.db.QueryRow(`
		SELECT  id,
				name,
				phone,
				is_active,
				is_blocked,
				shipper_id,
				branch_id,
				user_role_id,
				fcm_token,
				created_at,
				updated_at,
				platform_id
		FROM branch_users
		WHERE id=$1 and
		deleted_at=0`, id,
	)

	err := row.Scan(
		&branchUser.Id,
		&branchUser.Name,
		&branchUser.Phone,
		&branchUser.IsActive,
		&branchUser.IsBlocked,
		&branchUser.ShipperId,
		&branchUser.BranchId,
		&branchUser.UserRoleId,
		&fcmToken,
		&createdAt,
		&updatedAt,
		&platformID,
	)
	if err != nil {
		return nil, err
	}

	branchUser.CreatedAt = createdAt.Format(layoutDate)
	branchUser.UpdatedAt = updatedAt.Format(layoutDate)
	branchUser.FcmToken = etc.StringValue(fcmToken)
	branchUser.PlatformId = etc.StringValue(platformID)

	if err != nil {
		return nil, err
	}

	return &branchUser, nil
}

func (branchUsrRep *BranchUserRepo) GetAll(page, limit uint64, shipperId, search, branchId, userRoleID string) ([]*pb.BranchUser, uint64, error) {
	var (
		count       uint64
		createdAt   time.Time
		updatedAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
		branchUsers []*pb.BranchUser
		filter      string
		fcmToken    sql.NullString
	)

	offset := (page - 1) * limit
	params := map[string]interface{}{
		"shipper_id":   shipperId,
		"limit":        limit,
		"offset":       offset,
		"search":       search,
		"branch_id":    branchId,
		"user_role_id": userRoleID,
	}

	if search != "" {
		filter += ` AND (name ILIKE '%' || :search || '%' OR
			phone ILIKE '%' || :search || '%')  `
	}

	if branchId != "" {
		filter += ` AND branch_id=:branch_id `
	}

	if userRoleID != "" {
		filter += ` AND user_role_id=:user_role_id `
	}

	query := `
		SELECT  id,
				name,
				phone,
				is_active,
				is_blocked,
				shipper_id,
				branch_id,
				user_role_id,
				fcm_token,
				created_at,
				updated_at
		FROM branch_users
		WHERE shipper_id=:shipper_id AND deleted_at=0
		` + filter + `
		ORDER BY created_at DESC LIMIT :limit OFFSET :offset`

	rows, err := branchUsrRep.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var branchUser pb.BranchUser

		err = rows.Scan(
			&branchUser.Id,
			&branchUser.Name,
			&branchUser.Phone,
			&branchUser.IsActive,
			&branchUser.IsBlocked,
			&branchUser.ShipperId,
			&branchUser.BranchId,
			&branchUser.UserRoleId,
			&fcmToken,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		branchUser.CreatedAt = createdAt.Format(layoutDate)
		branchUser.UpdatedAt = updatedAt.Format(layoutDate)
		branchUser.FcmToken = etc.StringValue(fcmToken)

		branchUsers = append(branchUsers, &branchUser)
	}

	countQuery := `SELECT count(1) FROM branch_users
		WHERE shipper_id=:shipper_id AND deleted_at=0 ` + filter

	row, err := branchUsrRep.db.NamedQuery(countQuery, params)
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

	return branchUsers, count, nil
}

func (branchUsrRep *BranchUserRepo) Update(branchUser *pb.BranchUser) error {
	updateQuery :=
		`UPDATE branch_users
		 SET
			name=$1,
			phone=$2,
			branch_id=$3,
			user_role_id=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=$5`

	result, err := branchUsrRep.db.Exec(
		updateQuery,
		branchUser.GetName(),
		branchUser.GetPhone(),
		branchUser.GetBranchId(),
		branchUser.GetUserRoleId(),
		branchUser.GetId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (branchUsrRep *BranchUserRepo) Delete(id, shipperId string) error {
	result, err := branchUsrRep.db.Exec(`UPDATE branch_users SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where shipper_id=$1 AND id=$2 AND deleted_at=0 `, shipperId, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (branchUsrRep *BranchUserRepo) GetByPhone(phone, shipperId string) (*pb.BranchUser, error) {
	var (
		createdAt  time.Time
		updatedAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		branchUser pb.BranchUser
		fcmToken   sql.NullString
	)

	row := branchUsrRep.db.QueryRow(`
		SELECT  id,
				name,
				phone,
				is_active,
				is_blocked,
				shipper_id,
				branch_id,
				user_role_id,
				fcm_token,
				created_at,
				updated_at
		FROM branch_users
		WHERE shipper_id=$1 
		AND phone=$2
		AND deleted_at=0`, shipperId, phone,
	)

	err := row.Scan(
		&branchUser.Id,
		&branchUser.Name,
		&branchUser.Phone,

		&branchUser.IsActive,
		&branchUser.IsBlocked,

		&branchUser.ShipperId,
		&branchUser.BranchId,
		&branchUser.UserRoleId,
		&fcmToken,

		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	branchUser.CreatedAt = createdAt.Format(layoutDate)
	branchUser.UpdatedAt = updatedAt.Format(layoutDate)
	branchUser.FcmToken = etc.StringValue(fcmToken)

	if err != nil {
		return nil, err
	}

	return &branchUser, nil
}

func (branchUsrRep *BranchUserRepo) UpdateFcmToken(id, shipperId, fcmToken, platformID string) error {
	result, err := branchUsrRep.db.Exec(`
		UPDATE branch_users SET 
			fcm_token = $1,
			platform_id = $2
		WHERE shipper_id=$3 
		AND id = $4`,
		fcmToken,
		platformID,
		shipperId,
		id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (branchUsrRep *BranchUserRepo) DeleteFcmToken(id string) error {
	result, err := branchUsrRep.db.Exec(`
		UPDATE branch_users SET 
			fcm_token=Null,
			platform_id=Null
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
