#!/bin/bash
set -eo pipefail

#  -nographic -append "console=ttyS0 debug"

qemu-system-x86_64 --enable-kvm \
    -kernel linux/arch/x86/boot/bzImage -initrd initramfs.cpio.gz \
    -m 1024 -net nic,model=virtio -net user \
    -serial stdio \
    -display none \
    -append "root=/dev/ram console=ttyS0 GOROOT=/lib/goroot rdinit=/bin/jsos-init -- /lib/system/init.js"
# -append "root=/dev/ram console=ttyS0 rdinit=/bin/node -- --napi-modules /lib/system/init.mjs" \
