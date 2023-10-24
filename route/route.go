package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"plan/controller"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("frontend/*")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		//新增清单
		v1Group.POST("/todo", todos.CreateTodo)

		//查询所有
		v1Group.GET("/todo", todos.ListTodos)

		//更新状态
		v1Group.PUT("/todo/:id", todos.UpdateTodosStatus)

		//删除
		v1Group.DELETE("/todo/:id", todos.DeleteTodoById)
	}

	return r

}
