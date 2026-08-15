import { use, useEffect, useState } from 'react';
import PageHeader from '../components/PageHeader'
import './Page.css'
import api from '../config/axios';
import { useAuth } from '../context/AuthContext';
import TourToken from '../components/TourToken';

const NoActiveTourPage = ({user}) => {
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
				{tokens?.length !== 0 && tokens.map(t => (
					<TourToken token={t} />
				))}
			</div>
		</>
	);
};

const ActiveTourPage = ({tour}) => {
	return (
		<>
		</>
	);
};

const SimPage = () => {
	const {user} = useAuth();
	const [activeTour, setActiveTour] = useState(null);

	useEffect(() => {
		const fetchActiveTour = async () => {
			try {
				const res = await api.get(`/purchase/tokens/active/user/${user.id}`);
				console.log(res.data);
				setActiveTour(res.data);
			} catch(err) {
				console.log(err);
			}
			fetchActiveTour();
		}
	}, []);

	return (
		<>
			{activeTour ? <ActiveTourPage tour={activeTour} /> : <NoActiveTourPage user={user} />}
		</>
	);
	
};


export default SimPage;
