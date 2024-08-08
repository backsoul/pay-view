const puppeteer = require('puppeteer');
const initializeBrowser = async (proxy) => {
    const browser = await puppeteer.launch({
        headless: true, 
        args: [`--proxy-server=${proxy}`],
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