// Загрузка списка кинотеатров
function loadCinemas() {
    fetch("http://localhost:8080/api/cinemas")
        .then(response => response.json())
        .then(data => {
            const tbody = document.querySelector("#cinemas-table tbody");
            tbody.innerHTML = "";
            data.forEach(cinema => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${cinema.name}</td>
                    <td>${cinema.address}</td>
                    <td>${cinema.category}</td>
                    <td>${cinema.halls}</td>
                    <td>${cinema.seats}</td>
                    <td>${cinema.status}</td>
                `;
                tbody.appendChild(row);
            });
        })
        .catch(() => {
            alert("Не удалось загрузить список кинотеатров.");
        });
}

// Отправка нового кинотеатра
function createCinema(event) {
    event.preventDefault();

    const cinema = {
        name: document.getElementById("cinema-name").value,
        address: document.getElementById("cinema-address").value,
        category: document.getElementById("cinema-category").value,
        halls: parseInt(document.getElementById("cinema-halls").value),
        seats: parseInt(document.getElementById("cinema-seats").value),
        status: document.getElementById("cinema-status").value
    };

    fetch("http://localhost:8080/api/cinemas", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(cinema)
    })
        .then(response => {
            if (!response.ok) throw new Error();
            alert("Кинотеатр успешно добавлен!");
            document.querySelector("form").reset();
            loadCinemas();
        })
        .catch(() => {
            alert("Ошибка при добавлении кинотеатра.");
        });
}
