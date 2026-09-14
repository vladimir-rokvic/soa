import { useEffect, useRef, useState } from 'react';
import PageHeader from '../components/PageHeader'
import './Page.css'
import api from '../config/axios';
import { useAuth } from '../context/AuthContext';
import TourToken from '../components/TourToken';
import { MapContainer, Marker, TileLayer, useMapEvents } from 'react-leaflet';

const MapClickHandler = ({setCurrentPos}) => {
	useMapEvents({
		click(e) {
			setCurrentPos(e.latlng);
		}
	});
};

const NoActiveTourPage = ({user, currentPos, setCurrentPos}) => {
	const [tokens, setTokens] = useState([]);

	useEffect(() => {
		const fetchTokens = async () => {
			try {
				const res = await api.get(`/purchase/tokens/user/${user.id}`);
				console.log(res.data);
				setTokens(res.data);
			} catch(err) {
				console.log(err);
			}
		}
		fetchTokens();
	}, []);
	return (
		<>
			<PageHeader />
			<div className='sim-page'>
				{tokens?.length !== 0 && tokens.map((t, i) => (
					<TourToken tour={t} key={t.token_id}/>
				))}
			</div>
			<div className='no-tour-map'>
				<MapContainer scrollWheelZoom={false} style={{width: '100%', height: '100%'}} zoom={13} center={[45.2671, 19.8335]}>
					<TileLayer 
						attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      					url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
					/>
					<MapClickHandler setCurrentPos={setCurrentPos} />
					{currentPos && (<Marker position={currentPos}/>)}
				</MapContainer>
			</div>
		</>
	);
};

const ActiveTourPage = ({tour, currentPos, setCurrentPos}) => {
	const currentPosRef = useRef(currentPos);
	
	useEffect(() => {
		currentPosRef.current = currentPos;
	}, [currentPos]);
	
	useEffect(() => {
		const fetchActivity = async () => {
			const pos = currentPosRef.current;
			if (!pos) return;
	
			const body = {
				current_lat: pos.lat,
				current_lng: pos.lng
			};
			try {
				const res = await api.put(`/purchase/tokens/te/${tour.id}`, body);
				console.log(res.data);
			} catch(err) {
				console.log(err);
			}
		};
	
		const interval = setInterval(fetchActivity, 10000);
		return () => clearInterval(interval);
	}, [tour.id]);

	return (
		<>
			<PageHeader />
			<div className='sim-page'>
				<h1>{tour.title}</h1>
			</div>
			<div className='no-tour-map'>
				<MapContainer style={{width: '100%', height: '100%'}} zoom={13} scrollWheelZoom={false} 
					center={[tour.points[0].Lat, tour.points[0].Lng]}>
					<TileLayer 
						attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      					url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
					/>
					<MapClickHandler setCurrentPos={setCurrentPos}/>
					{currentPos && (<Marker position={currentPos}/>)}
					{tour.points.length !== 0 && tour.points.map(p => (
							<Marker position={[p.Lat, p.Lng]}>
							</Marker>
						)
					)}
				</MapContainer>
			</div>
		</>
	);
};

const SimPage = () => {
	const {user} = useAuth();
	const [activeTour, setActiveTour] = useState(null);
	const [currentPos, setCurrentPos] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		const fetchActiveTour = async () => {
			try {
				const res = await api.get(`/purchase/tokens/active/user/${user.id}`);
				console.log(res.data);
				setActiveTour(res.data);
				setCurrentPos({lat: res.data.current_lat, lng: res.data.current_lng});
			} catch(err) {
				console.log(err);
			} finally {
				setLoading(false);
			}
		}
		fetchActiveTour();
	}, []);

	if(loading) {
		return <p>Loading</p>;
	}

	return (
		<>
			{activeTour ? 
				(<ActiveTourPage 
					tour={activeTour}
					currentPos={currentPos}
					setCurrentPos={setCurrentPos}
				/>)
				: (<NoActiveTourPage 
						currentPos={currentPos}
						setCurrentPos={setCurrentPos}
						user={user}
				/>)}
		</>
	);
	
};


export default SimPage;
