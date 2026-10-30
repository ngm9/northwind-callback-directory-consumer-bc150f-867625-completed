class DirectoryService:
    def __init__(self, client):
        self._client = client

    def getPhoneNumbers(self, country: str, phone: str) -> str:
        payload = self._client.fetch_country(country)
        record = payload["data"][0]
        calling_code = record["callingCodes"][0]
        return "+" + calling_code + " " + phone
