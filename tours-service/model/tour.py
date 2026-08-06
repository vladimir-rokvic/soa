from pydantic import BaseModel
from enum import Enum
#from tag import Tag


class Difficulty(str, Enum):
    EASY = "Easy"
    MEDIUM = "Medium"
    HARD = "Hard"


class Status(str, Enum):
    DRAFT = "Draft"
    PUBLISHED = "Published"
    ARCHIVED = "Archived"


class Tour(BaseModel):
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    status: Status
    price: float
    author_id: str


class CreateTour(BaseModel):
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    author_id: str


class TourResponse(BaseModel):
    id: str
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    status: Status
    price: float
    author_id: str
