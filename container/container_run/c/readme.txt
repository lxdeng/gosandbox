https://blog.yadutaf.fr/2014/01/12/introduction-to-linux-namespaces-part-4-ns-fs/

CLONE_NEWNS cannot achieve the isolation of /proc as the above blog. Tested on Lubuntu 16.04. #mount -t proc proc /proc still messed up the default/parent namespace.
