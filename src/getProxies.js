// getProxies.js
const axios = require('axios');
const cheerio = require('cheerio');

const getProxies = async () => {
  try {
    const response = await axios.get('https://www.sslproxies.org/');
    const $ = cheerio.load(response.data);
    const proxies = [];

    $('.table tbody tr').each((index, element) => {
      const ip = $(element).find('td:nth-child(1)').text().trim();
      const port = $(element).find('td:nth-child(2)').text().trim();
      if (ip && port) {
        proxies.push(`http://${ip}:${port}`);
      }
    });

    return proxies;
  } catch (error) {
    console.error('Error al obtener proxies:', error);
    return [];
  }
};

module.exports = getProxies;
