package comment

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	belajar_golang "go-database"
	"go-database/entity"
	"testing"
)

func TestInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Komen Perdanaku",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Comment Terbaruku",
	}

	result, err := commentRepository.Update(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "repository@test.com",
	}

	result, err := commentRepository.Delete(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
