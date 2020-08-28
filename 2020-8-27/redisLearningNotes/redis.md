### Redis：

键值存储系统，数据结构服务器，开源，原子操作，保存在内存中

#### 特点：

数据持久化，保存在磁盘中，重启再次加载

存储多数据类型数据

#### 启动 ：

```bash
redis-cli --raw
```

#### 数据类型

##### String(字符串) 

* ##### 设置指定值(覆盖旧值，且无视类型)

```bash
set key value
```

* ##### 获取指定key的值

```bash
get key
```

* ##### 设置指定值，返回旧值

```bash
getset key value
```

* ##### 获取指定key所存储的字符串值的长度

```bash
strlen key 
```

##### Hash(哈希)  一个string类型的field和value的映射表 存储对象  field-value键值对

* ##### 设置

```bash
hset key field value
```

* ##### 删除 field

```bash
hdel key field
```

* ##### 计算field的个数

```bash
hlen key
```

* ##### 获取所有字段名

```bash
hkeys key
```

* ##### 获取哈希表所有字段的值

```bash
hvals key
```

* ##### 获取哈希表中指定字段的值

```bash
hget key field
```

##### List(列表)

##### Set(集合)

#### 指令

* ##### 删除已经存在的键	

```bash
del key_name 
```

* ##### 检查值是否存在,存在返回1

```bash
exists key_name
```

* ##### 设置过期时间(s)

```bash
expire key_name time
```

* ##### 通过键查找

```bash
keys key_name 
```

* ##### 所有存在的键	

```bash
keys *
```

* ##### 查看值的类型

```bash
type key_name
```

* ##### 返回Redis服务器信息

```bash
info 
```

* ##### 追加指定值

```bash
append key value 
```