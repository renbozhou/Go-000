学习笔记

### 总结
`error` && `panic` 
是Go的异常处理逻辑和思想和其他的编程有区别。

需要掌握以下几点
##`error`几种类型
##程序如何处理`error`  重在业务代码中尽量用 `wrap(err,"error message")`,最上层处理`error`打印`error stack`



### 作业概述
`dao`层发生的错误需要`wrap`往上抛,对于`error`的`最终处理`在`controller`层处理.
 

