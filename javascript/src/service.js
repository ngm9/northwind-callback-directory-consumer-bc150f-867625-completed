function createDirectoryService(client) {
  return {
    async getPhoneNumbers(country, phone) {
      const payload = await client.fetchCountry(country);
      const record = payload.data[0];
      const callingCode = record.callingCodes[0];
      return '+' + callingCode + ' ' + phone;
    },
  };
}

module.exports = { createDirectoryService };
