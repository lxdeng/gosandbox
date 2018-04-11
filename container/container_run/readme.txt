lxdeng@lu1:~/dev/go/src/github.com/lxdeng/gosandbox/container/container_run$ sudo ./crun run bash
Running child as 31994
Running [bash] as 1
root@lu1:~/dev/go/src/github.com/lxdeng/gosandbox/container/container_run# ps 
  PID TTY          TIME CMD
31993 pts/0    00:00:00 sudo
31994 pts/0    00:00:00 crun
31997 pts/0    00:00:00 exe
32000 pts/0    00:00:00 bash
32010 pts/0    00:00:00 ps

ps command inside the container does not show the PID 1. Because the command impl uses /proc file system to get its data.

Inside the container, the hostname command can change the hostname for the container.


get a root fs:
docker run -it busybox sh
docker export 15634c92fb3e > ~/contents.tar
cd /home/lxdeng/rt_fs
tar xvf ~/contents.tar

before run it, make sure /home/lxdeng/rt_fs/proc is not mounted from the last run.

$ sudo umount /home/lxdeng/rt_fs/proc

sudo ./crunfs run sh

# umount /proc
# ls proc

need to remount proc ???
# mount -t proc none /proc
# ps
PID   USER     TIME  COMMAND
    1 root      0:00 /proc/self/exe child sh
    4 root      0:00 sh
   12 root      0:00 ps
/ # 
