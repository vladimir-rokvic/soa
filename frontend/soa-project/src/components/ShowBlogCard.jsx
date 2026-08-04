import './CreateBlogCard.css'

const ShowBlogCard = ({title, description}) => {
	return(
		<div className='blog-card'>
			<div style={{width: '100%', display: 'flex', justifyContent: 'space-between'}}>
				<h2>{title}</h2>
			</div>
			<p>{description}</p>
		</div>
	);
};

export default ShowBlogCard;
