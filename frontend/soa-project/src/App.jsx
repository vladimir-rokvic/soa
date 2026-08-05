import { BrowserRouter, Route, Routes } from "react-router-dom";
import './App.css'
import HomePage from "./pages/HomePage";
import LogInPage from "./pages/LogInPage";
import { AuthProvider, useAuth } from "./context/AuthContext";
import RegisterPage from "./pages/RegisterPage";
import BlogPage from "./pages/BlogPage";
import MyProfilePage from "./pages/MyProfilePage";
import ToursPage from "./pages/ToursPage";

function AppRoutes() {
  const { user } = useAuth();
  return (
    <Routes>
      <Route path="/" element={user ? <HomePage /> : <LogInPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/blog" element={<BlogPage />} />
      <Route path="/profile" element={<MyProfilePage />} />
      <Route path="/tours" element={<ToursPage />} />
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
