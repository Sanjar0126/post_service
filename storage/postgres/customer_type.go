package postgres

import (
	"database/sql"
	"fmt"
	pb "genproto/user_service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type customerTypeRepo struct {
	db *sqlx.DB
}

func NewCustomerTypeRepo(db *sqlx.DB) repo.CustomerTypeI {
	return &customerTypeRepo{
		db: db,
	}
}

func (r *customerTypeRepo) Create(customer_type *pb.CustomerType) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO customer_types(
            id,
            name,
            phone_number
            )
        VALUES ($1, $2, $3)`

	_, err = r.db.Exec(
		insertNew,
		id.String(),
		customer_type.GetName(),
		customer_type.GetPhoneNumber(),
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *customerTypeRepo) Update(customer_type *pb.CustomerType) error {
	query := `UPDATE customer_types 
          SET
            name = $1,
            phone_number = $2,
            updated_at = current_timestamp
          WHERE id = $3 `

	result, err := r.db.Exec(
		query,
		customer_type.Name,
		customer_type.PhoneNumber,
		customer_type.Id,
	)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *customerTypeRepo) Get(id string) (*pb.CustomerType, error) {
	var customer_type pb.CustomerType

	query := `SELECT
          id,
          name,
          phone_number
        FROM customer_types WHERE id = $1`

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&customer_type.Id,
		&customer_type.Name,
		&customer_type.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &customer_type, nil
}

func (r *customerTypeRepo) GetAll(page, limit int64, name string) ([]*pb.CustomerType, int64, error) {
	var (
		customer_types []*pb.CustomerType
		filter         string
		count          int64
	)

	if name != "" {
		filter += fmt.Sprintf(` AND name ilike '%s' `, "%"+name+"%")
	}

	err := r.db.Get(&count, fmt.Sprintf(`SELECT count(1) FROM customer_types WHERE true %s `, filter))
	if err != nil {
		return nil, 0, err
	}

	filter += fmt.Sprintf(" OFFSET %d  LIMIT %d ", limit*(page-1), limit)

	query := `SELECT 
          id,
          name,
          phone_number
        FROM customer_types WHERE deleted_at is null  %s `

	rows, err := r.db.Query(fmt.Sprintf(query, filter))
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var customer_type pb.CustomerType

		err = rows.Scan(
			&customer_type.Id,
			&customer_type.Name,
			&customer_type.PhoneNumber,
		)

		if err != nil {
			return nil, 0, err
		}

		customer_types = append(customer_types, &customer_type)
	}

	return customer_types, count, err
}

func (r *customerTypeRepo) Delete(id string) error {
	_, err := r.db.Exec(`UPDATE customer_types  SET deleted_at = current_timestamp WHERE id = $1`, id)
	return err
}
