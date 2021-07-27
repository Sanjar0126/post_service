package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"gitlab.udevs.io/delever/delever_user_service/pkg/etc"
	"gitlab.udevs.io/delever/delever_user_service/pkg/helper"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type paymeInfoRepo struct {
	db *sqlx.DB
}

func NewPaymeInfoRepo(db *sqlx.DB) repo.PaymeInfoStorageI {
	return &paymeInfoRepo{
		db: db,
	}
}

func (pi *paymeInfoRepo) Create(payme *pb.Payme) (string, error) {
	token := etc.GeneratePaymeToken(payme.Login, payme.Key)
	nullableBranchID := helper.NullString(payme.BranchId)

	insertNew :=
		`INSERT INTO payme_info
		(
			shipper_id,
			branch_id,
			merchant_id,
			login,
			key,
			token
		)
		
		VALUES 
		($1, $2, $3, $4, $5, $6)`

	_, err := pi.db.Exec(
		insertNew,
		payme.ShipperId,
		nullableBranchID,
		payme.MerchantId,
		payme.Login,
		payme.Key,
		token,
	)

	if err != nil {
		return "", err
	}

	return payme.ShipperId, nil
}

func (pi *paymeInfoRepo) Get(id string, branchID string) (*pb.Payme, error) {
	var (
		createdAt  time.Time
		layoutDate string = "2006-01-02 15:04:05"
		payme      pb.Payme
		column     string
	)

	params := map[string]interface{}{
		"shipper_id": id,
		"branch_id":  branchID,
	}

	if branchID != "" {
		column += `and branch_id = :branch_id`
	}

	query := `
		SELECT 
			merchant_id,
			login,
			key,
			created_at
		FROM payme_info
		WHERE shipper_id= :shipper_id
	` + column

	stmt, err := pi.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(params)

	err = row.Scan(
		&payme.MerchantId,
		&payme.Login,
		&payme.Key,
		&createdAt,
	)
	if err == sql.ErrNoRows {
		query := pi.db.QueryRow(`
		SELECT 
			merchant_id,
			login,
			key,
			created_at
		FROM payme_info
		WHERE shipper_id=$1
	`, id)

		err = query.Scan(
			&payme.MerchantId,
			&payme.Login,
			&payme.Key,
			&createdAt,
		)
		branchID = ""
		if err != nil {
			return nil, err
		}
	}

	payme.CreatedAt = createdAt.Format(layoutDate)
	payme.BranchId = branchID
	payme.ShipperId = id

	return &payme, nil
}

func (pi *paymeInfoRepo) GetAll(shipperID string, branchIDs []string, page, limit uint64) ([]*pb.Payme, uint64, error) {
	var (
		column     string
		paymes     []*pb.Payme
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
			login,
			key,
			created_at
		FROM payme_info
		WHERE shipper_id=:shipper_id` + column + `
		OFFSET :offset
		LIMIT :limit `

	rows, err := pi.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var payme pb.Payme

		err := rows.Scan(
			&payme.BranchId,
			&payme.MerchantId,
			&payme.Login,
			&payme.Key,
			&createdAt,
		)

		if err != nil {
			return nil, 0, err
		}

		payme.CreatedAt = createdAt.Format(layoutDate)

		paymes = append(paymes, &payme)
	}

	stmt, err := pi.db.PrepareNamed(`SELECT count(1) FROM payme_info WHERE shipper_id=:shipper_id ` + column)
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

	return paymes, count, nil
}

func (pi *paymeInfoRepo) Update(payme *pb.Payme) error {
	var query string
	token := etc.GeneratePaymeToken(payme.Login, payme.Key)
	nullableBranchID := helper.NullString(payme.BranchId)

	if payme.BranchId == "" {
		query += ` shipper_id=$5 and branch_id is null or branch_id=$6 `
	} else {
		query += ` shipper_id=$5 and branch_id=$6 `
	}

	updateQuery := `
		UPDATE payme_info
		SET
			merchant_id=$1,
			login=$2,
			key=$3,
			token=$4,
			updated_at=CURRENT_TIMESTAMP
		WHERE 
	` + query

	result, err := pi.db.Exec(
		updateQuery,
		payme.MerchantId,
		payme.Login,
		payme.Key,
		token,
		payme.ShipperId,
		nullableBranchID,
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pi *paymeInfoRepo) Delete(id, branchID string) error {
	var query string
	nullableBranchID := helper.NullString(branchID)

	if branchID == "" {
		query += `shipper_id=$1 and branch_id is null or branch_id=$2`
	} else {
		query += `shipper_id=$1 and branch_id=$2`
	}

	result, err := pi.db.Exec(`DELETE FROM payme_info WHERE `+query, id, nullableBranchID)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pi *paymeInfoRepo) GetShipperByCredentials(token string) (string, error) {
	var shipperID string

	row := pi.db.QueryRow(`
		SELECT 
			shipper_id
		FROM payme_info
		WHERE token=$1
	`, token)

	err := row.Scan(
		&shipperID,
	)

	if err != nil {
		return "", err
	}

	return shipperID, nil
}
