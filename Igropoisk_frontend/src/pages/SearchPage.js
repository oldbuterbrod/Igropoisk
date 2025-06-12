import { useState } from "react";
import { Link } from "react-router-dom";


const SearchPage = () => {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);

  const handleSearch = () => {
    if (!query.trim()) return;

    fetch(`http://localhost:8080/api/games/search?query=${encodeURIComponent(query)}`)
      .then((res) => {
        if (!res.ok) throw new Error("Ошибка при поиске");
        return res.json();
      })
      .then((data) => setResults(data))
      .catch(() => setResults([]));
  };

  return (
    <div style={{ padding: "24px", maxWidth: "800px", margin: "0 auto" }}>
      
      <div style={{ marginBottom: "20px" }}>
        <label htmlFor="searchInput" style={{ display: "block", marginBottom: "16px" }}>
          Поиск
        </label>
        <input
          id="searchInput"
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="Введите название игры: "
          style={{
            padding: "12px",
            width: "100%",
            fontSize: "16px",
            borderRadius: "8px",
            border: "1px solid #ccc",
            marginBottom: "12px",
          }}
        />
        <button
          onClick={handleSearch}
          style={{
            width: "100%",
            height: "60px",
            fontSize: "18px",
            fontWeight: "bold",
            border: "none",
            color: "white",
            cursor: "pointer",
            borderRadius: "40px",
            background: "linear-gradient(90deg, rgba(15, 240, 252, 0.733) 50%, rgba(108, 99, 255, 0.733) 100%)",
          }}
        >
          Найти
        </button>
      </div>

     <ul style={{ listStyle: "none", padding: 0 }}>
  {results.map((game, index) => (
    <li
      key={game.id}
      style={{
        marginBottom: "20px",
        borderBottom: "1px solid #eee",
        paddingBottom: "12px",
      }}
    >
      <h3 style={{ margin: 0 }}>
        {index + 1}.{" "}
        <Link to={`/games/${game.id}`} style={{ textDecoration: "none", color: "#2c3e50" }}>
          {game.title}
        </Link>
      </h3>
    </li>
  ))}
</ul>

    </div>
  );
};

export default SearchPage;
