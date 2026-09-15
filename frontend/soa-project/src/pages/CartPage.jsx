import { useEffect, useState } from 'react';
import { useAuth } from '../context/AuthContext';
import './Page.css'
import api from '../config/axios';
import PageHeader from '../components/PageHeader';
import ItemCard from '../components/ItemCard';

const CartPage = () => {
	const {user} = useAuth();
	const [items, setItems] = useState([]);
	const [cart, setCart] = useState(null);

	useEffect(() => {
		const fetchCart = async () => {
			try {
				const res = await api.get(`/purchase/sc/user/${user.id}`);
				console.log(res.data);
				setItems(res.data.items);
				setCart(res.data);
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

  	const handleRemove = async (itemId) => {
		try {
			const res = await api.delete(`/purchase/sc/removeItem/${itemId}`);
			console.log(res.data);
			setCart(res.data);
		} catch(err) {
			console.log(err);
		}
      	setItems(prev => prev.filter(item => item.id !== itemId));
  	};

	return(
		<>
			<PageHeader />
			<div className='cart-page'>
				<div style={{display: 'flex', justifyContent: 'space-between'}}>
					<h2 style={{marginLeft: '5px'}}>Cart price: {cart?.price}</h2>
					<button 
						className='btn-save'
						style={{width: '100px', margin: '10px'}}
						onClick={handlePurchase}
					>Purchase</button>
				</div>
				<div className='cart-page-content'>
					{items.length !== 0 && items.map(item => (
						<ItemCard 
							id={item.tour_id} 
							key={item.id}
							itemId={item.id}
							onRemove={handleRemove}/>
					))}
				</div>
			</div>
		</>
	);
};


export default CartPage;
