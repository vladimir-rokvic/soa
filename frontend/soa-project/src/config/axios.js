import axios from "axios";

const api = axios.create({
baseURL: 'http://localhost:8000',
});

api.interceptors.request.use((config) => {
	const userStr = localStorage.getItem('user');
	if(userStr != null){
		const user = JSON.parse(userStr);
		config.headers.Authorization = `Bearer ${user.token}`;
	}

	return config;
});

export default api;
