学习笔记

### 总结
`error` && `panic` 
是Go的异常处理逻辑和思想和其他的编程有区别。

需要掌握以下几点
##`error`几种类型
##程序如何处理`error`  重在业务代码中尽量用 `wrap(err,"error message")`,最上层处理`error`打印`error stack`



### 作业概述
为了更好的体会课上所讲的不同层级之间error的处理,所以没有采用伪代码的方式,全部都是真实可运行的代码

### week02 作业结构
```
├── README.md  
├── controller ---------------  处理web request
│   └── usercontroller.go ---   userhandler   
├── dao ----------------------  数据操作层
│   └── userdao.go -----------  userdao
├── db -----------------------  数据库现先关
│   └── daoconnector.go ------  初始化db连接
│   └── ddl ------------------  mock数据脚本
│       └── person.sql -------  测试表sql
├── init.sh ------------------  go mod 脚本
├── model --------------------  db 相关model
│   └── usermodel.go ---------  用户model
├── service ------------------  具体业务处理层
│   └── userservice.go -------  user业务出来
└── userbootapplaction.go ----  应用启动
```

### 作业思路
对于 `ErrNoRows` 这个错误是不是应该`wrap`往上抛我是这么理解的,首先对于`dao层`查询的错误都是直接`wrap`往上抛,
到了`service`层之后,我认为有两种情况,一种是单纯的查询直接返回,一种查询之后还会有其他的逻辑操作,这两种情况需要是
不一样的,第一种情况是不做处理直接将`dao`层的结果和错误晚上抛,由`controller(web)`层去最终处理这个error,第二种
情况是后续还会有数据逻辑操作,就需要先判断`dao`层查询的数据是否正确没有错误,如果有错误直接往上抛,如果没有错误继续
后面的逻辑,当`service`层后面的逻辑出错的时候,这个时候需要对`service`的错误进行`wrap`,因为第一次错误发生在了
`service`层,对于错误的处理,最终是在`controller(web)`层,我在`controller`层对于error的处理是这样的,对于
 `ErrNoRows`的错误不会记录error日志,因为找不到记录不能当做error级别的日志进行处理,直接返回给用户一个找不到数据
 的提示即可,对于非`ErrNoRows`的error,会进行堆栈日志的记录,例如db发生问题突然连接不上,这种直接返回服务器错误的
 友好提示.
 所以单从回单问题的来看,`dao`层发生的错误需要`wrap`往上抛,对于`error`的`最终处理`在`controller`层处理.
 
 ### 其他
 具体细节请看代码
