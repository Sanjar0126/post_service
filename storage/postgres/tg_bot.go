package postgres

import (
	"database/sql"
	pb "genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/delever/delever_user_service/storage/repo"
)

type tgBotRepo struct {
	db *sqlx.DB
}

func NewTgBotRepo(db *sqlx.DB) repo.TgBotsStorageI {
	return &tgBotRepo{
		db: db,
	}
}

func (tg *tgBotRepo) Create(tgBot *pb.TgBot) (string, error) {
	query := `INSERT INTO tgbots (
		shipper_id, 
		bot_token,
		url,
		access_token
		) VALUES ($1, $2, $3, $4)`
	_, err := tg.db.Exec(query,
		tgBot.GetShipperId(),
		tgBot.GetBotToken(),
		tgBot.GetUrl(),
		tgBot.GetAccessToken(),
	)

	if err != nil {
		return "", err
	}

	return tgBot.GetShipperId(), nil
}

func (tg *tgBotRepo) Update(tgBot *pb.TgBot) error {
	query := `UPDATE tgbots SET 
		bot_token=$1, 
		url=$2,
		access_token=$3,
		updated_at=$4 
		WHERE shipper_id=$5 and deleted_at = 0`
	res, err := tg.db.Exec(
		query,
		tgBot.GetBotToken(),
		tgBot.GetUrl(),
		tgBot.GetAccessToken(),
		time.Now(),
		tgBot.GetShipperId(),
	)

	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (tg *tgBotRepo) Delete(shipperId *pb.ShipperId) error {
	query := `UPDATE tgbots 
			  SET deleted_at=$1 
			  WHERE shipper_id=$2 and deleted_at = 0`
	res, err := tg.db.Exec(query, time.Now().Unix(), shipperId.GetId())

	if err != nil {
		return err
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (tg *tgBotRepo) Get(shipperId string) (*pb.TgBot, error) {
	var (
		tgBot      pb.TgBot
		created_at time.Time
		updated_at time.Time
		layout     string = "2006-01-02 15:04:05"
	)
	query := `SELECT shipper_id, 
			bot_token, 
			url,
			access_token,
			created_at, 
			updated_at 
		FROM tgBots 
		WHERE shipper_id=$1 and deleted_at = 0`
	row := tg.db.QueryRow(query, shipperId)

	err := row.Scan(
		&tgBot.ShipperId,
		&tgBot.BotToken,
		&tgBot.Url,
		&tgBot.AccessToken,
		&created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}
	tgBot.CreatedAt = created_at.Format(layout)
	tgBot.UpdatedAt = updated_at.Format(layout)

	return &tgBot, nil

}

func (tg *tgBotRepo) GetAll(page, limit int64) ([]*pb.TgBot, int64, error) {
	var (
		count                  int64
		created_at, updated_at time.Time
		layout                 string = "2006-01-02 15:04:05"
		tgBots                 []*pb.TgBot
	)
	offset := (page - 1) * limit

	query := `SELECT 
			shipper_id,
			bot_token,
			url,
			access_token,
			created_at, 
			updated_at
		FROM tgBots
		WHERE deleted_at = 0
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	rows, err := tg.db.Queryx(query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var tgBot pb.TgBot

		err = rows.Scan(
			&tgBot.ShipperId,
			&tgBot.BotToken,
			&tgBot.Url,
			&tgBot.AccessToken,
			&created_at,
			&updated_at,
		)

		if err != nil {
			return nil, 0, err
		}

		tgBot.CreatedAt = created_at.Format(layout)
		tgBot.UpdatedAt = updated_at.Format(layout)

		tgBots = append(tgBots, &tgBot)
	}

	row := tg.db.QueryRow(`SELECT count(1) FROM tgBots WHERE deleted_at = 0`)

	err = row.Scan(&count)

	return tgBots, count, nil
}
