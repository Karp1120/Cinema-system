<template>
  <div id="app">
    <nav v-if="isAuthenticated">
      <router-link to="/films">🎬 Фильмы</router-link>
      <router-link to="/cinemas">🏢 Кинотеатры</router-link>
      <router-link to="/sessions">🕒 Сеансы</router-link>
      <router-link to="/reports">📊 Отчёты</router-link>
      <a href="#" @click="logout">🚪 Выйти</a>
    </nav>

    <router-view />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const isAuthenticated = computed(() => {
  return !!localStorage.getItem('user')
})

const logout = () => {
  localStorage.removeItem('user')
  router.push('/login')
}
</script>
