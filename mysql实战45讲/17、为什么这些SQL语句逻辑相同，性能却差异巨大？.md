# 问题
1. 条件字段函数操作
2. 隐式类型转换
3. 隐式字符编码转换
4. 总结

# 条件字段函数操作
- MySQL的规定如果对字段做了函数计算，就用不上索引了，导致优化器就决定放弃走树搜索功能(全索引扫描)
- 对索引字段做函数操作，可能会破坏索引值的有序性，因此优化器就决定放弃走树搜索功能
- 用不上索引：对条件字段使用了函数操作：select count(*) from tradelog where month(t_modified)=7;
- 能用上索引：select count(*) from tradelog where
    -> (t_modified >= '2016-7-1' and t_modified<'2016-8-1') or
    -> (t_modified >= '2017-7-1' and t_modified<'2017-8-1') or 
    -> (t_modified >= '2018-7-1' and t_modified<'2018-8-1');

# 隐式类型转换
- MySQL中，字符串和数字做比较的话，是将字符串转换成数字，下面的tradeid的varchar类型，输入的参数是int类型
- 用不上索引：隐式转换导致对条件字段使用了函数操作
  - select * from tradelog where tradeid=110717; 对优化器来说等价于select * from tradelog where  CAST(tradid AS signed int) = 110717;
- 能用上索引：select * from tradelog where tradeid="110717"

# 隐式字符编码转换
- MySQL中,2个表做joun查询，字符集不同只是条件之一，连接过程中要求在被驱动表的索引字段上加函数操作
- 用不上索引：隐式转换字符编码导致对条件字段使用了函数操作
  - 字符集 utf8mb4 是 utf8 的超集，这两个类型的字符串在做比较的时候，MySQL 内部的操作是，先把 utf8 字符串转成 utf8mb4 字符集，再做比较
  - select d.* from tradelog l, trade_detail d where d.tradeid=l.tradeid and l.id=2; /* 语句 Q1*/
  - 能用上索引：select d.* from tradelog l , trade_detail d where d.tradeid=CONVERT(l.tradeid USING utf8) and l.id=2; 

# 总结
- 对索引字段做函数操作，可能会破坏索引值的有序性，因此优化器就决定放弃走树搜索功能
- 把隐式/直接对字段做函数、加减操作的转到参数上可以解决这个问题