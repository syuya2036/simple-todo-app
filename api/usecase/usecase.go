package usecase

import (
	"fmt"
	"todo/model"
	"todo/repository"
)

type TodoUsecase interface {
	GetAll() ([]model.Task, error)
	GetOne(id int) (model.Task, error)
	CreateOne(task model.Task) (model.Task, error)
	UpdateOne(task model.Task) (model.Task, error)
}

type todoUsecase struct {
	repo repository.Repository
}

func NewTodoUsecase(repo repository.Repository) TodoUsecase {
	return &todoUsecase{repo}
}

func (usecase *todoUsecase) GetAll() ([]model.Task, error) {
	tasks, err := usecase.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (usecase *todoUsecase) GetOne(id int) (model.Task, error) {
	task, err := usecase.repo.GetOne(id)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (usecase *todoUsecase) CreateOne(task model.Task) (model.Task, error) {
	// 未完了で作成する
	task.Completed = false

	// titleが空の場合はエラー
	if task.Title == "" {
		return model.Task{}, fmt.Errorf("title is empty")
	}

	task, err := usecase.repo.CreateOne(task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (usecase *todoUsecase) UpdateOne(task model.Task) (model.Task, error) {
	task, err := usecase.repo.UpdateOne(task)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
