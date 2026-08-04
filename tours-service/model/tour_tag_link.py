import uuid
from sqlmodel import SQLModel, Field


class TourTagLink(SQLModel, table=True):
    tour_id: uuid.UUID = Field(foreign_key="tour.id", primary_key=True)
    tag_id: uuid.UUID = Field(foreign_key="tag.id", primary_key=True)
