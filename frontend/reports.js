// 1. Репертуар по кинотеатру
function getRepertoireByCinema(event) {
    event.preventDefault();
    const name = document.getElementById("cinema-name-rep").value;

    fetch(`http://localhost:8080/api/repertoire?cinema=${encodeURIComponent(name)}`)
        .then(res => res.json())
        .then(data => {
            const list = document.getElementById("repertoire-list");
            list.innerHTML = "";
            data.forEach(item => {
                const li = document.createElement("li");
                li.textContent = `${item.film_title} — ${item.date} ${item.time}`;
                list.appendChild(li);
            });
        })
        .catch(() => alert("Ошибка при получении репертуара."));
}

// 2. Кинотеатры по жанру (боевики, комедии и т.д.)
function getCinemasByGenre(genre) {
    const target = genre === "боевик" ? "action-cinemas-list" : "comedy-cinemas-list";

    fetch(`http://localhost:8080/api/cinemas-by-genre?genre=${encodeURIComponent(genre)}`)
        .then(res => res.json())
        .then(data => {
            const list = document.getElementById(target);
            list.innerHTML = "";
            data.forEach(cinema => {
                const li = document.createElement("li");
                li.textContent = `${cinema.name} (${cinema.address})`;
                list.appendChild(li);
            });
        })
        .catch(() => alert("Ошибка при получении кинотеатров по жанру."));
}

// 3. Свободные места и цена по ID сеанса
function getSeatsAndPrice(event) {
    event.preventDefault();
    const sessionId = document.getElementById("session-id").value;

    fetch(`http://localhost:8080/api/session-info?id=${sessionId}`)
        .then(res => res.json())
        .then(data => {
            const info = document.getElementById("seats-price-info");
            info.innerHTML = `Свободных мест: <strong>${data.free_seats}</strong><br>Цена билета: <strong>${data.price} ₽</strong>`;
        })
        .catch(() => alert("Ошибка при получении информации о сеансе."));
}

// 4. Фильмы по режиссёру
function getFilmsByDirector(event) {
    event.preventDefault();
    const director = document.getElementById("director-name").value;

    fetch(`http://localhost:8080/api/films-by-director?name=${encodeURIComponent(director)}`)
        .then(res => res.json())
        .then(data => {
            const list = document.getElementById("director-films-list");
            list.innerHTML = "";
            data.forEach(film => {
                const li = document.createElement("li");
                li.textContent = `${film.title} (${film.genre})`;
                list.appendChild(li);
            });
        })
        .catch(() => alert("Ошибка при получении фильмов по режиссёру."));
}
