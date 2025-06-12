const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk3MzA2MDksInVzZXJfaWQiOjd9.ZceIj-9pQa99vHtdmdDt0njn2BcQZV4p0kQEc3UbDCQ';

fetch('http://localhost:8080/api/user/ratings', {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json',
  },
})
  .then(response => {
    if (!response.ok) {
      throw new Error('Ошибка при получении отзывов пользователя');
    }
    return response.json();
  })
  .then(data => {
    console.log('Отзывы пользователя:', data);
    // data — массив объектов отзывов с полями, например:
    // [{ id, userID, gameID, text, createdAt, gameTitle }, ...]
  })
  .catch(error => {
    console.error('Ошибка:', error);
  });
