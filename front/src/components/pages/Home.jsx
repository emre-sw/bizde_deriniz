import { Link } from "react-router-dom";
import React from "react";

const Home = () => {
  return (
    <div>
      <h1>Home</h1>
      <p>Welcome to the home page</p>
      <div>
        <Link to="/auth/login">Login</Link>
        <br />
        <Link to="/auth/register">Register</Link>
      </div>
    </div>
  );
};

export default Home;
