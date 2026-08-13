import { BrowserRouter, Route, Routes } from "react-router-dom";
import './App.css'
import HomePage from "./pages/HomePage";
import LogInPage from "./pages/LogInPage";
import { AuthProvider, useAuth } from "./context/AuthContext";
import RegisterPage from "./pages/RegisterPage";
import BlogPage from "./pages/BlogPage";
import MyProfilePage from "./pages/MyProfilePage";
import ToursPage from "./pages/ToursPage";
import AddTourPage from "./pages/AddTourPage";
import EditTourPage from "./pages/EditTourPage";
import CartPage from "./pages/CartPage";
import SimPage from "./pages/SimPage";

function AppRoutes() {
  const { user } = useAuth();
  return (
    <Routes>
      <Route path="/" element={user ? <HomePage /> : <LogInPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/blog" element={<BlogPage />} />
      <Route path="/profile" element={<MyProfilePage />} />
      <Route path="/tours" element={<ToursPage />} />
      <Route path="/tours/add" element={<AddTourPage />} />
      <Route path="/tours/:id/edit" element={<EditTourPage />} />
      <Route path="/cart" element={<CartPage />} />
      <Route path="/simulation" element={<SimPage />} />
    </Routes>
  );
}

function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <AppRoutes />
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
