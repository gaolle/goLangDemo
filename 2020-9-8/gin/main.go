package main

import (
	"github.com/gin-gonic/gin"
)

// get 获取资源
// post 新建资源

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 查询数据
	// // 客户端通过GET请求获取/hello时，执行匿名函数请求数据
	// r.GET(`/hello`, func(c *gin.Context) {
	// 	// 200 statusOk
	// // 	c.JSON(200, gin.H{
	// // 		// 返回网页的数据格式
	// // 		"message": "hello world",
	// // 	})
	// 	c.JSON(500, gin.H{
	// 		"name": "gaolle",
	// 	})
	// })

	// 启动服务
	r.Run()
}

