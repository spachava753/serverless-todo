package db

import (
	"context"
	"serverless-todo/model"
)

const keyRepository = "Repository"

type Repository interface {
	Close()
	Insert(todo *model.Item) (int, error)
	Delete(id int) error
	GetAll() ([]model.Item, error)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Close(ctx context.Context) {
	getRepository(ctx).Close()
}

func Insert(ctx context.Context, todo *model.Item) (int, error) {
	return getRepository(ctx).Insert(todo)
}

func Delete(ctx context.Context, id int) error {
	return getRepository(ctx).Delete(id)
}

func GetAll(ctx context.Context) ([]model.Item, error) {
	return getRepository(ctx).GetAll()
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
