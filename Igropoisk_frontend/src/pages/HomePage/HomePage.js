import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import './HomePage.css';


const GameGrid = ({ title, games, loading, error }) => (
  <div className="section">
    <h2 className="section-title">{title}</h2>
    {loading && <p className="loading">Загрузка...</p>}
    {error && <p className="error">{error}</p>}
    {!loading && !error && games.length === 0 && <p className="no-data">Нет данных</p>}
    {!loading && !error && games.length > 0 && (
      <ul className="grid-list">
        {games.map((game) => (
          <li key={game.id} className="grid-item">
            <Link to={`/games/${game.id}`}>
              <img
                src={game.cover_url}
                alt={game.title}
                className="grid-image"
              />
            </Link>
          </li>
        ))}
      </ul>
    )}
  </div>
);

const TopGamesList = ({ games, loading, error }) => (
  <div className="section">
    <h2 className="section-title"> Топ 5 игр по версии Игропоиска</h2>
    {loading && <p className="loading">Загрузка...</p>}
    {error && <p className="error">{error}</p>}
    {!loading && !error && games.length === 0 && <p className="no-data">Нет данных</p>}
    {!loading && !error && games.length > 0 && (
      <ol className="top-list">
        {games.map((game) => (
          <li key={game.id} className="top-item">
            <img src={game.cover_url} alt={game.title} className="top-image" />
            <Link to={`/games/${game.id}`} className="top-link">
              {game.title}
            </Link>
            <span className="top-score">⭐ {game.average_score.toFixed(1)}</span>
          </li>
        ))}
      </ol>
    )}
  </div>
);

const HomePage = () => {
  const [popularGames, setPopularGames] = useState([]);
  const [topGames, setTopGames] = useState([]);
  const [recentGames, setRecentGames] = useState([]);
  const [upcomingGames, setUpcomingGames] = useState([]);

  const [loadingPopular, setLoadingPopular] = useState(true);
  const [loadingTop, setLoadingTop] = useState(true);
  const [loadingRecent, setLoadingRecent] = useState(true);
  const [loadingUpcoming, setLoadingUpcoming] = useState(true);

  const [errorPopular, setErrorPopular] = useState(null);
  const [errorTop, setErrorTop] = useState(null);
  const [errorRecent, setErrorRecent] = useState(null);
  const [errorUpcoming, setErrorUpcoming] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/games/popular")
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при загрузке популярных игр"))
      .then(setPopularGames)
      .catch(setErrorPopular)
      .finally(() => setLoadingPopular(false));

    fetch("http://localhost:8080/api/games/top")
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при загрузке топ игр"))
      .then(setTopGames)
      .catch(setErrorTop)
      .finally(() => setLoadingTop(false));

    fetch("http://localhost:8080/api/games/recent")
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при загрузке новинок"))
      .then(setRecentGames)
      .catch(setErrorRecent)
      .finally(() => setLoadingRecent(false));

    fetch("http://localhost:8080/api/games/upcoming")
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при загрузке предстоящих релизов"))
      .then(setUpcomingGames)
      .catch(setErrorUpcoming)
      .finally(() => setLoadingUpcoming(false));
  }, []);

   return (
    <div className="homepage">
      <GameGrid
        title="Популярные игры"
        games={popularGames}
        loading={loadingPopular}
        error={errorPopular}
      />
      <TopGamesList
        games={topGames}
        loading={loadingTop}
        error={errorTop}
      />
      <GameGrid
        title="Новинки"
        games={recentGames}
        loading={loadingRecent}
        error={errorRecent}
      />
      <GameGrid
        title="Предстоящие релизы"
        games={upcomingGames}
        loading={loadingUpcoming}
        error={errorUpcoming}
      />
    </div>
  );
};

export default HomePage;
