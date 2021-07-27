package postgres

import (
	"database/sql"
	"fmt"
	pb "genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"gitlab.udevs.io/delever/delever_user_service/config"
	"gitlab.udevs.io/delever/delever_user_service/pkg/etc"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type shipperRepo struct {
	db *sqlx.DB
}

// NewShipperRepo ...
func NewShipperRepo(db *sqlx.DB) repo.ShipperStorageI {
	return &shipperRepo{
		db: db,
	}
}

func (shp *shipperRepo) Create(shipper *pb.Shipper) (string, error) {
	var (
		callCenterTg sql.NullString = etc.NullString(shipper.CallCenterTg)
		menuImage    sql.NullString = etc.NullString(shipper.MenuImage)
	)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	if shipper.GetCrm() == "" {
		shipper.Crm = "none"
	}

	insertNew :=
		`INSERT INTO
		shippers
		(
			id,
			name,
			logo,
			description,
			phone,
			call_center_tg,
			work_hour_start,
			work_hour_end,
			menu_image,
			crm,
			courier_accepts_first,
			check_courier_action_radius,
			courier_action_radius,
			max_delivery_time,
			max_courier_orders,
			process_only_paid_orders,
			show_location_before_accepting,
			distance,
			enable_courier_working_hours
		)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`
	distance := shipper.GetDistance().Value
	_, err = shp.db.Exec(
		insertNew,
		id.String(),
		shipper.GetName(),
		shipper.GetLogo(),
		shipper.GetDescription(),
		pq.Array(shipper.GetPhone()),
		callCenterTg,
		shipper.WorkHourStart,
		shipper.WorkHourEnd,
		menuImage,
		shipper.Crm,
		shipper.CourierAcceptsFirst,
		shipper.CheckCourierActionRadius,
		shipper.CourierActionRadius,
		shipper.MaxDeliveryTime,
		shipper.MaxCourierOrders,
		shipper.ProcessOnlyPaidOrders,
		shipper.ShowLocationBeforeAccepting,
		distance,
		shipper.EnableCourierWorkingHours,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (shp *shipperRepo) Get(id string) (*pb.Shipper, error) {
	var (
		createdAt, workHourStart, workHourEnd time.Time
		layoutDate                            string = "2006-01-02 15:04:05"
		shipper                               pb.Shipper
		column                                string
		phones                                []string
		callCenterTg                          sql.NullString
		menuImage                             sql.NullString
		layoutTime                            string = "15:04"
		cfg                                          = config.Load()
		distance                              sql.NullFloat64
	)

	_, err := uuid.Parse(id)

	if err != nil {
		column = " phone "
	} else {
		column = " id "
	}

	row := shp.db.QueryRow(`
		SELECT  id,
				name,
				logo,
				description,
				phone,
				is_active,
				created_at,
				call_center_tg,
				work_hour_start,
				work_hour_end,
				menu_image,
				crm,
				courier_accepts_first,
				check_courier_action_radius,
				courier_action_radius,
				max_delivery_time,
				max_courier_orders,
				process_only_paid_orders,
				show_location_before_accepting,
				distance,
				enable_courier_working_hours
		FROM shippers
		WHERE `+column+`=$1 and
		deleted_at = 0`, id,
	)

	err = row.Scan(
		&shipper.Id,
		&shipper.Name,
		&shipper.Logo,
		&shipper.Description,
		pq.Array(&phones),
		&shipper.IsActive,
		&createdAt,
		&callCenterTg,
		&workHourStart,
		&workHourEnd,
		&menuImage,
		&shipper.Crm,
		&shipper.CourierAcceptsFirst,
		&shipper.CheckCourierActionRadius,
		&shipper.CourierActionRadius,
		&shipper.MaxDeliveryTime,
		&shipper.MaxCourierOrders,
		&shipper.ProcessOnlyPaidOrders,
		&shipper.ShowLocationBeforeAccepting,
		&distance,
		&shipper.EnableCourierWorkingHours,
	)

	if err != nil {
		return nil, err
	}

	if shipper.Logo != "" {
		shipper.Logo = fmt.Sprintf("%s%s", cfg.CDN, shipper.Logo)
	}

	shipper.CreatedAt = createdAt.Format(layoutDate)
	shipper.Phone = phones
	shipper.CallCenterTg = etc.StringValue(callCenterTg)
	shipper.WorkHourStart = workHourStart.Format(layoutTime)
	shipper.WorkHourEnd = workHourEnd.Format(layoutTime)
	shipper.MenuImage = etc.StringValue(menuImage)
	shipper.Distance = etc.FloatValue(distance)

	if shipper.MenuImage.GetValue() != "" {
		shipper.MenuImage.Value = fmt.Sprintf("%s%s", cfg.CDN, shipper.MenuImage.GetValue())
	}

	return &shipper, nil
}

func (shp *shipperRepo) GetAll(page, limit uint64, hasIiko bool) ([]*pb.Shipper, uint64, error) {
	var (
		filter                                string
		count                                 uint64
		createdAt, workHourStart, workHourEnd time.Time
		layoutDate                            string = "2006-01-02 15:04:05"
		layoutTime                            string = "15:04"
		shippers                              []*pb.Shipper
		phones                                []string
		callCenterTg                          sql.NullString
		menuImage                             sql.NullString
		cfg                                   = config.Load()
		distance                              sql.NullFloat64
	)

	offset := (page - 1) * limit

	if hasIiko {
		filter += ` AND crm='iiko' `
	}

	query := `
		SELECT  id,
				name,
				logo,
				description,
				phone,
				is_active,
				created_at,
				call_center_tg,
				work_hour_start,
				work_hour_end,
				menu_image,
				crm,
				courier_accepts_first,
				check_courier_action_radius,
				courier_action_radius,
				max_delivery_time,
				max_courier_orders,
				process_only_paid_orders,
				show_location_before_accepting,
				distance,
				enable_courier_working_hours
		FROM shippers
		WHERE deleted_at = 0 ` + filter + ` 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2`
	rows, err := shp.db.Queryx(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var shipper pb.Shipper
		err = rows.Scan(
			&shipper.Id,
			&shipper.Name,
			&shipper.Logo,
			&shipper.Description,
			pq.Array(&phones),
			&shipper.IsActive,
			&createdAt,
			&callCenterTg,
			&workHourStart,
			&workHourEnd,
			&menuImage,
			&shipper.Crm,
			&shipper.CourierAcceptsFirst,
			&shipper.CheckCourierActionRadius,
			&shipper.CourierActionRadius,
			&shipper.MaxDeliveryTime,
			&shipper.MaxCourierOrders,
			&shipper.ProcessOnlyPaidOrders,
			&shipper.ShowLocationBeforeAccepting,
			&distance,
			&shipper.EnableCourierWorkingHours,
		)

		if err != nil {
			return nil, 0, err
		}

		if shipper.Logo != "" {
			shipper.Logo = fmt.Sprintf("%s%s", cfg.CDN, shipper.Logo)
		}

		shipper.CreatedAt = createdAt.Format(layoutDate)
		shipper.Phone = phones
		shipper.CallCenterTg = etc.StringValue(callCenterTg)
		shipper.WorkHourStart = workHourStart.Format(layoutTime)
		shipper.WorkHourEnd = workHourEnd.Format(layoutTime)
		shipper.MenuImage = etc.StringValue(menuImage)
		shipper.Distance = etc.FloatValue(distance)

		if shipper.MenuImage.GetValue() != "" {
			shipper.MenuImage.Value = fmt.Sprintf("%s%s", cfg.CDN, shipper.MenuImage.GetValue())
		}

		shippers = append(shippers, &shipper)
	}

	row := shp.db.QueryRow(`
		SELECT count(1) 
		FROM shippers
		WHERE deleted_at = 0 ` + filter,
	)

	err = row.Scan(
		&count,
	)
	if err != nil {
		return nil, 0, nil
	}

	return shippers, count, nil
}

func (shp *shipperRepo) Update(shipper *pb.Shipper) error {
	var (
		callCenterTg sql.NullString  = etc.NullString(shipper.CallCenterTg)
		menuImage    sql.NullString  = etc.NullString(shipper.MenuImage)
		distance     sql.NullFloat64 = etc.ToNullFloat64(shipper.Distance)
	)

	if shipper.GetCrm() == "" {
		shipper.Crm = "none"
	}

	updateQuery :=
		`UPDATE shippers
		 SET
			name=$1,
			phone=$2,
			logo=$3,
			description=$4,
			call_center_tg=$5,
			work_hour_start=$6,
			work_hour_end=$7,
			updated_at=CURRENT_TIMESTAMP,
			is_active=$8,
			menu_image=$9,
			crm=$10,
			courier_accepts_first=$11,
			check_courier_action_radius=$12,
			courier_action_radius=$13,
			max_delivery_time=$14,
			max_courier_orders=$15,
			process_only_paid_orders=$16,
			show_location_before_accepting=$17,
			distance=$18,
			enable_courier_working_hours = $19
		WHERE id=$20`

	result, err := shp.db.Exec(
		updateQuery,
		shipper.GetName(),
		pq.Array(shipper.GetPhone()),
		shipper.GetLogo(),
		shipper.GetDescription(),
		callCenterTg,
		shipper.GetWorkHourStart(),
		shipper.GetWorkHourEnd(),
		shipper.GetIsActive(),
		menuImage,
		shipper.GetCrm(),
		shipper.GetCourierAcceptsFirst(),
		shipper.GetCheckCourierActionRadius(),
		shipper.GetCourierActionRadius(),
		shipper.GetMaxDeliveryTime(),
		shipper.GetMaxCourierOrders(),
		shipper.GetProcessOnlyPaidOrders(),
		shipper.GetShowLocationBeforeAccepting(),
		distance,
		shipper.GetEnableCourierWorkingHours(),
		shipper.GetId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (shp *shipperRepo) Delete(id string) error {
	result, err := shp.db.Exec(`UPDATE shippers SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int WHERE id=$1 AND deleted_at = 0`, id)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (shp *shipperRepo) GetByName(name string) (*pb.Shipper, error) {
	var (
		createdAt, workHourStart, workHourEnd time.Time
		layoutDate                            string = "2006-01-02 15:04:05"
		shipper                               pb.Shipper
		phones                                []string
		callCenterTg                          sql.NullString
		menuImage                             sql.NullString
		layoutTime                            string = "15:04"
		cfg                                          = config.Load()
		distance                              sql.NullFloat64
	)

	row := shp.db.QueryRow(`
		SELECT  id,
				name,
				logo,
				description,
				phone,
				is_active,
				created_at,
				call_center_tg,
				work_hour_start,
				work_hour_end,
				menu_image,
				crm,
				courier_accepts_first,
				check_courier_action_radius,
				courier_action_radius,
				max_delivery_time,
				max_courier_orders,
				process_only_paid_orders,
				show_location_before_accepting,
				distance,
				enable_courier_working_hours
		FROM shippers
		WHERE name = $1 
		AND deleted_at = 0`, name,
	)

	err := row.Scan(
		&shipper.Id,
		&shipper.Name,
		&shipper.Logo,
		&shipper.Description,
		pq.Array(&phones),
		&shipper.IsActive,
		&createdAt,
		&callCenterTg,
		&workHourStart,
		&workHourEnd,
		&menuImage,
		&shipper.Crm,
		&shipper.CourierAcceptsFirst,
		&shipper.CheckCourierActionRadius,
		&shipper.CourierActionRadius,
		&shipper.MaxDeliveryTime,
		&shipper.MaxCourierOrders,
		&shipper.ProcessOnlyPaidOrders,
		&shipper.ShowLocationBeforeAccepting,
		&distance,
		&shipper.EnableCourierWorkingHours,
	)

	if err != nil {
		return nil, err
	}

	if shipper.Logo != "" {
		shipper.Logo = fmt.Sprintf("%s%s", cfg.CDN, shipper.Logo)
	}

	shipper.CreatedAt = createdAt.Format(layoutDate)
	shipper.Phone = phones
	shipper.CallCenterTg = etc.StringValue(callCenterTg)
	shipper.WorkHourStart = workHourStart.Format(layoutTime)
	shipper.WorkHourEnd = workHourEnd.Format(layoutTime)
	shipper.MenuImage = etc.StringValue(menuImage)
	shipper.Distance = etc.FloatValue(distance)

	if err != nil {
		return nil, err
	}

	return &shipper, nil
}
