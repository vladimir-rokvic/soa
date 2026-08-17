import api from '../config/axios';
import './TourToken.css'

const TourToken = ({token, currentPos}) => {
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
			const res = await api.post(`/purchase/tokens/start/${token.ID}`, body);
			console.log(res.data);
		} catch(err) {
			console.log(err);
		}
	};
	return (
		<div className='tour-token'>
			<div style={{
					width: '100%',
					display: 'flex',
					justifyContent: 'space-between'
				}}>
				<h3>{token.ID}</h3>
				<button 
					className='btn-save'
					onClick={handleStart}
					style={{marginTop: '7px'}}>Start</button>
			</div>
		</div>
	);
};


export default TourToken;
