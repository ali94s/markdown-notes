# 设置虚拟机与windows的文件共享

windows下开发，在虚拟机的centos环境中测试，smaba是一种实现方式，也可以使用vmware的共享文件夹。

## 配置共享文件步骤

1. 在虚拟机设置中打开共享文件夹，选择需要共享的目录。

2. 在虚拟机中使用` vmware-hgfsclient `命令查看共享目录情况

   ```
   [root@localhost ~]# vmware-hgfsclient
   go-code
   ```

   

3. 这个时候还没有出现共享目录，需要挂载

   ```
   mkdir /root/src
   vmhgfs-fuse .host:/go-code /root/src
   
   [root@localhost ~]# pwd
   /root
   [root@localhost ~]# ls
   anaconda-ks.cfg  go-code
   ```

4. 在fstab文件中记录挂载信息，每次开启都挂载

   ```
   vi /etc/fstab
   # 添加挂载信息
   .host:/go-code          /root/go-code           fuse.vmhgfs-fuse allow_other,defaults 0 0
   ```

   