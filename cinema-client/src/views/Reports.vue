<template>
  <div>
    <h2>Репертуар по кинотеатру</h2>
    <form @submit.prevent="getRepertoireByCinema">
      <input v-model="cinemaName" placeholder="Название кинотеатра" required />
      <button type="submit">Показать</button>
    </form>
    <ul>
      <li v-for="item in repertoire" :key="item.film_title + item.date">
        {{ item.film_title }} — {{ item.date }} {{ item.time }}
      </li>
    </ul>

    <h2>Кинотеатры по жанру</h2>
    <select v-model="selectedGenre">
      <option v-for="genre in genres" :key="genre" :value="genre">{{ genre }}</option>
    </select>
    <button @click="getCinemasByGenre">Показать</button>
    <ul>
      <li v-for="cinema in genreCinemas" :key="cinema.name">
        {{ cinema.name }} ({{ cinema.address }})
      </li>
    </ul>

    <h2>Свободные места и цена по ID сеанса</h2>
    <form @submit.prevent="getSessionInfo">
      <input v-model="sessionId" type="number" placeholder="ID сеанса" required />
      <button type="submit">Показать</button>
    </form>
    <div v-if="sessionInfo">
      Свободных мест: <strong>{{ sessionInfo.free_seats }}</strong><br />
      Цена билета: <strong>{{ sessionInfo.price }} ₽</strong>
    </div>

    <h2>Фильмы по режиссёру</h2>
    <form @submit.prevent="getFilmsByDirector">
      <input v-model="directorName" placeholder="Имя режиссёра" required />
      <button type="submit">Показать</button>
    </form>
    <ul>
      <li v-for="film in directorFilms" :key="film.title">{{ film.title }} ({{ film.genre }})</li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const cinemaName = ref('')
const repertoire = ref([])

const selectedGenre = ref('')
const genres = ref([])
const genreCinemas = ref([])

const sessionId = ref('')
const sessionInfo = ref(null)

const directorName = ref('')
const directorFilms = ref([])

// Загрузка жанров при загрузке компонента
onMounted(() => {
  fetch('http://localhost:8080/api/genres')
    .then(res => res.json())
    .then(data => {
      genres.value = data
      if (data.length > 0) selectedGenre.value = data[0]
    })
})

// Репертуар по кинотеатру
function getRepertoireByCinema() {
  fetch(`http://localhost:8080/api/repertoire?cinema=${encodeURIComponent(cinemaName.value)}`)
    .then(res => res.json())
    .then(data => (repertoire.value = data))
    .catch(() => alert('Ошибка при получении репертуара'))
}

// Кинотеатры по жанру
function getCinemasByGenre() {
  fetch(`http://localhost:8080/api/cinemas-by-genre?genre=${encodeURIComponent(selectedGenre.value)}`)
    .then(res => res.json())
    .then(data => (genreCinemas.value = data))
    .catch(() => alert('Ошибка при получении кинотеатров по жанру'))
}

// Свободные места и цена
function getSessionInfo() {
  fetch(`http://localhost:8080/api/session-info?id=${sessionId.value}`)
    .then(res => res.json())
    .then(data => (sessionInfo.value = data))
    .catch(() => alert('Ошибка при получении информации о сеансе'))
}

// Фильмы по режиссёру
function getFilmsByDirector() {
  fetch(`http://localhost:8080/api/films-by-director?name=${encodeURIComponent(directorName.value)}`)
    .then(res => res.json())
    .then(data => (directorFilms.value = data))
    .catch(() => alert('Ошибка при получении фильмов по режиссёру'))
}
</script>

<style scoped>
form {
  margin: 10px 0;
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}
ul {
  margin-bottom: 20px;
}
</style>
