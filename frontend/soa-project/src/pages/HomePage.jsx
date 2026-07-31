import { useAuth } from "../context/AuthContext";

const HomePage = () => {
	const {logout} = useAuth();
	const handleLogout = () => {
		logout();
	}
	return(
		<>
			<h1>Home page</h1>
			<button onClick={handleLogout}>Log out</button>
		</>
	);
}

export default HomePage;

