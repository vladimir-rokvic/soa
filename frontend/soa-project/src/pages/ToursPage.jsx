import { useAuth } from '../context/AuthContext';
import GuidePage from './GuidePage';
import './Page.css'
import TouristPage from './TouristPage';

const ToursPage = () => {
	const {user} = useAuth();
	return user.role === "vodic" ? <GuidePage /> : <TouristPage />
};

export default ToursPage;
