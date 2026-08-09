from fastapi import FastAPI
import routers.tours as tours
from fastapi.staticfiles import StaticFiles

app = FastAPI()
app.mount("/tours/uploads", StaticFiles(directory="uploads"), name="uploads")
app.include_router(tours.router)


@app.on_event("startup")
def on_startup():
    pass
