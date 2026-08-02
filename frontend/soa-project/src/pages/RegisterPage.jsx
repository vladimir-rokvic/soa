import { useState } from 'react';
import './Page.css'
import api from '../config/axios';
import { useNavigate } from 'react-router-dom';

const RegisterPage = () => {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
	const [confirmPassword, setConfirmPassword] = useState('');
	const [email, setEmail] = useState('');
	const [category, setCategory] = useState('');

	const navigate = useNavigate();

	const handleRegister = async () => {
		if(password !== confirmPassword) return;
		if(category === "0") return;

		const body = {
			username: username,
			email: email,
			password: password,
			role: parseInt(category)
		};

		try {
			const res = await api.post('/users/add', body);
			console.log(res.data);
			navigate('/');
		} catch (err) {
			console.log(err);
		};
	};

	return(
		<div style={{display: 'flex', flexDirection: 'column', padding: '300px'}}>
			<label style={{margin: '0 auto', fontSize: '24px'}}>REGISTER</label>
			<div className="register-form">
					<label>USERNAME</label>
					<input
						type='text' 
						placeholder='Enter username here' 
						onChange={(e) => {setUsername(e.target.value)}}
					/>

					<label>EMAIL</label>
					<input
						type='text' 
						placeholder='Enter email here' 
						onChange={(e) => {setEmail(e.target.value)}}
					/>

					<label>CATEGORY</label>
    				<select className='select-box'
							value={category} 
							onChange={e => setCategory(e.target.value)}>
    				  <option value="">-- Select a role --</option>
    				  <option value="1">Tourist</option>
    				  <option value="2">Guide</option>
    				</select>

					<label>PASSWORD</label>
					<input
						type='password' 
						placeholder='Enter password here' 
						onChange={(e) => {setPassword(e.target.value)}}
					/>

					<label>CONFIRM PASSWORD</label>
					<input
						type='password' 
						placeholder='Confirm password here' 
						onChange={(e) => {setConfirmPassword(e.target.value)}}
					/>
					{password !== confirmPassword && 
					(<label className='error-msg'>Passwords must match</label>)}
					
					<button
						style={{
							width: '150px',
							height: '30px',
							margin: '20px auto',
						}}
						onClick={handleRegister}
					>Register
					</button>
			</div>
		</div>
	);
};


export default RegisterPage;
