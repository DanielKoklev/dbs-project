#!/bin/bash
set -e
qemu-system-x86_64 \
  -m 1024 -enable-kvm \
  -drive file=nbd:localhost:10809,if=virtio,cache=none,aio=native,format=raw