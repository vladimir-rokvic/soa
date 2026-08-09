from pydantic import BaseModel
from enum import Enum
from .point import Point, CreatePoint


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
    start_point: Point
    end_point: Point


class CreateTour(BaseModel):
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    author_id: str
    start_point: CreatePoint | None
    end_point: CreatePoint | None


class TourResponse(BaseModel):
    id: str
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    status: Status
    price: float
    author_id: str
    start_point: Point
    end_point: Point


class UpdateTour(BaseModel):
    title: str
    description: str
    difficulty: Difficulty
    tags: list[str]
    start_point: CreatePoint | None
    end_point: CreatePoint | None


class PublishTour(BaseModel):
    price: float
