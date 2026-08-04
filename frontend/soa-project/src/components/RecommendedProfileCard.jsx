import api from '../config/axios';
import { useAuth } from '../context/AuthContext';
import './ProfileCard.css'

const RecommendedProfileCard = ({recommendedUser}) => {
	const {user} = useAuth();

	const handleFollow = async () => {
		const body = {
			clientId: user.id,
			userId: recommendedUser.id
		};
		console.log(body);

		try {
			const res = await api.post('/followers/users/follow', body)
			console.log(res);
		} catch(err) {
			console.log(err)
		}
	};

	return(
		<div className='recommended-card'>
				<h3>{recommendedUser.username}</h3>
				<button 
					onClick={handleFollow}
					className='btn-save'>Follow</button>
		</div>
	);
};

export default RecommendedProfileCard;
