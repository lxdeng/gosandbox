Default namespaces:

root@Ubu1404:/proc/1/ns# ls -l
total 0
lrwxrwxrwx 1 root root 0 Aug 18 11:16 ipc -> ipc:[4026531839]
lrwxrwxrwx 1 root root 0 Aug 18 11:16 mnt -> mnt:[4026531840]
lrwxrwxrwx 1 root root 0 Aug 18 11:16 net -> net:[4026531957]
lrwxrwxrwx 1 root root 0 Aug 18 11:16 pid -> pid:[4026531836]
lrwxrwxrwx 1 root root 0 Aug 18 11:16 user -> user:[4026531837]
lrwxrwxrwx 1 root root 0 Aug 18 11:16 uts -> uts:[4026531838]

Start to run:
----------------
$ sudo ./container run ls
[sudo] password for lxdeng: 
starting child: ....


When the "child" process running,

pid, mnt, uts now is of different namespaces
root@Ubu1404:/proc/3197/ns# ls -l
total 0

lrwxrwxrwx 1 root root 0 Aug 18 11:24 ipc -> ipc:[4026531839]
lrwxrwxrwx 1 root root 0 Aug 18 11:24 mnt -> mnt:[4026532198]
lrwxrwxrwx 1 root root 0 Aug 18 11:24 net -> net:[4026531957]
lrwxrwxrwx 1 root root 0 Aug 18 11:24 pid -> pid:[4026532201]
lrwxrwxrwx 1 root root 0 Aug 18 11:24 user -> user:[4026531837]
lrwxrwxrwx 1 root root 0 Aug 18 11:24 uts -> uts:[4026532199]

When the "grandchild" process running,

root@Ubu1404:/proc/3278/ns# ls -l
total 0
lrwxrwxrwx 1 root root 0 Aug 18 11:26 ipc -> ipc:[4026531839]
lrwxrwxrwx 1 root root 0 Aug 18 11:26 mnt -> mnt:[4026532198]
lrwxrwxrwx 1 root root 0 Aug 18 11:26 net -> net:[4026531957]
lrwxrwxrwx 1 root root 0 Aug 18 11:26 pid -> pid:[4026532201]
lrwxrwxrwx 1 root root 0 Aug 18 11:26 user -> user:[4026531837]
lrwxrwxrwx 1 root root 0 Aug 18 11:26 uts -> uts:[4026532199]

As it shows, both grandchild and child are of the same namespaces.
