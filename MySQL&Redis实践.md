# MySQL主从复制

## 修改主库配置文件

1. 开启bin log日志功能
2. 设置server-id
3. 指定需要同步的数据库

```
[mysqld] 
## 设置server_id，同一局域网中需要唯一 
server_id=101  
## 指定不需要同步的数据库名称 
binlog-ignore-db=mysql   
## 开启二进制日志功能 
log-bin=mall-mysql-bin   
## 设置二进制日志使用内存大小（事务） 
binlog_cache_size=1M   
## 设置使用的二进制日志格式（mixed,statement,row） 
binlog_format=mixed   
## 二进制日志过期清理时间。默认值为0，表示不自动清理。 
expire_logs_days=7   
## 跳过主从复制中遇到的所有错误或指定类型的错误，避免slave端复制中断。 
## 如：1062错误是指一些主键重复，1032错误是因为主从数据库数据不一致 
slave_skip_errors=1062 
## 设置utf8
collation_server = utf8_general_ci 
## 设置server字符集
character_set_server = utf8 
[client]
default_character_set=utf8 
```

## 在主库配置slave访问master的用户的IP权限

```
# 创建同步用户
CREATE USER 'slave'@'%' IDENTIFIED WITH mysql_native_password by '123456';

# 同步用户授权
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';

#在主库中查看主从同步状态
show master status;
```

## 修改从库配置文件

在配置文件中指定唯一id，即server-id

```
# 添加配置文件
[mysqld] 
## 设置server_id，同一局域网中需要唯一 
server_id=102 
## 指定不需要同步的数据库名称 
binlog-ignore-db=mysql   
## 开启二进制日志功能，以备Slave作为其它数据库实例的Master时使用 
log-bin=mall-mysql-slave1-bin   
## 设置二进制日志使用内存大小（事务） 
binlog_cache_size=1M   
## 设置使用的二进制日志格式（mixed,statement,row） 
binlog_format=mixed   
## 二进制日志过期清理时间。默认值为0，表示不自动清理。 
expire_logs_days=7   
## 跳过主从复制中遇到的所有错误或指定类型的错误，避免slave端复制中断。 
## 如：1062错误是指一些主键重复，1032错误是因为主从数据库数据不一致 
slave_skip_errors=1062   
## relay_log配置中继日志 
relay_log=mall-mysql-relay-bin   
## log_slave_updates表示slave将复制事件写进自己的二进制日志 
log_slave_updates=1   
## slave设置为只读（具有super权限的用户除外） 
read_only=1 
## 设置utf8
collation_server = utf8_general_ci 
## 设置server字符集
character_set_server = utf8 
[client]
default_character_set=utf8 
```

## 从库配置关联主库

```
change master to master_host='主库IP',master_port=主库端口号,master_user='主库用户名',master_password='主库密码',master_log_file='主库日志文件名',master_log_pos=主库position;
master_log_pos可配置为0，系统会自动匹配，需验证
```

## 在从数据库中查看主从同步状态

```
show slave status\G;
```

## 在从库中开启主从同步

```
start slave;
```

## 查看从数据库状态

![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E9%AB%98%E7%BA%A7%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/2.png)

## 主从复制测试

```
1.主机新建数据库 --->  使用数据库 ---> 新建表 --->插入数据 ， ok
2.从机使用库 ---> 查看记录 ok
```

# Redis集群

## cluster模式-哈希槽分区进行亿级数据存储

单机单台不可能，肯定是分布式存储，用redis如何落地？

1. 哈希取余分区

   > hash(key) % N个机器台数，计算出哈希值，用来决定数据映射到哪一个节点上。 
   >
   > 优点：
   >   简单粗暴，直接有效，只需要预估好数据规划好节点，例如3台、8台、10台，就能保证一段时间的数据支撑。使用Hash算法让固定的一部分请求落到同一台服务器上，这样每台服务器固定处理一部分请求（并维护这些请求的信息），起到负载均衡+分而治之的作用。 
   >
   > 缺点：
   >    原来规划好的节点，进行扩容或者缩容就比较麻烦了额，不管扩缩，每次数据变动导致节点有变动，映射关系需要重新进行计算，在服务器个数固定不变时没有问题，如果需要弹性扩容或故障停机的情况下，原来的取模公式就会发生变化：Hash(key)/3会变成Hash(key) /?。此时地址经过取余运算的结果将发生很大变化，根据公式获取的服务器也会变得不可控。 
   > 某个redis机器宕机了，由于台数数量变化，会导致hash取余全部数据重新洗牌

   

2. 一致性哈希算法分区

   > ？当服务器个数发生变动时，尽量减少影响客户端到服务器的映射关系
   >
   >  一致性哈希算法必然有个hash函数并按照算法产生hash值，这个算法的所有可能哈希值会构成一个全量集，这个集合可以成为一个hash空间[0,2^32-1]，这个是一个线性空间，但是在算法中，我们通过适当的逻辑控制将它首尾相连(0 = 2^32),这样让它逻辑上形成了一个环形空间，这个由2^32个点组成的圆环称为Hash环。
   >
   > ![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E9%AB%98%E7%BA%A7%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/3.png)
   >
   > 【服务器IP节点映射】选择服务器的IP或主机名作为关键字进行哈希(hash(ip))，这样每台机器就能确定其在哈希环上的位置。
   >
   > 【key落到服务器的落键规则】 当我们需要存储一个kv键值对时，首先计算key的hash值，hash(key)，将这个key使用相同的函数Hash计算出哈希值并确定此数据在环上的位置， 从此位置沿环顺时针“行走” ，第一台遇到的服务器就是其应该定位到的服务器，并将该键值对存储在该节点上。 如我们有Object A、Object B、Object C、Object D四个数据对象，经过哈希计算后，在环空间上的位置如下：根据一致性Hash算法，数据A会被定为到Node A上，B被定为到Node B上，C被定为到Node C上，D被定为到Node D上。
   >
   > ![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E9%AB%98%E7%BA%A7%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/5.png)
   >
   > 优点：
   >
   > 容错性
   >
   > ```
   > 假设Node C宕机，可以看到此时对象A、B、D不会受到影响，只有C对象被重定位到Node D。一般的，在一致性Hash算法中，如果一台服务器不可用，则 受影响的数据仅仅是此服务器到其环空间中前一台服务器（即沿着逆时针方向行走遇到的第一台服务器）之间数据 ，其它不会受到影响。简单说，就是C挂了，受到影响的只是B、C之间的数据，并且这些数据会转移到D进行存储。 
   > ```
   >
   > 扩展性
   >
   > ```
   > 数据量增加了，需要增加一台节点NodeX，X的位置在A和B之间，那收到影响的也就是A到X之间的数据，重新把A到X的数据录入到X上即可， 
   > 不会导致hash取余全部数据重新洗牌。 
   > ```
   >
   > 缺点：
   >
   > 数据倾斜问题
   >
   > ```
   > 一致性Hash算法在服务节点太少时 ，容易因为节点分布不均匀而造成 数据倾斜 （被缓存的对象大部分集中缓存在某一台服务器上）问题， 
   > 例如系统中只有两台服务器：
   > ```
   >
   > ![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E9%AB%98%E7%BA%A7%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/8.png)
   >
   > 总结：
   >
   > ```
   > 为了在节点数目发生改变时尽可能少的迁移数据 
   >   
   > 将所有的存储节点排列在首尾相接的Hash环上，每个key在计算Hash后会 顺时针找到临近的存储节点存放。 
   > 而当有节点加入或退出时仅影响该节点在Hash环上 顺时针相邻的后续节点 。     
   > 优点 
   > 加入和删除节点只影响哈希环中顺时针方向的相邻的节点，对其他节点无影响。 
   >   
   > 缺点  
   > 数据的分布和节点的位置有关，因为这些节点不是均匀的分布在哈希环上的，所以数据在进行存储时达不到均匀分布的效果。 
   > ```

3. 哈希槽分区

> ​     ？解决一致性哈希算法数据倾斜问题
>
> 哈希槽就是一个数组，数组[0,2^14-1]形成的hash slot 空间
>
> 解决均匀分配，在数据和节点之间加入了一层，称为哈希槽（slot)，用于管理数据和节点之间的关系，相当于节点上放的是槽
>
> ![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E9%AB%98%E7%BA%A7%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/9.png)
>
> ```
> 槽解决的是粒度问题，相当于把粒度变大了，这样便于数据移动。 
> 哈希解决的是映射问题，使用key的哈希值来计算所在的槽，便于数据分配。
> 多少个hash槽 ？
> 一个集群只能有16384个槽，编号0-16383（0-2^14-1）。这些槽会分配给集群中的所有主节点，分配策略没有要求。可以指定哪些编号的槽分配给哪个主节点。集群会记录节点和槽的对应关系。解决了节点和槽的关系后，接下来就需要对key求哈希值，然后对16384取余，余数是几key就落入对应的槽里。slot = CRC16(key) % 16384。以槽为单位移动数据，因为槽的数目是固定的，处理起来比较容易，这样数据移动问题就解决了。 
> ```

## Redis集群3主3从扩缩容配置

1. 环境准备

   ```
   #关闭防火墙
   systemctl stop firewalld
   #查看防火墙状态
   systemctl status firewalld
   #拉去镜像
   docker pull redis:6.0.8
   ```

   ![img](https://img2022.cnblogs.com/blog/741283/202210/741283-20221031110850161-173945626.png)

2. 新建6个docker容器redis实例

   ```
   docker run -d --name redis-node-1 --net host --privileged=true -v /data/redis/share/redis-node-1:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6381 
     
   docker run -d --name redis-node-2 --net host --privileged=true -v /data/redis/share/redis-node-2:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6382 
     
   docker run -d --name redis-node-3 --net host --privileged=true -v /data/redis/share/redis-node-3:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6383 
     
   docker run -d --name redis-node-4 --net host --privileged=true -v /data/redis/share/redis-node-4:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6384 
     
   docker run -d --name redis-node-5 --net host --privileged=true -v /data/redis/share/redis-node-5:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6385 
     
   docker run -d --name redis-node-6 --net host --privileged=true -v /data/redis/share/redis-node-6:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6386 
   
   # 创建并运行docker容器实例
   docker run 
   # 容器名字
   --name redis-node-6
   # 使用宿主机的IP和端口，默认
   --net host
   # 获取宿主机root用户权限
   --privileged=true
   # 容器卷，宿主机地址:docker内部地址
   -v /data/redis/share/redis-node-6:/data
   # redis镜像和版本号
   redis:6.0.8
   # 开启redis集群
   --cluster-enabled yes 
   # 开启持久化
   --applendonly yes 
   ```

3. 进入容器构建6台机器主从集群关系

   ```
   docker exec -it redis-node-1 /bin/bash
   
   #构建关系
   redis-cli --cluster create 192.168.206.10:6381 192.168.206.10:6382 192.168.206.10:6383 192.168.206.10:6384 192.168.206.10:6385 192.168.206.10:6386 --cluster-replicas 1
   --clusterr-replicas 1表示为每个master创建一个slave节点
   ```

   

4. 登录某个阶段，查看集群和节点状态

   ```
   docker exec -it redis-node-6 /bin/bash 
   #登录redis
   #普通方式登录，存储数据时会出现moved重定向操作
   redis-cli -p 6386
   -c 采用集群策略连接，设置数据会自动切换到相应的写主机，优化路由
   
   #查看集群状态
   cluster info
   #查看节点状态
   cluster nodes
   ```

## 主从容错切换迁移

```
docker exec -it redis-node-6 /bin/bash

#检查集群情况，确定主从关系
redeis-cli --cluster check 192.168.206.10:6386

docker stop redis-node-1
docker exec -it redis-node-2 /bin/bash
redis-cli -p 6382 -c

#查看节点状态，node-4成为master
cluster nodes

#node1恢复，node1成为slave
docker start redis-node-1
cluster nodes
```

## 集群主从扩容

```
#新增2主机
docker run -d --name redis-node-7 --net host --privileged=true -v /data/redis/share/redis-node-7:/data --cluster-enabed yes --appendonly yes --port 6387
docker run -d --name redis-node-8 --net host --privileged=true -v /data/redis/share/redis-node-8:/data --cluster-enabled yes appendonly yes --port 6388
docker ps
docker exec -it redis-node-7 /bin/bash

#新增节点6387作为master节点加入原集群
redis-cli --cluster add-node 192.168.206.10:6387 192.168.206.10:6381
redis-cli --cluster check 192.168.206.10:6381

#重新分派槽号
redis-cli --cluster reshard 192.168.206.10:6381
redis-cli --cluster check 192.168.206.10:6381

#为新主节点分配从节点
redis-cli --cluster add-node IP:新slave端口 IP:新master端口 --cluster-slave --cluster-master-id 新master id

redis-cli --cluster add-node 192.168.206.10:6388 192.168.206.10:6387 --cluster-slave --cluster-master-id 新master id

redis-cli --cluster check 192.168.206.10:6381
```

## 主从缩容

1. 从集群中删除从节点

   ```
   redis-cli --cluster del-node ip:从机端口 从机在集群中的id
   ```

2. check

   ```
   redis-cli --cluster check ip:集群内某端口
   ```

3. 清空主节点槽号，重新分配

   ```
   redis-cli --cluster reshard ip:主节点IP
   redis-cli --cluster reshard 192.168.206.10:6381
   ```

4. 删除要删除的从节点的master

   ```
   redis-cli --cluster del-node ip:端口 slave的master节点在集群中id
   ```

5. check

   