package model

// 数据库实体类
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//后面的json为格式和对应的字段，前端传递过来需要携带的字段
	Status bool `json:"status"`
}
