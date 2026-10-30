const express = require('express');
const { createRouter } = require('./routes');
const { createDirectoryService } = require('./service');
const { createCountriesClient } = require('./client');
const { loadConfig } = require('./config');

function createApp(service) {
  const config = loadConfig();
  const resolvedService = service || createDirectoryService(createCountriesClient(config));
  const app = express();
  app.use(createRouter(resolvedService));
  return app;
}

module.exports = { createApp };
