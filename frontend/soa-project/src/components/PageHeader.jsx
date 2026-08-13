import { useNavigate } from 'react-router-dom';
import './PageHeader.css'
import { useAuth } from '../context/AuthContext';

const StripButton = ({onClick, children}) => {
	return (
		<div className='strip-button' onClick={onClick}>
			<p>{children}</p>
		</div>
	);
};

const PageHeader = () => {
	const navigate = useNavigate();
	const { logout } = useAuth();

	const handleLogout = () => {
		logout();
		navigate('/');
	};
	const goToBlog = () => {
		navigate('/blog');
	};
	const goToProfile = () => {
		navigate('/profile');
	};
	const goToTours = () => {
		navigate('/tours');
	};
	const goToHome = () => {
		navigate('/');
	};
	const goToSim = () => {
		navigate('/simulation');
	};
	const goToCart = () => {
		navigate('/cart');
	};
	const btnList = {
		"HOME": goToHome, 
		"BLOG": goToBlog, 
		"PROFILE": goToProfile, 
		"TOURS": goToTours,
		"SIM": goToSim,
		"CART": goToCart
	};
	return (
		<div className='page-header-container'>
			<div className='page-header'>
				<h1>APP NAME</h1>
			</div>

			<div className='page-header-strip'>
				<div style={{display: 'flex'}}>
					{Object.entries(btnList).map(([k, v], i) => <StripButton 
					onClick={v} key={i}>{k}</StripButton>)}
				</div>
				<div>
					<StripButton onClick={handleLogout}>LOGOUT</StripButton>
				</div>
			</div>
		</div>
	);
};

export default PageHeader;
