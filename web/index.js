const http = require('http')

async function fetch(url) {
    return new Promise((resolve, reject) => {
        http.get(url, res => {
            let data = ''
        
            res.on('data', (chunk) => {
                data += chunk
            })
        
            res.on('end', () => {
                resolve(JSON.parse(data))
            })
        }).on('error', err => {
            reject(err)
        })
    })
}

http.createServer(async function(req, res){
    const films = await fetch('http://films-api')
    const content = films
        .map(({ title, year, director }) => 
            `${title} (${year}) - Directed by ${director.name}`)
        .reduce((list, str) => `${list}<li>${str}</li>`, '')

    res.writeHead(200, { 'Content-Type': 'text/html' })
    res.write(`<html><body><h1>Films:</h1><ul>${content}</ul></body></html>`)
    res.end()
  }).listen(80)
