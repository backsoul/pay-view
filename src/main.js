const express = require('express')
const app = express()
const port = 3000

const browser = require('/app/src/browser.js');

app.get('/bot', (req, res) => {
    console.log(req.query)
    const userTwitch = req.query.user;
    const limit = req.query.limit;
    if(userTwitch && limit){
        console.log(`initialize user: ${userTwitch} limit: ${limit}`)
        browser(userTwitch, limit);
        res.send("BOT SEND")
    } else {
        res.send("params missing user or limit")
    }
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
