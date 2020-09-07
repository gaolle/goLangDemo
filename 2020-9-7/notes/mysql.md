#### 主键（primarily key）

能够唯一标志表中某一行的属性，一个表只能有一个主键

#### 外键（foreign key）

建立和加强两个表之间的链接，约束两个表之间数据的一致性。表的外键为另一个表的主键

创建父表A（只有一个）

```mysql
CREATE TABLE A  
(A_ID int NOT NULL, 
A_NAME VARCHAR(100) NOT NULL,
PRIMARY KEY (A_ID)；
```

创建子表B（一个或者多个）

```mysql
CREATE TABLE B 
( B_ID int NOT NULL, 
B_NAME VARCHAR(100) NOT NULL, 
A_ID int NOT NULL, 
PRIMARY KEY (B_ID)); 
```

A_ID为主键，B_ID为外键

```mysql
ALTER TABLE B ADD FOREIGN KEY (A_ID) REFERENCES A(A_ID);
```

A添加数据

| A_ID | A_NAME |
| ---- | ------ |
| 1    | A      |
| 2    | B      |
| 3    | C      |

B添加数据

| B_ID | B_NAME | A_ID |
| ---- | ------ | ---- |
| 1    | C      | 1    |
| 2    | C      | 2    |

向B中添加数据父表中A_ID不存在的值时强制不执行

```mysql
INSERT INTO B (B_ID, B_NAME, A_ID) VALUES (3, "C", 4);
ERROR 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`users`.`B`, CONSTRAINT `B_ibfk_1` FOREIGN KEY (`A_ID`) REFERENCES `A` (`A_ID`))
```

#### 索引

可以用来快速定位具有特征值的数据

#### 唯一索引

索引列的值必须唯一，允许空值

创建唯一索引

```mysql
CREATE UNIQUE INDEX A_INDEX ON A(A_NAME)
```

