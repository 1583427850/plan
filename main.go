package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//后面的json为格式和对应的字段，前端传递过来需要携带的字段
	Status bool `json:"status"`
}

var (
	DB *gorm.DB
)

// 连接数据库
func initMysql() {
	var err error
	dsn := "root:abc8909389@tcp(127.0.0.1:3306)/golearn?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		panic(err)
	}
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("frontend/*")
	r.Static("/static", "static")

	initMysql()
	//自动映射和创建对应结构体
	DB.AutoMigrate(&Todo{})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("/v1")
	{
		//新增清单
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			//插入对应数据，判断是否有错误
			if err := DB.Create(&todo).Error; err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, "创建成功")
		})

		v1Group.GET("/todo", func(c *gin.Context) {
			var todos []Todo
			if err := DB.Find(&todos).Error; err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, todos)
		})

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			var todo Todo
			//判断是否有对应记录
			if err := DB.Where("id=?", id).First(&todo).Error; err != nil {
				//如果没找到也会返回错误，并且从无为notfounderr
				if errors.Is(err, gorm.ErrRecordNotFound) {
					panic("找不到对应数据")
				}
				panic(err)
			}
			c.BindJSON(&todo)
			DB.Save(todo)
			c.JSON(http.StatusOK, "修改成功")

		})

		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			if len(id) == 0 {
				panic("id不能为空")
			}

			if err := DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, "删除成功")

		})
	}
	r.Run()
}
