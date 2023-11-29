import repl from "node:repl";

console.log("Booting JsOS...");

const init = repl.start("$ ");

init.on("exit", () => {
	console.log("Exited repl, rebooting...");
	process.exit();
});
