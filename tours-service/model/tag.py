from sqlmodel import SQLModel, Field, Relationship
import uuid
from tour_tag_link import TourTagLink
from tour import Tour


class Tag(SQLModel, table=True):
    id: uuid.UUID = Field(default_factory=uuid.uuid4, primary_key=True)
    title: str
    description: str
    tours: list[Tour] = Relationship(back_populates="tags",
                                     link_model=TourTagLink)
