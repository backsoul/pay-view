const express = require('express')
const app = express()
const port = 3000

const browser = require('/app/src/browser.js');

app.get('/bot', (req, res) => {
    console.log(req.query)
    const platform = req.query.platform;
    const id = req.query.id;
    const limit = req.query.limit;
    switch (platform) {
        case "yt":
            url = "https://www.youtube.com/watch?v=" + id;
            break;
        case "twitch":
            url = "https://www.twitch.tv/" + id;
            break;
        default:
            break;
    }
    
    if(url && limit){
        console.log(`initialize id: ${url} limit: ${limit}`)
        browser(url, limit);
        res.send("BOT SEND")
    } else {
        res.send("params missing user or limit")
    }
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
