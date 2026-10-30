const request = require('supertest');
const { createApp } = require('../src/app');
const { createDirectoryService } = require('../src/service');

function stubClient(payload) {
  return { fetchCountry: async () => payload };
}

function buildApp(payload) {
  const service = createDirectoryService(stubClient(payload));
  return createApp(service);
}

test('endpoint returns 200 for known country', async () => {
  const app = buildApp({ data: [{ name: 'Afghanistan', callingCodes: ['93'] }] });
  const response = await request(app)
    .get('/phone-numbers')
    .query({ country: 'Afghanistan', phone: '656445445' });
  expect(response.status).toBe(200);
});
