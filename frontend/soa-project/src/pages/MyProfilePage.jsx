import { useEffect, useState } from 'react'
import PageHeader from '../components/PageHeader'
import './Page.css'
import api from '../config/axios';
import { useAuth } from '../context/AuthContext';

const MyProfilePage = () => {
	const [profile, setProfile] = useState();
	const [editing, setEditing] = useState(false);

	const [username, setUsername] = useState('');
	const [email, setEmail] = useState('');
	const [firstName, setFirstName] = useState('');
	const [lastName, setLastName] = useState('');
	const [bio, setBio] = useState('');
	const [motto, setMotto] = useState('');
	const [category, setCategory] = useState('');

	const [file, setFile] = useState(null);
	const [rawFile, setRawFile] = useState(null);

	const {user, login, logout} = useAuth();

	useEffect(() => {
		const fetchUserData = async () => {
			const res = await api.get('/users/myProfile');
			console.log(res.data);
			setProfile(res.data);
		};
		fetchUserData();
	}, []);

	const roleToCategory = { turista: '1', vodic: '2' };
	const categoryToRole = { '1': 'turista', '2': 'vodic' };

	const startEditing = () => {
		setUsername(profile.username || '');
		setEmail(profile.email || '');
		setFirstName(profile.first_name || '');
		setLastName(profile.last_name || '');
		setBio(profile.bio || '');
		setMotto(profile.motto || '');
		setCategory(roleToCategory[profile.role] || '');
		setEditing(true);
	};

	const cancelEditing = () => {
		setEditing(false);
	};

	const handleSave = async () => {
		const role = categoryToRole[category] || profile.role;

		const formData = new FormData();

		const body = {
			id: profile.id,
			username: username,
			email: email,
			role: role,
			first_name: firstName,
			last_name: lastName,
			bio: bio,
			motto: motto
		};


		formData.append('body', JSON.stringify(body));
		if (rawFile) {
			formData.append('profile_image', rawFile);
		}

		try {
			const res = await api.put('/users/myProfile', formData);
			console.log(res.data);
			setProfile(res.data);
			const token = user.token;
			logout();
			const logInUser = {
				id: res.data.id,
				username: username,
				token: token,
				role: res.data.role
			};
			login(logInUser);
		} catch (err) {
			console.log(err);
			return;
		}

		setEditing(false);
	};

	const handleChange = (e) => {
		setFile(URL.createObjectURL(e.target.files[0]));
		setRawFile(e.target.files[0]);
	};

	const roleMap = new Map([
		['vodic', 'Guide'],
		['turista', 'Tourist']
	]);

	if(!profile) return <p>Loading</p>;

	return(
		<>
			<PageHeader />
			<div className='profile-page'>
				<div className='user-info'>
					<div style={{display: 'flex', flexDirection: 'column'}}>
						<div className='profile-img'>
							{editing ? (<img src={file} class='pr-img'/>) :
							(<img className='pr-img' src={profile.image_path === '' ? 
							null : `http://localhost:8080/users/${profile.image_path}`} />)}
						</div>
						{editing ? 
						(<>
						<label 
							style={{cursor: 'pointer', margin: '10px auto'}}
							htmlFor='profile-img-input'>Choose profile image</label>
							<input id='profile-img-input' type='file'
							onChange={handleChange}
							accept='image/png, image/jpeg' style={{display: 'none'}}/>
						</>) :
						(<label style={{margin: '10px auto'}}>Profile image</label>)}
					</div>
					<div className='profile-info'>
						<div style={{display: 'flex'}}>
							<div className='first-name'>
								{editing ? (
									<input
										placeholder='First name'
										value={firstName}
										onChange={(e) => setFirstName(e.target.value)}
									/>
								) : (
									profile.first_name ? 
									(<h2>{profile.first_name}</h2>) :
									(<h2>No first name set</h2>)
								)}
							</div>
							<div className='last-name'>
								{editing ? (
									<input
										placeholder='Last name'
										value={lastName}
										onChange={(e) => setLastName(e.target.value)}
									/>
								) : (
									profile.last_name ? 
									(<h2>{profile.last_name}</h2>) :
									(<h2>No last name set</h2>)
								)}
							</div>
						</div>
						<div style={{display: 'flex'}}>
							<div id='username-border'>
								{editing ? (
									<input
										placeholder='Username'
										value={username}
										onChange={(e) => setUsername(e.target.value)}
									/>
								) : (
									<p id='username'>@{profile.username}</p>
								)}
							</div>
							{editing ? (
								<input
									className='email'
									placeholder='Email'
									value={email}
									onChange={(e) => setEmail(e.target.value)}
								/>
							) : (
								<p className='email'>Email: {profile.email}</p>
							)}
							{editing ? (
								<select className='email select-box'
										value={category} 
										onChange={e => setCategory(e.target.value)}>
								  <option value="">-- Select a role --</option>
								  <option value="1">Tourist</option>
								  <option value="2">Guide</option>
								</select>
							) : (
								<p className='email'>Role: {roleMap.get(profile.role)}</p>
							)}
						</div>
						{editing ? (
							<input
								className='motto-input'
								style={{marginTop: '15px'}}
								placeholder='Motto'
								value={motto}
								onChange={(e) => setMotto(e.target.value)}
							/>
						) : (
							<p style={{
								margin: '0',
								marginTop: '15px',
								fontSize: '20px'}}>
								Motto: {profile.motto ? profile.motto : 'No motto set'}</p>
						)}
						<div className='bio'>
							{editing ? (
								<textarea
									placeholder='Bio'
									value={bio}
									onChange={(e) => setBio(e.target.value)}
								/>
							) : (
								profile.bio ? <p>{profile.bio}</p> : <p>No bio set</p>
							)}
						</div>
					</div>
				</div>
				<div style={{display: 'flex', justifyContent: 'right'}}>
					{editing ? (
						<>
							<button 
								onClick={cancelEditing}
								style={{marginTop: '10px', marginRight: '10px'}}>Cancel</button>
							<button 
								onClick={handleSave}
								style={{marginTop: '10px', marginRight: '10px'}}>Save</button>
						</>
					) : (
						<button 
							onClick={startEditing}
							style={{marginTop: '10px', marginRight: '10px'}}>Edit</button>
					)}
				</div>
			</div>
		</>
	);
};

export default MyProfilePage;
