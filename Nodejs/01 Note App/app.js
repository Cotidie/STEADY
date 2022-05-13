import validator from 'validator';

const add = require('./notes-app/utils.js')
const getNotes = require('./notes-app/notes.js')

const msg = getNotes()
console.log(msg)

const email = 'wsj@gmai'
console.log(validator.isEmail(email))