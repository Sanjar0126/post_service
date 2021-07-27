package postgres

import (
	"database/sql"
	"fmt"
	pb "genproto/user_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"gitlab.udevs.io/delever/delever_user_service/config"
	"gitlab.udevs.io/delever/delever_user_service/pkg/etc"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type branchRepo struct {
	db *sqlx.DB
}

// NewBranchRepo ...
func NewBranchRepo(db *sqlx.DB) repo.BranchStorageI {
	return &branchRepo{
		db: db,
	}
}

func (br *branchRepo) Create(branch *pb.Branch) (string, error) {
	var (
		destination sql.NullString = etc.NullString(branch.Destination)
		image       sql.NullString = etc.NullString(branch.Image)
		fareID      sql.NullString = etc.NullString(branch.FareId)
		tgChatID    sql.NullString = etc.NullString(branch.TgChatId)
	)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		branches
		(
			id,
			shipper_id,
			name,
			phone,
			address,
			location,
			destination,
			image,
			work_hour_start,
			work_hour_end,
			fare_id,
			tg_chat_id
		)
		VALUES
		($1, $2, $3, $4, $5, st_makepoint($6, $7), $8, $9, $10, $11, $12, $13)`

	_, err = br.db.Exec(
		insertNew,
		id.String(),
		branch.GetShipperId(),
		branch.GetName(),
		branch.GetPhone(),
		branch.GetAddress(),
		branch.GetLocation().GetLong(),
		branch.GetLocation().GetLat(),
		destination,
		image,
		branch.GetWorkHourStart(),
		branch.GetWorkHourEnd(),
		fareID,
		tgChatID,
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (br *branchRepo) Get(id string) (*pb.Branch, error) {
	var (
		createdAt, workHourStart, workHourEnd                                time.Time
		layoutDate                                                           string = "2006-01-02 15:04:05"
		branch                                                               pb.Branch
		location                                                             pb.Location
		column                                                               string
		destination, image, jowiID, iikoID, fareID, iikoTerminalID, tgChatID sql.NullString
		cfg                                                                         = config.Load()
		layoutTime                                                           string = "15:04"
	)

	_, err := uuid.Parse(id)
	if err != nil {
		column = " phone "
	} else {
		column = " id "
	}

	row := br.db.QueryRow(`
		SELECT  id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end,
				jowi_id,
				iiko_id,
				fare_id,
				iiko_terminal_id,
				tg_chat_id 
		FROM branches
		WHERE `+column+`=$1 and
		deleted_at = 0 `, id,
	)

	err = row.Scan(
		&branch.Id,
		&branch.ShipperId,
		&branch.Name,
		&branch.Phone,
		&branch.IsActive,
		&branch.Address,
		&location.Long,
		&location.Lat,
		&createdAt,
		&destination,
		&image,
		&workHourStart,
		&workHourEnd,
		&jowiID,
		&iikoID,
		&fareID,
		&iikoTerminalID,
		&tgChatID,
	)
	if err != nil {
		return nil, err
	}

	if image.String != "" {
		image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
	}

	branch.CreatedAt = createdAt.Format(layoutDate)
	branch.Location = &location
	branch.Destination = etc.StringValue(destination)
	branch.Image = etc.StringValue(image)
	branch.JowiId = etc.StringValue(jowiID)
	branch.IikoId = etc.StringValue(iikoID)
	branch.FareId = etc.StringValue(fareID)
	branch.IikoTerminalId = etc.StringValue(iikoTerminalID)
	branch.WorkHourStart = workHourStart.Format(layoutTime)
	branch.WorkHourEnd = workHourEnd.Format(layoutTime)
	branch.TgChatId = etc.StringValue(tgChatID)

	return &branch, nil
}

func (br *branchRepo) GetAll(shipperId, search, fareId string, page, limit uint64, jowi, iiko bool) ([]*pb.Branch, uint64, error) {
	var (
		filter                                                               string
		count                                                                uint64
		createdAt, workHourStart, workHourEnd                                time.Time
		layoutDate                                                           string = "2006-01-02 15:04:05"
		branches                                                             []*pb.Branch
		destination, image, jowiID, iikoID, fareID, iikoTerminalID, tgChatID sql.NullString
		cfg                                                                         = config.Load()
		layoutTime                                                           string = "15:04"
	)

	offset := (page - 1) * limit
	params := map[string]interface{}{
		"shipper_id": shipperId,
		"limit":      limit,
		"offset":     offset,
		"search":     search,
		"fare_id":    fareId,
	}

	if search != "" {
		filter += ` AND (name ILIKE '%' || :search || '%') `
	}

	if fareId != "" {
		filter += ` AND fare_id = :fare_id `
	}

	if iiko {
		filter += ` AND iiko_id IS NULL `
	}

	if jowi {
		filter += ` AND jowi_id IS NULL `
	}

	query := `
		SELECT  id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end,
				jowi_id,
				iiko_id,
				fare_id,
				iiko_terminal_id,
				tg_chat_id
		FROM branches
		WHERE shipper_id=:shipper_id
		  AND deleted_at = 0` + filter + ` 
		ORDER BY created_at DESC 
		LIMIT :limit OFFSET :offset`
	rows, err := br.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var (
			b        pb.Branch
			location pb.Location
		)
		err = rows.Scan(
			&b.Id,
			&b.ShipperId,
			&b.Name,
			&b.Phone,
			&b.IsActive,
			&b.Address,
			&location.Long,
			&location.Lat,
			&createdAt,
			&destination,
			&image,
			&workHourStart,
			&workHourEnd,
			&jowiID,
			&iikoID,
			&fareID,
			&iikoTerminalID,
			&tgChatID,
		)

		if err != nil {
			return nil, 0, err
		}

		if image.String != "" {
			image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
		}

		b.CreatedAt = createdAt.Format(layoutDate)
		b.Location = &location
		b.Destination = etc.StringValue(destination)
		b.Image = etc.StringValue(image)
		b.JowiId = etc.StringValue(jowiID)
		b.IikoId = etc.StringValue(iikoID)
		b.FareId = etc.StringValue(fareID)
		b.IikoTerminalId = etc.StringValue(iikoTerminalID)
		b.WorkHourStart = workHourStart.Format(layoutTime)
		b.WorkHourEnd = workHourEnd.Format(layoutTime)
		b.TgChatId = etc.StringValue(tgChatID)
		branches = append(branches, &b)
	}

	rows, err = br.db.NamedQuery(`
		SELECT count(1) 
		FROM branches
		WHERE shipper_id=:shipper_id AND deleted_at = 0`+filter, params,
	)

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, 0, err
		}
		rows.Close()
	}

	return branches, count, nil
}

func (br *branchRepo) Update(branch *pb.Branch) error {
	var (
		destination    sql.NullString = etc.NullString(branch.Destination)
		iikoID         sql.NullString = etc.NullString(branch.IikoId)
		jowiID         sql.NullString = etc.NullString(branch.JowiId)
		fareID         sql.NullString = etc.NullString(branch.FareId)
		iikoTerminalID sql.NullString = etc.NullString(branch.IikoTerminalId)
		tgChatID       sql.NullString = etc.NullString(branch.TgChatId)
		image          sql.NullString = etc.NullString(branch.Image)
	)

	updateQuery :=
		`UPDATE branches
		 SET
			name=$1,
			phone=$2,
			address=$3,
			location=st_makepoint($4,$5),
			updated_at=CURRENT_TIMESTAMP,
			destination=$6,
			image=$7,
			is_active=$8,
			work_hour_start=$9,
			work_hour_end=$10,
			jowi_id=$11,
			iiko_id=$12,
			fare_id=$13,
			iiko_terminal_id=$14,
			tg_chat_id=$15
		WHERE id=$16`

	result, err := br.db.Exec(
		updateQuery,
		branch.GetName(),
		branch.GetPhone(),
		branch.GetAddress(),
		branch.GetLocation().GetLong(),
		branch.GetLocation().GetLat(),
		destination,
		image,
		branch.GetIsActive(),
		branch.GetWorkHourStart(),
		branch.GetWorkHourEnd(),
		jowiID,
		iikoID,
		fareID,
		iikoTerminalID,
		tgChatID,
		branch.GetId(),
	)

	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (br *branchRepo) Delete(id, shipperId string) error {
	result, err := br.db.Exec(`UPDATE branches SET deleted_at=date_part('epoch', CURRENT_TIMESTAMP)::int where shipper_id=$1 AND id=$2 and deleted_at = 0`, shipperId, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (br *branchRepo) GetNearestBranch(shipperId string, location *pb.Location) ([]*pb.Branch, error) {
	var (
		stDistance                                                           string
		layoutDate                                                           string = "2006-01-02 15:04:05"
		branches                                                             []*pb.Branch
		destination, image, jowiID, iikoID, fareID, iikoTerminalID, tgChatID sql.NullString
		createdAt, workHourStart, workHourEnd                                time.Time
		cfg                                                                         = config.Load()
		layoutTime                                                           string = "15:04"
	)
	rows, err := br.db.Queryx(
		`SELECT id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end,
				jowi_id,
				iiko_id,
				fare_id,
				iiko_terminal_id,
				tg_chat_id,
				ST_Distance (geometry(location),geometry(st_makepoint($1, $2))) as distance 
				FROM branches 
			where shipper_id=$3 and deleted_at = 0
			ORDER BY distance asc`, location.GetLong(), location.GetLat(), shipperId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var b pb.Branch
		var loc pb.Location
		err := rows.Scan(
			&b.Id,
			&b.ShipperId,
			&b.Name,
			&b.Phone,
			&b.IsActive,
			&b.Address,
			&loc.Long,
			&loc.Lat,
			&createdAt,
			&destination,
			&image,
			&workHourStart,
			&workHourEnd,
			&jowiID,
			&iikoID,
			&fareID,
			&iikoTerminalID,
			&tgChatID,
			&stDistance,
		)

		if err != nil {
			return nil, err
		}

		if image.String != "" {
			image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
		}

		b.CreatedAt = createdAt.Format(layoutDate)
		b.Location = &loc
		b.WorkHourStart = workHourStart.Format(layoutTime)
		b.WorkHourEnd = workHourEnd.Format(layoutTime)
		b.Destination = etc.StringValue(destination)
		b.Image = etc.StringValue(image)
		b.JowiId = etc.StringValue(jowiID)
		b.IikoId = etc.StringValue(iikoID)
		b.FareId = etc.StringValue(fareID)
		b.IikoTerminalId = etc.StringValue(iikoTerminalID)
		b.TgChatId = etc.StringValue(tgChatID)

		branches = append(branches, &b)
	}

	return branches, nil
}

func (br *branchRepo) GetByName(shipperId, name string) (*pb.Branch, error) {
	var (
		createdAt, workHourStart, workHourEnd                                time.Time
		layoutDate                                                           string = "2006-01-02 15:04:05"
		branch                                                               pb.Branch
		location                                                             pb.Location
		destination, image, jowiID, iikoID, fareID, iikoTerminalID, tgChatID sql.NullString
		cfg                                                                         = config.Load()
		layoutTime                                                           string = "15:04"
	)

	row := br.db.QueryRow(`
		SELECT  id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end,
				jowi_id,
				iiko_id,
				fare_id,
				iiko_terminal_id,
				tg_chat_id
		FROM branches
		WHERE shipper_id=$1
		AND name=$2
		AND deleted_at = 0 `,
		shipperId,
		name,
	)

	err := row.Scan(
		&branch.Id,
		&branch.ShipperId,
		&branch.Name,
		&branch.Phone,
		&branch.IsActive,
		&branch.Address,
		&location.Long,
		&location.Lat,
		&createdAt,
		&destination,
		&image,
		&workHourStart,
		&workHourEnd,
		&jowiID,
		&iikoID,
		&fareID,
		&iikoTerminalID,
		&tgChatID,
	)
	if err != nil {
		return nil, err
	}

	if image.String != "" {
		image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
	}

	branch.CreatedAt = createdAt.Format(layoutDate)
	branch.Location = &location
	branch.Destination = etc.StringValue(destination)
	branch.Image = etc.StringValue(image)
	branch.JowiId = etc.StringValue(jowiID)
	branch.IikoId = etc.StringValue(iikoID)
	branch.FareId = etc.StringValue(fareID)
	branch.TgChatId = etc.StringValue(tgChatID)
	branch.IikoTerminalId = etc.StringValue(iikoTerminalID)
	branch.WorkHourStart = workHourStart.Format(layoutTime)
	branch.WorkHourEnd = workHourEnd.Format(layoutTime)

	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (br *branchRepo) GetByJowiID(jowiID string) (*pb.Branch, error) {
	var (
		createdAt, workHourStart, workHourEnd time.Time
		layoutDate                            string = "2006-01-02 15:04:05"
		branch                                pb.Branch
		location                              pb.Location
		destination, image                    sql.NullString
		cfg                                          = config.Load()
		layoutTime                            string = "15:04"
	)

	row := br.db.QueryRow(`
		SELECT  id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end
		FROM branches
		WHERE jowi_id=$1 and
		deleted_at = 0 `, jowiID,
	)

	err := row.Scan(
		&branch.Id,
		&branch.ShipperId,
		&branch.Name,
		&branch.Phone,
		&branch.IsActive,
		&branch.Address,
		&location.Long,
		&location.Lat,
		&createdAt,
		&destination,
		&image,
		&workHourStart,
		&workHourEnd,
	)
	if err != nil {
		return nil, err
	}

	if image.String != "" {
		image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
	}

	branch.CreatedAt = createdAt.Format(layoutDate)
	branch.Location = &location
	branch.Destination = etc.StringValue(destination)
	branch.Image = etc.StringValue(image)
	branch.WorkHourStart = workHourStart.Format(layoutTime)
	branch.WorkHourEnd = workHourEnd.Format(layoutTime)

	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (br *branchRepo) GetByIikoID(iikoID string) (*pb.Branch, error) {
	var (
		createdAt, workHourStart, workHourEnd time.Time
		layoutDate                            string = "2006-01-02 15:04:05"
		branch                                pb.Branch
		location                              pb.Location
		destination, image                    sql.NullString
		cfg                                          = config.Load()
		layoutTime                            string = "15:04"
	)

	row := br.db.QueryRow(`
		SELECT  id,
				shipper_id,
				name,
				phone,
				is_active,
				address,
				st_x(location),
				st_y(location),
				created_at,
				destination,
				image,
				work_hour_start,
				work_hour_end
		FROM branches
		WHERE iiko_id=$1 and
		deleted_at = 0 `, iikoID,
	)

	err := row.Scan(
		&branch.Id,
		&branch.ShipperId,
		&branch.Name,
		&branch.Phone,
		&branch.IsActive,
		&branch.Address,
		&location.Long,
		&location.Lat,
		&createdAt,
		&destination,
		&image,
		&workHourStart,
		&workHourEnd,
	)
	if err != nil {
		return nil, err
	}

	if image.String != "" {
		image.String = fmt.Sprintf("%s%s", cfg.CDN, image.String)
	}

	branch.CreatedAt = createdAt.Format(layoutDate)
	branch.Location = &location
	branch.Destination = etc.StringValue(destination)
	branch.Image = etc.StringValue(image)
	branch.WorkHourStart = workHourStart.Format(layoutTime)
	branch.WorkHourEnd = workHourEnd.Format(layoutTime)

	if err != nil {
		return nil, err
	}

	return &branch, nil
}
