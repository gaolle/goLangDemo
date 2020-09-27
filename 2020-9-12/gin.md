```go
func main() {
   // 初始化引擎
    engine := gin.Default()
    // 注册一个路径，记载模板时从该目录找
    // 改目录下的资源可以按照改路径访问
    engine.LoadHTMLGlob("views/*/*")
    //记载主页
    engine.Get("html-test", func(c *gin.Context) {
        c.HTML(200, )
    })
    engine.Run(":8080")
}
```

