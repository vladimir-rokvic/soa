
const TourCard = ({tour}) => {
	return (
		<div className="tour-card">
			<h2>{tour.title}</h2>
			<div className="tour-card-metadata">
				<p>{tour.status}</p>
				<p>{tour.difficulty}</p>
				<p>{tour.price}</p>
			</div>
			<p>{tour.description}</p>
			<div className="tour-card-tags">
			</div>
		</div>
	);
};


export default TourCard;
