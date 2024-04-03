const http = require('http')
const port = 8000

const server = http.createServer(function (request, response) {
    response.end(`hello, node!`)
});

server.listen(port, (err) => {
    if (err) {
        return console.log('something bad happened', err)
    }

    console.log(`server is listening on ${port}`)
})
