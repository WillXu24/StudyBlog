正当我想要在刚装好的Manjaro上配置数据库环境时，发现MangoDB竟然已经被官方软件源移除了，官方文档赫然写着公告：

```
MongoDB has been removed from the official repositories due to its re-licensing 
issues（由于重新授权问题，MongoDB已从官方存储库中删除）
```

这这这!谁顶得住，官方还提到如果想要安装可以通过源码自编译的方式，动辄需要180GB +可用磁盘空间，并且可能需要几个小时才能构建（在Intel i7上使用6.5小时，在使用高端NVMe的32 Xeon核心上运行1小时。），所以决定通过官方的Linux通用包安装，所以有了如下的记录。

<!--more-->

### 一、下载软件包

- 方法一：通过浏览器进入[官网](https://www.mongodb.com/download-center/community)选择版本与系统后点击下载，我这里版本选择最新的4.0.11，系统选择Linux 64-bit legacy x64。
- 方法二：通过命令行下载到指定目录

```bash
$ wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-4.0.11.tgz
```

### 二、解压软件包

解压到文件夹：～/tools/mangodb/

### 三、配置终端环境

打开用户根目录下的`.bashrc`文件

```bash
$ cd
$ vim .bashrc #没有装vim可用自带的vi或者gedit代替
```

在文件**开头**添加

```bash
#添加mongodb，路径替换为你解压的路径
export PATH=$PATH:/home/manjaro/tools/mongodb/bin
```

### 四、测试运行服务

```bash
$ source ~/.bashrc
$ mongo version
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27017/version?gssapiServiceName=mongodb
2019-07-28T16:30:57.705+0800 E QUERY    [js] Error: couldn't connect to server 127.0.0.1:27017, connection attempt failed: SocketException: Error connecting to 127.0.0.1:27017 :: caused by :: Connection refused :
connect@src/mongo/shell/mongo.js:344:17
@(connect):2:6
exception: connect failed
```

有版本号出现，说明已经成功，否则请自行检查排错。

### 五、指定数据目录

如果现在直接执行`mongod`，会发现报错，因为没有指定存储位置。

可以自己创建并指定，我选择放在mongodb的软件包同一目录下

```bash
$ cd ~/tools/mongodb
$ mkdir data
$ mongod --dbpath /home/willxu/tools/mongodb/data  #运行mongodb服务，初始化
```

### 六、命令行管理器

这时候已经可以正常启动，默认监听本地：27017端口，这时候再另开一个终端，运行自带命令行管理工具

```bash
$ mongo
MongoDB shell version v4.0.11
connecting to: mongodb://127.0.0.1:27017/?gssapiServiceName=mongodb
MongoDB server version: 4.0.11
Welcome to the MongoDB shell.
For interactive help, type "help".
For more comprehensive documentation, see
	http://docs.mongodb.org/
Questions? Try the support group
	http://groups.google.com/group/mongodb-user
Server has startup warnings: 
2019-07-28T16:40:21.295+0800 I STORAGE  [initandlisten] 
2019-07-28T16:40:21.295+0800 I STORAGE  [initandlisten] ** WARNING: Using the XFS filesystem is strongly recommended with the WiredTiger storage engine
2019-07-28T16:40:21.295+0800 I STORAGE  [initandlisten] **          See http://dochub.mongodb.org/core/prodnotes-filesystem
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] ** WARNING: This server is bound to localhost.
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          Remote systems will be unable to connect to this server. 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          Start the server with --bind_ip <address> to specify which IP 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          addresses it should serve responses from, or with --bind_ip_all to
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          bind to all interfaces. If this behavior is desired, start the
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] **          server with --bind_ip 127.0.0.1 to disable this warning.
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] 
2019-07-28T16:40:22.164+0800 I CONTROL  [initandlisten] ** WARNING: soft rlimits too low. rlimits set to 31125 processes, 1048576 files. Number of processes should be at least 524288 : 0.5 times number of files.
> 

```

可以通过命令行试着插入文档等操作了。

### 备注

- 如果命令失效，注意执行一下`source ~/.bashrc`
- 推荐GUI管理工具：[robo3t](https://robomongo.org/)