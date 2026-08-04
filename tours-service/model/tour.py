from sqlmodel import SQLModel, Field, Relationship
from enum import IntEnum
import uuid
from tag import Tag
from tour_tag_link import TourTagLink


class Difficulty(IntEnum):
    EASY = 0
    MEDIUM = 1
    HARD = 2


class Status(IntEnum):
    DRAFT = 0
    PUBLISHED = 1
    ARCHIVED = 2


class Tour(SQLModel, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    title: str
    description: str
    authorId: uuid.UUID
    cost: float
    difficulty: Difficulty
    tags: list[Tag] = Relationship(back_populates="tours",
                                   link_model=TourTagLink)
    status: Status
