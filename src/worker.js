const { parentPort } = require('worker_threads');
const fs = require('fs');
const { initializeBrowser } = require('/app/src/utils/utils');

parentPort.on('message', async (data) => {
  const {proxy, url } = data;
  const dir = '/app/src/images/' + proxy.split("//")[1];
  try {
    console.log("initialize proxy:", proxy);
    const { browser, page } = await initializeBrowser(proxy);
    
    console.log("goto proxy:", proxy);
    await page.goto(url, { waitUntil: "domcontentloaded" });

    if (!fs.existsSync(dir)){
      fs.mkdirSync(dir);
    }
    
    await page.screenshot({
      path: dir + '/twitch0.jpg',
      type: 'jpeg',
      quality: 70
    });
    console.log("entry page and sleep 90s proxy:", proxy);
    await new Promise(resolve => setTimeout(resolve, 90000));
    await page.screenshot({
      path: dir + '/twitch90.jpg',
      type: 'jpeg',
      quality: 70
    });
    console.log("finish view proxy:", proxy);
    await browser.close();
    console.log("success open browser: ", proxy);
    parentPort.postMessage({ success: true });
  } catch (e) {
    const errorMessage = e.toString();
    if (errorMessage.includes("ms exceeded")) {
      console.log(`error try open with ip ${proxy}, error: ip ms exceeded`);
      parentPort.postMessage({ success: false, error: "ip ms exceeded" });
    } else {
      parentPort.postMessage({ success: false, error: e });
    }
  }
});
