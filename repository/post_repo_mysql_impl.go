package repository

import (
	"context"
	"database/sql"
	"web_blog/model"
)

type mysqlPostRepo struct {
	Db *sql.DB
}

func NewMysqlPostRepo(db *sql.DB) *mysqlPostRepo {
	return &mysqlPostRepo{Db: db}
}

func (repo *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Post, error) {
	//context 是什么
	rows, err := repo.Db.QueryContext(ctx, query, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	data := make([]*model.Post, 0)
	for rows.Next() {
		row := new(model.Post)
		err := rows.Scan(
			&row.ID,
			&row.Title,
			&row.Content,
		)

		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	return data, nil
}

func (repo *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*model.Post, error) {
	query := "Select id, title, content From posts limit ?"
	return repo.fetch(ctx, query, num)
}

func (repo *mysqlPostRepo) FindById(ctx context.Context, id int64) (*model.Post, error) {
	query := "Select id, title, content From posts where id=?"
	rows, e := repo.fetch(ctx, query)
	if e != nil {
		return nil, e
	}
	if len(rows) > 0 {
		return rows[0], nil
	} else {
		return nil, model.ErrNotFound
	}
}

func (repo *mysqlPostRepo) Create(ctx context.Context, post *model.Post) (int64, error) {
	query := "Insert posts Set title=?, content=?"
	stmt, err := repo.Db.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}
	result, err := stmt.ExecContext(ctx, post.Title, post.Content)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (repo *mysqlPostRepo) Update(ctx context.Context, post *model.Post) (*model.Post, error) {
	query := "Update posts Set title=?, content=? where id=?"
	stmt, err := repo.Db.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		post.Title,
		post.Content,
		post.ID,
	)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From posts Where id=?"
	stmt, err := repo.Db.PrepareContext(ctx, query)
	defer stmt.Close()
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
