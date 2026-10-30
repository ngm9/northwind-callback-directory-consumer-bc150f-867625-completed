from dataclasses import dataclass


@dataclass(frozen=True)
class Config:
    upstream_base_url: str = "https://jsonmock.hackerrank.com/api/countries"
    request_timeout_seconds: float = 5.0


def load_config() -> Config:
    return Config()
