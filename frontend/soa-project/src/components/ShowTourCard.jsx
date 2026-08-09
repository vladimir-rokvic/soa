import './TourCard.css'


const ShowTourCard = ({tour}) => {
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

	const handleBuy = async () => {
		try {
			//TODO
		} catch (err) {
			console.log(err);
		};
	};

	return (
		<div className="tour-card">
			<div className='tour-card-header'>
				<h2>{tour.title}</h2>
				<div className='tour-card-header-buttons'>
					<button onClick={handleBuy}>Buy</button>
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
	);
};


export default ShowTourCard;
