import { Link } from 'react-router-dom';
import './Header.css';
import logo from '../../logo.png'

const Header = () => {
  return (
    <header>
  <div className="header-logo-title">
    <img src={logo} alt="Логотип" className="logo" />
    <h1 className="header-text">ИГРОПОИСК</h1>
  </div>

  <nav className="header-nav">
    <div className="spacer" />
    <div className="header-buttons">
      <Link to="/">Главная</Link>
      <Link to="/profile">Профиль</Link>
      <Link to="/search">Поиск</Link>
    </div>
    <div className="spacer" />
  </nav>
</header>

  );
};

export default Header;

