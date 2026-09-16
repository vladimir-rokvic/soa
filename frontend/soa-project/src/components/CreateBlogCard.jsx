import { useState } from 'react';
import './CreateBlogCard.css'
import { useAuth } from '../context/AuthContext';
import api from '../config/axios';

const CreateBlogCard = ({blogCreated}) => {
	const [blogTitle, setBlogTitle] = useState('');
	const [blogDescription, setBlogDescription] = useState('');
	const [previews, setPreviews] = useState([]);
	const [rawFiles, setRawFiles] = useState([]);

	const {user} = useAuth();

	const handleFileChange = (e) => {
		const files = Array.from(e.target.files);
		setRawFiles(files);
		setPreviews(files.map(f => URL.createObjectURL(f)));
	};

	const handleSave = async () => {
		const body = {
			title: blogTitle,
			description: blogDescription,
			author_id: user.id
		};

		const formData = new FormData();
		formData.append('body', JSON.stringify(body));
		rawFiles.forEach(file => {
			formData.append('images', file);
		});

		try {
			const res = await api.post('/blog/', formData);
			console.log(res.data);
			setBlogDescription('');
			setBlogTitle('');
			setRawFiles([]);
			setPreviews([]);
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
			<div style={{display: 'flex', flexDirection: 'column', marginTop: '10px'}}>
				<label 
					style={{cursor: 'pointer', margin: '5px auto'}}
					htmlFor='blog-image-upload'>Add images</label>
				<input 
					id='blog-image-upload'
					type='file'
					accept='image/png, image/jpeg'
					multiple
					onChange={handleFileChange}
					style={{display: 'none'}}
				/>
				{previews.length > 0 && (
					<div style={{display: 'flex', flexWrap: 'wrap', gap: '5px'}}>
						{previews.map((src, i) => (
							<img key={i} src={src} style={{maxWidth: '80px', maxHeight: '80px'}} />
						))}
					</div>
				)}
			</div>
		</div>
	);
};

export default CreateBlogCard;
