// Получить список фильмов и отобразить в таблице
function loadFilms() {
    fetch('http://localhost:8080/api/films')
        .then(res => res.json())
        .then(data => {
            const tbody = document.querySelector('#films-table tbody');
            tbody.innerHTML = '';
            data.forEach(film => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${film.title}</td>
                    <td>${film.director}</td>
                    <td>${film.genre}</td>
                    <td>${film.studio}</td>
                `;
                tbody.appendChild(row);
            });
        })
        .catch(error => {
            console.error('Ошибка загрузки фильмов:', error);
            alert('Не удалось загрузить фильмы.');
        });
}

// Отправить новый фильм на сервер
function createFilm(event) {
    event.preventDefault();

    const film = {
        title: document.getElementById('title').value,
        director: document.getElementById('director').value,
        genre: document.getElementById('genre').value,
        studio: document.getElementById('studio').value,
        operator: '',
        actors: '',
        duration: 120,
        age_limit: 12
    };

    fetch('http://localhost:8080/api/films', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(film)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Ошибка при добавлении фильма');
            }
            return response.text();
        })
        .then(() => {
            alert('Фильм добавлен!');
            document.querySelector('form').reset();
            loadFilms();
        })
        .catch(error => {
            console.error(error);
            alert('Не удалось добавить фильм.');
        });
}

function showTab(id) {
    document.querySelectorAll('.tab-content').forEach(el => el.classList.remove('active'));
    document.querySelectorAll('.tab').forEach(el => el.classList.remove('active'));

    document.getElementById(id).classList.add('tab-content', 'active');
    document.querySelector(`.tab[onclick="showTab('${id}')"]`).classList.add('active');
}
