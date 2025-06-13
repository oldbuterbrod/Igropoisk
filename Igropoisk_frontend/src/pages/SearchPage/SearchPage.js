import { useState } from "react";
import { Link } from "react-router-dom";
import "./SearchPage.css";

const SearchPage = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [searched, setSearched] = useState(false); // 👈

  const handleSearch = () => {
    if (!query.trim()) return;

    setSearched(true); 

    fetch(`http://localhost:8080/api/games/search?query=${encodeURIComponent(query)}`)
      .then((res) => {
        if (!res.ok) throw new Error("Ошибка при поиске");
        return res.json();
      })
      .then((data) => setResults(data))
      .catch(() => setResults([]));
  };

  return (
    <div className="search-page">
      <div style={{ marginBottom: "20px" }}>
        <label htmlFor="searchInput" className="search-label">Поиск</label>
        <input
          id="searchInput"
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="Введите название игры: "
          className="search-input"
        />
        <button onClick={handleSearch} className="search-button">Найти</button>
      </div>

      {Array.isArray(results) && results.length > 0 ? (
        <ul className="search-results">
          {results.map((game, index) => (
            <li key={game.id} style={{ animationDelay: `${index * 0.1}s` }}>
              <h3>
                {index + 1}.{" "}
                <Link to={`/games/${game.id}`}>
                  {game.title}
                </Link>
              </h3>
            </li>
          ))}
        </ul>
      ) : (
        searched && <p className="no-results">Нет результатов</p> 
      )}
    </div>
  );
};

export default SearchPage;
