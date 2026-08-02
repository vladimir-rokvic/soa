import PageHeader from "../components/PageHeader";
import { useAuth } from "../context/AuthContext";
import './Page.css';

const HomePage = () => {
	const {logout} = useAuth();
	const handleLogout = () => {
		logout();
	}
	return(
		<div>
			<PageHeader />
		</div>
	);
}

export default HomePage;

