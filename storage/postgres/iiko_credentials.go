package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type iikoCredentialsRepo struct {
	db *sqlx.DB
}

func NewIikoCredentialsRepo(db *sqlx.DB) repo.IikoCredentialsStorageI {
	return &iikoCredentialsRepo{
		db: db,
	}
}

func (ic *iikoCredentialsRepo) Create(iikoC *pb.IikoCredentials) (string, error) {
	insertNew :=
		`INSERT INTO
		iiko_credentials
		(
			shipper_id,
			api_login,
			dispatcher_id
		)
		VALUES
		($1, $2, $3)`

	_, err := ic.db.Exec(
		insertNew,
		iikoC.GetShipperId(),
		iikoC.GetApiLogin(),
		iikoC.GetDispatcherId(),
	)
	if err != nil {
		return "", err
	}

	return iikoC.GetShipperId(), nil
}

func (ic *iikoCredentialsRepo) Get(shipperID string) (*pb.IikoCredentials, error) {
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		iikoC      pb.IikoCredentials
	)

	row := ic.db.QueryRow(`
		SELECT  shipper_id,
				api_login,
				dispatcher_id,
				created_at
		FROM iiko_credentials
		WHERE shipper_id=$1
		AND deleted_at = 0`, shipperID,
	)

	err := row.Scan(
		&iikoC.ShipperId,
		&iikoC.ApiLogin,
		&iikoC.DispatcherId,
		&createdAt,
	)
	if err != nil {
		return nil, err
	}

	iikoC.CreatedAt = createdAt.Format(layoutDate)

	return &iikoC, nil
}

func (ic *iikoCredentialsRepo) Update(iikoC *pb.IikoCredentials) error {
	updateQuery :=
		`UPDATE iiko_credentials
		 SET
			api_login=$1,
			dispatcher_id=$2,
			updated_at=CURRENT_TIMESTAMP
		WHERE shipper_id=$3
		AND deleted_at=0`

	result, err := ic.db.Exec(
		updateQuery,
		iikoC.GetApiLogin(),
		iikoC.GetDispatcherId(),
		iikoC.GetShipperId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ic *iikoCredentialsRepo) Delete(shipperID string) error {
	result, err := ic.db.Exec(`UPDATE iiko_credentials SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where shipper_id=$1 and deleted_at=0`, shipperID)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
