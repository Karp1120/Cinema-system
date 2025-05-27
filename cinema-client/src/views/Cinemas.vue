<template>
  <div>
    <h2>Список кинотеатров</h2>
    <button @click="loadCinemas">Обновить</button>

    <table>
      <thead>
        <tr>
          <th>Название</th>
          <th>Адрес</th>
          <th>Категория</th>
          <th>Залов</th>
          <th>Мест</th>
          <th>Состояние</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="cinema in cinemas" :key="cinema.id">
          <td>{{ cinema.name }}</td>
          <td>{{ cinema.address }}</td>
          <td>{{ cinema.category }}</td>
          <td>{{ cinema.halls }}</td>
          <td>{{ cinema.seats }}</td>
          <td>{{ cinema.status }}</td>
        </tr>
      </tbody>
    </table>

    <h3>Добавить кинотеатр</h3>
    <form @submit.prevent="createCinema">
      <input v-model="newCinema.name" placeholder="Название" required />
      <input v-model="newCinema.address" placeholder="Адрес" required />
      <input v-model="newCinema.category" placeholder="Категория" required />
      <input v-model.number="newCinema.halls" placeholder="Залов" required />
      <input v-model.number="newCinema.seats" placeholder="Мест" required />
      <input v-model="newCinema.status" placeholder="Состояние" required />
      <button type="submit">Добавить</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const cinemas = ref([])
const newCinema = ref({
  name: '',
  address: '',
  category: '',
  halls: 1,
  seats: 100,
  status: ''
})

const loadCinemas = () => {
  fetch('http://localhost:8080/api/cinemas')
    .then(res => res.json())
    .then(data => (cinemas.value = data))
    .catch(() => alert('Ошибка при загрузке кинотеатров'))
}

const createCinema = () => {
  fetch('http://localhost:8080/api/cinemas', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(newCinema.value)
  })
    .then(() => {
      loadCinemas()
      newCinema.value = {
        name: '',
        address: '',
        category: '',
        halls: 1,
        seats: 100,
        status: ''
      }
    })
    .catch(() => alert('Ошибка при добавлении кинотеатра'))
}

onMounted(loadCinemas)
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
