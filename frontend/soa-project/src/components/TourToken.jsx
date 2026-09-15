import api from '../config/axios';
import './TourToken.css'

const TourToken = ({tour, currentPos, setActiveTour}) => {
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
	const handleStart = async () => {
		var body = {};
		if (currentPos) {
			body = {
				current_lat: currentPos.lat,
				current_lng: currentPos.lng
			};
		} else {
			body = {
				current_lat: null,
				current_lng: null
			};
		}

		try {
			const res = await api.post(`/purchase/tokens/start/${tour.token_id}`, body);
			console.log(res.data);
			setActiveTour(res.data);
		} catch(err) {
			console.log(err);
		}
	};
	return (
		<div className="show-tour-card" style={{marginBottom: '5px'}}>
			<div style={{width: '100%'}}>
			<div className='tour-card-header'>
				<h2>{tour.title}</h2>
				<div className='tour-card-header-buttons'>
					<button 
						className='btn-save'
						onClick={handleStart}
						style={{marginTop: '7px'}}>Start</button>
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
				{tour.tags.length !== 0 && tour.tags.map((t, i) => 
					<p style={{marginLeft: '5px'}} key={i}>{t}</p>)}
			</div>
			</div>
		</div>
	);
};


export default TourToken;
