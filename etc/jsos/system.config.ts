export default {
	init: {
		mode: "repl"
	},
	kernel: {
		moduleCompression: "zstd",
		initramfs: {
			compression: "zstd"
		},
		filesystems: [
			{
				id: "btrfs",
				acl: true
			},
			{
				id: "fuse"
			},
			{
				id: "fat"
			},
			{
				id: "tmpfs",
				acl: true,
				xattr: true
			},
			{
				id: "efivar"
			},
			{
				id: "nfsv4",
				swap: true,
				root: true,
				server: true
			}
		]
	},
};
