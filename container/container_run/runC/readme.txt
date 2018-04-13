https://github.com/opencontainers/runc

1. runc spec
update the config.json to add CAP_SYS_ADMIN


2. mkdir rootfs
docker export $(docker create busybox) | tar -C rootfs -xvf -

3. chown root:root rootfs

4. sudo runc run mycontainerid

5. mkdir /ramroot

6. mount -n -t tmpfs -o size=500M none /ramroot

7. df -k 

8. on the host df -k

The host and the container have different mounts
