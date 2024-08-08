const puppeteer = require('puppeteer');
const initializeBrowser = async (proxy) => {
    const browser = await puppeteer.launch({
        headless: true, 
        args: ['--disable-gpu', '--disable-setuid-sandbox', '--no-sandbox', '--no-zygote',`--proxy-server=${proxy}`],
        ignoreHTTPSErrors: true,
        dumpio: false
      });
    const page = await browser.newPage();
    await page.setRequestInterception(true);
    page.on('request', request => {
        request.continue({
            proxy: proxy
        });
    });
    return {browser, page};
}

module.exports = {
    initializeBrowser
};