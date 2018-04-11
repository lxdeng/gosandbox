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
