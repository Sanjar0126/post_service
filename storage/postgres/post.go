package postgres

import (
	"database/sql"
	pb "genproto/post_service"
	"github.com/Sanjar0126/post_service/storage/repo"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (pr *postRepo) Create(post *pb.Post) (string, error) {
	var (
		title  string
		body   string
		author string
	)

	id, err := uuid.NewRandom()
	if err != nil {
		return "", nil
	}

	insertNew :=
		`INSERT INTO
		posts (id, title, body, author, created_at) 
		values ($1, $2, $3, $4, $5)`

	_, err = pr.db.Exec(
		insertNew,
		id.String(),
		title,
		body,
		author,
		post.GetCreatedAt(),
	)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (pr *postRepo) Get(id string) (*pb.Post, error) {
	var (
		post pb.Post
	)

	row := pr.db.QueryRow(
			`select * from posts where id=$1`, id,
		)

	err := row.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.Author,
			&post.CreatedAt,
		)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (pr *postRepo) GetAll(page, limit uint64) ([]*pb.Post, uint64, error) {
	var (
		count uint64
		posts []*pb.Post
	)

	offset := (page - 1) * limit
	params := map[string]interface{}{
		"limit": limit,
		"offset": offset,
	}

	query := `
			select * from posts
			order by created_at desc 
			limit :limit offset :offset
		`

	rows, err := pr.db.NamedQuery(query, params)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var p pb.Post
		err = rows.Scan(
				&p.Id,
				&p.Title,
				&p.Body,
				&p.Author,
				&p.CreatedAt,
			)

		if err != nil {
			return nil, 0, err
		}
	}

	rows, err = pr.db.NamedQuery(`
		SELECT count(1) 
		FROM posts`, params,
	)

	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, 0, err
		}

		err = rows.Close()
		if err != nil {
			return nil, 0, err
		}
	}

	return posts, count, nil
}

func (pr *postRepo) Update(post *pb.Post) error {

	updateQuery := `
			UPDATE posts set title=$1, body=$2, author=$3 where id=&4
		`

	result, err := pr.db.Exec(updateQuery, post.GetTitle(), post.GetBody(), post.GetAuthor(), post.GetId())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (pr *postRepo) Delete(id string) error {
	result, err := pr.db.Exec("delete from posts where id="+id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

