# 问题
1. count(*) 的实现方式
2. 计数的方式
3. 不同count的用法

# count(*) 的实现方式
- myisam引擎把一个表的总行数存在了磁盘上，执行 count(*) 的时候会直接返回这个数，效率很高
- InnoDB 引擎就麻烦了，它执行 count(*) 的时候，需要把数据一行一行地从引擎里面读出来，然后累积计数
  - 由于多版本并发控制(MVCC)的原因，InnoDB 表“应该返回多少行”也是不确定的，所以需要一行行的读，不能像myisam一样把总行数存起来

# 计数的方式
- 采用统计，show table status：索引统计的值是通过采样来估算的，show table status 命令显示的行数也不能直接使用
- 用缓存保持计数，存在丢失更新的问题、值在逻辑上不精准（两个不同的存储构成的系统，不支持分布式事务，无法拿到精确一致的视图）
- 用数据库保持计数，在计数的地方使用事务解决精准的问题（拿到精确一致的视图）
```
  CREATE TABLE `rows_stat` (
    `table_name` varchar(64) NOT NULL,
    `row_count` int(10) unsigned NOT NULL,
    PRIMARY KEY (`table_name`)
  ) ENGINE=InnoDB;
```

# 不同count的用法
- count(字段)：InnoDB 引擎会遍历整张表，把每一行的字段值取出来，server 层要什么字段，InnoDB 就返回什么字段
- count(主键ID)：InnoDB 引擎会遍历整张表，把每一行的 id 值都取出来，返回给 server 层。server 层拿到 id 后，判断是不可能为空的，就按行累加
- count(1)：InnoDB 引擎遍历整张表，但不取值。server 层对于返回的每一行，放一个数字“1”进去，判断是不可能为空的，按行累加
- count(*)：是个例外，并不会把全部字段取出来，而是优化器专门做了优化，不取值
- 按照效率排序的话，count(字段)<count(主键 id)<count(1)≈count(*)，所以我建议你，尽量使用 count(*)。