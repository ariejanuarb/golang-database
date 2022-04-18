package comment

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email,comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

// service = test, repo = impl
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "UPDATE comments SET comment = ? WHERE email = ?"
	result, err := repository.DB.ExecContext(ctx, script, comment.Comment, comment.Email)
	if err != nil {
		return comment, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rowCnt == 0 {
		return comment, err
	}
	return comment, err
}

func (repository *commentRepositoryImpl) Delete(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "DELETE FROM comments WHERE email = ?"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email)
	if err != nil {
		return comment, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return comment, err
	}
	if rowCnt == 0 {
		return comment, err
	}
	return comment, nil
}
