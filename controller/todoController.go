package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plan/model"
	todos "plan/service"
)

func CreateTodo(c *gin.Context) {
	//查询数据
	todo := new(model.Todo)
	c.BindJSON(&todo)
	err := todos.CreateTodo(todo)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "新增成功")
}

func ListTodos(c *gin.Context) {
	todos, err := todos.ListTodos()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, todos)
}

func UpdateTodosStatus(c *gin.Context) {
	todo := new(model.Todo)
	id := c.Param("id")
	c.BindJSON(&todo)
	err := todos.UpdateStatus(id, todo)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "修改成功")

}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	//todo 后面可以添加一些参数校验
	err := todos.DeleteTodo(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "删除成功")
}
