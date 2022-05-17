import fs from 'fs'

export const config = {
    weatherStack: {
        endpoint: "",
        apiKey: "",
        location: {
            lat: "",
            lng: ""
        }
    }
}

function parseConfig() {
    const raw = readJSON('config.json')
    config.weatherStack = raw.weatherStack
}

function readJSON(path) {
    const buffer = fs.readFileSync(path)
    const jsonStirng = buffer.toString()
    const object = JSON.parse(jsonStirng)
    return object
}

parseConfig()