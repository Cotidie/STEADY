const yargs = require('yargs')

// Customize yargs version, help document
yargs.version('1.1.0')

// 명령어 추가
yargs.command({
    command: "add",
    describe: "Add a new note",
    builder: {
        title: {
            describe: "Set a note's title",
            demandOption: true,
            type: 'string'
        },
        body: {
            describe: "Set a note's body",
            demandOption: false,
            type: 'string'
        }
    },
    handler: function(argv) {
        console.log("Adding a new note!")
        console.log("title: ", argv.title)
        console.log("body: ", argv.body)
    }
})

yargs.command({
    command: 'remove',
    describe: "Remove a note",
    handler: function() {
        console.log("Removing the note")
    }
})

yargs.command({
    command: 'list',
    describe: "List notes",
    handler: function() {
        console.log("Listing notes")
    }
})

yargs.command({
    command: 'read',
    describe: "Read a note",
    handler: function() {
        console.log("Reading the note")
    }
})

yargs.parse()