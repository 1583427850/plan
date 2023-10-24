package todos

import (
	"errors"
	"gorm.io/gorm"
	"plan/config"
	"plan/dao"
	"plan/model"
)

func CreateTodo(todo *model.Todo) error {
	return dao.CreateTodo(todo)
	//todo 可以判断是否已经有对应记录了
}

func UpdateStatus(id string, todo *model.Todo) error {
	//判断对应数据是否存在
	oldTdo, err := dao.GetTodoById(id)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return err
	}
	//更新
	todo.Title = oldTdo.Title
	return config.DB.Model(&model.Todo{}).Where("id=?", oldTdo.ID).Update("status", todo.Status).Error
}

func ListTodos() (todos []model.Todo, err error) {
	return dao.ListTodo()
}

func DeleteTodo(id string) error {
	return dao.DeleteTodoById(id)
}
