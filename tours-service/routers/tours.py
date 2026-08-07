from fastapi import APIRouter, HTTPException, Form, UploadFile, File
from bson import ObjectId
from bson.errors import InvalidId
from database import tours_collection
from model.tour import CreateTour, UpdateTour, Tour, TourResponse, Status, Difficulty
from model.point import Point
import json
import os
import uuid

router = APIRouter(
    prefix="/tours",
    tags=["Tours"]
)


def to_image_url(image_path: str | None) -> str | None:
    if not image_path:
        return None
    filename = os.path.basename(image_path)
    return f"/tours/uploads/{filename}"


def to_response(doc: dict) -> TourResponse:
    start_point = dict(doc["start_point"])
    start_point["image_path"] = to_image_url(start_point.get("image_path"))

    end_point = dict(doc["end_point"])
    end_point["image_path"] = to_image_url(end_point.get("image_path"))

    return TourResponse(
        id=str(doc["_id"]),
        title=doc["title"],
        description=doc["description"],
        difficulty=doc["difficulty"],
        tags=doc["tags"],
        status=doc["status"],
        price=doc["price"],
        author_id=doc["author_id"],
        start_point=Point(**doc["start_point"]),
        end_point=Point(**doc["end_point"])
    )


@router.get("/", response_model=list[TourResponse])
def get_all():
    results = tours_collection.find({})

    return [
        to_response(result) for result in results
    ]


@router.get("/user/{user_id}", response_model=list[TourResponse])
def get_by_author_id(user_id: str):
    results = tours_collection.find({"author_id": user_id})

    return [
        to_response(result) for result in results
    ]


@router.get("/{tour_id}", response_model=TourResponse)
def get_by_id(tour_id: str):
    try:
        t_id = ObjectId(tour_id)
    except InvalidId:
        raise HTTPException(status_code=400, detail="Invalid id")

    results = tours_collection.find_one({"_id": t_id})
    if results is None:
        raise HTTPException(
            status_code=404,
            detail=f"No tour found by id: {tour_id}"
        )

    return to_response(results)


UPLOAD_DIR = "./uploads"
os.makedirs(UPLOAD_DIR, exist_ok=True)


async def save_image(image: UploadFile) -> str:
    ext = image.filename.split(".")[-1]
    filename = f"{uuid.uuid4()}.{ext}"
    path = os.path.join(UPLOAD_DIR, filename)

    with open(path, "wb") as f:
        con = await image.read()
        f.write(con)

    return path


@router.post("/", response_model=TourResponse)
async def add_tour(
    tour_body: str = Form(...),
    start_point_image: UploadFile | None = File(None),
    end_point_image: UploadFile | None = File(None)
):

    start_image_path = None
    if start_point_image:
        start_image_path = await save_image(start_point_image)

    end_image_path = None
    if end_point_image:
        end_image_path = await save_image(end_point_image)

    tour = CreateTour(**json.loads(tour_body))
    new_tour = Tour(
        title=tour.title,
        description=tour.description,
        difficulty=tour.difficulty,
        tags=tour.tags,
        status=Status.DRAFT,
        price=0,
        author_id=tour.author_id,
        start_point=Point(
            title=tour.start_point.title,
            description=tour.start_point.description,
            lat=tour.start_point.lat,
            lng=tour.start_point.lng,
            image_path=start_image_path
        ),
        end_point=Point(
            title=tour.end_point.title,
            description=tour.end_point.description,
            lat=tour.end_point.lat,
            lng=tour.end_point.lng,
            image_path=end_image_path
        )
    )

    result = tours_collection.insert_one(new_tour.model_dump(mode="json"))

    doc = new_tour.model_dump(mode="json")
    doc["_id"] = result.inserted_id

    return to_response(doc)


@router.put("/{tour_id}", response_model=TourResponse)
async def update_tour(
    tour_id: str,
    tour_body: str = Form(...),
    start_point_image: UploadFile | None = File(None),
    end_point_image: UploadFile | None = File(None)
):
    try:
        t_id = ObjectId(tour_id)
    except InvalidId:
        raise HTTPException(status_code=400, detail="Invalid id")

    existing = tours_collection.find_one({"_id": t_id})
    if existing is None:
        raise HTTPException(
            status_code=404,
            detail=f"No tour found by id: {tour_id}"
        )

    tour = UpdateTour(**json.loads(tour_body))

    start_image_path = existing["start_point"].get("image_path")
    if start_point_image:
        start_image_path = await save_image(start_point_image)

    end_image_path = existing["end_point"].get("image_path")
    if end_point_image:
        end_image_path = await save_image(end_point_image)

    update_doc = {
        "title": tour.title,
        "description": tour.description,
        "difficulty": tour.difficulty,
        "tags": tour.tags,
        "start_point": Point(
            title=tour.start_point.title,
            description=tour.start_point.description,
            lat=tour.start_point.lat,
            lng=tour.start_point.lng,
            image_path=start_image_path
        ).model_dump(mode="json"),
        "end_point": Point(
            title=tour.end_point.title,
            description=tour.end_point.description,
            lat=tour.end_point.lat,
            lng=tour.end_point.lng,
            image_path=end_image_path
        ).model_dump(mode="json"),
    }

    tours_collection.update_one({"_id": t_id}, {"$set": update_doc})

    updated = tours_collection.find_one({"_id": t_id})
    return to_response(updated)
