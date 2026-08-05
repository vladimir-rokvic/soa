import { useNavigate } from "react-router-dom";
import PageHeader from "../components/PageHeader"
import './Page.css'

const GuidePage = () => {
	const navigate = useNavigate();

	return(
		<>
			<PageHeader />
			<div className="guide-page">
				<button className="btn-save"
					style={{
						margin: '10px',
						width: '115px',
						justifySelf: 'right'
					}}
					onClick={() => {navigate('/tours/add')}}
				>+ Add tour</button>
				<div className="guide-content">
				</div>
			</div>
		</>
	);
};


export default GuidePage;
