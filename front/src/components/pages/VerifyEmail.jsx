import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { post } from "../api/method";
import { useAuth } from "../../context/AuthContext";

const VerifyEmail = () => {
  const [verificationCode, setVerificationCode] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { email, login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setError("");
      const response = await post("/auth/verify-email", {
        email: email,
        verification_code: verificationCode,
      });

      if (response) {
        if (response.access_token && response.refresh_token) {
          login(response);
        } else {
          navigate("/auth/login", { replace: true });
          return;
        }
        navigate("/dashboard", { replace: true });
      }
    } catch (err) {
      console.error("Verification error:", err);
      if (err.response) {
        setError(err.response.data?.error || "Verification failed.");
      } else if (err.request) {
        setError(
          "Server connection failed. Please check your internet connection."
        );
      } else {
        setError("Verification failed. Please try again.");
      }
    }
  };

  return (
    <div>
      <div>
        <div>
          <h2>Email Verification</h2>
          <p>Please enter the verification code sent to {email}.</p>
        </div>
        {error && <div>{error}</div>}
        <form onSubmit={handleSubmit}>
          <div>
            <input
              type="text"
              placeholder="Verification Code"
              value={verificationCode}
              onChange={(e) => setVerificationCode(e.target.value)}
              required
            />
          </div>
          <div>
            <button type="submit">Verify and Continue</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default VerifyEmail;
