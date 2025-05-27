function loadSessions() {
    fetch("http://localhost:8080/api/sessions")
        .then(response => response.json())
        .then(data => {
            const tbody = document.querySelector("#sessions-table tbody");
            tbody.innerHTML = "";
            data.forEach(session => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${session.date}</td>
                    <td>${session.time}</td>
                    <td>${session.cinema_id}</td>
                    <td>${session.film_id}</td>
                    <td>${session.hall_number}</td>
                    <td>${session.price}</td>
                    <td>${session.available_seats}</td>
                `;
                tbody.appendChild(row);
            });
        })
        .catch(() => {
            alert("Не удалось загрузить сеансы.");
        });
}

function createSession(event) {
    event.preventDefault();

    const session = {
        date: document.getElementById("session-date").value,
        time: document.getElementById("session-time").value,
        cinema_id: parseInt(document.getElementById("session-cinema").value),
        film_id: parseInt(document.getElementById("session-film").value),
        hall_number: parseInt(document.getElementById("session-hall").value),
        price: parseFloat(document.getElementById("session-price").value),
        available_seats: parseInt(document.getElementById("session-free-seats").value)
    };

    fetch("http://localhost:8080/api/sessions", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(session)
    })
        .then(response => {
            if (!response.ok) throw new Error();
            alert("Сеанс успешно добавлен!");
            document.querySelector("form").reset();
            loadSessions();
        })
        .catch(() => {
            alert("Ошибка при добавлении сеанса.");
        });
}

// Загрузка при открытии страницы
window.onload = loadSessions;
