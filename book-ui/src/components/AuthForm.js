import React, { useState } from "react";
import axios from "axios";

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
        console.error(`Error ${type}:`, error.response.data);
      });
  };

  return (
    <div>
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
      <br />
      <label>Password:</label>
      <input
        type="password"
        name="password"
        value={formData.password}
        onChange={handleChange}
      />
      <br />
      <button onClick={handleSubmit}>
        {type === "login" ? "Login" : "Signup"}
      </button>
    </div>
  );
};

export default AuthForm;
