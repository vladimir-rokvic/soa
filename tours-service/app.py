from fastapi import FastAPI
import routers.tours as tours

app = FastAPI()
app.include_router(tours.router)


@app.on_event("startup")
def on_startup():
    pass
