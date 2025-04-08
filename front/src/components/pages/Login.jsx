import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { post } from "../api/method";
import { useAuth } from "../../context/AuthContext";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { login, setEmail: setAuthEmail } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setError("");
      const response = await post("/auth/login", {
        email,
        password,
      });

      if (response && response.access_token) {
        setAuthEmail(email);
        login(response);
        navigate("/dashboard");
      }
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        setError(err.response.data.error);
      } else {
        setError("Login failed. Please try again.");
      }
    }
  };

  return (
    <div>
      <div>
        <div>
          <h2>Login</h2>
        </div>
        {error && <div>{error}</div>}
        <form onSubmit={handleSubmit}>
          <div>
            <div>
              <input
                type="email"
                required
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                autoComplete="username"
              />
            </div>
            <div>
              <input
                type="password"
                required
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                autoComplete="current-password"
              />
            </div>
          </div>

          <div>
            <button type="submit">Giriş Yap</button>
          </div>
        </form>

        <div>
          <p>Do you have an account? </p>
          <Link to="/auth/register">
            <span>Register</span>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Login;
