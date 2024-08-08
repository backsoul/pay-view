const getProxies = require('./getProxies');
const { Worker } = require('worker_threads');

function runWorker(proxy) {
  return new Promise((resolve, reject) => {
    const worker = new Worker('./worker.js');
    worker.postMessage(proxy);
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

const browser = async () => {
  // Obtener la lista de proxies gratuitos
  const proxies = await getProxies();
  if (proxies.length === 0) {
    console.error('No se encontraron proxies.');
    return;
  }
  const proxiesShort = [];
  for (let index = 0; index < 10; index++) {
    const proxy = proxies[index];
    proxiesShort.push(proxy)
  }
  const promises = proxiesShort.map(proxy => runWorker(proxy));
  const results = await Promise.all(promises);
  console.log('All workers completed:', results);
}

module.exports = browser;