from app.service import DirectoryService


class StubClient:
    def __init__(self, payload):
        self._payload = payload

    def fetch_country(self, country):
        return self._payload


def test_single_calling_code():
    client = StubClient({"data": [{"name": "Afghanistan", "callingCodes": ["93"]}]})
    service = DirectoryService(client)
    assert service.getPhoneNumbers("Afghanistan", "656445445") == "+93 656445445"
