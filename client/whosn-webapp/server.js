const express = require('express')
// eslint-disable-next-line no-unused-vars
// const bodyParser = require('body-parser');
const path = require('path')
const app = express()
const port = process.env.PORT || 3000

app.use(express.static(path.join(__dirname, 'build')))

// This route serves the React app
if (process.env.ENVIRONMENT === 'production') {
    app.get('/*', (req, res) =>
        res.sendFile(path.join(__dirname, 'build', 'index.html'))
    )
} else {
    app.get('/', (req, res) =>
        res.sendFile(path.resolve(__dirname, 'build', 'index.html'))
    )
}

app.listen(port, () => console.log(`Server listening on port ${port}`))
