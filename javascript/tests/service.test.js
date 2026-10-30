const { createDirectoryService } = require('../src/service');

function stubClient(payload) {
  return { fetchCountry: async () => payload };
}

test('single calling code resolves', async () => {
  const service = createDirectoryService(
    stubClient({ data: [{ name: 'Afghanistan', callingCodes: ['93'] }] })
  );
  const result = await service.getPhoneNumbers('Afghanistan', '656445445');
  expect(result).toBe('+93 656445445');
});
