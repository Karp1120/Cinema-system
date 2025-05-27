<template>
  <div>
    <h2>Список фильмов</h2>
    <button @click="loadFilms">Обновить</button>

    <table>
      <thead>
        <tr>
          <th>Название</th>
          <th>Режиссёр</th>
          <th>Жанр</th>
          <th>Киностудия</th>
          <th>Удалить</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="film in films" :key="film.id">
          <td>{{ film.title }}</td>
          <td>{{ film.director }}</td>
          <td>{{ film.genre }}</td>
          <td>{{ film.studio }}</td>
          <td><button @click="deleteFilm(film.id)">🗑</button></td>
        </tr>
      </tbody>
    </table>

    <h3>Добавить фильм</h3>
    <form @submit.prevent="createFilm">
      <input v-model="newFilm.title" placeholder="Название" required />
      <input v-model="newFilm.director" placeholder="Режиссёр" required />
      <input v-model="newFilm.genre" placeholder="Жанр" required />
      <input v-model="newFilm.studio" placeholder="Киностудия" required />
      <button type="submit">Добавить</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const films = ref([])
const newFilm = ref({
  title: '',
  director: '',
  genre: '',
  studio: ''
})

const loadFilms = () => {
  fetch('http://localhost:8080/api/films')
    .then(res => res.json())
    .then(data => (films.value = data))
    .catch(() => alert('Ошибка при загрузке фильмов'))
}

const createFilm = () => {
  fetch('http://localhost:8080/api/films', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(newFilm.value)
  })
    .then(() => {
      loadFilms()
      newFilm.value = { title: '', director: '', genre: '', studio: '' }
    })
    .catch(() => alert('Ошибка при добавлении фильма'))
}

const deleteFilm = (id) => {
  fetch(`http://localhost:8080/api/films/${id}`, {
    method: 'DELETE'
  })
    .then(() => loadFilms())
    .catch(() => alert('Ошибка при удалении фильма'))
}

onMounted(loadFilms)
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
  gap: 10px;
}
</style>
