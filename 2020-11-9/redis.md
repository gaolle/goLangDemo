## docker安装redis

查看可用版本

```shell
docker search redis
```

拉取最新

```shell
docker pull redis:latest
```

查看本地镜像

```shell
docker images
```

运行redis容器

```shell
docker run -itd --name redis-learn -p 6379:6379 redis
```

查看容器的运行信息

```shell
docker ps 
```

连接redis服务

```shell
docker exec -it redis-learn /bin/bash
```

进入redis

```shell
redis-cli
```

### 字符串

#### SET

```shell
#不存在key,创建
127.0.0.1:6379> set key val
OK
127.0.0.1:6379> get key
"val"
#存在key,覆盖
127.0.0.1:6379> set key value
OK
127.0.0.1:6379> get key
"value"
#设置过期时间 单位秒 EX
127.0.0.1:6379> set key value EX 1000
OK
127.0.0.1:6379> TTL key
(integer) 993
#设置过期时间 单位毫秒 PX
127.0.0.1:6379> set key value PX 100000000
OK
127.0.0.1:6379> TTL key
(integer) 99993
#当key不存在时才执行 NX
127.0.0.1:6379> set key-noexists val NX
OK
127.0.0.1:6379> GET key-noexists
"val"
#当key存在时返回nil
127.0.0.1:6379> set key value NX
(nil)
127.0.0.1:6379> GET key
"value"
#判断key是否存在
127.0.0.1:6379> exists key1
(integer) 0
#当key存在时才执行 XX
127.0.0.1:6379> set key1 val XX
(nil)
127.0.0.1:6379> set key1 val
OK
127.0.0.1:6379> SET key1 newvalue XX
OK
127.0.0.1:6379> GET key1
"newvalue"
```

#### GET

```shell
#key存在
127.0.0.1:6379> GET key
"value"
#key不存在
127.0.0.1:6379> get key2
(nil)
#get 只能用于字符串 用于其他类型时报错
```

#### GETSET

```shell
#key存在，修改并返回旧值
127.0.0.1:6379> GETSET key newvalue
"value"
127.0.0.1:6379> get key
"newvalue"
#key不存在，返回nil
127.0.0.1:6379> GETSET key3 newvalue
(nil)
#get 只能用于字符串 用于其他类型时报错
```

#### STRLEN

```shell
#key存在返回对应字符串的长度
127.0.0.1:6379> STRLEN key
(integer) 8
#不存在返回0
127.0.0.1:6379> STRLEN key4
(integer) 0
#get 只能用于字符串 用于其他类型时报错
```

#### APPEND

```shell
#key存在增加至尾部，返回字符串的长度
127.0.0.1:6379> APPEND key 001
(integer) 11
127.0.0.1:6379> get key
"newvalue001"
#key不存在，创建新的，返回字符串的长度
127.0.0.1:6379> APPEND key5 001
(integer) 3
127.0.0.1:6379> get key5
"001"
```

#### SETRANGE

```shell
#key存在修改偏移量起始的值，返回字符串长度
127.0.0.1:6379> get key
"newvalue001"
127.0.0.1:6379> SETRANGE key 6 setrange
(integer) 14
127.0.0.1:6379> get key
"newvalsetrange"
#key不存在，偏移量之前用空白符填充 返回字符长长度
127.0.0.1:6379> SETRANGE keysetrange 6 setrange
(integer) 14
127.0.0.1:6379> get keysetrange
"\x00\x00\x00\x00\x00\x00setrange"
```

#### GETRANGE

```shell
#检索一定返回的值 闭区间
127.0.0.1:6379> GETRANGE KEY 0 1
"\x00\x00"
#-1 从末尾开始
127.0.0.1:6379> GETRANGE key 0 -1
"newvalsetrange"
#超过长度，自动截断
127.0.0.1:6379> GETRANGE key 0 100
"newvalsetrange"
127.0.0.1:6379> GETRANGE key -1 100
"e"
#只能正顺序获取
127.0.0.1:6379> GETRANGE key -1 -5
""
```

#### INCR

```shell
#key 不存在返回，初始化为0 然后执行 返回执行后的结果 自增 
127.0.0.1:6379> INCR keyincr
(integer) 1
127.0.0.1:6379> get keyincr
"1"
127.0.0.1:6379> INCR keyincr
(integer) 2
127.0.0.1:6379> get keyincr
"2"
#当key对应的值不能解析为数字时，报错
127.0.0.1:6379> set keyincr aa
127.0.0.1:6379> INCR keyincr
(error) ERR value is not an integer or out of range
```

#### INCRBY

```shell
#key 不存在返回，初始化为0 然后执行 返回执行后的结果 以步长增加值 
127.0.0.1:6379> INCRBY keyincrby 20
(integer) 20
127.0.0.1:6379> get keyincrby
"20"
127.0.0.1:6379> set keyincrby aaa
OK
#当key对应的值不能解析为数字时，报错
127.0.0.1:6379> incrby keyincrby
(error) ERR wrong number of arguments for 'incrby' command
```

##### INCRBYFLOAT

浮点数

#### DECR

自减

#### DECRBY

自减步长

#### MSET

```shell
#不存在时创建，存在时覆盖 原子操作
127.0.0.1:6379> MSET k1 v1 k2 v2 k3 v3
OK
127.0.0.1:6379> MSET k1 v11 k2 v22
OK
```

#### MSETNX

```shell
#都不存在时才执行 成功返回1 不成功返回0 原子操作 
127.0.0.1:6379> MSETNX k1 v1 k4 v4
(integer) 0
127.0.0.1:6379> MSETNX k4 v4 k5 v5
(integer) 1
127.0.0.1:6379>
```

http://redisdoc.com/string/msetnx.html