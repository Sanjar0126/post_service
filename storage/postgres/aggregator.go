package postgres

import (
	"database/sql"
	"fmt"

	pb "genproto/user_service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type aggregatorRepo struct {
	db *sqlx.DB
}

// NewAggregatorRepo ...
func NewAggregatorRepo(db *sqlx.DB) repo.AggregatorStorageI {
	return &aggregatorRepo{
		db: db,
	}
}

func (br *aggregatorRepo) Create(aggregator *pb.Aggregator) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO aggregators(
            id,
            name,
            phone_number,
            shipper_id
            )
        VALUES ($1, $2, $3, $4)`

	_, err = br.db.Exec(
		insertNew,
		id.String(),
		aggregator.GetName(),
		aggregator.GetPhoneNumber(),
		aggregator.GetShipperId(),
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (br *aggregatorRepo) Update(aggregator *pb.Aggregator) error {
	query := `UPDATE aggregators 
          SET
            name = $1,
            phone_number = $2,
            shipper_id = $3
          WHERE id = $4 `

	result, err := br.db.Exec(
		query,
		aggregator.Name,
		aggregator.PhoneNumber,
		aggregator.ShipperId,
		aggregator.Id,
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (br *aggregatorRepo) Get(id string) (*pb.Aggregator, error) {
	var aggregator pb.Aggregator

	query := `SELECT
          id,
          name,
          phone_number,
          shipper_id
        FROM aggregators WHERE id = $1`

	row := br.db.QueryRow(query, id)
	err := row.Scan(
		&aggregator.Id,
		&aggregator.Name,
		&aggregator.PhoneNumber,
		&aggregator.ShipperId,
	)

	if err != nil {
		return nil, err
	}

	return &aggregator, nil
}

func (br *aggregatorRepo) GetAll(page, limit int64, name string, shipper_id string) ([]*pb.Aggregator, int64, error) {
	var (
		aggregators []*pb.Aggregator
		filter      string
		count       int64
	)

	if name != "" {
		filter += fmt.Sprintf(` AND name ilike '%s' `, "%"+name+"%")
	}

	if shipper_id != "" {
		filter += fmt.Sprintf(` AND shipper_id = '%s' `, shipper_id)
	}

	err := br.db.Get(&count, fmt.Sprintf(`SELECT count(1) FROM aggregators WHERE true %s `, filter))
	if err != nil {
		return nil, 0, err
	}

	filter += fmt.Sprintf(" OFFSET %d  LIMIT %d ", limit*(page-1), limit)

	query := `SELECT 
          id,
          name,
          phone_number,
          shipper_id
        FROM aggregators WHERE true %s `

	rows, err := br.db.Query(fmt.Sprintf(query, filter))
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var aggregator pb.Aggregator

		err = rows.Scan(
			&aggregator.Id,
			&aggregator.Name,
			&aggregator.PhoneNumber,
			&aggregator.ShipperId,
		)

		if err != nil {
			return nil, 0, err
		}

		aggregators = append(aggregators, &aggregator)
	}

	return aggregators, count, err
}

func (br *aggregatorRepo) Delete(id string) error {
	_, err := br.db.Exec(`DELETE FROM aggregators WHERE id = $1`, id)
	return err
}
