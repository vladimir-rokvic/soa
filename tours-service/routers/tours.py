from fastapi import APIRouter, HTTPException
from bson import ObjectId
from bson.errors import InvalidId
from database import tours_collection
from model.tour import CreateTour, Tour, TourResponse, Status, Difficulty

router = APIRouter(
    prefix="/tours",
    tags=["Tours"]
)


def to_response(doc: dict) -> TourResponse:
    return TourResponse(
        id=str(doc["_id"]),
        title=doc["title"],
        description=doc["description"],
        difficulty=doc["difficulty"],
        tags=doc["tags"],
        status=doc["status"],
        price=doc["price"],
        author_id=doc["author_id"]
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


@router.post("/", response_model=TourResponse)
def add_tour(tour: CreateTour):
    new_tour = Tour(
        title=tour.title,
        description=tour.description,
        difficulty=tour.difficulty,
        tags=tour.tags,
        status=Status.DRAFT,
        price=0,
        author_id=tour.author_id
    )

    result = tours_collection.insert_one(new_tour.model_dump(mode="json"))

    doc = new_tour.model_dump(mode="json")
    doc["_id"] = result.inserted_id

    return to_response(doc)
