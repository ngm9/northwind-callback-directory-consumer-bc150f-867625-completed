const express = require('express');

function createRouter(service) {
  const router = express.Router();

  router.get('/phone-numbers', async (req, res) => {
    const { country, phone } = req.query;
    const resolved = await service.getPhoneNumbers(country, phone);
    res.send(resolved);
  });

  return router;
}

module.exports = { createRouter };
