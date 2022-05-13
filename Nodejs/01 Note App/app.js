import validator from 'validator';

const add = require('./libraries/utils.js')
const getNotes = require('./libraries/notes.js')

const msg = getNotes()
console.log(msg)

const email = 'wsj@gmai'
console.log(validator.isEmail(email))