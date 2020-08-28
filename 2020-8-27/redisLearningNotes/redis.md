
​	键值存储系统，数据结构服务器，开源，原子操作，保存在内存中

##### 特点：

​	数据持久化，保存在磁盘中，重启再次加载

​	存储多数据类型数据

##### 启动 ： 

​	redis-cli --raw

##### 数据类型

​	String(字符串)

​		

​	Hash(哈希)  一个string类型的field和value的映射表 存储对象

​		//设置

​		hset key field value //field-value键值对

​		//删除 field

​		hdel key field

​		//计算field的个数

​		hlen key

​	List(列表)

​	Set(集合)

##### 指令

​	//删除已经存在的键	

​	del key_name 

​	//检查值是否存在,存在返回1

​	exists key_name

​	//设置过期时间(s)

​	expire key_name time

​	//通过键查找

​	keys key_name 

​	keys * //所有存在的键

​	//查看值的类型

​	type key_name

​	
