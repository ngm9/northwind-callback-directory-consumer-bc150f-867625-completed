from fastapi import APIRouter, Query

from app.client import CountriesClient
from app.config import load_config
from app.service import DirectoryService

router = APIRouter()

_config = load_config()
_service = DirectoryService(CountriesClient(_config))


def get_service() -> DirectoryService:
    return _service


@router.get("/phone-numbers")
def phone_numbers(country: str = Query(...), phone: str = Query(...)):
    service = get_service()
    resolved = service.getPhoneNumbers(country, phone)
    return resolved
