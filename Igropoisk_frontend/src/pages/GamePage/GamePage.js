import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import "./GamePage.css";

const formatDate = (dateString) => {
  const options = { day: "2-digit", month: "long", year: "numeric" };
  return new Date(dateString).toLocaleDateString("ru-RU", options);
};

const GamePage = () => {
  const { id } = useParams();
  const [game, setGame] = useState(null);
  const [genre, setGenre] = useState(null);
  const [similarGames, setSimilarGames] = useState([]);
  const [gameRating, setGameRating] = useState(null);
  const [reviews, setReviews] = useState([]);
  const [reviewText, setReviewText] = useState("");
  const [reviewMessage, setReviewMessage] = useState("");
  const [userScore, setUserScore] = useState(0);
  const [ratingMessage, setRatingMessage] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const isReleased = (releaseDate) => new Date(releaseDate) <= new Date();

  useEffect(() => {
    setLoading(true);
    Promise.all([
      fetch(`http://localhost:8080/api/games/${id}`).then(res => {
        if (!res.ok) throw new Error("Не удалось загрузить игру");
        return res.json();
      }),
      fetch(`http://localhost:8080/api/games/${id}/genres`).then(res => {
        if (!res.ok) throw new Error("Не удалось загрузить жанр");
        return res.json();
      }),
      fetch(`http://localhost:8080/api/games/${id}/similar`).then(res => {
        if (!res.ok) throw new Error("Не удалось загрузить похожие игры");
        return res.json();
      }),
      fetch(`http://localhost:8080/api/games/${id}/rating`)
        .then(res => res.ok ? res.json() : { average: 0 })
        .catch(() => ({ average: 0 })),
      fetch(`http://localhost:8080/api/games/${id}/reviews`)
        .then(res => res.ok ? res.json() : [])
        .catch(() => [])
    ])
      .then(([gameData, genreData, similarData, ratingData, reviewsData]) => {
        const filteredSimilar = similarData.filter(g => isReleased(g.release_date));
        setGame(gameData);
        setGenre(genreData);
        setSimilarGames(filteredSimilar);
        setGameRating(ratingData);
        setReviews(Array.isArray(reviewsData) ? reviewsData : []);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [id]);

  const handleRatingSubmit = async () => {
    const token = localStorage.getItem("token");
    if (!token) {
      setRatingMessage("Вы не авторизованы");
      return;
    }

    try {
      const res = await fetch(`http://localhost:8080/api/games/${id}/rate`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ score: userScore }),
      });

      if (!res.ok) throw new Error("Не удалось отправить оценку");

      setRatingMessage("Оценка успешно отправлена");
    } catch (err) {
      setRatingMessage(err.message || "Ошибка при отправке");
    }
  };

  const handleReviewSubmit = async () => {
    const token = localStorage.getItem("token");
    if (!token) {
      setReviewMessage("Вы не авторизованы");
      return;
    }

    try {
      const res = await fetch(`http://localhost:8080/api/games/${id}/reviews`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ text: reviewText }),
      });

      if (!res.ok) throw new Error("Не удалось отправить отзыв");

      const newReview = await res.json();
      setReviews(prev => [newReview, ...prev]);
      setReviewText("");
      setReviewMessage("Отзыв отправлен!");
    } catch (err) {
      setReviewMessage(err.message || "Ошибка при отправке");
    }
  };

  const handleAddToFavorites = async () => {
    const token = localStorage.getItem("token");
    if (!token) {
      alert("Вы не авторизованы");
      return;
    }

    try {
      const res = await fetch(`http://localhost:8080/api/games/${id}/favorite`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) throw new Error("Не удалось добавить в избранное");

      alert("Игра добавлена в избранное!");
    } catch (err) {
      alert(err.message || "Ошибка при добавлении в избранное");
    }
  };

  if (loading) return <p className="game-page">Загрузка...</p>;
  if (error) return <p className="game-page" style={{ color: "red" }}>Ошибка: {error}</p>;
  if (!game) return <p className="game-page">Игра не найдена</p>;

  const released = isReleased(game.release_date);

  return (
    <div className="game-page">
      <div className="game-header">
        <img src={game.cover_url} alt={game.title} className="game-cover" />
        <div className="game-details">
          <h1 className="game-title">{game.title}</h1>

          {released && (
            <p className="game-rating">
              <strong>{gameRating.average.toFixed(1)}</strong>
            </p>
          )}

          {released && (
            <div className="rating-form">
              <label>
                <strong>Поставить оценку:</strong>{" "}
                <select value={userScore} onChange={(e) => setUserScore(Number(e.target.value))}>
                  <option value={0}>Выберите оценку</option>
                  {[...Array(10)].map((_, i) => (
                    <option key={i + 1} value={i + 1}>{i + 1}</option>
                  ))}
                </select>
              </label>
              <button
                onClick={handleRatingSubmit}
                disabled={userScore === 0}
                className="rating-button"
              >
                Поставить
              </button>
              {ratingMessage && <p>{ratingMessage}</p>}
            </div>
          )}

          {released ? (
            <button onClick={handleAddToFavorites} className="favorite-button">
              Добавить в любимое
            </button>
          ) : (
            <p className="future-release">Игра выйдет в будущем</p>
          )}

          <p className="game-description">{game.description}</p>
          <p><strong>Разработчик:</strong> {game.developer}</p>
          <p><strong>Издатель:</strong> {game.publisher}</p>
          <p><strong>Платформы:</strong> {game.platforms}</p>
          <p><strong>Жанр:</strong> {genre?.genres?.join(", ")}</p>
          <p><strong>Дата релиза:</strong> {formatDate(game.release_date)}</p>
        </div>
      </div>

      {released && (
        <div>
          <h2>Похожие игры</h2>
          {similarGames.length > 0 ? (
            <ul className="similar-games">
              {similarGames.map((similar) => (
                <li key={similar.id} className="similar-game-item">
                  <Link to={`/games/${similar.id}`}>
                    <img src={similar.cover_url} alt={similar.title} className="similar-game-img" />
                  </Link>
                </li>
              ))}
            </ul>
          ) : (
            <p>Нет похожих игр</p>
          )}
        </div>
      )}

      {released && (
        <div>
          <h2>Отзывы</h2>
          {reviews.length > 0 ? (
            <ul className="review-list">
              {reviews.map((review, index) => (
                <li key={index} className="review-item">
                  <strong>{review.user?.username}:</strong> {review.text}
                </li>
              ))}
            </ul>
          ) : (
            <p>Отзывы отсутствуют</p>
          )}

          <div className="review-form">
            <textarea
              value={reviewText}
              onChange={(e) => setReviewText(e.target.value)}
              placeholder="Оставьте отзыв..."
              rows={3}
              className="review-textarea"
            />
            <button
              onClick={handleReviewSubmit}
              className="review-button"
              disabled={!reviewText.trim()}
            >
              Отправить отзыв
            </button>
            {reviewMessage && <p>{reviewMessage}</p>}
          </div>
        </div>
      )}
    </div>
  );
};

export default GamePage;
