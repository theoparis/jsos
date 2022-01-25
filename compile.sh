#!/bin/bash
set -eo pipefail

echo_info() {
    echo -e "\033[1;32m[INFO]\033[0m $1"
}

# Variables
export ARCH=${ARCH:-x86_64}

echo_info "Compiling go files..."

go build -ldflags "-linkmode external -extldflags -static" -o build/jsos-init -a cmd/init/init.go

mkdir -p rootfs/{bin,lib,tmp,dev}
mkdir -p rootfs/lib/system || true

echo_info "Building the rootfs..."

cp build/* rootfs/bin/
cp -r src/* rootfs/lib/

cd rootfs
mkdir dev || true
sudo mknod dev/ram b 1 0 || true
sudo mknod dev/console c 5 1 || true
sudo mknod dev/ttyS0 c 4 64 || true
sudo mknod dev/tty0 c 4 0 || true
find . | cpio -o -H newc | gzip -9 >../initramfs.cpio.gz
cd ..

echo_info "Setting up the kernel sources..."
# if linux is not cloned, clone it
if [ ! -d "linux" ]; then
    git clone https://github.com/torvalds/linux --depth 1 --single-branch
fi

cd linux

echo_info "Creating the kernel config..."

make defconfig

# merge the config with ../kernel.config
if [ -f ../kernel.config ]; then
    echo_info "Merging the kernel config..."
    make oldconfig
    cat ../kernel.config >>.config
fi

echo_info "Building the linux kernel..."
make -j"$(nproc)"
cd ..

# gcc -static -o rootfs/init ../src/init.c

# grub-mkrescue -o timos.iso ./initrd

echo_info "Done!"
