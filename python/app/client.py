import httpx


class CountriesClient:
    def __init__(self, config):
        self._config = config

    def fetch_country(self, country: str) -> dict:
        url = "https://jsonmock.hackerrank.com/api/countries?name=" + country
        response = httpx.get(url)
        response.raise_for_status()
        return response.json()
