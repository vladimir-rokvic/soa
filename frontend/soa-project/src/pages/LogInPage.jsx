import { useState } from 'react';
import './Page.css'
import { useNavigate } from 'react-router-dom';
import api from '../config/axios';
import { useAuth } from '../context/AuthContext';

const LogInPage = () => {
	const [username, setUsername] = useState(null);
	const [password, setPassword] = useState(null);

	const { login } = useAuth();
	const navigate = useNavigate();

	const handleLogIn = async () => {
		const body = {
			username: username,
			password: password
		}
		
		try {
			const res = await api.post('/users/login', body);
			console.log(res.data);
			const user = {
				id: res.data.id,
				username: username,
				token: res.data.token,
				role: res.data.role
			};
			login(user);
			navigate('/');
		} catch (err) {
			console.log(err);
		}
	}
	return(
		<div>
			<div className="loginForm">
				<label style={{fontSize: '24px'}}>LOG IN</label>
				<div className='idk'>
					<label>EMAIL OR USERNAME</label>
					<input
						type='text' 
						style={{
							width: '300px',
							height: '25px',
							margin: '0 auto',
						}}
						placeholder='Enter username or email here' 
						onChange={(e) => {setUsername(e.target.value)}}
					/>
					<label>PASSWORD</label>
					<input
						type='password' 
						style={{
							width: '300px',
							height: '25px',
							margin: '0 auto',
						}}
						placeholder='Enter password here' 
						onChange={(e) => {setPassword(e.target.value)}}
					/>
					<button
						style={{
							width: '150px',
							height: '30px',
							margin: '20px auto',
						}}
						onClick={handleLogIn}
					>Log in
					</button>
					<p style={{color: 'gray'}}>Don't have an account,
						<a href='/register' style={{color: 'gray'}}> create one</a>
					</p>
				</div>
			</div>
		</div>
	);
};

export default LogInPage;
