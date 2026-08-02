import api from '../config/axios';
import './CreateBlogCard.css'

const BlogCard = ({id, title, description, blogDelete}) => {
	const handleDelete = async () => {
		try {
			const res = await api.delete(`/blog/${id}`);
			console.log(res.data);
			blogDelete(res.data);
		} catch (err) {
			console.log(err);
		};
	};

	//TODO
	const handleEdit = async () => {
	};

	return(
		<div className='blog-card'>
			<div style={{width: '100%', display: 'flex', justifyContent: 'space-between'}}>
				<h2>{title}</h2>
				<div>
					<button
						onClick={handleEdit}
						className='btn-save'
						style={{marginRight: '5px'}}
					>Edit</button>
					<button onClick={handleDelete} className='btn-save'>Delete</button>
				</div>
			</div>
			<p>{description}</p>
		</div>
	);
};

export default BlogCard;
