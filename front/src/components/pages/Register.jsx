import React, { useState } from "react";
import { post } from "../api/method";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../../context/AuthContext";

const Register = () => {
  const [error, setError] = useState("");
  const { email, setEmail } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setError("");
      const password = e.target.password.value;
      const response = await post("/auth/register", {
        email: email,
        password: password,
      });
      if (response && response.message) {
        navigate("/auth/verify-email", { replace: true });
      }
    } catch (err) {
      if (err.response && err.response.data && err.response.data.error) {
        setError(err.response.data.error);
      } else {
        setError("Registration failed. Please try again.");
      }
    }
  };

  return (
    <div>
      <div>
        <h2>Create Account</h2>
      </div>
      {error && <div>{error}</div>}
      <form onSubmit={handleSubmit}>
        <div>
          <div>
            <input
              type="email"
              name="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              autoComplete="username"
              required
            />
          </div>
          <div>
            <input
              type="password"
              name="password"
              placeholder="Password"
              autoComplete="new-password"
              required
            />
          </div>
        </div>
        <div>
          <button type="submit">Register</button>
        </div>
      </form>
      <div>
        <p>Do you have an account? </p>
        <Link to="/auth/login">
          <span>Login</span>
        </Link>
      </div>
    </div>
  );
};

export default Register;
