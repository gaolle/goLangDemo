### MySQL：关系数据库管理系统（RDBMS）(将数据保存在不同的表中)

#### 特点：

* 数据以表格的形式出现
* 每行为各种记录名称
* 每列为记录名称所对应的数据域
* 许多的行和列组成一张表单
* 若干的表单组成database

#### RDBMS：

* 数据库：关联表的集合
* 数据表：数据的矩阵
* 冗余：存储两倍数据，让系统速度更快
* 主键：唯一，一个数据表只能包含一个主键
* 外键：关联两个表
* 复合键：将多个列作为一个索引键
* 索引：快速索引特定信息。对数据库表中一列或多列的值进行排序的一种结构
* 参照完整性：保证数据的一致性

* 创建数据库

  ```mysql
  mysqladmin -u root -p create users
  ```

* 删除数据库

  ```mysql
  mysqladmin -u root -p drop users
  ```

* 选择数据库

  ```mysql
  use users
  ```

* 创建数据表

  ```mysql
  CREATE TABLE users_tbl
  ```

* 删除数据表

  ```mysql
  DROP TABLE users_tbl
  ```

* 插入数据

  ```mysql
  INSERT INTO users_tbl
  (id, name)
  VALUES
  (1, "A")
  ```

* 查询数据库

  ```mysql
  SELECT * from users_tbl
  ```

### docker安装MySql

* 拉取镜像

  ```shell
  docker pull mysql:latest
  ```

* 查看本地镜像

  ```shell
  docker images
  ```

* 运行一个容器

  ```shell
  docker run  --name Master -p 3306:3306 --restart=always  -e MYSQL_ROOT_PASSWORD=123456 -d mariadb:10.1
  ```

* 进入容器

  ```shell
  docker exec -it Master bash
  ```

* 登录mysql

  ```shell
  mysql -u root -p
  ```

  



