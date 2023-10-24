package dao

import (
	"plan/config"
	"plan/model"
)

// CreateTodo 创建一条
func CreateTodo(todo *model.Todo) error {
	return config.DB.Create(&todo).Error
}

// ListTodo 遍历所有
func ListTodo() (todos []model.Todo, err error) {
	err = config.DB.Find(&todos).Error
	return

}

// GetTodoById 根据id获取一条信息
func GetTodoById(id string) (*model.Todo, error) {
	todo := new(model.Todo)
	err := config.DB.Where("id=?", id).First(&todo).Error
	return todo, err
}

// UpdateTodo 更新
func UpdateTodo(todo *model.Todo) error {
	return config.DB.Save(todo).Error
}

// DeleteTodoById 删除
func DeleteTodoById(id string) error {
	return config.DB.Where("id=?", id).Delete(model.Todo{}).Error
}
