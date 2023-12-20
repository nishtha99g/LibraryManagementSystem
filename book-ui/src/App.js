import React, { useState } from "react";
import BookList from "../src/components/BookList";
import AuthForm from "../src/components/AuthForm";
import BookForm from "./components/BookForm";
import { Button } from "@mui/material";
import "./App.css"; // Import the CSS file

function App() {
  const [token, setToken] = useState(""); // Token to track user authentication
  const [authType, setAuthType] = useState("login"); // Authentication type

  const handleLogin = (userToken) => {
    setToken(userToken);
  };

  const handleLogout = () => {
    setToken("");
  };

  const handleAuthTypeChange = (type) => {
    setAuthType(type);
  };

  return (
    <div className="app-container">
      {token ? (
        <div className="app-container">
          <BookList onLogout={handleLogout} />
          <BookForm />
        </div>
      ) : (
        <div align="center">
          {/* User is not authenticated, render the AuthForm component */}
          <AuthForm type={authType} onSuccess={handleLogin} />
          <div className="auth-buttons">
            <Button
              variant="contained"
              color="secondary"
              onClick={() => handleAuthTypeChange("login")}
            >
              Login
            </Button>
            <Button
              variant="contained"
              color="secondary"
              onClick={() => handleAuthTypeChange("signup")}
            >
              Signup
            </Button>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
