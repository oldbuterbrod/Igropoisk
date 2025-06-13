import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
import './RegisterPage.css'

const RegisterPage = () => {
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: ""
  });
  const [error, setError] = useState(null);
  const navigate = useNavigate();

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
      const res = await fetch("http://localhost:8080/api/auth/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.error || "Ошибка регистрации");
        return;
      }



      if (data.token) {
        localStorage.setItem("token", data.token);
        navigate("/profile");
      } else {
       
        await loginAfterRegister();
      }
    } catch (err) {
      setError("Сетевая ошибка");
    }
  }

  
  async function loginAfterRegister() {
    try {
      const res = await fetch("http://localhost:8080/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          email: formData.email,
          password: formData.password
        })
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data.error || "Ошибка логина");
        return;
      }

      localStorage.setItem("token", data.token);
      navigate("/profile");
    } catch {
      setError("Сетевая ошибка при логине");
    }
  }

  return (
    <div className="register-container">
  <h2>Регистрация</h2>
  <form onSubmit={handleSubmit}>
    <label>
      Имя пользователя:
      <input
        type="text"
        name="username"
        value={formData.username}
        onChange={handleChange}
        required
        minLength={3}
      />
    </label>
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
    <label>
  Повторите пароль:
  <input
    type="password"
    name="confirmPassword"
    value={formData.confirmPassword || ""}
    onChange={handleChange}
    required
    minLength={6}
  />
</label>


    {error && (
      <div className="error-message">{error}</div>
    )}

    <button type="submit">Зарегистрироваться</button>
  </form>

  <Link to="/login">
    <button type="button" className="link-button">
      У меня уже есть аккаунт!
    </button>
  </Link>
</div>

  );
}

export default RegisterPage;