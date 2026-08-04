import { useEffect, useState } from "react";
import PageHeader from "../components/PageHeader";
import './Page.css';
import { useAuth } from "../context/AuthContext";
import api from '../config/axios';
import RecommendedProfileCard from "../components/RecommendedProfileCard";
import ShowBlogCard from "../components/ShowBlogCard";

const HomePage = () => {
	//mental note ovo su blogovi koji su pisali drugi ljudi
	const [blogs, setBlogs] = useState([]);
	const {user} = useAuth();
	const [recommended, setRecommended] = useState([]);
	useEffect(() => {
		const fetchBlogs = async () => {
			try {
				const res = await api.get(`/blog/forUser/${user.id}`);
				console.log(res.data);
				setBlogs(res.data);
			} catch(err) {
				console.log(err);
			}
		};

		const fetchRecommended = async () => {
			try {
				const res = await api.get(`/followers/users/${user.id}/recommendations`);
				console.log(res.data);
				setRecommended(res.data);
			} catch(err) {
				console.log(err);
			}
		};
		fetchBlogs();
		fetchRecommended();
	}, []);

	return(
		<div>
			<PageHeader />
			<div className="homepage-content">
				<div className="blog-content">
					{blogs.map(b => (
						<ShowBlogCard title={b.title} description={b.description}/>
					))}
				</div>
				<div className="recommended-users">
					<div style={{height: '50px', display: 'flex'}}>
						<h2 style={{margin: '0 auto'}}>Recommended</h2>
					</div>
					{recommended.map(r => (
						<RecommendedProfileCard recommendedUser={r} key={r.id} />
					))}
				</div>
			</div>
		</div>
	);
}

export default HomePage;

