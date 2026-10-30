const axios = require('axios');

function createCountriesClient(config) {
  return {
    async fetchCountry(country) {
      const url = 'https://jsonmock.hackerrank.com/api/countries?name=' + country;
      const response = await axios.get(url);
      return response.data;
    },
  };
}

module.exports = { createCountriesClient };
