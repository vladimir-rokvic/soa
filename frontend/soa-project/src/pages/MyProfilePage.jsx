import { useEffect, useState } from 'react'
import PageHeader from '../components/PageHeader'
import './Page.css'
import api from '../config/axios';

const MyProfilePage = () => {
	const [profile, setProfile] = useState();

	useEffect(() => {
		const fetchUserData = async () => {
			const res = await api.get('/users/myProfile');
			console.log(res.data);
			setProfile(res.data);
		};
		fetchUserData();
	}, []);

	const [editing, setEditing] = useState(false);
	const changeEditing = () => {
		setEditing(!editing);
	};

	const [category, setCategory] = useState('');
	const handleSave = async () => {
		setEditing(!editing);

		if(category === "1") profile.role = "turista";
		if(category === "2") profile.role = "vodic";

		const body = {
			id: profile.id,
			username: profile.username,
			email: profile.email,
			role: profile.role,
			//TODO
			//password: "vlada123"
		};

		try {
			const res = await api.put('/users/myProfile', body);
			console.log(res.data);
		} catch (err) {
			console.log(err);
		};
	};

	if(!profile) return <p>Loading</p>;
	return(
		<>
			<PageHeader />
			<div className='profile-page'>
				{!editing ? (
				<>
				<div className='user-info'>
					<h1>{profile.username}</h1>
					<h2>{profile.email}</h2>
					<h2>{profile.role}</h2>
				</div>
				<div style={{display: 'flex', justifyContent: 'right'}}>
					<button 
						onClick={changeEditing}
						style={{marginTop: '10px', marginRight: '10px'}}>Edit</button>
				</div>
				</>
				) : (
				<>
				<div className='user-info-edit'>
					<input
						placeholder={profile.username}
						onChange={(e) => profile.username = e.target.value}
					/>
					<input
						placeholder={profile.email}
						onChange={(e) => profile.email = e.target.value}
					/>
    				<select className='select-box'
							value={category} 
							onChange={e => setCategory(e.target.value)}>
    				  <option value="">-- Select a role --</option>
    				  <option value="1">Tourist</option>
    				  <option value="2">Guide</option>
    				</select>
				</div>
				<div style={{display: 'flex', justifyContent: 'right'}}>
					<button 
						onClick={handleSave}
						style={{marginTop: '10px', marginRight: '10px'}}>Save</button>
				</div>
				</>
				)}
			</div>
		</>
	);
};

export default MyProfilePage;
