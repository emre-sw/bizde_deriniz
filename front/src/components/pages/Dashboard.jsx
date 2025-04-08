import { useAuth } from "../../context/AuthContext";
import { useNavigate, Navigate } from "react-router-dom";
import { post } from "../api/method";
import { useEffect } from "react";

const Dashboard = () => {
  const { email, logout, accessToken, refreshToken } = useAuth();
  const navigate = useNavigate();

  // Token kontrolü
  if (!accessToken || !refreshToken) {
    return <Navigate to="/auth/login" replace />;
  }

  const handleLogout = async () => {
    try {
      await post("/auth/logout", {
        access_token: accessToken,
        refresh_token: refreshToken,
      });

      logout();
      navigate("/auth/login", { replace: true });
    } catch (error) {
      console.error("Logout error:", error);
      logout();
      navigate("/auth/login", { replace: true });
    }
  };

  return (
    <div>
      <nav>
        <div>
          <div>
            <div>
              <h1>Dashboard</h1>
            </div>
            <div>
              <span>{email}</span>
              <button onClick={handleLogout}>Logout</button>
            </div>
          </div>
        </div>
      </nav>

      <main>
        <div>
          <div>
            <div>
              <h2>Hoş Geldiniz!</h2>
              <p>Dashboard içeriği burada olacak...</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
