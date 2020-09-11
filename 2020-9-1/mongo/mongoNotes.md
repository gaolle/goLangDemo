### docker安装mongo  

拉取最新版本镜像

```shell
docker pull mongo:latest
```

运行容器

```shell
docker run -itd --name mongo -p 27017:27017 mongo
```

进入mongo

```shell
docker exec -it mongo mongo admin
```

创建/切换数据库

```shell
use student
```

删除数据库（切换至数据库中）

```shell
db.dropDatabase()
```

插入集合中文档

```shell
db.student.inset({name:'xiaoming', age:18})
```

查看插入文档

```shell
db.student.find()
```

定义常量插入

```shell
stu=({name:'xiaoming', age:18})
db.student.insert(stu)
```

更新数据

```shell
db.student.update({'name':'g'}, {$set:{'age':18}})
```

删除数据

```shell
db.student.remove({name:'g'})
```

查询数据

```shell
db.student.find()
```

格式化查询数据

```shell
db.student.find().pretty()
```

