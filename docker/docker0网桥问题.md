# docker0网桥问题

在平时开发测试中经常需要连接各种各样的环境，环境的ip也都是千奇百怪，使用虚拟机的时候需要注意docker0网桥网段冲突问题。

一般的默认的docker0网桥为`172.17`网段，所以当连接一个其他环境同网段ip时往往会出现冲突导致网络不通。

## 修改docker0网桥

```
vim /etc/docker/daemon.json

{
	"bip":"192.178.1.100/24"    //随便写个不冲突的
}

systemctl restart docker
```

或者直接down调这个网桥(临时的)

```
1.ifconfig docker0 down
2.ifdown docker0
3.ip link set docker0 down
```

三种方式都可以暂时的关闭这个网桥

如果要删除网桥的话

```
brctl delbr docker0
```

brctl需要安装

```
yum install -y bridge-utils 
```



