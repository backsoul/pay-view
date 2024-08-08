const getProxies = require('/app/src/getProxies');
const { Worker } = require('worker_threads');

function runWorker(proxy, url) {
  return new Promise((resolve, reject) => {
    const worker = new Worker('/app/src/worker.js');
    const obj = {proxy, url};
    worker.postMessage(obj);
    worker.on('message', (message) => {
      resolve(message);
    });
    worker.on('error', (error) => {
      reject(error);
    });
    worker.on('exit', (code) => {
      if (code !== 0) {
        reject(new Error(`Worker stopped with exit code ${code}`));
      }
    });
  });
}

const browser = async (user, limit) => {
  // Obtener la lista de proxies gratuitos
  const proxies = await getProxies();
  if (proxies.length === 0) {
    console.error('No se encontraron proxies.');
    return;
  }
  const proxiesShort = [];
  for (let index = 0; index < limit; index++) {
    const proxy = proxies[index];
    proxiesShort.push(proxy)
  }
  const promises = proxiesShort.map(proxy => runWorker(proxy, user));
  const results = await Promise.all(promises);
  console.log('All workers completed:', results);
}

module.exports = browser;