import { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';
import './Page.css'
import api from '../config/axios';
import PageHeader from '../components/PageHeader';

const CartPage = () => {
	const {user} = useAuth();
	const [items, setItems] = useState([]);

	useEffect(() => {
		const fetchCart = async () => {
			try {
				const res = await api.get(`/purchase/sc/user/${user.id}`);
				console.log(res.data);
				setItems(res.data);
			} catch (err) {
				console.log(err);
			}
		};
		fetchCart();
	}, []);

	const handlePurchase = async () => {
		try {
			const res = await api.post(`/purchase/sc/user/${user.id}`, {});
			console.log(res.data);
		} catch(err) {
			console.log(err);
		}
	};

	return(
		<>
			<PageHeader />
			<div className='cart-page'>
				<div style={{display: 'flex', justifyContent: 'right'}}>
					<button 
						className='btn-save'
						style={{width: '100px', margin: '10px'}}
						onClick={handlePurchase}
					>Purchase</button>
				</div>
				<div className='cart-page-content'>
				</div>
			</div>
		</>
	);
};


export default CartPage;
