package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"gitlab.udevs.io/delever/delever_user_service/pkg/helper"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type clickInfoRepo struct {
	db *sqlx.DB
}

func NewClickInfoRepo(db *sqlx.DB) repo.ClickInfoStorageI {
	return &clickInfoRepo{
		db: db,
	}
}

func (ci *clickInfoRepo) Create(click *pb.Click) (string, error) {
	nullableBranchID := helper.NullString(click.BranchId)
	insertNew :=
		`INSERT INTO click_info
		(
			shipper_id,
			branch_id,
			merchant_id,
			service_id,
			merchant_user_id,
			key
		)
		
		VALUES 
		($1, $2, $3, $4, $5, $6)`

	_, err := ci.db.Exec(
		insertNew,
		click.ShipperId,
		nullableBranchID,
		click.MerchantId,
		click.ServiceId,
		click.GetMerchantUserId(),
		click.Key,
	)

	if err != nil {
		return "", err
	}

	return click.ShipperId, nil
}

func (ci *clickInfoRepo) Get(id, branchID string) (*pb.Click, error) {
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		click      pb.Click
		column     string
	)

	params := map[string]interface{}{
		"shipper_id": id,
		"branch_id":  branchID,
	}

	if branchID != "" {
		column += ` and branch_id=:branch_id `
	}

	query := `
		SELECT 
			merchant_id,
			service_id,
			merchant_user_id,
			key,
			created_at
		FROM click_info
		WHERE shipper_id=:shipper_id
	` + column

	stmt, err := ci.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(params)

	err = row.Scan(
		&click.MerchantId,
		&click.ServiceId,
		&click.MerchantUserId,
		&click.Key,
		&createdAt,
	)
	if err == sql.ErrNoRows {
		query := ci.db.QueryRow(`
		SELECT 
			merchant_id,
			service_id,
			merchant_user_id,
			key,
			created_at
		FROM click_info
		WHERE shipper_id=$1 and branch_id is null
	`, id)

		err = query.Scan(
			&click.MerchantId,
			&click.ServiceId,
			&click.MerchantUserId,
			&click.Key,
			&createdAt,
		)
		branchID = ""
		if err != nil {
			return nil, err
		}
	}

	click.BranchId = branchID
	click.ShipperId = id
	click.CreatedAt = createdAt.Format(layoutDate)

	return &click, nil
}

func (ci *clickInfoRepo) GetAll(shipperID string, branchIDs []string, page, limit uint64) ([]*pb.Click, uint64, error) {
	var (
		column     string
		clicks     []*pb.Click
		count      uint64
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
	)

	offset := (page - 1) * limit

	params := map[string]interface{}{
		"shipper_id": shipperID,
		"branch_ids": pq.Array(branchIDs),
		"offset":     offset,
		"limit":      limit,
	}

	if len(branchIDs) > 0 {
		column += ` and branch_id = ANY(:branch_ids) `
	}

	query := `
		SELECT 
			branch_id,
			merchant_id,
			service_id,
			merchant_user_id,
			key,
			created_at
		FROM click_info
		WHERE shipper_id=:shipper_id` + column + `
		OFFSET :offset
		LIMIT :limit `

	rows, err := ci.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var click pb.Click

		err := rows.Scan(
			&click.BranchId,
			&click.MerchantId,
			&click.ServiceId,
			&click.MerchantUserId,
			&click.Key,
			&createdAt,
		)

		if err != nil {
			return nil, 0, err
		}

		click.CreatedAt = createdAt.Format(layoutDate)

		clicks = append(clicks, &click)
	}

	stmt, err := ci.db.PrepareNamed(`SELECT count(1) FROM click_info WHERE shipper_id=:shipper_id ` + column)
	if err != nil {
		return nil, 0, err
	}

	row := stmt.QueryRow(params)

	err = row.Scan(
		&count,
	)
	if err != nil {
		return nil, 0, err
	}

	return clicks, count, nil
}

func (ci *clickInfoRepo) Update(click *pb.Click) error {
	var query string
	branchID := helper.NullString(click.BranchId)

	if click.BranchId == "" {
		query += `shipper_id=$5 and branch_id is null or branch_id=$6`
	} else {
		query += `shipper_id=$5 and branch_id=$6`
	}

	updateQuery := `
		UPDATE click_info
		SET
			merchant_id=$1,
			service_id=$2,
			merchant_user_id=$3,
			key=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE
	` + query

	result, err := ci.db.Exec(
		updateQuery,
		click.MerchantId,
		click.ServiceId,
		click.MerchantUserId,
		click.Key,
		click.ShipperId,
		branchID,
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ci *clickInfoRepo) Delete(id, branchID string) error {
	var query string
	nullableBranchID := helper.NullString(branchID)

	if branchID == "" {
		query += `shipper_id=$1 and branch_id is null or branch_id=$2`
	} else {
		query += `shipper_id=$1 and branch_id=$2`
	}

	result, err := ci.db.Exec(`DELETE FROM click_info WHERE `+query, id, nullableBranchID)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ci *clickInfoRepo) GetShipperAndKeyByCredentials(serviceID int64) (string, string, error) {
	var shipperID, key string

	row := ci.db.QueryRow(`
		SELECT 
			shipper_id,
			key
		FROM click_info
		WHERE service_id=$1
	`, serviceID)

	err := row.Scan(
		&shipperID,
		&key,
	)

	if err != nil {
		return "", "", err
	}

	return shipperID, key, nil
}
