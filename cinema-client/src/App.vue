<template>
  <div id="app">
    <nav>
      <router-link to="/films">🎬 Фильмы</router-link>
      <router-link to="/cinemas">🏢 Кинотеатры</router-link>
      <router-link to="/sessions">🕒 Сеансы</router-link>
      <router-link to="/reports">📊 Отчёты</router-link>
    </nav>
    <router-view />
  </div>
</template>

<style>
nav {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  background: #eef;
  font-weight: bold;
}
</style>
