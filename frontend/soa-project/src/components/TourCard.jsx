import { useNavigate } from 'react-router-dom';
import api from '../config/axios';
import './TourCard.css';

const TourCard = ({tour}) => {
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

	const navigate = useNavigate();

	const handlePublish = async () => {
		try {
			const res = await api.put(`/tours/publish/${tour.id}`);
			console.log(res.data);
			window.location.reload(false);
		} catch(err) {
			console.log(err);
		};
	};

	return (
		<div className="tour-card">
			<div className='tour-card-header'>
				<h2>{tour.title}</h2>
				<div className='tour-card-header-buttons'>
					<button onClick={() => {navigate(`/tours/${tour.id}/edit`)}}>Edit</button>
					{tour.status === 'Draft' && <button onClick={handlePublish}>Publish</button>}
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
				{tour.tags.length !== 0 && tour.tags.map(t => <p>{t}</p>)}
			</div>
			<div className="tour-card-price">
				<label>Price: </label>
				<p>{tour.price}</p>
			</div>
		</div>
	);
};


export default TourCard;
