# VK SERVICE

Проект представляет собой сервис, реализующее REST-API условного маркетплейса:

- Регистрация и авторизация.
- Создание и управление пунктами выдачи заказов.
- Размещение нового объявления.
- Отображение ленты объявлений.

## Установка и запуск

### Требования

- Docker и Docker Compose
- Go 1.24 или выше (для локальной разработки)

### Через Docker Compose

1. Клонируйте репозиторий:

```bash
git clone https://github.com/pmi-119/vk_test_project.git
```

Убедитесь что у вас установлен docker и docker-compose

2. Запустите docker:

```bash
docker-compose up
```

3. Запустите проект

```
go run cmd/service/main.go
```

4. Сервис будет доступен по адресу: http://localhost:8080

## Модель данных

Проект использует следующую структуру базы данных:

- **user** - пользователи системы
- **product** - объявления
