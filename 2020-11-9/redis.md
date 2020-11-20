# Redis

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

#### INCRBYFLOAT

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

### 哈希

#### HSET

```shell
#不存在创建，返回创建的个数 设置多个field
127.0.0.1:6379> HSET keyhset f v f1 v1
(integer) 2
127.0.0.1:6379> HGET keyhset f
"v"
#field	存在，覆盖，返回创建的个数即0
127.0.0.1:6379> HSET keyhset f val
(integer) 0
127.0.0.1:6379> HGET keyhset f
"val"
```

#### HSETNX

```shell
127.0.0.1:6379> HGET keyhset f
"val"
#field存在，放弃 返回0
127.0.0.1:6379> HSETNX keyhset f v
(integer) 0
127.0.0.1:6379> hget keyhset f
"val"
#field不存在，创建并返回1 
127.0.0.1:6379> HSETNX keyhsetnx f v
(integer) 1
127.0.0.1:6379> HGET keyhsetnx f
"v"
```

#### HEXISTS

```shell
#检查 filed是否存在 存在返回1
127.0.0.1:6379> HEXISTS keyhset f3
(integer) 1
#检查 filed是否存在 不存在返回0
127.0.0.1:6379> HEXISTS keyhset f2
(integer) 0
```

#### HEGT

```shell
#field存在返回值
127.0.0.1:6379> HGET keyhset f
"val"
#不存在返回nil
127.0.0.1:6379> HGET keyhset ff
(nil)
```

#### HDEL

```shell
127.0.0.1:6379> HGET keyhset f
"val"
127.0.0.1:6379> HGET keyhset f1
"v1"
#field存在，返回删除field的个数
127.0.0.1:6379> HDEL keyhset f f1
(integer) 2
127.0.0.1:6379> HGET keyhset f1
(nil)
127.0.0.1:6379> HGET keyhset f
(nil)
#field不存在，返回0
127.0.0.1:6379> HDEL keyhset fff
(integer) 0
```

#### HLEN

```shell
#key存在返回域的个数
127.0.0.1:6379> HLEN keyhset
(integer) 1
127.0.0.1:6379> HKEYS keyhset
1) "f3"
#key不存在返回0
127.0.0.1:6379> HLEN keyhlen
(integer) 0
```

#### HSTRLEN

```shell
#key或者field不存在 返回0
127.0.0.1:6379> HSTRLEN keyset f3
(integer) 0
#key中的field存在 返回对应字符串长度
127.0.0.1:6379> HSTRLEN keyhset f3
(integer) 1
```

#### HINCRBY

```shell
127.0.0.1:6379> HGET keyhset f3
"v"
#value为字符串报错
127.0.0.1:6379> HINCRBY keyhset f3 1
(error) ERR hash value is not an integer
#不存在 创建 返回field的值
127.0.0.1:6379> HINCRBY keyhincrby f 1
(integer) 1
127.0.0.1:6379> HGET keyhincrby f
"1"
127.0.0.1:6379> HINCRBY keyhincrby f 2
(integer) 3
127.0.0.1:6379> HGET keyhincrby f
"3"
```

#### HINCRBYFLOAT

浮点数自增

#### HMSET

```shell
#设置多个field-value 不存在创建返回ok 存在时覆盖返回ok 不要求全匹配
127.0.0.1:6379> HMSET keyhmset f v f1 v1
OK
127.0.0.1:6379> HGET keyhmset f
"v"
127.0.0.1:6379> HGET keyhmset f1
"v1"
127.0.0.1:6379> HMSET keyhmset f val f2 v2
OK
127.0.0.1:6379> HGET keyhmset f
"val"
127.0.0.1:6379> HGET keyhmset f2
"v2"
```

#### HMGET

```shell
#返回key中field的值，不存在返回nil
127.0.0.1:6379> HMGET keyhmset f f1 f2 f3
1) "val"
2) "v1"
3) "v2"
4) (nil)
```

#### HKEYS

```shell
#存在 返回key中所有的域
127.0.0.1:6379> HKEYS keyhmset
1) "f"
2) "f1"
3) "f2"
#不存在 返回空
127.0.0.1:6379> HKEYS keyhkeys
(empty array)
```

#### HVALS

```shell
#存在 返回key中的所有值
127.0.0.1:6379> HVALS keyhmset
1) "val"
2) "v1"
3) "v2"
#不存在 返回空
127.0.0.1:6379> HVALS keyhvals
(empty array)
127.0.0.1:6379>
```

#### HGETALL

```shell
#存在 返回key中所有的filed value
127.0.0.1:6379> HGETALL keyhmset
1) "f"
2) "val"
3) "f1"
4) "v1"
5) "f2"
6) "v2"
#不存在 返回空列表
127.0.0.1:6379> HGETALL keyhgetall
(empty array)
```

### 列表

#### LPUSH

```shell
#不存在创建，返回列表中所有元素的个数  压入顺序从左到右 每次压入列表头
127.0.0.1:6379> LPUSH keypush e e1 e2
(integer) 3
127.0.0.1:6379> LPUSH keypush e
(integer) 4
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e2"
3) "e1"
4) "e"
```

#### LPUSHX

```shell
#key存在 压入 返回列表所有元素的个数 压入头部
127.0.0.1:6379> LPUSHX keypush epush
(integer) 5
127.0.0.1:6379> LRANGE keypush 0 -1
1) "epush"
2) "e"
3) "e2"
4) "e1"
5) "e"
#key不存在 不操作 返回0
127.0.0.1:6379> LPUSHX keypushx epush
(integer) 0
```

#### RPUSH

```shell
#压入尾部
127.0.0.1:6379> RPUSH keypush er
(integer) 6
127.0.0.1:6379> LRANGE keypush 0 -1
1) "epush"
2) "e"
3) "e2"
4) "e1"
5) "e"
6) "er"
```

#### RPUSHX

与LPUSHX只有压入顺序不同

#### LPOP

```shell
#移除并返回列表的头元素
127.0.0.1:6379> LPOP keypush
"epush"
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e2"
3) "e1"
4) "e"
5) "er"
#key 不存在返回nil
127.0.0.1:6379> LPOP keypush1
(nil)
```

#### RPOP

尾部移除

#### RPOPLPUSH

移除并返回列表的尾元素

```shell
#取出第一个key的尾部压入第二个key的头部
127.0.0.1:6379> RPOPLPUSH keypush keydes
"er"
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e2"
3) "e1"
4) "e"
127.0.0.1:6379> LRANGE keydes 0 -1
1) "er"
#自旋状态 取出尾部压入头部 可以用于遍历 时间复杂度为0(1)
127.0.0.1:6379> RPOPLPUSH keypush keypush
"e"
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e"
3) "e2"
4) "e1"
127.0.0.1:6379> RPOPLPUSH keypush keypush
"e1"
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e1"
2) "e"
3) "e"
4) "e2"
#当第一个key不存在时，返回nil
127.0.0.1:6379> RPOPLPUSH keypush1 keypush
(nil)
```

#### LREM

```shell
#key不存在 返回0
127.0.0.1:6379> LREM keypush1 1 e
(integer) 0
#key存在 count > 0 从头部移除count个element 返回移除个数
127.0.0.1:6379> LREM keypush 1 e
(integer) 1
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e"
3) "e"
4) "e"
5) "e"
6) "e1"
7) "e"
8) "e"
9) "e2"
#key存在 count < 0 从尾部移除count个element 返回移除个数
127.0.0.1:6379> LREM keypush -1 e
(integer) 1
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e"
3) "e"
4) "e"
5) "e"
6) "e1"
7) "e"
8) "e2"
```

#### LLEN

```shell
#key 存在 返回列表的个数
127.0.0.1:6379> LLEN keypush
(integer) 8
#key 不存在 返回0
127.0.0.1:6379> LLEN keypush1
(integer) 0
```

#### LINDEX

```shell
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "e"
3) "e"
4) "e"
5) "e"
6) "e1"
7) "e"
8) "e2"
#返回 列表中index位置的元素 从0开始 负数从0开始
127.0.0.1:6379> LINDEX keypush 0
"e"
127.0.0.1:6379> LINDEX keypush -1
"e2"
#当index超过列表数时，返回nil
127.0.0.1:6379> LINDEX keypush 10
(nil)
```

#### LINSERT

```shell
#key存在 AFTER 在e的之后插入eee 只会检索到第一个
127.0.0.1:6379> LINSERT keypush AFTER e eee
(integer) 9
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "eee"
3) "e"
4) "e"
5) "e"
6) "e"
7) "e1"
8) "e"
9) "e2"
#key存在 BEFORE 在e的之前插入eee 只会检索到第一个
127.0.0.1:6379> LINSERT keypush BEFORE e eee
(integer) 10
127.0.0.1:6379> LRANGE keypush 0 -1
 1) "eee"
 2) "e"
 3) "eee"
 4) "e"
 5) "e"
 6) "e"
 7) "e"
 8) "e1"
 9) "e"
10) "e2"
#key存在 检索值不存在 返回-1
127.0.0.1:6379> LINSERT keypush BEFORE ee eee
(integer) -1
#key不存在 返回0
127.0.0.1:6379> LINSERT keypush1 BEFORE ee eee
(integer) 0
```

#### LSET

```shell
#key 存在 将index=0位置元素设为eset
127.0.0.1:6379> LSET keypush 0 eset
OK
127.0.0.1:6379> LRANGE keypush 0 -1
 1) "eset"
 2) "e"
 3) "eee"
 4) "e"
 5) "e"
 6) "e"
 7) "e"
 8) "e1"
 9) "e"
10) "e2"
#key存在 index超出或不存在 返回错误
127.0.0.1:6379> LSET keypush 11 eset
(error) ERR index out of range
#key不存在 返回错误
127.0.0.1:6379> LSET keypush1 11 eset
(error) ERR no such key
```

#### LRANGE

```shell
#检索一定范围的值 从0开始 -1 代表尾部 闭区间 只能正向检索
127.0.0.1:6379> LRANGE keypush 0 -1
 1) "eset"
 2) "e"
 3) "eee"
 4) "e"
 5) "e"
 6) "e"
 7) "e"
 8) "e1"
 9) "e"
10) "e2"
127.0.0.1:6379> LRANGE keypush 0 100
 1) "eset"
 2) "e"
 3) "eee"
 4) "e"
 5) "e"
 6) "e"
 7) "e"
 8) "e1"
 9) "e"
10) "e2"
127.0.0.1:6379> LRANGE keypush -1 100
1) "e2"
#只能正向索引
127.0.0.1:6379> LRANGE keypush 11 100
(empty array)
```

#### LTRIM

```shell
127.0.0.1:6379> LRANGE keypush 0 100
 1) "eset"
 2) "e"
 3) "eee"
 4) "e"
 5) "e"
 6) "e"
 7) "e"
 8) "e1"
 9) "e"
10) "e2"
#key 存在 保留范围之间的 闭区间 成功返回ok
#注意 索引从0开始
127.0.0.1:6379> LTRIM keypush  1 10
OK
127.0.0.1:6379> LRANGE keypush 0 -1
1) "e"
2) "eee"
3) "e"
4) "e"
5) "e"
6) "e"
7) "e1"
8) "e"
9) "e2"
```

#### BLPOP

#### BRPOP

#### BRPOPLPUSH

### 集合

#### SADD

```shell
#将元素加入到集合key中， 如果已经存在忽略 返回新加入元素的个数
127.0.0.1:6379> SADD keyset a b c d
(integer) 4
127.0.0.1:6379> SMEMBERS keyset
1) "d"
2) "c"
3) "b"
4) "a"
127.0.0.1:6379> sadd keyset a e f
(integer) 2
127.0.0.1:6379> SMEMBERS keyset
1) "a"
2) "f"
3) "e"
4) "c"
5) "d"
6) "b"
127.0.0.1:6379> sadd keyset f a
(integer) 0
127.0.0.1:6379> SMEMBERS keyset
1) "f"
2) "e"
3) "c"
4) "d"
5) "b"
6) "a"
```

#### SISMEMBER

```shell
#判断元素是否存在于集合中，存在返回1 不存在返回0
127.0.0.1:6379> SISMEMBER keyset a
(integer) 1
127.0.0.1:6379> SISMEMBER keyset g
(integer) 0
```

#### SPOP

```shell
#移除并返回集合中一个或多个元素 count可选
127.0.0.1:6379> SPOP keyset 0
(empty array)
127.0.0.1:6379> SPOP keyset 1
1) "b"
#key不存在时返回nil
127.0.0.1:6379> SPOP keyset1
(nil)
127.0.0.1:6379> SPOP keyset1 1
(empty array)
```

#### SRANDMEMBER

```shell
#key 存在 随机返回一个元素
127.0.0.1:6379> SRANDMEMBER keyset
"d"
127.0.0.1:6379> SMEMBERS keyset
1) "e"
2) "c"
3) "d"
4) "a"
#key 存在 随机返回一个数组（可能会存在重复元素） 数组个数为 可选参数count的绝对值
127.0.0.1:6379> SRANDMEMBER keyset -3
1) "e"
2) "d"
3) "e"
127.0.0.1:6379> SRANDMEMBER keyset -5
1) "c"
2) "c"
3) "d"
4) "a"
5) "a"
127.0.0.1:6379> SRANDMEMBER keyset 5
1) "e"
2) "d"
3) "c"
4) "a"
```

#### SREM

```shell
#key存在 移除指定元素 不存在忽略 返回移除成功的个数
127.0.0.1:6379> SREM keyset g
(integer) 0
127.0.0.1:6379> SREM keyset g a
(integer) 1
127.0.0.1:6379> SMEMBERS keyset
1) "e"
2) "c"
3) "d"
```

#### SMOVE

```shell
127.0.0.1:6379> SMEMBERS keyset
1) "e"
2) "c"
3) "d"
#key类型不匹配时报错
127.0.0.1:6379> SMOVE keyset keydes a
(error) WRONGTYPE Operation against a key holding the wrong kind of value
#将key中的元素转移到指定的key中 成功返回1 失败返回0
127.0.0.1:6379> SMOVE keyset keymovedes e
(integer) 1
127.0.0.1:6379> SMOVE keyset keymovedes a
(integer) 0
127.0.0.1:6379> SMEMBERS keyset
1) "c"
2) "d"
127.0.0.1:6379> SMEMBERS keymovedes
1) "e"
```

#### SCARD

```shell
#key存在 返回集合中的元素数量
127.0.0.1:6379> SCARD keyset
(integer) 2
#key不存在 返回0
127.0.0.1:6379> SCARD keyset1
(integer) 0
```

#### SMEMBERS

```shell
#key存在 返回集合的所有元素
127.0.0.1:6379> SMEMBERS keyset
1) "c"
2) "d"
#key不存在 返回空
127.0.0.1:6379> SMEMBERS keyset1
(empty array)
```

#### SINTER

```shell
127.0.0.1:6379> SMEMBERS keyset
1) "c"
2) "d"
127.0.0.1:6379> SMEMBERS keyset1
1) "d"
2) "c"
3) "b"
4) "a"
#返回指定集合中的交集 当key不存在时当空集合处理
127.0.0.1:6379> SINTER keyset keyset1
1) "c"
2) "d"
```

#### SINTERSTORE

```shell
127.0.0.1:6379> SMEMBERS keydesc
1) "a"
#将两个集合的交集存储在指定的集合中 返回交集的个数
127.0.0.1:6379> SINTERSTORE keydesc keyset keyset1
(integer) 2
127.0.0.1:6379> SMEMBERS keydesc
1) "d"
2) "c"
```

#### SUNION

并集

#### SUNIONSTORE

并集

#### SDIFF

```shell
#返回给定集合的差集 以第一个key作为基准 判断第二个key
127.0.0.1:6379> SMEMBERS keyset
1) "c"
2) "d"
3) "a"
127.0.0.1:6379> SMEMBERS keydesc
1) "e"
2) "d"
3) "c"
4) "a"
127.0.0.1:6379> SDIFF keyset keydesc
(empty array)
127.0.0.1:6379> SDIFF keydesc keyset
1) "e"
```

#### SDIFFSTORE

```shell
#将指定集合的差集存储到指定集合 res 中
127.0.0.1:6379> SDIFFSTORE res keydesc keyset
(integer) 1
127.0.0.1:6379> SMEMBERS res
1) "e"
```

### 有序集合

#### ZADD

```shell
#key不存在创建 并返回新添加的元素 score可以是整数值或者浮点数值 有序集合score可以重复 member不能重复
127.0.0.1:6379> ZADD keyzadd 1 a
(integer) 1
127.0.0.1:6379> ZADD keyzadd 2 b
(integer) 1
127.0.0.1:6379> ZADD keyzadd 4 d
(integer) 1
127.0.0.1:6379> ZADD keyzadd 3 c
(integer) 1
127.0.0.1:6379> ZRANGE keyzadd 0 -1
1) "a"
2) "b"
3) "c"
4) "d"
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "a"
2) "1"
3) "b"
4) "2"
5) "c"
6) "3"
7) "d"
8) "4"
127.0.0.1:6379> ZADD keyzaa 6 a
(integer) 1
#key不存在创建 score存在根据score更新位置并重新插入不计入返回结果（新增） socre存在并score不变忽略
127.0.0.1:6379> ZADD keyzaa 6 a 2 b
(integer) 1
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "a"
2) "1"
3) "b"
4) "2"
5) "c"
6) "3"
7) "d"
8) "4"
127.0.0.1:6379> ZADD keyzadd 6 a 2 b
(integer) 0
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "b"
2) "2"
3) "c"
4) "3"
5) "d"
6) "4"
7) "a"
8) "6"
```

#### ZSCORE

```shell
#返回key中member对应的score
127.0.0.1:6379> ZSCORE keyzadd a
"6"
#key不存在或者member不存在 返回nil
127.0.0.1:6379> ZSCORE keyzadd f
(nil)
```

#### ZINCRBY

```shell
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "b"
2) "2"
3) "c"
4) "3"
5) "d"
6) "4"
7) "a"
8) "6"
#key存在 member存在 给member对应的score增加上增量 返回新的score 字符串格式
127.0.0.1:6379> ZINCRBY keyzadd 5 b
"7"
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "c"
2) "3"
3) "d"
4) "4"
5) "a"
6) "6"
7) "b"
8) "7"
127.0.0.1:6379> ZINCRBY keyzadd 5 f
"5"
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
127.0.0.1:6379> ZINCRBY keyzadd 5 e
"5"
#score可以存在多个
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "e"
 6) "5"
 7) "f"
 8) "5"
 9) "a"
10) "6"
11) "b"
12) "7"
#member不能同时存在多个,只能唯一
127.0.0.1:6379> ZINCRBY keyzadd 6 e
"11"
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
11) "e"
12) "11"
```

#### ZCARD

```shell
#key存在 返回基数（个数）
127.0.0.1:6379> ZCARD keyzadd
(integer) 6
#key不存在 返回0
127.0.0.1:6379> ZCARD keyzadd1
(integer) 0
```

#### ZCOUNT

```shell
#返回key中 score区间的member个数 闭区间
127.0.0.1:6379> ZCOUNT keyzadd 4 6
(integer) 3
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
11) "e"
12) "11"
```

#### ZRANGE

```shell
#key存在 返回score区间的成员 递增 WITHSCORES可选 依次返回成员和score 下标从0开始 -1代表尾部
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
11) "e"
12) "11"
127.0.0.1:6379> ZRANGE keyzadd 0 -1
1) "c"
2) "d"
3) "f"
4) "a"
5) "b"
6) "e"
```

#### ZREVRANGE

递减返回结果

#### ZRANGBYSCORE

递增返回score介于min和max之间的成员 闭区间

#### ZREVRANGESCORE

递减返回score介于min和max之间的成员 闭区间

#### ZRANK

```shell
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
11) "e"
12) "11"
#返回元素递增排序 排序从0开始 
127.0.0.1:6379> ZRANK keyzadd f
(integer) 2
#元素不存在 返回nil
127.0.0.1:6379> ZRANK keyzadd g
(nil)
```

#### ZREVRANK

递减输出排序位次

#### ZREM

```shell
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
 1) "c"
 2) "3"
 3) "d"
 4) "4"
 5) "f"
 6) "5"
 7) "a"
 8) "6"
 9) "b"
10) "7"
11) "e"
12) "11"
#移除key中元素 不存在元素忽略 返回成功移除元素的个数
127.0.0.1:6379> ZREM keyzadd c d
(integer) 2
127.0.0.1:6379> ZRANGE keyzadd 0 -1 WITHSCORES
1) "f"
2) "5"
3) "a"
4) "6"
5) "b"
6) "7"
7) "e"
8) "11"
127.0.0.1:6379> ZREM keyzadd g
(integer) 0
```

#### ZREMRANGEBYRANK

移除排名区间的元素 闭区间

#### ZREMRANGEBYSCORE

移除score区间的元素 闭区间

### 位图

#### SETBIT

```shell
#设置或清除key所存储字符串中指定偏移量上的位（bit） 1设置0清除
#偏移量必须大于0
#返回指定偏移量之前的位
#设置
127.0.0.1:6379> SETBIT keysetbit 100 1
(integer) 0
127.0.0.1:6379> GETBIT keysetbit 100
(integer) 1
#清除
127.0.0.1:6379> SETBIT keysetbit 100 0
(integer) 1
127.0.0.1:6379> GETBIT keysetbit 100
(integer) 0
```

#### GETBIT

```shell
#返回字符串中偏移量的位
127.0.0.1:6379> SETBIT keysetbit 100 1
(integer) 0
127.0.0.1:6379> GETBIT keysetbit 100
(integer) 1
```

#### BITCOUNT

```shell
#返回设置为1的位的数量
127.0.0.1:6379> BITCOUNT keysetbit
(integer) 2
```

#### BITPOS

```shell
#返回第一个位为值的二进制的位置
127.0.0.1:6379> BITPOS keysetbit 1
(integer) 100
```

#### BITOP

```shell
127.0.0.1:6379> SETBIT keysetbit1 100 1
(integer) 0
127.0.0.1:6379> SETBIT keysetbit1 101 1
(integer) 0
127.0.0.1:6379> GETBIT keysetbit1 100
(integer) 1
127.0.0.1:6379> GETBIT keysetbit1 101
(integer) 1
127.0.0.1:6379> GETBIT keysetbit 101
(integer) 1
127.0.0.1:6379> GETBIT keysetbit 100
(integer) 0
#对多个key进行元操作
#AND 与 OR 或 XOR 异或
#NOT 对key取反存到一个key中
#返回存取key的字符串的长度
127.0.0.1:6379> BITOP AND destkey keysetbit keysetbit1
(integer) 13
127.0.0.1:6379> GETBIT destkey 100
(integer) 0
127.0.0.1:6379> GETBIT destkey 101
(integer) 1
```

# mongoDB

查看可用版本

```shell
docker search mongo
```

拉取最新

```shell
docker pull mongo:latest
```

查看本地镜像

```shell
docker images
```

运行redis容器

```shell
docker run -itd --name mongo-learn -p 27017:27017 mongo
```

查看容器的运行信息

```shell
docker ps 
```

连接redis服务

```shell
docker exec -it mongo-learn mongo
```

#### 查询数据库

db

#### 选择数据库

use dbname

#### 查询数据库中的集合

show collections

插入数据时，_id 没有选择性插入，数据库将自己创建一个唯一 （类似主键）

集合通过_id排序

数据库名不区分大小写

BSON 最大 16MB 嵌套不超过100层

更新数据时，_id字段不能更新 设置upsert:true不存在时插入新的collection

索引 除主键外还可以创建其他的索引 底层数据结构为B树 索引创建不能重命名，只能删除当前索引然后重新创建新的索引

1：升序 -1：降序

多个索引匹配时，顺序必须一致 多索引不超过32个



