from fastapi.testclient import TestClient

from app.main import create_app
from app.routes import get_service
from app.service import DirectoryService


class StubClient:
    def __init__(self, payload):
        self._payload = payload

    def fetch_country(self, country):
        return self._payload


def build_client(payload):
    app = create_app()
    service = DirectoryService(StubClient(payload))
    app.dependency_overrides[get_service] = lambda: service
    return TestClient(app)


def test_endpoint_single_calling_code():
    client = build_client({"data": [{"name": "Afghanistan", "callingCodes": ["93"]}]})
    response = client.get("/phone-numbers", params={"country": "Afghanistan", "phone": "656445445"})
    assert response.status_code == 200
