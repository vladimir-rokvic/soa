import { useEffect, useState } from 'react';
import PageHeader from '../components/PageHeader'
import './Page.css'
import { useAuth } from '../context/AuthContext';
import CreateBlogCard from '../components/CreateBlogCard';
import api from '../config/axios';
import BlogCard from '../components/BlogCard';

const BlogPage = () => {
	const {user} = useAuth();
	const [blogs, setBlogs] = useState([]);

	const blogCreated = (blog) => {
		setBlogs(prev => [...prev, blog])
	};

	const blogDeleted = (blog) => {
		setBlogs(blogs.filter(b => b.id !== blog.id));
	};

	useEffect(() => {
		const fetchBlogsForUser = async () => {
			try {
				const res = await api.get(`/blog/user/${user.id}`);
				console.log("Blogs: ", res.data);
				setBlogs(res.data);
			} catch (err) {
				console.log(err);
			}
		};
		fetchBlogsForUser();
	}, []);
	return(
		<>
			<PageHeader />
			<CreateBlogCard blogCreated={blogCreated}/>
			{blogs.length !== 0 && blogs.map((blog) => 
			<BlogCard key={blog.id} 
				id={blog.id} 
				title={blog.title} 
				description={blog.description}
				blogDelete={blogDeleted}
			/>)}
		</>
	);
};

export default BlogPage;
