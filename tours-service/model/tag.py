from pydantic import BaseModel
import uuid


class Tag(BaseModel):
    id: uuid.UUID
    title: str
    description: str
