// Загрузка списка фильмов
function loadFilms() {
    fetch("http://localhost:8080/api/films")
        .then(response => response.json())
        .then(data => {
            const tbody = document.querySelector("#films-table tbody");
            tbody.innerHTML = "";
            data.forEach(film => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${film.title}</td>
                    <td>${film.director}</td>
                    <td>${film.genre}</td>
                    <td>${film.studio}</td>
                    <td><button onclick="deleteFilm(${film.id})">🗑</button></td>
                `;
                tbody.appendChild(row);
            });
        })
        .catch(() => {
            alert("Не удалось загрузить список фильмов.");
        });
}

// Отправка нового фильма на сервер
function createFilm(event) {
    event.preventDefault();

    const film = {
        title: document.getElementById("title").value,
        director: document.getElementById("director").value,
        genre: document.getElementById("genre").value,
        studio: document.getElementById("studio").value,
        operator: "",
        actors: "",
        duration: 120,
        age_limit: 12
    };

    fetch("http://localhost:8080/api/films", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(film)
    })
        .then(response => {
            if (!response.ok) throw new Error();
            alert("Фильм успешно добавлен!");
            document.querySelector("form").reset();
            loadFilms();
        })
        .catch(() => {
            alert("Ошибка при добавлении фильма.");
        });
}

function deleteFilm(id) {
    if (!confirm("Удалить этот фильм?")) return;

    fetch(`http://localhost:8080/api/films/${id}`, {
        method: "DELETE"
    })
        .then(response => {
            if (!response.ok) throw new Error();
            alert("Фильм удалён.");
            loadFilms();
        })
        .catch(() => {
            alert("Не удалось удалить фильм.");
        });
}
