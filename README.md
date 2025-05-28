#  Cinema System

Проект для управления кинотеатрами: просмотр фильмов, расписаний сеансов, отчётов и работа с пользователями.

##  Структура проекта

cinema-system/
│
├── back/ # Бэкенд на Go
│ ├── handlers/ # Обработчики маршрутов
│ ├── models/ # Структуры данных
│ ├── main.go # Точка входа в сервер
│ └── ... # Другие файлы
│
├── cinema-client/ # Фронтенд на Vue 3 + Vite
│ ├── src/
│ │ ├── views/ # Страницы: Login, Films, Cinemas и т.д.
│ │ └── ...
│ ├── tests/ # Unit-тесты (Vitest)
│ ├── vite.config.js # Конфигурация Vite
│ └── package.json # Зависимости и скрипты


---

## !! Запуск проекта !!
 
### Бэкенд (Go)

1. Перейдите в папку `back/`:
    ```bash
    cd back
    ```

2. Убедитесь, что установлен Go и PostgreSQL.

3. Запустите сервер:
    ```bash
    go run main.go
    ```

> Сервер будет доступен по адресу: `http://localhost:8080`

---

### Фронтенд (Vue 3 + Vite)

1. Перейдите в папку `cinema-client/`:
    ```bash
    cd cinema-client
    ```

2. Установите зависимости:
    ```bash
    npm install
    ```

3. Запустите фронтенд:
    ```bash
    npm run dev
    ```

> Откроется страница по адресу: `http://localhost:5173`

---

## ✅ Unit-тесты

В проекте реализовано 11+ unit-тестов для основных компонентов (`Login.vue`, `Films.vue`, `Reports.vue`, и т.д.).

Для запуска:

```bash
npm run test
