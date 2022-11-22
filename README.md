[![CircleCI](https://dl.circleci.com/status-badge/img/gh/Tyrqvir/anti-brute-force/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/Tyrqvir/anti-brute-force/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/Tyrqvir/anti-brute-force)](https://goreportcard.com/report/github.com/Tyrqvir/anti-brute-force)

# Проектная работа - Сервис "Анти-брутфорс"

#### [ТЗ на разработку сервиса "Анти-брутфорс"](./docs/01-anti-bruteforce.md)

# Реализация:
Выбрал в качестве алгоритма для rate limit - lickyBucket. 

В качестве хранилища для bucket, black/white list выбрал redis как наиболее простой и быстрый способ для внедрения + это работа с памятью, что есть быстро и не нужно тратить время на переиндексацию листов.

# Описание используемой среды

- redis, 7.0.5 version.
- gui for redis
- go lang 1.19

[//]: # (- postgres, 14.4 version)

# Команды доступные на проекте
- Список доступных команд.
```bash
$ make help
```
- Запуск линтера. Если в системе нет `golangci-lint` он будет установлен.
```bash
$ make lint
```
- Запуск Unit тестов.
```bash
$ make test
```
- Разворачивание и запуск проекта в docker-compose.
```bash
$ make run
```
- Пересборка проекта в docker-compose.
```bash
$ make rebuild
```
- Остановка проекта в docker-compose.
```bash
$ make down
```
- Запуск набора команд перед commit.
```bash
$ make pre-commit
```
- Запуск кодогенерации.
```bash
$ make generate
```
- Запуск кодогенерации для DI.
```bash
$ make wire
```
- Запуск интеграционных тестов.
```bash
$ make integration-tests
```
- Запуск билда для сервиса.
```bash
$ make build
```