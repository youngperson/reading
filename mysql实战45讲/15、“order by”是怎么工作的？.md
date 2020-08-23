# 问题
1. sort_buffer
2. 全字段排序和rowid排序

# sort_buffer
```
使用explain命令查看带order by的查询语句执行情况
有时候可以看到Extra字段中的Using filesort表示就是需要排序，MySQL 会给每个线程分配一块内存用于排序，称为 sort_buffer
```

# 全字段排序和rowid排序
- sort_buffer_size，就是 MySQL 为排序开辟的内存（sort_buffer）的大小
- 如果要排序的数据量小于 sort_buffer_size，排序就在内存中完成
- 如果排序数据量太大，内存放不下，则不得不利用磁盘临时文件辅助排序
  - 可以用下面介绍的方法，来确定一个排序语句是否使用了临时文件
```
/* 打开 optimizer_trace，只对本线程有效 */
SET optimizer_trace='enabled=on'; 

/* @a 保存 Innodb_rows_read 的初始值 */
select VARIABLE_VALUE into @a from  performance_schema.session_status where variable_name = 'Innodb_rows_read';

/* 执行语句 */
select city, name,age from t where city='杭州' order by name limit 1000; 

/* 查看 OPTIMIZER_TRACE 输出 */
SELECT * FROM `information_schema`.`OPTIMIZER_TRACE`\G

/* @b 保存 Innodb_rows_read 的当前值 */
select VARIABLE_VALUE into @b from performance_schema.session_status where variable_name = 'Innodb_rows_read';

/* 计算 Innodb_rows_read 差值 */
select @b-@a;
``` 
- max_length_for_sort_data，是 MySQL 中专门控制用于排序的行数据的长度的一个参数
  - 如果内存够，就要多利用内存，尽量减少磁盘访问
  - MySQL 实在是担心排序内存太小，会影响排序效率，才会采用 rowid 排序算法，这样排序过程中一次可以排序更多行，但是需要再回到原表去取数据
  - MySQL 认为内存足够大，会优先选择全字段排序，把需要的字段都放到 sort_buffer 中，这样排序后就会直接从内存里面返回查询结果了，不用再回到原表去取数据