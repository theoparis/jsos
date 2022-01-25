const { exec } = require("std/sys");

const output = exec("/bin/echo", "Hello world");
console.log(output);
