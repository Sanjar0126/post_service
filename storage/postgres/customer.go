package postgres

import (
	"database/sql"
	"fmt"
	pb "genproto/user_service"
	"time"

	"gitlab.udevs.io/delever/delever_user_service/pkg/etc"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type customerRepo struct {
	db *sqlx.DB
}

// NewCustomerRepo ...
func NewCustomerRepo(db *sqlx.DB) repo.CustomerStorageI {
	return &customerRepo{
		db: db,
	}
}

func (cm *customerRepo) Create(customer *pb.Customer) (string, error) {
	var (
		fcmToken         sql.NullString = etc.NullString(customer.FcmToken)
		tgChatID         sql.NullString = etc.NullString(customer.TgChatId)
		dateOfBirth      sql.NullString = etc.NullString(customer.DateOfBirth)
		customer_type_id sql.NullString = etc.NullString(customer.CustomerTypeId)
	)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		customers
		(
			id,
			name,
			phone,
			shipper_id,
			fcm_token,
			tg_chat_id,
			date_of_birth,
			is_aggregate,
            customer_type_id
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = cm.db.Exec(
		insertNew,
		id.String(),
		customer.GetName(),
		customer.GetPhone(),
		customer.ShipperId,
		fcmToken,
		tgChatID,
		dateOfBirth,
		customer.IsAggregate,
		customer_type_id,
	)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (cm *customerRepo) Get(id string) (*pb.Customer, error) {
	var (
		createdAt        time.Time
		layoutDateTime   string = "2006-01-02 15:04:05"
		customer         pb.Customer
		column           string
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		platformID       sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
	)
	_, err := uuid.Parse(id)

	if err != nil {
		column = " phone"
	} else {
		column = " id"
	}

	row := cm.db.QueryRow(`
		SELECT  id,
				name,
				phone,
				is_blocked,
				created_at,
				fcm_token,
				tg_chat_id,
				shipper_id,
				bot_language,
				platform_id,
				to_char(date_of_birth, 'YYYY-MM-DD'),
				is_aggregate,
                customer_type_id 
		FROM customers
		WHERE deleted_at = 0 AND `+column+`=$1 `, id,
	)

	err = row.Scan(
		&customer.Id,
		&customer.Name,
		&customer.Phone,
		&customer.IsBlocked,
		&createdAt,
		&fcmToken,
		&tgChatID,
		&customer.ShipperId,
		&customer.BotLanguage,
		&platformID,
		&dateOfBirth,
		&customer.IsAggregate,
		&customer_type_id,
	)

	if err != nil {
		return nil, err
	}

	customer.CreatedAt = createdAt.Format(layoutDateTime)
	customer.FcmToken = etc.StringValue(fcmToken)
	customer.TgChatId = etc.StringValue(tgChatID)
	customer.PlatformId = etc.StringValue(platformID)
	customer.DateOfBirth = etc.StringValue(dateOfBirth)
	customer.CustomerTypeId = etc.StringValue(customer_type_id)

	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (cm *customerRepo) GetAll(shipperId, search, customerTypeId string, page, limit uint64) ([]*pb.Customer, uint64, error) {
	var (
		count            uint64
		column           string
		createdAt        time.Time
		layoutDateTime   string = "2006-01-02 15:04:05"
		customers        []*pb.Customer
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
	)

	offset := (page - 1) * limit

	if search != "" {
		search = "%" + search + "%"
	}

	params := map[string]interface{}{
		"shipper_id":       shipperId,
		"offset":           offset,
		"limit":            limit,
		"search":           search,
		"customer_type_id": customerTypeId,
	}

	if search != "" {
		column += ` AND (phone ILIKE :search OR name ILIKE :search) `
	}

	if customerTypeId != "" {
		column += ` AND customer_type_id = :customer_type_id `
	}

	queryCount := `
		SELECT count(1) FROM customers
		WHERE shipper_id=:shipper_id and deleted_at=0 ` + column

	row, err := cm.db.NamedQuery(queryCount, params)
	if err != nil {
		return nil, 0, err
	}

	if row.Next() {
		err = row.Scan(
			&count,
		)
		if err != nil {
			return nil, 0, err
		}
		row.Close()
	}

	q := `SELECT 
			id, 
			name,
			phone, 
			is_blocked,
			created_at,
			fcm_token,
			tg_chat_id,
			shipper_id,
			bot_language,
			date_of_birth,
			is_aggregate,
            customer_type_id
		FROM customers
		WHERE shipper_id=:shipper_id 
		AND deleted_at=0 ` + column + ` ORDER BY created_at DESC LIMIT :limit OFFSET :offset`
	rows, err := cm.db.NamedQuery(q, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var c pb.Customer
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Phone,
			&c.IsBlocked,
			&createdAt,
			&fcmToken,
			&tgChatID,
			&c.ShipperId,
			&c.BotLanguage,
			&dateOfBirth,
			&c.IsAggregate,
			&customer_type_id,
		)
		if err != nil {
			return nil, 0, err
		}

		c.CreatedAt = createdAt.Format(layoutDateTime)
		c.FcmToken = etc.StringValue(fcmToken)
		c.TgChatId = etc.StringValue(tgChatID)
		c.DateOfBirth = etc.StringValue(dateOfBirth)
		c.CustomerTypeId = etc.StringValue(customer_type_id)

		customers = append(customers, &c)
	}

	return customers, count, nil
}

func (cm *customerRepo) GetAggregate(shipperId, search string, page, limit uint64) ([]*pb.Customer, uint64, error) {
	var (
		count            uint64
		column           string
		createdAt        time.Time
		layoutDateTime   string = "2006-01-02 15:04:05"
		customers        []*pb.Customer
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
	)

	offset := (page - 1) * limit

	if search != "" {
		search = "%" + search + "%"
	}

	params := map[string]interface{}{
		"shipper_id": shipperId,
		"offset":     offset,
		"limit":      limit,
		"search":     search,
	}

	if search != "" {
		column += ` AND (phone ILIKE :search OR name ILIKE :search) `
	}

	queryCount := `
		SELECT count(1) FROM customers
		WHERE is_aggregate=true AND shipper_id=:shipper_id AND deleted_at=0 ` + column

	row, err := cm.db.NamedQuery(queryCount, params)
	if err != nil {
		return nil, 0, err
	}

	if row.Next() {
		err = row.Scan(
			&count,
		)
		if err != nil {
			return nil, 0, err
		}
		row.Close()
	}

	q := `SELECT 
			id, 
			name,
			phone, 
			is_blocked,
			created_at,
			fcm_token,
			tg_chat_id,
			shipper_id,
			bot_language,
			date_of_birth,
			is_aggregate,
            customer_type_id
		FROM customers
		WHERE is_aggregate=true AND shipper_id=:shipper_id
		AND deleted_at=0 ` + column + ` ORDER BY created_at DESC LIMIT :limit OFFSET :offset`
	rows, err := cm.db.NamedQuery(q, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var c pb.Customer
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Phone,
			&c.IsBlocked,
			&createdAt,
			&fcmToken,
			&tgChatID,
			&c.ShipperId,
			&c.BotLanguage,
			&dateOfBirth,
			&c.IsAggregate,
			&customer_type_id,
		)
		if err != nil {
			return nil, 0, err
		}

		c.CreatedAt = createdAt.Format(layoutDateTime)
		c.FcmToken = etc.StringValue(fcmToken)
		c.TgChatId = etc.StringValue(tgChatID)
		c.DateOfBirth = etc.StringValue(dateOfBirth)
		c.CustomerTypeId = etc.StringValue(customer_type_id)

		customers = append(customers, &c)
	}

	return customers, count, nil
}

func (cm *customerRepo) GetNonAggregate(shipperId, search string, page, limit uint64) ([]*pb.Customer, uint64, error) {
	var (
		count            uint64
		column           string
		createdAt        time.Time
		layoutDateTime   string = "2006-01-02 15:04:05"
		customers        []*pb.Customer
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
	)

	offset := (page - 1) * limit

	if search != "" {
		search = "%" + search + "%"
	}

	params := map[string]interface{}{
		"shipper_id": shipperId,
		"offset":     offset,
		"limit":      limit,
		"search":     search,
	}

	if search != "" {
		column += ` AND (phone ILIKE :search OR name ILIKE :search) `
	}

	queryCount := `
		SELECT count(1) FROM customers
		WHERE is_aggregate=false AND shipper_id=:shipper_id AND deleted_at=0 ` + column

	row, err := cm.db.NamedQuery(queryCount, params)
	if err != nil {
		return nil, 0, err
	}

	if row.Next() {
		err = row.Scan(
			&count,
		)
		if err != nil {
			return nil, 0, err
		}
		row.Close()
	}

	q := `SELECT 
			id, 
			name,
			phone, 
			is_blocked,
			created_at,
			fcm_token,
			tg_chat_id,
			shipper_id,
			bot_language,
			date_of_birth,
			is_aggregate,
            customer_type_id
		FROM customers
		WHERE is_aggregate=false AND shipper_id=:shipper_id
		AND deleted_at=0 ` + column + ` ORDER BY created_at DESC LIMIT :limit OFFSET :offset`
	rows, err := cm.db.NamedQuery(q, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var c pb.Customer
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Phone,
			&c.IsBlocked,
			&createdAt,
			&fcmToken,
			&tgChatID,
			&c.ShipperId,
			&c.BotLanguage,
			&dateOfBirth,
			&c.IsAggregate,
			&customer_type_id,
		)
		if err != nil {
			return nil, 0, err
		}

		c.CreatedAt = createdAt.Format(layoutDateTime)
		c.FcmToken = etc.StringValue(fcmToken)
		c.TgChatId = etc.StringValue(tgChatID)
		c.DateOfBirth = etc.StringValue(dateOfBirth)
		c.CustomerTypeId = etc.StringValue(customer_type_id)

		customers = append(customers, &c)
	}

	return customers, count, nil
}

func (cm *customerRepo) Update(customer *pb.Customer) error {
	var (
		dateOfBirth      sql.NullString = etc.NullString(customer.DateOfBirth)
		customer_type_id sql.NullString = etc.NullString(customer.CustomerTypeId)
	)
	updateQuery :=
		`UPDATE customers
		 SET
		  	name=$1,
			phone=$2,
			date_of_birth=$3,
			updated_at=CURRENT_TIMESTAMP,
			is_blocked=$4,
			is_aggregate=$5,
            customer_type_id=$6
		WHERE shipper_id = $7 AND id=$8 AND deleted_at=0`
	result, err := cm.db.Exec(
		updateQuery,
		customer.GetName(),
		customer.GetPhone(),
		dateOfBirth,
		customer.GetIsBlocked(),
		customer.GetIsAggregate(),
		customer_type_id,
		customer.ShipperId,
		customer.GetId(),
	)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (cm *customerRepo) Delete(shipperId, id string) error {
	result, err := cm.db.Exec(`UPDATE customers SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where shipper_id=$1 AND id=$2 and deleted_at=0`, shipperId, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (cm *customerRepo) SearchByPhone(shipperId, phone string, limit uint64, customerTypeId string) ([]*pb.Customer, error) {
	var (
		customers        []*pb.Customer
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
		column           string
	)

	if customerTypeId != "" {
		column = fmt.Sprintf(` AND customer_type_id = %s `, customerTypeId)
	}

	column += " AND phone LIKE '%" + phone + "%' "

	query := `
			SELECT  
			id, 
			name, 
			phone, 
			is_blocked, 
			created_at,
			fcm_token,
			tg_chat_id,
			date_of_birth,
			is_aggregate,
            customer_type_id 
			FROM customers
			WHERE shipper_id=$1 %s AND 
			deleted_at=0
			ORDER BY created_at DESC
			LIMIT $2 `
	rows, err := cm.db.Queryx(fmt.Sprintf(query, column), shipperId, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c pb.Customer
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Phone,
			&c.IsBlocked,
			&c.CreatedAt,
			&fcmToken,
			&tgChatID,
			&dateOfBirth,
			&c.IsAggregate,
			&customer_type_id,
		)

		if err != nil {
			return nil, err
		}

		c.FcmToken = etc.StringValue(fcmToken)
		c.TgChatId = etc.StringValue(tgChatID)
		c.DateOfBirth = etc.StringValue(dateOfBirth)

		customers = append(customers, &c)
	}

	return customers, nil
}

func (cm *customerRepo) GetOrInsert(shipperId, phone, name string) (string, error) {
	var (
		customerID   string
		customerName string
	)
	row := cm.db.QueryRow(`SELECT id, name FROM customers WHERE shipper_id = $1 AND phone=$2 AND deleted_at=0`, shipperId, phone)

	err := row.Scan(&customerID, &customerName)

	if err == sql.ErrNoRows {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}

		_, err = cm.db.Exec(`insert into customers (
			id,
			name,
			phone,
			shipper_id)
			values ($1, $2, $3, $4)`,
			id, name, phone, shipperId)

		if err != nil {
			return "", err
		}

		return id.String(), nil
	} else if err != nil {
		return "", err
	}

	if customerName == name {
		return customerID, nil
	} else {
		_, err = cm.db.Exec(`UPDATE customers SET name=$1, updated_at=current_timestamp WHERE shipper_id=$2 AND id=$3`, name, shipperId, customerID)

		if err != nil {
			return "", err
		}

		return customerID, nil
	}
}

func (cm *customerRepo) GetByPhone(shipperId, phone string) (*pb.Customer, error) {
	var (
		createdAt        time.Time
		layoutDateTime   string = "2006-01-02 15:04:05"
		customer         pb.Customer
		fcmToken         sql.NullString
		tgChatID         sql.NullString
		dateOfBirth      sql.NullString
		customer_type_id sql.NullString
	)

	row := cm.db.QueryRow(`
		SELECT id,
				name,
				phone,
				is_blocked,
				created_at,
				fcm_token,
				tg_chat_id,
				shipper_id,
				date_of_birth,
				is_aggregate,
                customer_type_id
		FROM customers
		WHERE shipper_id = $1 AND phone=$2 and
		deleted_at=0`, shipperId, phone,
	)

	err := row.Scan(
		&customer.Id,
		&customer.Name,
		&customer.Phone,
		&customer.IsBlocked,
		&createdAt,
		&fcmToken,
		&tgChatID,
		&customer.ShipperId,
		&dateOfBirth,
		&customer.IsAggregate,
		&customer_type_id,
	)
	if err != nil {
		return nil, err
	}

	customer.CreatedAt = createdAt.Format(layoutDateTime)
	customer.FcmToken = etc.StringValue(fcmToken)
	customer.TgChatId = etc.StringValue(tgChatID)
	customer.DateOfBirth = etc.StringValue(dateOfBirth)
	customer.CustomerTypeId = etc.StringValue(customer_type_id)

	return &customer, nil
}

func (cm *customerRepo) UpdateFcmToken(id, shipperId, fcmToken, platformID string) error {
	result, err := cm.db.Exec(`
		UPDATE customers
		SET fcm_token = $1,
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

func (cm *customerRepo) UpdateTgChatID(id, shipperId, tgChatID string) error {
	result, err := cm.db.Exec(`
		UPDATE customers
		SET tg_chat_id = $1
		WHERE shipper_id=$2
		AND id = $3`,
		tgChatID,
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

func (cm *customerRepo) GetRegisteredCustomersReport(shipperId string, year, month uint64) ([]*pb.MontlyRegisteredCustomersReport, error) {
	var monthlyReports []*pb.MontlyRegisteredCustomersReport

	currentLocation := time.Now().Location()
	firstOfMonth := time.Date(int(year), time.Month(month), 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).Day()

	query := `
		SELECT 
			count(1)
		FROM customers
		WHERE shipper_id = $1
		AND deleted_at=0 
		AND created_at between $2 AND $3
	`
	for d := 1; d <= lastOfMonth; d++ {
		var monthlyReport pb.MontlyRegisteredCustomersReport
		start := time.Date(int(year), time.Month(int(month)), d, 5, 0, 0, 0, currentLocation)
		end := start.AddDate(0, 0, 1)
		dayStart := start.Format("2006-01-02 15:04:05")
		dayEnd := end.Format("2006-01-02 15:04:05")

		row := cm.db.QueryRow(query, shipperId, dayStart, dayEnd)

		err := row.Scan(&monthlyReport.Count)
		if err != nil {
			return nil, err
		}

		monthlyReport.Day = uint64(d)
		monthlyReports = append(monthlyReports, &monthlyReport)
	}

	return monthlyReports, nil
}

func (cm *customerRepo) AttachBotLanguage(id, lang string) error {
	if lang == "" {
		return nil
	}

	updateQuery :=
		`UPDATE customers
		 SET
		    bot_language=$1,
			updated_at=CURRENT_TIMESTAMP
		WHERE id=$2 AND deleted_at=0`

	result, err := cm.db.Exec(
		updateQuery,
		lang,
		id,
	)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
