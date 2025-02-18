# GoURLShortener

GoURLShortener - это простой и эффективный сервис для сокращения URL-адресов, написанный на Go. Он позволяет пользователям сокращать длинные URL-адреса, управлять ими и перенаправлять на исходные URL-адреса.

## Особенности

- **Сокращение URL**: Преобразование длинных URL-адресов в более короткие и удобные ссылки.
- **Перенаправление**: Перенаправление пользователей с коротких URL-адресов на исходные.
- **Управление URL**: Сохранение, получение и удаление сокращенных URL-адресов.
- **Логирование**: Встроенное логирование для мониторинга и отладки.
- **Управление конфигурацией**: Отдельные файлы конфигурации для локальной и производственной среды.
- **Юнит-тесты**: Комплексные юнит-тесты для обеспечения надежности сервиса.
- **Готовность к развертыванию**: Включает конфигурацию развертывания для легкой настройки.
- **Авторизация**: Защита определенных маршрутов с помощью базовой аутентификации.
- **GitHub Actions**: Мануальный деплой по тегу.

## Примеры

- ThisProject: http://185.246.118.245:8082/my1pet
- YouTube: http://185.246.118.245:8082/yt
- Amazon: http://185.246.118.245:8082/hUhha3 - alias сгенерирован автоматически

## Структура проекта

- **cmd/URLShortener/main.go**: Точка входа в приложение.
- **config/**: Содержит файлы конфигурации для разных сред.
- **deployment/**: Содержит файлы конфигурации для развертывания.
- **internal/**:
  - **config/**: Управление конфигурацией.
  - **http-server/**: Содержит код, связанный с HTTP-сервером, включая обработчики и middleware.
  - **lib/**: Утилиты, такие как обработка API ответов, логирование и генерация случайных строк.
  - **storage/**: Управление операциями хранения, включая реализацию на SQLite.
- **tests/**: Содержит юнит-тесты для проекта.

## Хендлеры и маршруты

- **Без авторизации**: Обработка перенаправления по сокращенному URL.
- **С авторизацией**: Сохранение новых URL, а также удаление существующих URL с использованием базовой аутентификации.

## Методы API

- **GET /**: Перенаправление по сокращенному URL.
- **POST /url**: Сохранение нового URL (требуется базовая аутентификация).
- **DELETE /url/{alias}**: Удаление существующего URL по алиасу (требуется базовая аутентификация).

## GitHub Actions

Этот проект включает поддержку GitHub Actions для автоматизации деплоя. Развертывание происходит вручную по тегу.

## Запуск сервера без GitHub Actions

### Необходимые условия

- Go 1.16 или выше
- SQLite3
- gcc

### Установка

1. Клонирование репозитория:

   ```sh
   git clone https://github.com/AleBal-A/GoURLShortener.git
   cd GoURLShortener
   ```

2. Установите зависимости:

   ```sh
   go mod download
   ```

   Переменные окружения:
   ```sh
   export CONFIG_PATH=/root/apps/GoURLShortener/config/prod.yaml
   ```

### Запуск приложения

1. Запустите сервер:

   ```sh
   go run cmd/URLShortener/main.go
   ```

2. Сервис будет доступен по адресу `http://0.0.0.0:8082`.

### Конфигурация

Файлы конфигурации находятся в каталоге `config/`. Вы можете изменить файлы `local.yaml` и `prod.yaml` для настройки видимости логов.


### Тестирование

Чтобы запустить функциональные тесты:

```sh
go test ./test
```
