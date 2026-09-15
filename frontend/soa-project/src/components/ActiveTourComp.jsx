import { useEffect, useState } from 'react';
import './ActiveTourComp.css'
import api from '../config/axios';

const ActiveTourComp = ({te, setActiveTour}) => {
	const [tour, setTour] = useState(null);

	useEffect(() => {
		const fetchTour = async () => {
			try {
				const res = await api.get(`/purchase/tokens/${te.token_id}/tour`);
				setTour(res.data);
				console.log(res.data);
			} catch(err) {
				console.log(err);
			};
		};
		
		fetchTour();
	}, []);

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

	const handleAbandon = async () => {
		try {
			const res = await api.put(`/purchase/tokens/te/abandon/${te.id}`);
			console.log(res.data);
			setActiveTour(null);
		} catch(err) {
			console.log(err);
		}
	};

	if(!tour) return;

	return (
		<div className="show-tour-card" style={{marginBottom: '5px'}}>
			<div style={{width: '100%'}}>
			<div className='tour-card-header'>
				<h2>{tour.title}</h2>
				<div className='tour-card-header-buttons'>
					<button 
						className='btn-save'
						onClick={handleAbandon}
						style={{marginTop: '7px', width: '100px'}}>Abandon</button>
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


export default ActiveTourComp;
