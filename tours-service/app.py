from fastapi import FastAPI
from sqlmodel import SQLModel, create_engine, Session
from api.tours import router

app = FastAPI()
app.include_router(router)

DATABASE_URL = "postgresql://postgres:mypassword@database:5432/postgres"

engine = create_engine(DATABASE_URL, echo=True)


def get_session():
    with Session(engine) as session:
        yield session


@app.on_event("startup")
def on_startup():
    SQLModel.metadata.create_all(engine)
