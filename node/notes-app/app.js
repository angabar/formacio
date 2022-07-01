import chalk from "chalk";
import yargs from "yargs";
import { hideBin } from "yargs/helpers";

yargs(hideBin(process.argv))
    .command({
        command: "add",
        aliases: ["a"],
        desc: "Add a new note",
        handler: (argv) => {
            console.log("adding a note");
        },
    })
    .command({
        command: "remove",
        aliases: ["rm"],
        desc: "Remove a note",
        handler: (argv) => {
            console.log("removing a note");
        },
    })
    .command({
        command: "list",
        aliases: ["l"],
        desc: "List all notes",
        handler: (argv) => {
            console.log("Listing notes");
        },
    })
    .command({
        command: "read",
        aliases: ["r"],
        desc: "Read a note",
        handler: (argv) => {
            console.log("Reading a note");
        },
    })
    .demandCommand()
    .recommendCommands()
    .strict().argv;
