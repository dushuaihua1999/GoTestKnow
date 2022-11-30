# DTM学习日记

tip:

navicat for mysql 16激活

下载:[https://kkdaj.lanzouq.com/iCJfT04o3pkh](https://links.jianshu.com/go?to=https%3A%2F%2Fkkdaj.lanzouq.com%2FiCJfT04o3pkh) 密码:ec23

[AT vs XA | DTM开源项目文档](https://dtm.pub/practice/at.html#db-support)

## 1.go embed用法

![image-20221130223354102](.\dtm\image-20221130223354102.png)

使用 //go:embed 的时候，一定要引入 embed 包，可以使用 _ 来引入(import _"embed")，不然会报错：//go:embed only allowed in Go files that import "embed"
// 和 go:embed 之间不能有空格， // go:embed 这种写法是不能解析的。
//go:embed 指令只能用在包一级的变量中，不能用在函数或方法级别。



### 1.多个文件

![image-20221130223024065](.\dtm\image-20221130223024065.png)

```
//go:embed admin/dist
var admin embed.FS
```

![image-20221130223054244](.\dtm\image-20221130223054244.png)



### 2.string

![image-20221130223216010](.\dtm\image-20221130223216010.png)

### 3.[]byte

![image-20221130223236832](.\dtm\image-20221130223236832.png)

## 2.XA与AT分布式事务模式

外部XA与内部XA

外部XA用于多个数据库实例，由应用层来决定提交与回滚

内部XA用于统一数据库的多个引擎，由数据库内部机制来执行提交与回滚

### 1.XA 数据库二阶段提交

特点：1.简单易理解 2.易开发，由底层数据库自动完成 3.对资源进行了长时间的锁定，并发度低，不适合高并发的业务。

第一阶段：

各个RM(参与者)向 TM(事务管理器) 锁住资源并注册 ready状态

第二阶段:

TM收到所有RM的ready,向RM的发送commit命令，等所有RMcommit完成。全局事务commit完成，否则回滚。

### 2.AT模式

![image-20221130234827219](.\dtm\image-20221130234827219.png)

总结就是，利用  Image  + lockKey来实现回滚与一致性。

Image用来保证数据的可恢复性

lockKey用来定位发生改变的数据位置以及分布式一致性事务的确定。只有得到所有业务的lockKey保证后，各个RM才能执行本地提交。



XA 在数据库系统层面实现了行锁，原理与普通事务相同，因此一旦出现两个事务访问同一行数据，那么后一个事务会阻塞，完全不会有脏回滚的问题