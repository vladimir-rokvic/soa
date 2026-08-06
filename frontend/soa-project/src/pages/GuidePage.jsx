import { useNavigate } from "react-router-dom";
import PageHeader from "../components/PageHeader"
import './Page.css'
import { useEffect, useState } from "react";
import api from "../config/axios";
import { useAuth } from "../context/AuthContext";
import TourCard from "../components/TourCard";

const GuidePage = () => {
	const navigate = useNavigate();
	const {user} = useAuth();
	const [tours, setTours] = useState([]);

	useEffect(() => {
		const fetchTours = async () => {
			try {
				const res = await api.get(`/tours/user/${user.id}`);
				console.log(res.data);
				setTours(res.data);
			} catch(err) {
				console.log(err);
			};
		};
		fetchTours();
	}, []);

	return(
		<>
			<PageHeader />
			<div className="guide-page">
				<div style={{display: 'flex', justifyContent: 'right'}}>
				<button className="btn-save"
					style={{
						margin: '10px',
						width: '115px',
						justifySelf: 'right'
					}}
					onClick={() => {navigate('/tours/add')}}
				>+ Add tour</button>
				</div>
				<div className="guide-content">
					{tours?.length !== 0 && tours.map(t => (
						<TourCard key={t.id} tour={t} />
					))}
				</div>
			</div>
		</>
	);
};


export default GuidePage;
