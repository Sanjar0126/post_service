package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type jowiCredentialsRepo struct {
	db *sqlx.DB
}

// NewJowiCredentialsRepo ...
func NewJowiCredentialsRepo(db *sqlx.DB) repo.JowiCredentialsStorageI {
	return &jowiCredentialsRepo{
		db: db,
	}
}

func (jc *jowiCredentialsRepo) Create(jowiC *pb.JowiCredentials) (string, error) {
	insertNew :=
		`INSERT INTO
		jowi_credentials
		(
			shipper_id,
			dispatcher_id
		)
		VALUES
		($1, $2)`

	_, err := jc.db.Exec(
		insertNew,
		jowiC.GetShipperId(),
		jowiC.GetDispatcherId(),
	)
	if err != nil {
		return "", err
	}

	return jowiC.GetShipperId(), nil
}

func (jc *jowiCredentialsRepo) Get(shipperID string) (*pb.JowiCredentials, error) {
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		jowiC      pb.JowiCredentials
	)

	row := jc.db.QueryRow(`
		SELECT  shipper_id,
				dispatcher_id,
				created_at
		FROM jowi_credentials
		WHERE shipper_id=$1
		AND deleted_at = 0`, shipperID,
	)

	err := row.Scan(
		&jowiC.ShipperId,
		&jowiC.DispatcherId,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	jowiC.CreatedAt = createdAt.Format(layoutDate)

	return &jowiC, nil
}

func (jc *jowiCredentialsRepo) Update(jowiC *pb.JowiCredentials) error {
	updateQuery :=
		`UPDATE jowi_credentials
		 SET
			dispatcher_id=$1,
			updated_at=CURRENT_TIMESTAMP
		WHERE shipper_id=$2
		AND deleted_at=0`

	result, err := jc.db.Exec(
		updateQuery,
		jowiC.GetDispatcherId(),
		jowiC.GetShipperId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (jc *jowiCredentialsRepo) Delete(shipperID string) error {
	result, err := jc.db.Exec(`UPDATE jowi_credentials SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where shipper_id=$1 and deleted_at = 0`, shipperID)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
