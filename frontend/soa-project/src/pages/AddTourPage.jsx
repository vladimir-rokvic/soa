import { useState } from 'react';
import PageHeader from '../components/PageHeader'
import api from '../config/axios';
import './Page.css'
import { useAuth } from '../context/AuthContext';
import { MapContainer, Marker, Polyline, TileLayer, useMapEvents } from 'react-leaflet';

const MapClickHandler = ({start, end, setStart, setEnd}) => {
	useMapEvents({
		click(e) {
			if(!start) {
				setStart(e.latlng);
			} else if(!end){
				setEnd(e.latlng);
			}
		},
	});
};

const AddTourPage = () => {
	const [title, setTitle] = useState('');
	const [description, setDescription] = useState('');
	const [difficulty, setDifficulty] = useState('');
	const [tag, setTag] = useState('');
	const [tags, setTags] = useState([]);
	const [startPoint, setStartPoint] = useState(null);
	const [endPoint, setEndPoint] = useState(null);
	const {user} = useAuth();
	const handleSave = async () => {
		const body = {
			title: title,
			description: description,
			difficulty: difficulty,
			tags: tags,
			author_id: user.id,
			startPos: startPoint,
			endPos: endPoint
		};

		try {
			const res = await api.post('/tours/', body);
			console.log(res.data);
		} catch(err) {
			console.log(err);
		};
	};
	const handleAdd = () => {
	
		setTags(prev => [...prev, tag]);
		setTag('');
	};
	return(
		<>
			<PageHeader />
			<div className='add-tour-page'>
				<div style={{display: 'flex', justifyContent: 'right'}}>
					<button className='btn-save' onClick={handleSave}>Save</button>
				</div>
				<div className='add-tour-content'>
					<div className='majak'>
						<input 
							placeholder='Enter title here'
							value={title}
							onChange={(e) => setTitle(e.target.value)}
						/>
						<textarea 
							placeholder='Enter description here'
							value={description}
							onChange={(e) => setDescription(e.target.value)}
						/>
						<div>
    						<select className='select-box'
									value={difficulty} 
									onChange={e => setDifficulty(e.target.value)}>
    						  <option value="">-- Select a difficulty --</option>
    						  <option value="Easy">Easy</option>
    						  <option value="Medium">Medium</option>
    						  <option value="Hard">Hard</option>
    						</select>
						</div>
						<div className='add-tour-tags'>
							<input
								placeholder='Enter tag here'
								value={tag}
								onChange={e => setTag(e.target.value)}
							/>
							<button className='btn-save' onClick={handleAdd}>+Add</button>
							<div className='tags-container'>
								{tags?.length !== 0 && tags.map((t, i) => (
									<>
										<p key={i}>{t}</p>
									</>
								))}
							</div>
						</div>
					</div>

					<div className='tour-map'>
						<MapContainer
							center={[45.2671, 19.8335]}
							zoom={13}
							style={{width: '100%', height: '100%'}}
						>
            				<TileLayer
            				    attribution='&copy; OpenStreetMap contributors'
            				    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            				/>
							<MapClickHandler 
								start={startPoint}
								end={endPoint}
								setStart={setStartPoint}
								setEnd={setEndPoint}
							/>
							{startPoint && (
								<Marker
									position={startPoint}
									draggable={true}
									eventHandlers={{
										dragend: (e) => {
											setStartPoint(e.target.getLatLng());
										}
									}}
								/>
							)}
							{endPoint && (
								<Marker
									position={endPoint}
									draggable={true}
									eventHandlers={{
										dragend: (e) => {
											setEndPoint(e.target.getLatLng());
										}
									}}
								/>
							)}
							{startPoint && endPoint && (
								<Polyline positions={[startPoint, endPoint]} />
							)}
						</MapContainer>
					</div>
				</div>
			</div>
		</>
	);
};


export default AddTourPage;
