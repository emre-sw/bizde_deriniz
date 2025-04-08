import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Register from "./components/pages/Register";
import VerifyEmail from "./components/pages/VerifyEmail";
import Dashboard from "./components/pages/Dashboard";
import Login from "./components/pages/Login";
import { AuthProvider } from "./context/AuthContext";
import Home from "./components/pages/Home";

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/auth/register" element={<Register />} />
          <Route path="/auth/login" element={<Login />} />
          <Route path="/auth/verify-email" element={<VerifyEmail />} />
          <Route path="/dashboard" element={<Dashboard />} />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
