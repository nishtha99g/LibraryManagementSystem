import React, { useState } from "react";
import BookList from "../src/components/BookList";
import AuthForm from "../src/components/AuthForm";
import BookForm from "./components/BookForm";

function App() {
  const [token, setToken] = useState(""); // Token to track user authentication
  const [authType, setAuthType] = useState("login"); // Added state for authentication type

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
        <div>
          <div className="auth-buttons">
            <button onClick={() => handleAuthTypeChange("login")}>Login</button>
            <button onClick={() => handleAuthTypeChange("signup")}>
              Signup
            </button>
          </div>
          {/* User is not authenticated, render the AuthForm component */}
          <AuthForm type={authType} onSuccess={handleLogin} />
        </div>
      )}
    </div>
  );
}

export default App;
