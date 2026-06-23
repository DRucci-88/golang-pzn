package main

import (
	"context"
	"fmt"
	"go-mysql/entity"
	"go-mysql/repository"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	result, err := commentRepository.Insert(ctx, entity.Comment{
		Email:   "repo@gmail.com",
		Comment: "Test Repo",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range result {
		fmt.Println(comment, " ")
	}
}
