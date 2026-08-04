import uuid
from sqlmodel import SQLModel
from models.tour import Difficulty, Status


class TourCreate(SQLModel):
    title: str
    description: str
    cost: float
    difficulty: Difficulty
    status: Status
    tag_ids: list[uuid.UUID] = []


class TourRead(SQLModel):
    id: uuid.UUID
    title: str
    description: str
    authorId: uuid.UUID
    cost: float
    difficulty: Difficulty
    status: Status
    tags: list["TagRead"] = []
