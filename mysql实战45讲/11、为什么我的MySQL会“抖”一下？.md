# 问题
1. 脏页和干净页
2. 刷脏页的情况有哪些
3. 刷脏页的控制策略

# 脏页和干净页
- innodb在处理更新语句的时候使用了WAL机制，写内存后写redo log，返回给客户端，本次更新成功。接下来会涉及刷盘操作
- 当内存数据页和磁盘数据页内容不一致的时候，我们称这个内存页为脏页
- 当内存数据写到磁盘后，内存和磁盘上的数据页内容一致了，称为干净页

# 刷脏页的情况有哪些
- MySQL偶尔抖一下，那个瞬间可能就是在刷脏页
1. innodb的redolog写满了，把checkpoint往前推进刷脏页，redolog留出空间可以继续写
  - 需要避免的，会导致系统的压力增大。整个系统就不能再接受更新了，所有的更新都必须堵住

2. 系统内存不足，当需要更新的内存页，而内存页不够用。需要淘汰一些数据页，当淘汰的是脏页就要将脏页写入磁盘
  - 这种情况其实是常态，缓冲池buffer pool不够用开始淘汰一些数据页
  - 避免一个查询要淘汰的脏页个数太多，会导致查询的响应时间明显变长

3. 系统空闲的时候会进行刷脏页
  - 对系统没有什么压力

4. MySQL正常关闭的情况，会把脏页刷到磁盘上
  - 对系统没有什么压力

# 刷脏页的控制策略
```
1. 告诉innodb主机的IO能力：通过 fio 这个工具来测试出磁盘的IOPS，设置到innodb_io_capacity上
2. 关注脏页比例不要经常接近75%：
      mysql> select VARIABLE_VALUE into @a from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_dirty';
      select VARIABLE_VALUE into @b from global_status where VARIABLE_NAME = 'Innodb_buffer_pool_pages_total';
      select @a/@b;
3. 关掉连坐机制：innodb_flush_neighbors值为 0 时表示不找邻居，自己刷自己的。不要连带邻居的一起刷，有可能会很多
```
