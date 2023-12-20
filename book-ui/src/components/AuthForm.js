import React, { useState } from "react";
import axios from "axios";
import "../css/AuthForm.css";
import { Button } from "@mui/material";

const AuthForm = ({ type, onSuccess }) => {
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = () => {
    const apiEndpoint = type === "login" ? "/login" : "/signup";

    // Make a POST request to the appropriate endpoint
    axios
      .post(`http://localhost:8080${apiEndpoint}`, formData)
      .then((response) => {
        console.log(`${type} successful:`, response.data);

        // If successful, call the onSuccess callback
        if (onSuccess) {
          onSuccess(response.data);
        }
      })
      .catch((error) => {
        console.error(`Error ${type}:`, error.response.data.error);
        switch (error.response.data.error) {
          case "Invalid Password": {
            const errorMessage = document.createElement("p");
            errorMessage.textContent = "Invalid password";
            errorMessage.style.color = "red";
            document.querySelector(".auth-container").appendChild(errorMessage);
            break;
          }
          case "User does not exist": {
            const errorMessage = document.createElement("p");
            errorMessage.textContent = "Invalid email";
            errorMessage.style.color = "red";
            document.querySelector(".auth-container").appendChild(errorMessage);
            break;
          }
          case "User already exists": {
            const errorMessage = document.createElement("p");
            errorMessage.textContent = "User already exists, please login";
            errorMessage.style.color = "red";
            document.querySelector(".auth-container").appendChild(errorMessage);
            break;
          }
          default: {
          }
        }
      });
  };
  return (
    <div className="auth-container">
      <h2>{type === "login" ? "Login" : "Signup"}</h2>
      {type === "signup" && (
        <>
          <label>Name:</label>
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleChange}
          />
        </>
      )}
      <label>Email:</label>
      <input
        type="text"
        name="email"
        value={formData.email}
        onChange={handleChange}
      />
      <label>Password:</label>
      <input
        type="password"
        name="password"
        value={formData.password}
        onChange={handleChange}
      />
      <Button variant="contained" color="secondary" onClick={handleSubmit}>
        {type === "login" ? "Login" : "Signup"}
      </Button>
    </div>
  );
};

export default AuthForm;
