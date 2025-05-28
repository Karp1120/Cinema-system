<template>
  <div>
    <h2>Список сеансов</h2>
    <button @click="loadSessions">Обновить</button>

    <table>
      <thead>
        <tr>
          <th>Дата</th>
          <th>Время</th>
          <th>ID кинотеатра</th>
          <th>ID фильма</th>
          <th>Зал</th>
          <th>Цена</th>
          <th>Свободные места</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="session in sessions" :key="session.id">
          <td>{{ session.date }}</td>
          <td>{{ session.time }}</td>
          <td>{{ session.cinema_id }}</td>
          <td>{{ session.film_id }}</td>
          <td>{{ session.hall_number }}</td>
          <td>{{ session.price }}</td>
          <td>{{ session.available_seats }}</td>
        </tr>
      </tbody>
    </table>

    <h3>Добавить сеанс</h3>
    <form @submit.prevent="createSession">
      <input type="date" v-model="newSession.date" required />
      <input type="time" v-model="newSession.time" required />
      <input type="number" v-model.number="newSession.cinema_id" placeholder="ID кинотеатра" required />
      <input type="number" v-model.number="newSession.film_id" placeholder="ID фильма" required />
      <input type="number" v-model.number="newSession.hall_number" placeholder="Номер зала" required />
      <input type="number" v-model.number="newSession.price" placeholder="Цена" required />
      <input type="number" v-model.number="newSession.available_seats" placeholder="Свободных мест" required />
      <button type="submit">Добавить</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const sessions = ref([])
const newSession = ref({
  date: '',
  time: '',
  cinema_id: 1,
  film_id: 1,
  hall_number: 1,
  price: 400,
  available_seats: 50
})

const loadSessions = () => {
  fetch('http://localhost:8080/api/sessions')
    .then(res => res.json())
    .then(data => (sessions.value = data))
    .catch(() => alert('Ошибка при загрузке сеансов'))
}

const createSession = () => {
  fetch('http://localhost:8080/api/sessions', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(newSession.value)
  })
    .then(res => {
      if (res.status === 409) {
        alert('Ошибка: в этом зале уже есть сеанс на эту дату и время')
        return
      }
      if (!res.ok) throw new Error()
      loadSessions()
      newSession.value = {
        date: '',
        time: '',
        cinema_id: 1,
        film_id: 1,
        hall_number: 1,
        price: 400,
        available_seats: 50
      }
    })
    .catch(() => alert('Ошибка при добавлении сеанса'))
}

onMounted(loadSessions)
</script>

<style scoped>
table {
  border-collapse: collapse;
  margin-top: 10px;
}
td, th {
  border: 1px solid #ccc;
  padding: 5px 10px;
}
form {
  margin-top: 20px;
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
</style>
