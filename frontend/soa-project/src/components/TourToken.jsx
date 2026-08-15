import './TourToken.css'

const TourToken = ({token}) => {
	const handleStart = async () => {
	};
	return (
		<div className='tour-token'>
			<div style={{
					width: '100%',
					display: 'flex',
					justifyContent: 'space-between'
				}}>
				<h3>{token.ID}</h3>
				<button 
					className='btn-save'
					onClick={handleStart}
					style={{marginTop: '7px'}}>Start</button>
			</div>
		</div>
	);
};


export default TourToken;
