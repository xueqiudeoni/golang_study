# docker 命令

## 新建+启动容器

```
docker run [options] IMAGES [COMMAND][ARG...]
```

options说明（常用）：

--name="容器新名字"：为容器指定一个名称；

-d：后台运行容器并返回容器id，也即启动守护式容器（后台运行）；

-i：以交互模式运行容器，通常与-t同时使用；

-t：为容器重新分配一个伪输入终端，通常与-i同时使用，也即启动交互式容器(前台有伪终端，等待交互)；

-P：随机端口映射；

-p：指定端口映射

![32](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E5%9F%BA%E7%A1%80%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/32.png)

```
docker run -it centos /bin/bash
参数说明：
-i: 交互式操作。
-t: 终端。
centos : centos 镜像。
/bin/bash：放在镜像名后的是命令，这里我们希望有个交互式 Shell，因此用的是/bin/bash。 要退出终端，直接输入 exit
```

## 列出当前所有正在运行的容器

```
docker ps [options]
#options说明
-a:列出所有正在运行+历史上运行过的容器
-l:显示最近创建的容器
-n:显示最近创建的n个容器
-q:静默模式，只显示容器编号
```

## 退出容器

```
# 两种退出方式
# 1、run进去容器，exit退出，容器停止
exit 
# 2、run进去容器，ctrl+p+q退出，容器不停止
ctrl+p+q
```

## 启动已停止运行的容器

```
# 启动已停止运行的容器
docker start 容器ID或者容器名
# 重启容器
docker restart 容器ID或者容器名
# 停止容器
docker stop 容器ID或者容器名
# 强制停止容器
docker kill 容器ID或容器名
# 删除已停止的容器
docker rm 容器ID
# 一次性删除多个容器实例
docker rm -rf $(docker ps -a -q)

docker ps -a -q | xargs docker rm
```

## 查看容器日志

```
# 查看容器日志
docker logs 容器ID
```

## 查看容器内运行的进程

```
docker top 容器id
```

## 进入正在运行的容器并以命令行交互

```
docker exec -it 容器id bashshell
```

## 从容器内拷贝文件至主机

```
docker cp 容器id:容器内路径 目的主机路径
```

## 导入和导出容器

```
#导出容器内的内容留作为一个tar归档文件
docker export 容器id > 文件.tar
#从tar中的内容创建一个新的文件系统再导入为镜像
cat 文件名.tar|docker import -镜像用户/镜像名:镜像版本号
```

![40](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E5%9F%BA%E7%A1%80%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/40.png)

```
attach    Attach to a running container                 # 当前 shell 下 attach 连接指定运行镜像 
build     Build an image from a Dockerfile              # 通过 Dockerfile 定制镜像 
commit    Create a new image from a container changes   # 提交当前容器为新的镜像 
cp        Copy files/folders from the containers filesystem to the host path   #从容器中拷贝指定文件或者目录到宿主机中 
create    Create a new container                        # 创建一个新的容器，同 run，但不启动容器 
diff      Inspect changes on a container's filesystem   # 查看 docker 容器变化 
events    Get real time events from the server          # 从 docker 服务获取容器实时事件 
exec      Run a command in an existing container        # 在已存在的容器上运行命令 
export    Stream the contents of a container as a tar archive   # 导出容器的内容流作为一个 tar 归档文件[对应 import ] 
history   Show the history of an image                  # 展示一个镜像形成历史 
images    List images                                   # 列出系统当前镜像 
import    Create a new filesystem image from the contents of a tarball # 从tar包中的内容创建一个新的文件系统映像[对应export] 
info      Display system-wide information               # 显示系统相关信息 
inspect   Return low-level information on a container   # 查看容器详细信息 
kill      Kill a running container                      # kill 指定 docker 容器 
load      Load an image from a tar archive              # 从一个 tar 包中加载一个镜像[对应 save] 
login     Register or Login to the docker registry server    # 注册或者登陆一个 docker 源服务器 
logout    Log out from a Docker registry server          # 从当前 Docker registry 退出 
logs      Fetch the logs of a container                 # 输出当前容器日志信息 
port      Lookup the public-facing port which is NAT-ed to PRIVATE_PORT    # 查看映射端口对应的容器内部源端口 
pause     Pause all processes within a container        # 暂停容器 
ps        List containers                               # 列出容器列表 
pull      Pull an image or a repository from the docker registry server   # 从docker镜像源服务器拉取指定镜像或者库镜像 
push      Push an image or a repository to the docker registry server    # 推送指定镜像或者库镜像至docker源服务器 
restart   Restart a running container                   # 重启运行的容器 
rm        Remove one or more containers                 # 移除一个或者多个容器 
rmi       Remove one or more images       # 移除一个或多个镜像[无容器使用该镜像才可删除，否则需删除相关容器才可继续或 -f 强制删除] 
run       Run a command in a new container              # 创建一个新的容器并运行一个命令 
save      Save an image to a tar archive                # 保存一个镜像为一个 tar 包[对应 load] 
search    Search for an image on the Docker Hub         # 在 docker hub 中搜索镜像 
start     Start a stopped containers                    # 启动容器 
stop      Stop a running containers                     # 停止容器 
tag       Tag an image into a repository                # 给源中镜像打标签 
top       Lookup the running processes of a container   # 查看容器中运行的进程信息 
unpause   Unpause a paused container                    # 取消暂停容器 
version   Show the docker version information           # 查看 docker 版本号 
wait      Block until a container stops, then print its exit code   # 截取容器停止时的退出状态值 
```

# docker镜像

## docker镜像加载原理

docker的镜像实际上由一层一层的文件系统组成，这种层级的文件系统UnionFS。

bootfs(boot file system)主要包含bootloader和kernel, bootloader主要是引导加载kernel, Linux刚启动时会加载bootfs文件系统， 在Docker镜像的最底层是引导文件系统bootfs。 这一层与我们典型的Linux/Unix系统是一样的，包含boot加载器和内核。当boot加载完成之后整个内核就都在内存中了，此时内存的使用权已由bootfs转交给内核，此时系统也会卸载bootfs。

rootfs (root file system) ，在bootfs之上 。包含的就是典型 Linux 系统中的 /dev, /proc, /bin, /etc 等标准目录和文件。rootfs就是各种不同的操作系统发行版，比如Ubuntu，Centos等等。

![img](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E5%9F%BA%E7%A1%80%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/41.png)

> 对于一个精简的OS，rootfs可以很小，只需要包括最基本的命令、工具和程序库就可以了，因为底层直接用Host的kernel，自己只需要提供 rootfs 就行了。由此可见对于不同的linux发行版, bootfs基本是一致的, rootfs会有差别, 因此不同的发行版可以公用bootfs。

**为什么Docker镜像要采用这种分层结构呢**

> 镜像分层最大的一个好处就是共享资源，方便复制迁移，就是为了复用

#### 重点理解

> Docker镜像层都是只读的，容器层是可写的，当容器启动时，一个新的可写层被加载到镜像的顶部。这一层通常被称作"容器层"，"容器层"之下的都叫"镜像层"。
>
> 所有对容器的改动 - 无论添加、删除、还是修改文件都只会发生在容器层中。只有容器层是可写的，容器层下面的所有镜像层都是只读的。

![43](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E5%9F%BA%E7%A1%80%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/43.png)

## docker镜像生成

```
docker commit [OPTIONS]容器id [REPOSOTORY][:TAG]
OPTIONS说明
-a ：提交的镜像作者
-m：提交时的文字说明
```

![49](https://gitee.com/yooome/golang/raw/main/Docker%E8%AF%A6%E7%BB%86%E6%95%99%E7%A8%8B/Docker%E5%9F%BA%E7%A1%80%E8%AF%A6%E7%BB%86%E7%AC%94%E8%AE%B0/images/49.png)

## 本地镜像发布到私有库

```
#1、下载镜像Docker Registry
docker pull registry
#2、运行私有库，相当于本地有个私有库docker hub
docker run -d -p 5000:5000 -v /zzyyuse/myregistry/:/tmp/registry --privileged=true registry
默认情况，仓库被创建在容器的/var/lib/registry目录下，建议自行用容器卷映射，方便于宿主机联调
（/zzyyuse/myregistry/是宿主机绝对路径目录，/tmp/registry为容器内目录 ）

-v /zzyyuse/myregistry/:/tmp/registry --privileged=true registry ???
一句话：有点类似我们Redis里面的rdb和aof文件
将docker容器内的数据保存进宿主机的磁盘中
运行一个带有容器卷存储功能的容器实例
docker run -it --privileged=true -v /宿主机绝对路径目录:/容器内目录 镜像名
```

# docker容器数据卷

## 宿主vs容器之间映射添加容器卷

```
docker run -it --privileged=true -v /宿主机绝对路径目录:/容器内目录 镜像名
```

## 查看容器卷是否挂载成功

```
docker inspect 容器id
#如果挂载成功，则inspect命令下mount字段会有对应路径
```

## 容器卷ro和rw读写规则

在默认情况下，使用第一节中：宿主和容器之间映射添加容器卷的命令时，默认的读写方式是主机和容器内都可以进行读写操作，即默认为如下公式，只是省略了后面的冒号，即

```bash
docker run -it --privileged=true -v /宿主机绝对路径目录:/容器内目录:rw 镜像名
```

但如果想容器实例内部被限制，只能读取宿主机上的文件，但不能进行写操作，需要如下命令：

```
docker run -it --privileged=true -v /宿主机绝对路径：/容器内目录:ro 镜像名
```

即此时如果宿主机写入内容，可以同步给容器，容器可以读取到，但不能进行写操作

## 卷的继承与共享

按照上述操作u1完成了和宿主机的映射，而容器u2如想继承u1的卷规则，可以使用如下命令：

```
docker run -it --privileged=true --volumns-from=父类u1 --name=u2 镜像名
```

 通过上述操作可以实现主机和两个容器之间的数据共享，比如在u2新建一个文件，主机和u1均会进行同步。如果我们停掉了u2的容器，然后再宿主机上新建一个文件，则u1会直接同步，那当u2重启之后进入也能查看到该文件。本质上宿主机和多个容器能够共享所有文件，尽管某一容器停止，当其他容器和宿主机进行文件操作时，该容器重启进入后文件还会进行同步和更新。

# DockerFile

## RUN 指令

`RUN` 指令是在容器内执行 shell 命令，默认会是用 `/bin/sh -c` 的方式执行

### 执行命令的两种方式

- `RUN <command>`（*shell*形式，该命令在shell中运行）
- `RUN ["executable", "param1", "param2"]`（*exec*形式）

Dockerfile 中每一个指令都会建立一层，`RUN` 也不例外。每一个 `RUN` 的行为，就和刚才我们手工建立镜像的过程一样：新建立一层，在其上执行这些命令，执行结束后，`commit` 这一层的修改，构成新的镜像

所以，在使用 *shell* 方式，尽量多的使用续行符`\`

```text
RUN /bin/bash -c 'source $HOME/.bashrc; \
 echo $HOME'
```

当使用 `exec` 方式时，需要明确指定 `shell` 路径，否则变量可能不会生效

```text
FROM centos
ENV name="yangge"
RUN ["/bin/echo", "$name"]
```

![img](https://pic4.zhimg.com/v2-0ac5e51068d8b7fc7e9481d85dd2ccc7_r.jpg)

> 可以看到 `$name` 被作为普通的字符串输出了，因为 `$name` 是 shell 中的用法，而这里里并没有 使用到 shell

**下面是正确的做法**

```text
FROM alpine
ENV name="yangge"
RUN ["/bin/sh", "-c", "/bin/echo $name"]
```

> **注意:** *exec*的方式下，列表中的内容会被解析为JSON数组，这意味着您必须在单词周围使用双引号（“） 而非单引号（'）。

![img](https://pic3.zhimg.com/v2-a58f85bedd99d350a83932096ddbf3da_r.jpg)

## CMD

`Dockerfile` 中只能有一条`CMD`指令。如果列出多个，`CMD` 则只有最后一个`CMD`会生效。

**CMD 主要目的是为运行容器时提供默认值**

指定默认的容器主进程的启动命令。在启动(运行)一个容器时可以指定新的命令来替代镜像设置中的这个默认命令

- `shell` 格式：`CMD <命令>`
- `exec` 格式：`CMD ["可执行文件", "参数1", "参数2"...]`
- 参数列表格式：`CMD ["参数1", "参数2"...]`。在指定了 `ENTRYPOINT` 指令后，用 `CMD` 指定具体的参数。

> **注意**：不要混淆`RUN` 和 `CMD`。`RUN`实际上运行一个命令并提交结果; `CMD`在构建时不执行任何操作，但指定镜像的默认命令。

Docker 不是虚拟机，容器内没有后台服务的概念。

不要期望这样启动一个程序到后台:

```text
CMD systemctl start nginx
```

这行被 `Docker` 理解为：

```text
CMD ["sh" "-c" "systemctl start nginx"]
```

对于容器而言，其启动程序就是容器的应用进程，容器就是为了主进程而存在的，主进程退出，容器就失去了存在的意义，从而退出，其它辅助进程不是它需要关心的东西。

就像上面的示例中，主进程是 `sh` , 那么当 `service nginx start` 命令结束后，`sh` 也就结束了，`sh` 作为主进程退出了，自然就会使容器退出。

正确的做法是直接执行 `nginx` 这个可执行文件，并且关闭后台守护的方式，使程序在前台运行。

```text
CMD ["nginx", "-g", "daemon off;"]
```

## ENTRYPOINT

和 `CMD` 一样，都是在指定容器的启动程序及参数。

`ENTRYPOINT` 在运行时也可以被替代，不过比 `CMD` 要略显繁琐，需要通过 `docker run` 的参数 `--entrypoint` 来指定。

`ENTRYPOINT` 的格式和 `RUN` 指令格式一样，也分为 `exec` 格式和 `shell` 格式。

当指定了 `ENTRYPOINT` 后，`CMD` 的含义就发生了改变，不再是直接的运行其命令，而是将 `CMD` 的内容作为参数传给 `ENTRYPOINT` 指令，也就是实际执行时，将变为：

```bash
<ENTRYPOINT> "<CMD>"
```

有了 `CMD` 后，为什么还要有 `ENTRYPOINT` 呢？

这种 `<ENTRYPOINT> "<CMD>"` 给我们带来了什么好处么？

让我们来看几个场景。

- 场景一：让镜像变成像命令一样使用

**CMD 方式**

```text
FROM centos
RUN yum update \
    && yum install -y curl
CMD [ "curl", "-s", "http://ip.cn" ]
```

构建镜像后, 运行容器

```text
# docker run --rm centos-echo-ip-cmd
```

执行下面命令会报错

```text
# docker run --rm centos-echo-ip-cmd -i
```

我们可以看到报错，`executable file not found`。之前我们说过，跟在镜像名后面的是 `command`，运行时会替换 `CMD` 的默认值。因此这里的 `-i` 并不是添加在原来的 `curl -s http://ip.cn` 后面。 而是替换了原来的 `CMD`，变成了 `CMD ["-i"]`，而 `-i` 根本不是命令，所以报了`可执行文件找不到`。

**所以应该使用 `ENTRYPOINT` 方式**

```text
FROM centos
RUN yum install -y curl
ENTRYPOINT ["curl", "-s", "http://ip.cn"]
```

再次构建镜像后, 运行容器

```text
# docker run --rm centos-echo-ip-entrypoint
# docker run --rm centos-echo-ip-entrypoint -i
```

这样的话, 最终的指令就变成 `ENTRYPOINT ["curl", "-s", "http://ip.cn", "-i"]`

- 场景二：应用运行前的准备工作

启动容器就是启动主进程，但有些时候，启动主进程前，需要一些准备工作。

官方镜像 `redis` 中的示例：

```text
FROM alpine:3.4
...
RUN addgroup -S redis && adduser -S -G redis redis
...
ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 6379
CMD [ "redis-server" ]
```

可以看到其中为 redis 服务创建了 redis 用户，并在最后指定了 `ENTRYPOINT` 为 `docker-entrypoint.sh` 脚本。

```text
#!/bin/sh
...
# allow the container to be started with `--user`
if [ "$1" = 'redis-server' -a "$(id -u)" = '0' ]; then
    chown -R redis .
    exec gosu redis "$0" "$@"
fi

exec "$@"
```

该脚本的内容就是根据 `CMD` 的内容来判断，如果是 `redis-server` 的话，则切换到 `redis` 用户身份启动服务器，否则依旧使用 `root` 身份执行。比如：

```text
$ docker run -it redis id
uid=0(root) gid=0(root) groups=0(root)
```

还有 `ENTRYPOINT` 指令不会被 `RUN` 指令覆盖，而 `CMD` 指令会被 `RUN` 指令覆盖

## RUN&&CMD

> RUN是构件容器时就运行的命令以及提交运行结果 CMD是容器启动时执行的命令，在构件时并不运行，构件时紧紧指定了这个命令到底是个什么样子

## ADD

`COPY` 和 `ADD` 指令中选择的原则，所有的文件复制均使用 `COPY` 指令，仅在需要自动解压缩的场合使用 `ADD`。

支持自动解压缩，压缩格式支持： `gzip`, `bzip2` 以及 `xz`

## USER

`USER` 则是改变执行 `RUN`, `CMD` 以及 `ENTRYPOINT` 这类命令的身份。

这个用户必须是事先在容器内存在(建立好)的，否则无法切换。

如果以 `root` 执行的脚本，在执行期间希望改变身份，比如希望以某个已经建立好的用户来运行某个服务进程，不要使用 `su` 或者 `sudo`，这些都需要比较麻烦的配置，而且在 TTY 缺失的环境下经常出错。建议使用 `gosu`。

```text
# 建立 redis 用户，并使用 gosu 换另一个用户执行命令
RUN groupadd -r redis && useradd -r -g redis redis
# 下载 gosu
RUN wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/1.7/gosu-amd64" \
    && chmod +x /usr/local/bin/gosu \
    && gosu nobody true
# 设置 CMD，并以另外的用户执行
CMD [ "exec", "gosu", "redis", "redis-server" ]
```

## HEALTHCHECK 健康检查指令

语法：

```
 HEALTHCHECK [OPTIONS] CMD command
 HEALTHCHECK NONE
```

第一个的功能是在容器内部运行一个命令来检查容器的健康状况

第二个的功能是在基础镜像中取消健康检查命令

[OPTIONS]的选项支持以下三中选项：

–interval=DURATION 两次检查默认的时间间隔为30秒

–timeout=DURATION 健康检查命令运行超时时长，默认30秒

–retries=N 当连续失败指定次数后，则容器被认为是不健康的，状态为unhealthy，默认次数是3

注意：

HEALTHCHECK命令只能出现一次，如果出现了多次，只有最后一个生效。

CMD后边的命令的返回值决定了本次健康检查是否成功，具体的返回值如下：

0: success - 表示容器是健康的

1: unhealthy - 表示容器已经不能工作了

2: reserved - 保留值

例子：

```
HEALTHCHECK --interval=5m --timeout=3s \
CMD curl -f http://localhost/ || exit 1
```

健康检查命令是：

```
curl -f http://localhost/ || exit 1
```

两次检查的间隔时间是5秒

命令超时时间为3秒

## ONBUILD

语法：

```
ONBUILD [INSTRUCTION]
```

这个命令只对当前镜像的子镜像生效。

比如当前镜像为A，在Dockerfile种添加：

```
ONBUILD RUN ls -al
```

这个 ls -al 命令不会在A镜像构建或启动的时候执行
此时有一个镜像B是基于A镜像构建的，那么这个ls -al 命令会在B镜像构建的时候被执行

## ARG

语法：

```
ARG <name>[=<default value>]
```

设置变量命令，ARG命令定义了一个变量，在docker build创建镜像的时候，使用 --build-arg =来指定参数

如果用户在build镜像时指定了一个参数没有定义在Dockerfile种，那么将有一个Warning

提示如下：

[Warning] One or more build-args [foo] were not consumed.

我们可以定义一个或多个参数，如下：

```
FROM busybox
ARG user1
ARG buildno
```

也可以给参数一个默认值：

```
FROM busybox
ARG user1=someuser
ARG buildno=1
```

# Docker-compose

> Docker-Compose是Docker官方的开源项目，负责实现对Docker容器集群的快速编排。
>
> docker建议我们每一个容器中只运行一个服务,因为docker容器本身占用资源极少,所以最好是将每个服务单独的分割开来但是这样我们又面临了一个问题？ 
>
> 如果我需要同时部署好多个服务,难道要每个服务单独写Dockerfile然后在构建镜像,构建容器,这样累都累死了,所以docker官方给我们提供了docker-compose多服务部署的工具 
>
> 例如要实现一个Web微服务项目，除了Web服务容器本身，往往还需要再加上后端的数据库mysql服务容器，redis服务器，注册中心eureka，甚至还包括负载均衡容器等等。。。。。。 
>
> Compose允许用户通过一个单独的 docker-compose.yml模板文件 （YAML 格式）来定义 一组相关联的应用容器为一个项目（project）。 
>
> 可以很容易地用一个配置文件定义一个多容器的应用，然后使用一条指令安装这个应用的所有依赖，完成构建。Docker-Compose 解决了容器与容器之间如何管理编排的问题。

## 两要素

1. 服务(service):
   一个个应用容器实例，比如订单微服务、库存微服务、mysql容器、nginx容器或者redis容器。
2. 工程(project):
   由一组关联的应用容器组成的一个完整业务单元，在 docker-compose.yml 文件中定义。

## 三步骤

1. 编写Dockerfile定义各个微服务应用并构建出对应的镜像文件
2. 使用 docker-compose.yml 定义一个完整业务单元，安排好整体应用中的各个容器服务
3. 执行docker-compose up命令 来启动并运行整个应用程序，完成一键部署上线

## 常用命令

```
Compose 常用命令 
docker-compose -h                           #  查看帮助 
docker-compose up                           #  启动所有 docker-compose服务 
docker-compose up -d                        #  启动所有 docker-compose服务 并后台运行 
docker-compose down                         #  停止并删除容器、网络、卷、镜像。 
docker-compose exec  yml里面的服务id                 # 进入容器实例内部  docker-compose exec  docker-compose.yml文件中写的服务id  /bin/bash 
docker-compose ps                      # 展示当前docker-compose编排过的运行的所有容器 
docker-compose top                     # 展示当前docker-compose编排过的容器进程 
 
docker-compose logs  yml里面的服务id     #  查看容器输出日志 
docker-compose config     #  检查配置 
docker-compose config -q  #  检查配置，有问题才有输出 
docker-compose restart   #  重启服务 
docker-compose start     #  启动服务 
docker-compose stop      #  停止服务 
```