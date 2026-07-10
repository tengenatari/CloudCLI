# CLOUD CLI 

![Go Version](https://img.shields.io/badge/Go-1.25-blue)


## Команды


Сборка утилиты

Для Windows
```bash
go build -o mws.exe ./cmd/app
```
Для Linux
```bash
go build -o mws ./cmd/app
```

Запуск всех тестов
```bash
go test -cover ./...
```

## Описание

CLI позволяющий управлять профилями юзеров

Создание профиля
```bash
profile create --name --user --project
```

Удаление профиля
```bash
profile delete --name
```

Получение информации о профиле
```bash
profile get --name
```

Список всех профилей
```bash
profile list
```

Справка
```bash
profile help
```


## Архитектура
 - **CLI layer** - Cobra
 - **Service layer** - Golang
 - **Repository layer** - Yaml

## Структура репозитория

```
CloudCLI/
├── cmd/app/               # Точка входа приложения
├── config/                # Конфигурация приложения
├── internal/              # Внутренние пакеты
│   ├── cli/               # Интерфейс командной строки
│   ├── models/            # Модели данных
│   ├── repository/        # Взаимодействие с хранилищем
│   └── services/          # Бизнес-логика
└── scripts/               # Вспомогательные скрипты
```

## Допущения
При наличии посторонних yaml файлов 
    - при создании будет писаться о наличии посторонних файлов и необходимости их переместить или удалить
    - при удалении, получении будет выводиться ошибка о том, что такого юзера не существует

Символы аргументов: name, project, user - должны соответствовать регулярному выражению `^[a-zA-Z0-9\s\-_]+$`

За консистентность данных отвечает слой repository

