import React, { useState, useEffect } from "react";
import { useNavigate, Link } from "react-router-dom";
import './LoginPage.css';


const LoginPage = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: ""
  });
  const [error, setError] = useState(null);
  const navigate = useNavigate();


  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      navigate("/profile");
    }
  }, [navigate]);

  function handleChange(e) {
    setFormData(prev => ({
      ...prev,
      [e.target.name]: e.target.value
    }));
  }

  async function handleSubmit(e) {
    e.preventDefault();
    setError(null);

    try {
      const res = await fetch("http://localhost:8080/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.error || "Ошибка авторизации");
        return;
      }

      if (data.token) {
        localStorage.setItem("token", data.token);
        navigate("/profile");
      } else {
        setError("Токен не получен");
      }
    } catch (err) {
      setError("Сетевая ошибка");
    }
  }

  return (
    <div className="login-container">
  <h2>Авторизация</h2>
  <form onSubmit={handleSubmit}>
    <label>
      Email:
      <input
        type="email"
        name="email"
        value={formData.email}
        onChange={handleChange}
        required
      />
    </label>

    <label>
      Пароль:
      <input
        type="password"
        name="password"
        value={formData.password}
        onChange={handleChange}
        required
        minLength={6}
      />
    </label>

    {error && <div className="error-message">{error}</div>}

    <button type="submit" className="login-button">Войти</button>
  </form>

  <Link to="/register">
    <button type="button" className="register-button">
      Нет аккаунта? Зарегистрироваться
    </button>
  </Link>
</div>

  );
};

export default LoginPage;
