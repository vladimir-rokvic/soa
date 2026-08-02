import { useState } from 'react';
import './CreateBlogCard.css'
import { useAuth } from '../context/AuthContext';
import api from '../config/axios';

const CreateBlogCard = ({blogCreated}) => {
	const [blogTitle, setBlogTitle] = useState('');
	const [blogDescription, setBlogDescription] = useState('');

	const {user} = useAuth();

	const handleSave = async () => {
		const body = {
			title: blogTitle,
			description: blogDescription,
			author_id: user.id
		};

		try {
			const res = await api.post('/blog/', body);
			console.log(res.data);
			setBlogDescription('');
			setBlogTitle('');
			blogCreated(res.data);
		} catch (err) {
			console.log(err);
		};
	};

	return(
		<div className='create-blog-card'>
			<div style={{width: '100%', display: 'flex', justifyContent: 'space-between'}}>
				<input className='blog-title-input'
					type="text" 
					placeholder='Enter title here'
					value={blogTitle}
					onChange={(e) => {setBlogTitle(e.target.value)}}
					/>
				<button onClick={handleSave} className='btn-save'>Save</button>
			</div>
			<textarea className='blog-description-input'
				type="text" 
				value={blogDescription}
				placeholder='Enter description here'
				onChange={(e) => {setBlogDescription(e.target.value)}}
				/>	
		</div>
	);
};

export default CreateBlogCard;
