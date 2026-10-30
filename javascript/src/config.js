function loadConfig() {
  return {
    upstreamBaseUrl: 'https://jsonmock.hackerrank.com/api/countries',
    requestTimeoutMs: 5000,
  };
}

module.exports = { loadConfig };
