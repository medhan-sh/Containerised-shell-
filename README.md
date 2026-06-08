# Containerised Shell:

A container runtime built from scratch in Go.

Implements Linux namespaces (UTS, PID, MNT, NET), cgroups, 
chroot, and an Alpine Linux rootfs — no Docker, no libraries.

There's no particular reason for anyone to use this neither anyone is gonna use this coz making a container and making it run a shell will do more than this but this looks cool and is a learning project plus i got way too much time so why not.

## Run it

git clone https://github.com/medhan-sh/mycontainer
cd mycontainer
make all
make run

Requires Linux and sudo. Genuinely does not work anywhere else🥲

## How it works

- `run` — forks a child process with isolated namespaces
- `child` — sets up chroot into Alpine rootfs, mounts /proc, applies cgroup limits, execs the shell


