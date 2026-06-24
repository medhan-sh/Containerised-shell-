# containerized-shell

A container runtime built from scratch in Go

Implements the core Linux primitives that container runtimes are actually built on: namespaces, cgroups, and chroot. Runs an isolated Alpine Linux shell with its own process tree, filesystem, and network stack.

## What it does

- Forks a child process into isolated **UTS, PID, MNT, and NET namespaces**
- Sets up **chroot** into an Alpine Linux rootfs
- Mounts `/proc` inside the container
- Drops you into a real shell — fully isolated from the host

## Why this is interesting

Docker at its core utilizes Linux Namespaces so tried tinkering a little with it.
The Shell is still pretty imature but i'll add features which seem interesting enough to me, right now its fetching from your current FS which kinda defeats the whole point of having a custom shell.

## Run it

```bash
git clone https://github.com/medhan-sh/Containerised-shell-
cd Containerised-shell-
make all
make run
```

> Requires Linux and `sudo`. Does not work on macOS or Windows — that's not a bug, that's the point.

## How it works

```
make run
  └── run cmd
        └── clone() with CLONE_NEWUTS | CLONE_NEWPID | CLONE_NEWNS | CLONE_NEWNET
              └── child cmd
                    ├── chroot → Alpine rootfs
                    ├── mount /proc
                    ├── apply cgroup limits
                    └── exec /bin/sh
```

| Step | Syscall / Mechanism | What it isolates |
|------|-------------------|-----------------|
| Namespace fork | `clone()` with flags | PID tree, hostname, mounts, network |
| Filesystem isolation | `chroot()` | Root filesystem |
| Process visibility | `mount /proc` | Only sees its own PIDs |
| Resource limits | cgroup v2 (`/sys/fs/cgroup`) | CPU and memory |

## Tech

- **Go** — `os/exec`, `syscall` package for namespace flags and chroot
- **Linux namespaces** — UTS, PID, MNT, NET via `CLONE_NEW*` flags
- **cgroups v2** — resource limiting via `/sys/fs/cgroup`
- **Alpine Linux rootfs** — minimal filesystem for the container environment

## What's next

- [ ] Add bubbletea TUI for managing running containers
- [ ] Network namespace with veth pair + NAT
- [ ] Image layer support (overlay filesystem)
- [ ] Resource usage dashboard

## References

- [Liz Rice — Containers from Scratch (GopherCon 2016)](https://www.youtube.com/watch?v=8fi7uSYlOdc)
- Linux `man 7 namespaces`, `man 7 cgroups`
- _Container Security_ by Liz Rice
