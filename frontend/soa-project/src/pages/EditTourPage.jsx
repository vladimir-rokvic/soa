import { useParams, useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';
import PageHeader from '../components/PageHeader'
import api from '../config/axios';
import './Page.css'
import { MapContainer, Marker, Polyline, Popup, TileLayer, useMapEvents }
from 'react-leaflet';

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

const PointPopup = ({children, p, initialTitle, initialDescription, initialImage, setPoint, setPointFile}) => {
	const [title, setTitle] = useState(initialTitle || '');
	const [description, setDescription] = useState(initialDescription || '');
	const [file, setFile] = useState(initialImage || null);
	const [rawFile, setRawFile] = useState(null);

	const handleFileChange = (e) => {
		if(e.target.files) {
			setFile(URL.createObjectURL(e.target.files[0]));
			setRawFile(e.target.files[0]);
		}
	};
	const handleSave = () => {
		const point = {
			title: title,
			description: description,
			lat: p.lat,
			lng: p.lng
		};
		setPoint(point);
		setPointFile(rawFile);
	};
	return(
		<Popup className='add-tour-popup'>
			<div className='add-tour-popup-content'>
				<div style={{display: 'flex', marginBottom: '5px'}}>
					<h3>{children}</h3>
					<button onClick={handleSave}>Save</button>
				</div>
				<div className='add-tour-popup-metadata'>
					<div style={{display: 'flex', flexDirection: 'column'}}>
						<div className='popup-image'>
							{file && <img src={file} />}
						</div>
						<input
							accept="image/png, image/jpeg" 
							type='file' 
							onChange={handleFileChange}
							id="file-upload"
						/>
						<label 
							style={{cursor: 'pointer', margin: '0 auto'}}
							htmlFor='file-upload'>Choose</label>
					</div>
					<div style={{
							display: 'flex',
							flexDirection: 'column',
							marginLeft: '10px'
					}}>
						<input 
							placeholder='Enter title here'
							value={title}
							onChange={e => setTitle(e.target.value)}
						/>
						<textarea 
							placeholder='Enter description here'
							value={description}
							onChange={e => setDescription(e.target.value)}
						/>
					</div>
				</div>
			</div>
		</Popup>
	);
};

const EditTourPage = () => {
	const {id} = useParams();
	const navigate = useNavigate();

	const [tour, setTour] = useState(null);

	const [title, setTitle] = useState('');
	const [description, setDescription] = useState('');
	const [difficulty, setDifficulty] = useState('');
	const [tag, setTag] = useState('');
	const [tags, setTags] = useState([]);
	const [startPoint, setStartPoint] = useState(null);
	const [endPoint, setEndPoint] = useState(null);

	const [sp, setSp] = useState(null);
	const [ep, setEp] = useState(null);
	const [spFile, setSpFile] = useState(null);
	const [epFile, setEpFile] = useState(null);

	useEffect(() => {
		const fetchTourData = async () => {
			try {
				const res = await api.get(`/tours/${id}`);
				const data = res.data;
				setTour(data);
				console.log(data);

				setTitle(data.title || '');
				setDescription(data.description || '');
				setDifficulty(data.difficulty || '');
				setTags(data.tags || []);

				if (data.start_point) {
					setStartPoint({lat: data.start_point.lat, lng: data.start_point.lng});
					setSp(data.start_point);
				}
				if (data.end_point) {
					setEndPoint({lat: data.end_point.lat, lng: data.end_point.lng});
					setEp(data.end_point);
				}
			} catch(err) {
				console.log(err);
			};
		};
		fetchTourData();
	}, [id]);

	const handleSave = async () => {
		const formData = new FormData();

		const body = {
			title: title,
			description: description,
			difficulty: difficulty,
			tags: tags,
			start_point: sp,
			end_point: ep
		};

		formData.append('tour_body', JSON.stringify(body));

		if (spFile) formData.append('start_point_image', spFile);
		if (epFile) formData.append('end_point_image', epFile);

		try {
			const res = await api.put(`/tours/${id}`, formData);
			console.log(res.data);
		} catch(err) {
			console.log(err);
		};

		navigate('/tours');
	};

	const handleAdd = () => {
		setTags(prev => [...prev, tag]);
		setTag('');
	};

	if (!tour) return null;

	const mapCenter = startPoint
		? [startPoint.lat, startPoint.lng]
		: [45.2671, 19.8335];

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
									<p key={i}>{t}</p>
								))}
							</div>
						</div>
					</div>

					<div className='tour-map'>
						<MapContainer
							center={mapCenter}
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
								>
									<PointPopup 
										p={startPoint}
										initialTitle={sp?.title}
										initialDescription={sp?.description}
										initialImage={sp?.image_path ? 
										`http://localhost:8084/tours/${sp.image_path}` 
										: null}
										setPoint={setSp}
										setPointFile={setSpFile}
									>Start point</PointPopup>
								</Marker>
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
								>
									<PointPopup 
										p={endPoint}
										initialTitle={ep?.title}
										initialDescription={ep?.description}
										initialImage={ep?.image_path ? 
										`http://localhost:8084/tours/${ep.image_path}` 
										: null}
										setPoint={setEp}
										setPointFile={setEpFile}
									>End point</PointPopup>
								</Marker>
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

export default EditTourPage;
