#!/bin/bash
set -e
fio --name=4k-randwrite --rw=randwrite --bs=4k --iodepth=32 --numjobs=1 \
    --filename=nbd:localhost:10809 --direct=1 --time_based --runtime=30