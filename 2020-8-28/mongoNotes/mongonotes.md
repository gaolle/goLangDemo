### MongoDB

基于分布式文件存储的数据库

##### 文档

一组键值对（BSON），有序，不需要设置相同的字段，并且相同的字段不需要相同的数据类型

##### 集合

##### 数据库

默认数据库为“db",存储在data目录中

* 显示所有数据的列表

  ```shell
  show dbs
  ```

* 显示当前数据库对象或集合

  ```shell
  db
  ```

* 连接到一个指定的数据库

  ```shell
  use dataserver
  ```

##### 创建数据库

```shell
use dataserver
```

##### 查看所有数据库

```shell
show dbs
```

##### 删除当前数据库

```shell
db.dropDatabase()
```