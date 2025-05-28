<template>
  <div class="auth-container">
    <h2>{{ isLoginMode ? 'Вход' : 'Регистрация' }}</h2>

    <form @submit.prevent="handleSubmit">
      <input v-model="form.username" placeholder="Логин" required />
      <input v-model="form.password" type="password" placeholder="Пароль" required />
      <button type="submit">{{ isLoginMode ? 'Войти' : 'Зарегистрироваться' }}</button>
    </form>

    <p class="switch-mode">
      {{ isLoginMode ? 'Нет аккаунта?' : 'Уже есть аккаунт?' }}
      <a href="#" @click.prevent="toggleMode">{{ isLoginMode ? 'Зарегистрироваться' : 'Войти' }}</a>
    </p>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const isLoginMode = ref(true)

const form = ref({
  username: '',
  password: ''
})

const toggleMode = () => {
  isLoginMode.value = !isLoginMode.value
}

const handleSubmit = () => {
  const url = isLoginMode.value
    ? 'http://localhost:8080/api/login'
    : 'http://localhost:8080/api/register'

  fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(form.value)
  })
    .then(res => {
      if (!res.ok) throw new Error()
      return res.json()
    })
    .then(data => {
      localStorage.setItem('user', form.value.username)
      router.push('/films') // Перенаправляем после входа
    })
    .catch(() => {
      alert('Ошибка авторизации / регистрации')
    })
}
</script>

<style scoped>
.auth-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 30px;
  background: #f3f3f3;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
input, button {
  padding: 10px;
  font-size: 1rem;
}
button {
  background: #409eff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
button:hover {
  background: #2d8cf0;
}
.switch-mode {
  margin-top: 15px;
  text-align: center;
}
.switch-mode a {
  color: #409eff;
  cursor: pointer;
  margin-left: 5px;
}
</style>
