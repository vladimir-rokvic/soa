import { useEffect, useState } from 'react';
import './CreateBlogCard.css'
import { useAuth } from '../context/AuthContext';
import api from '../config/axios';

const CommentCard = ({comment}) => {
	const [user, setUser] = useState();
	useEffect(() => {
		const fetchCommenterData = async () => {
			try {
				const res = await api.get(`/users/${comment.author_id}`)
				console.log(res.data);
				setUser(res.data);
			} catch(err) {
				console.log(err);
			};
		};
		fetchCommenterData();
	}, []);

	return(
		<>
			<h3>{user?.username}</h3>
			<p>{comment.text}</p>
		</>
	);
};

const ShowBlogCard = ({blog}) => {
	const [comment, setComment] = useState('');
	const [comments, setComments] = useState(blog.comments);
	const {user} = useAuth();

	const addComment = (comment) => {
		setComments(prev => [...prev, comment]);
	};

	const handleComment = async () => {
		const body = {
			text: comment,
			author_id: user.id,
			blog_id: blog.id
		};

		try {
			const res = await api.post('/blog/comment/', body);
			console.log(res.data);
			addComment(res.data);
		} catch(err) {
			console.log(err);
		}
	};

	return(
		<div className='blog-card'>
			<div style={{width: '100%'}}>
				<h2>{blog.title}</h2>
				<p>{blog.description}</p>
			</div>
			<div className='add-comment-section'>
				<input onChange={e => setComment(e.target.value)} 
				type='text' placeholder='Enter comment here'/>
				<button onClick={handleComment}>Comment</button>
			</div>
			{comments.length !== 0 && (<div className='comment-section'>
				{comments.map(c => (<CommentCard key={c.id} comment={c} />))}
			</div>)}
		</div>
	);
};

export default ShowBlogCard;
