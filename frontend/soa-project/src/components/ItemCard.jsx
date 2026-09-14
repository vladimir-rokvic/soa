import { useEffect, useState } from 'react';
import { MapContainer, Marker, Popup, TileLayer } from 'react-leaflet';
import './TourCard.css'
import api from '../config/axios';

const StartPopup = ({point}) => {
	return (
		<Popup className='add-tour-popup'>
			<div className='add-tour-popup-content'>
				<div style={{display: 'flex', marginBottom: '5px'}}>
					<h3>Start point</h3>
				</div>
				<div className='add-tour-popup-metadata'>
					<div style={{display: 'flex', flexDirection: 'column'}}>
						<div className='popup-image'>
							{point.image_path && <img src={"http://localhost:8000" + point.image_path} />}
						</div>
					</div>
					<div style={{
							display: 'flex',
							flexDirection: 'column',
							marginLeft: '10px'
					}}>
						<label>{point.title}</label>
						<p>{point.description}</p>
					</div>
				</div>
			</div>
		</Popup>
	);
};

const ItemCard = ({id, itemId, onRemove}) => {
	const [tour, setTour] = useState(null);

	const diffColor = new Map([
		['Easy', '#00CC00'],
		['Medium', '#FF9900'], 
		['Hard', '#CC3300']
	]);
	const statusColor = new Map([
		['Draft', '#00CC00'],
		['Published', '#FF9900'], 
		['Archived', '#CC3300']
	]);

	useEffect(() => {
		const fetchTour = async () => {
			try {
				const res = await api.get(`/tours/${id}`);
				console.log(res.data);
				setTour(res.data);
			} catch (err) {
				console.log(err);
			}
		};
		fetchTour();
	}, [id]);

	const handleRemove = async () => {
		try {
			const res = await api.delete(`/purchase/oi/${itemId}`);
			console.log(res.data);
			onRemove(itemId);
		} catch (err) {
			console.log(err);
		}
	};

	if (!tour) return null;

	return (
		<div className="show-tour-card">
			<div style={{width: '80%'}}>
			<div className='tour-card-header'>
				<h2>{tour.title}</h2>
				<div className='tour-card-header-buttons'>
					<button style={{width: '100px'}} onClick={handleRemove}>Remove</button>
				</div>
			</div>
			<p>{tour.description}</p>
			<div className="tour-card-metadata">
				<label>Status: </label>
				<p style={{color: statusColor.get(tour.status)}}>{tour.status}</p>
				<label>Difficulty: </label>
				<p style={{color: diffColor.get(tour.difficulty)}}>{tour.difficulty}</p>
			</div>
			<br />
			<label style={{marginTop: '10px'}}>Tags: </label>
			<div className="tour-card-tags">
				{tour.tags.length !== 0 && tour.tags.map((t, i) => <p key={i}>{t}</p>)}
			</div>
			<div className="tour-card-price">
				<div style={{display: 'flex'}}>
					<label>Price: </label>
					<p>{tour.price}</p>
				</div>
			</div>
			</div>
			<div className='show-card-map'>
				<MapContainer 
					center={[tour.start_point.lat, tour.start_point.lng]} 
					zoom={12}
					style={{width: '100%', height: '100%'}}
					scrollWheelZoom={false}>
    				<TileLayer
						attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      					url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    				/>
					<Marker position={[tour.start_point.lat, tour.start_point.lng]}>
						<StartPopup point={tour.start_point}/>
					</Marker>
				</MapContainer>
			</div>
		</div>
	);
};

export default ItemCard;
