from fastapi import APIRouter, Depends
from sqlmodel import Session, select
from app import get_session
from models.tour import Tour

router = APIRouter(prefix="/tours", tags=["tours"])


@router.post("", response_model=Tour)
def create_tour(tour: Tour, session: Session = Depends(get_session)):
    session.add(tour)
    session.commit()
    session.refresh(tour)
    return tour


@router.get("", response_model=list[Tour])
def get_tours(session: Session = Depends(get_session)):
    return session.exec(select(Tour)).all()
