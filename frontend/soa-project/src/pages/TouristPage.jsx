import './Page.css'
import PageHeader from '../components/PageHeader'
import { useEffect, useState } from 'react';
import api from '../config/axios';
import ShowTourCard from '../components/ShowTourCard';

const TouristPage = () => {
	const [tours, setTours] = useState([]);
	
	useEffect(() => {
		const fetchPublishedTours = async () => {
			try {
				const res = await api.get('tours/published');
				console.log(res.data);
				setTours(res.data);
			} catch(err) {
				console.log(err);
			};
		};
		fetchPublishedTours();
	}, []);
	return(
		<>
			<PageHeader />
			<div className="guide-page">
				<div className="guide-content">
					{tours?.length !== 0 && tours.map(t => (
						<ShowTourCard key={t.id} tour={t}/>
					))}
				</div>
			</div>
		</>
	);
};


export default TouristPage;
