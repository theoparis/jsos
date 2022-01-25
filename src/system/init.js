const { mount } = require("std/sys");

console.log("Booting up...");
mount("proc", "proc", "/proc");
mount("sysfs", "sysfs", "/sys");
mount("tmpfs", "tmpfs", "/tmp");

setInterval(() => {
    console.log("Hello world");
}, 5000);
