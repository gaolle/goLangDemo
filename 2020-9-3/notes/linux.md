#### awk：格式化处理文本文件

```shell
cat text.txt
a b c d e f 
```

以空格分隔，打印第1、4项

```shell
awk '{print $1 $4}' text.txt
ad 
```

以空格分隔，打印第1、4项（按照指定格式）

```shell
awk '{printf "%-8s %-10s\n",$1,$4}' text.txt
a        d                
```

设置变量并打印出来

```shell
awk -va=1 '{print $2 $1+a}' text.txt
b1 
```

运行脚本

```shell
awk -f {awk脚本} {文件名}
```

```shell
cat text.txt
1 A                                                                                       2 B                                                                                       3 C                                                                                       4 D                                                                                       5 E                                                                                       6 F  
```

过滤第一列大于2并且第二列等于“E”的行

```shell
awk '$1>2 && $2=="E" {print $1,$2,$3}' text.txt
5 E    
```

输出第二列包含“E”，并打印第一列

```shell
awk '$2 ~ /E/ {print $1}' text.txt
5       
```

模式取反（输出第二列包含“E”，并打印第一列）

```shell
/text# awk '$2 !~ /E/ {print $1}' text.txt
1                                                                                         2                                                                                         3                                                                                         4                                                                                         6 
```

从文件中找出长度大于1的行

```shell
awk 'length>1' text.txt
1 A                                                                                       2 B                                                                                       3 C                                                                                       4 D                                                                                       5 E                                                                                       6 F                                                                         
```

#### sed利用脚本处理文件

```shell
cat sedtext                                                                               1 a                                                                                       2 b                                                                                       3 c                                                                                       4 d                          
```

在第4行之后添加5e

```shell
sed -e 4a\5e sedtext                                                                     1 a                                                                                       2 b                                                                                       3 c                                                                                       4 d                                                                                       5e                     
```

以行号显示文件，并删除第5行

```shell
nl sedtext | sed '5d'                                                                     1  1 a                                                                                   2  2 b                                                                                   3  3 c                                                                                   4  4 d                
```

以行号显示文件，并删除4到最后一行

```shell
nl sedtext | sed '4,$d'                                                                 
1  1 a                                                                         
2  2 b                                                                               
3  3 c                
```

将2-5行的内容替换为2 b

```shell
nl sedtext | sed '2,5c  2 b'                                                             1  1 a                                                                                   2 b                         
```

列出2-3行

```shell
nl sedtext | sed -n '2,3p'                                                               2  2 b                                                                                   3  3 c                  
```

搜索与b关键字有关的行

```shell
nl sedtext | sed -n '/b/p'                                                               2  2 b                
```

搜索与b关键字有关的行,并删除

```shell
nl sedtext | sed '/b/d'                                                                   1  1 a                                                                                   3  3 c                                                                                   4  4 d                                                                                   5  5e                  
```

直接修改文件中的内容

```shell
sed -i '2a 2 b' sedtext                                                                   cat sedtext                                                                               1 a                                                                                       2 b                                                                                       2 b                                                                                       3 c                                                                                       4 d                                                                                       5e                                    
```

#### ls：列出指定目录下的文件

* -a 显示所有文件
* -l 详细信息
* -t 时间顺序
* -R 层序文件输出

uniq：检查反复出现的行列

```shell
cat uniqtest                                                                             A                                                                                          A                                                                                         A                                                                                          B                                                                                         B                                                                                          C                            
```

删除重复的行

```shell
uniq uniqtest
A
B
C
```

删除重复的行,并统计次数

```shell
uniq -c uniqtext                                                                         3 A                                                                                       2 B                                                                                       1 C  
```

统计重复的行

```shell
uniq -d  uniqtext                                                                         A                                                                                         B                               
```

#### wc：用于计算字数

行数、字数、字节数

```shell
wc uniqtext                                                                             
6  6 12 uniqtext          
```

#### ip：用于显示或设置网络设备

* link：网络设备
* address：设备上的地址
* route：路由表条目
* rule：路由规则

#### vim快捷方式

##### Normal->Insert

* i键：当前光标插入
* a键：光标下一处插入

##### Normal

光标快速移动：

* hjkl 左下上右
* w 下一单词首位
* b 当前首位
* gg 文本首位
* shift + g 文本末尾
* number + shift + g 转至number行
* o 当前光标下方插入一行
* shift + o 当前光表插入一行
* x 删除光标字符
* shift + x 删除光标之前字符
* dd 删除当前行并将内容防止剪贴板
* de 删除当前光标之后并将内容防止剪贴板
* daw 删除光标所在的整个单词
* u 还原上一个操作
* shift + 4 行末
* shift + 6 行首

v 进入可视模式

* y 复制选定块
* yy 整行赋值
* p 粘贴

查找与替换

* f + 键 光标之后查找键
* F + 键 光标之前查找键
* r + 键  将光标值替换为键
