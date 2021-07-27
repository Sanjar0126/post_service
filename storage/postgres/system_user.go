package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type systemUserRepo struct {
	db *sqlx.DB
}

// NewSystemUserRepo ...
func NewSystemUserRepo(db *sqlx.DB) repo.SystemUserStorageI {
	return &systemUserRepo{
		db: db,
	}
}

func (sysUsrRep *systemUserRepo) Create(systemUser *pb.SystemUser) (string, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		system_users
		(
			id,
			name,
			username,
			password,
			phone
		)
		VALUES
		($1, $2, $3, $4, $5)`
	_, err = sysUsrRep.db.Query(
		insertNew,
		id.String(),
		systemUser.GetName(),
		systemUser.GetUsername(),
		systemUser.GetPassword(),
		systemUser.GetPhone(),
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (sysUsrRep *systemUserRepo) Get(id string) (*pb.SystemUser, error) {
	var (
		createdAt  time.Time
		updatedAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		systemUser pb.SystemUser
		column     string
	)

	_, err := uuid.Parse(id)

	if err != nil {
		column = " username "
	} else {
		column = " id "
	}

	row := sysUsrRep.db.QueryRow(`
		SELECT  id,
				name,
				username,
				password,
				phone,
				is_blocked,
				created_at,
				updated_at
		FROM system_users
		WHERE `+column+`=$1 and
		deleted_at=0`, id,
	)

	err = row.Scan(
		&systemUser.Id,
		&systemUser.Name,
		&systemUser.Username,
		&systemUser.Password,
		&systemUser.Phone,
		&systemUser.IsBlocked,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	systemUser.CreatedAt = createdAt.Format(layoutDate)
	systemUser.UpdatedAt = updatedAt.Format(layoutDate)

	if err != nil {
		return nil, err
	}

	return &systemUser, nil
}

func (sysUsrRep *systemUserRepo) GetAll(page, limit uint64, search string) ([]*pb.SystemUser, uint64, error) {
	var (
		count       uint64
		createdAt   time.Time
		updatedAt   time.Time
		layoutDate  string = "2006-01-02 15:04:05"
		systemUsers []*pb.SystemUser
		filter      string
	)

	offset := (page - 1) * limit
	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"search": search,
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
				created_at,
				updated_at
		FROM system_users
		WHERE deleted_at=0
		` + filter + ` 
		ORDER BY created_at DESC LIMIT :limit OFFSET :offset`

	rows, err := sysUsrRep.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var systemUser pb.SystemUser
		err = rows.Scan(
			&systemUser.Id,
			&systemUser.Name,
			&systemUser.Username,
			&systemUser.Password,
			&systemUser.Phone,
			&systemUser.IsBlocked,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		systemUser.CreatedAt = createdAt.Format(layoutDate)
		systemUser.UpdatedAt = updatedAt.Format(layoutDate)

		systemUsers = append(systemUsers, &systemUser)
	}

	countQuery := `SELECT count(1) FROM system_users 
		WHERE deleted_at=0 ` + filter

	row, err := sysUsrRep.db.NamedQuery(countQuery, params)
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

	return systemUsers, count, nil
}

func (sysUsrRep *systemUserRepo) Update(systemUser *pb.SystemUser) error {
	updateQuery :=
		`UPDATE system_users
		 SET
			name=$1,
			username=$2,
			phone=$3,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=$4`

	result, err := sysUsrRep.db.Exec(
		updateQuery,
		systemUser.GetName(),
		systemUser.GetUsername(),
		systemUser.GetPhone(),
		systemUser.GetId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (sysUsrRep *systemUserRepo) Delete(id string) error {
	result, err := sysUsrRep.db.Exec(`UPDATE system_users SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where id=$1 and deleted_at=0`, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (sysUsrRep *systemUserRepo) GetByUsername(username string) (*pb.SystemUser, error) {
	var (
		createdAt  time.Time
		updatedAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		systemUser pb.SystemUser
	)

	row := sysUsrRep.db.QueryRow(`select id,
			name,
			username,
			password,
			phone,
			is_blocked,
			created_at,
			updated_at
		from system_users where username=$1 and deleted_at=0`, username)

	err := row.Scan(
		&systemUser.Id,
		&systemUser.Name,
		&systemUser.Username,
		&systemUser.Password,
		&systemUser.Phone,
		&systemUser.IsBlocked,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	systemUser.CreatedAt = createdAt.Format(layoutDate)
	systemUser.UpdatedAt = updatedAt.Format(layoutDate)

	if err != nil {
		return nil, err
	}

	return &systemUser, nil
}

func (sysUsrRep *systemUserRepo) ChangePassword(id, password string) error {
	result, err := sysUsrRep.db.Exec(`update system_users set password = $1 where id=$2`, password, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
