import { useState } from 'react';
import api from '../config/axios';
import './CreateBlogCard.css'

const EditingCard = ({id, title, description, dateCreated, onSave, onCancel}) => {
	const [editedTitle, setEditedTitle] = useState(title);
	const [editedDescription, setEditedDescription] = useState(description);

	const handleSave = async () => {
		const body = {
			id: id,
			title: editedTitle,
			description: editedDescription
		};
		try{
			const res = await api.put('/blog/', body);
			console.log(res.data);
			onSave({title: editedTitle, description: editedDescription});
		}catch (err) {
			console.log(err);
		}
	};

	return(
		<div className='blog-card'>
			<div style=
					{{width: '100%', display: 'flex', justifyContent: 'space-between'}}>
				<div style={{display: 'flex'}}>
					<input
						className='blog-title-input'
						value={editedTitle}
						onChange={(e) => setEditedTitle(e.target.value)}
					/>
					<p style={{
							marginLeft: '10px',
							marginTop: '5px',
							color: '#a0a0a0'}}>Date posted: {dateCreated}</p>
				</div>
				<div>
					<button
						onClick={onCancel}
						className='btn-save'
						style={{marginRight: '5px'}}
					>Cancel</button>
					<button onClick={handleSave} className='btn-save'>Save</button>
				</div>
			</div>
			<textarea
				className='blog-description-input'
				value={editedDescription}
				onChange={(e) => setEditedDescription(e.target.value)}
			/>
		</div>
	);
};

const BlogCard = ({id, title, description, dateCreated, blogDelete}) => {
	const [editing, setEditing] = useState(false);
	const [cardTitle, setCardTitle] = useState(title);
	const [cardDescription, setCardDescription] = useState(description);

	const handleDelete = async () => {
		try {
			const res = await api.delete(`/blog/${id}`);
			console.log(res.data);
			blogDelete(res.data);
		} catch (err) {
			console.log(err);
		};
	};

	const handleEdit = () => {
		setEditing(true);
	};

	const handleCancelEdit = () => {
		setEditing(false);
	};

	const handleSaveEdit = ({title: newTitle, description: newDescription}) => {
		setCardTitle(newTitle);
		setCardDescription(newDescription);
		setEditing(false);
	};

	return(
		<>
		{editing ? (
			<EditingCard
				id={id}
				title={cardTitle}
				description={cardDescription}
				dateCreated={dateCreated}
				onSave={handleSaveEdit}
				onCancel={handleCancelEdit}
			/>
		) : (
		<div className='blog-card'>
			<div style=
					{{width: '100%', display: 'flex', justifyContent: 'space-between'}}>
				<div style={{display: 'flex'}}>
					<h2>{cardTitle}</h2>
					<p style={{
							marginLeft: '10px',
							marginTop: '15px',
							color: '#a0a0a0'}}>Date posted: {dateCreated}</p>
				</div>
				<div>
					<button
						onClick={handleEdit}
						className='btn-save'
						style={{marginRight: '5px'}}
					>Edit</button>
					<button onClick={handleDelete} className='btn-save'>Delete</button>
				</div>
			</div>
			<p>{cardDescription}</p>
		</div>)}
		</>
	);
};

export default BlogCard;
