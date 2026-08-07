from pydantic import BaseModel
from typing import Optional


class CreatePoint(BaseModel):
    title: str
    description: str
    lat: float
    lng: float


class Point(BaseModel):
    title: str
    description: str
    image_path: Optional[str] = None
    lat: float
    lng: float
