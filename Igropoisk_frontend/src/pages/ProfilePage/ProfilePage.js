import { useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import './ProfilePage.css'
const ProfilePage = () => {
  const [user, setUser] = useState(null);

  const [ratings, setRatings] = useState({});
  const [ratingsError, setRatingsError] = useState(null);
  const [loadingRatings, setLoadingRatings] = useState(false);

  const [reviews, setReviews] = useState([]);
  const [reviewsError, setReviewsError] = useState(null);
  const [loadingReviews, setLoadingReviews] = useState(false);

  const [favorites, setFavorites] = useState([]);
  const [loadingFavorites, setLoadingFavorites] = useState(false);
  const [favoritesError, setFavoritesError] = useState(null);

  const [recommendations, setRecommendations] = useState([]);
  const [loadingRecommendations, setLoadingRecommendations] = useState(false);
  const [recommendationsError, setRecommendationsError] = useState(null);

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      navigate("/register");
      return;
    }

    fetch("http://localhost:8080/api/auth/me", {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    })
      .then((res) => {
        if (res.status === 401) {
          localStorage.removeItem("token");
          navigate("/login");
          return null;
        }
        return res.json();
      })
      .then((data) => {
        if (data) setUser(data);
      })
      .catch(() => {
        localStorage.removeItem("token");
        navigate("/login");
      });
  }, [navigate]);

  useEffect(() => {
    if (!user) return;
    const token = localStorage.getItem("token");
    if (!token) return;

    setLoadingRatings(true);
    fetch("http://localhost:8080/api/user/ratings", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при получении рейтингов"))
      .then((data) => setRatings(data || {}))
      .catch((err) => setRatingsError(err))
      .finally(() => setLoadingRatings(false));
  }, [user]);

  useEffect(() => {
    if (!user) return;
    const token = localStorage.getItem("token");
    if (!token) return;

    setLoadingReviews(true);
    fetch("http://localhost:8080/api/user/reviews", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при получении отзывов"))
      .then((data) => setReviews(data || []))
      .catch((err) => setReviewsError(err))
      .finally(() => setLoadingReviews(false));
  }, [user]);

  useEffect(() => {
    if (!user) return;
    const token = localStorage.getItem("token");
    if (!token) return;

    setLoadingFavorites(true);
    fetch("http://localhost:8080/api/user/favorites", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при получении любимых игр"))
      .then((data) => setFavorites(data || []))
      .catch((err) => setFavoritesError(err))
      .finally(() => setLoadingFavorites(false));
  }, [user]);

  useEffect(() => {
    if (!user) return;
    const token = localStorage.getItem("token");
    if (!token) return;

    setLoadingRecommendations(true);
    fetch("http://localhost:8080/api/games/recommendations", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => res.ok ? res.json() : Promise.reject("Ошибка при получении рекомендаций"))
      .then((data) => setRecommendations(data || []))
      .catch((err) => setRecommendationsError(err))
      .finally(() => setLoadingRecommendations(false));
  }, [user]);

  const handleRemoveFavorite = async (gameId) => {
    const token = localStorage.getItem("token");
    if (!token) return;

    try {
      const res = await fetch(`http://localhost:8080/api/games/${gameId}/favorite`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) throw new Error("Не удалось удалить из избранного");

      setFavorites((prev) => prev.filter((game) => game.id !== gameId));
    } catch (err) {
      alert(err.message || "Ошибка при удалении из избранного");
    }
  };

  const handleDeleteRating = async (gameId) => {
    const token = localStorage.getItem("token");
    if (!token) return;

    if (!window.confirm("Вы уверены, что хотите удалить эту оценку?")) return;

    try {
      const res = await fetch(`http://localhost:8080/api/games/${gameId}/rating`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) throw new Error("Не удалось удалить оценку");

      setRatings((prev) => {
        const newRatings = { ...prev };
        delete newRatings[gameId];
        return newRatings;
      });
    } catch (err) {
      alert(err.message || "Ошибка при удалении оценки");
    }
  };

  const handleDeleteReview = async (gameId) => {
    const token = localStorage.getItem("token");
    if (!token) return;

    if (!window.confirm("Вы уверены, что хотите удалить этот отзыв?")) return;

    try {
      const res = await fetch(`http://localhost:8080/api/games/${gameId}/review`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) throw new Error("Не удалось удалить отзыв");

      setReviews((prev) => prev.filter((review) => review.gameId !== gameId));
    } catch (err) {
      alert(err.message || "Ошибка при удалении отзыва");
    }
  };
  const handleLogout = () => {
  localStorage.removeItem("token");
  navigate("/login");
};


  return (
    <div className="profile-page">
      <h1>Профиль</h1>
      {user ? (
        <>
          <p>Имя пользователя: <strong>{user.username}</strong></p>
          <p>Email: <strong>{user.email}</strong></p>
          <div className="logout-container">
          <button className="logout-button" onClick={handleLogout}>Выйти из профиля</button>
         </div>

          <div className="profile-section">
            <h2>Мои оценки</h2>
            {loadingRatings && <p>Загрузка оценок...</p>}
            {ratingsError && <p style={{ color: "red" }}>{ratingsError}</p>}
            {ratings && Object.keys(ratings).length === 0 && <p>Оценок нет</p>}
            {ratings && Object.entries(ratings).length > 0 && (
              <ul className="profile-list">
  {Object.entries(ratings).map(([gameId, ratingInfo]) => (
    <li key={gameId}>
      <div>
        Игра: <strong>{ratingInfo.title}</strong>
        <p></p>Оценка: <strong>{ratingInfo.score}</strong>
      </div>
      <button className="profile-button" onClick={() => handleDeleteRating(gameId)}>Удалить</button>
    </li>
  ))}
</ul>

            )}
          </div>

          <div className="profile-section">
            <h2>Мои отзывы</h2>
            {loadingReviews && <p>Загрузка отзывов...</p>}
            {reviewsError && <p style={{ color: "red" }}>{reviewsError}</p>}
            {reviews.length === 0 && <p>Отзывов нет</p>}
            {reviews.length > 0 && (
              <ul className="profile-list">
                {reviews.map((review) => (
                  <li key={review.id}>
                    <div>
                      <strong>{review.gameTitle}</strong> — {review.text}
                      <br />
                      <small>Дата: {new Date(review.created_at).toLocaleDateString("ru-RU")}</small>
                    </div>
                    <button className="profile-button" onClick={() => handleDeleteReview(review.gameId)}>Удалить</button>
                  </li>
                ))}
              </ul>
            )}
          </div>

          <div className="profile-section">
            <h2>Любимые игры</h2>
            {loadingFavorites && <p>Загрузка...</p>}
            {favoritesError && <p style={{ color: "red" }}>{favoritesError}</p>}
            {favorites.length === 0 && <p>Нет любимых игр</p>}
            {favorites.length > 0 && (
              <ul className="profile-list">
                {favorites.map((game) => (
                  <li key={game.id}>
                    <Link className="profile-link" to={`/games/${game.id}`}><strong>{game.title}</strong></Link>
                    <button className="profile-button" onClick={() => handleRemoveFavorite(game.id)}>Удалить</button>
                  </li>
                ))}
              </ul>
            )}
          </div>

          <div className="profile-section">
            <h2>Рекомендации</h2>
            {loadingRecommendations && <p>Загрузка рекомендаций...</p>}
            {recommendationsError && <p style={{ color: "red" }}>{recommendationsError}</p>}
            {recommendations.length === 0 && <p>Нет рекомендаций</p>}
            {recommendations.length > 0 && (
              <ul className="profile-list">
                {recommendations.map((game) => (
                  <li key={game.id}>
                    <Link className="profile-link" to={`/games/${game.id}`}><strong>{game.title}</strong></Link>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </>
      ) : (
        <p>Загрузка профиля...</p>
      )}
    </div>
  );
};

export default ProfilePage;
