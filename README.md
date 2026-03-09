# loglint

Go-линтер для проверки лог-записей. Совместим с [golangci-lint](https://golangci-lint.run) через Module Plugin System.

## Правила

| Правило | Описание | Пример ошибки |
|---|---|---|
| `lowercase` | Сообщение должно начинаться со строчной буквы | `"Starting server"` |
| `english` | Только английский язык | `"запуск сервера"` |
| `specialchars` | Нет спецсимволов и эмодзи | `"failed!!!"`, `"error🚀"` |
| `sensitive` | Нет чувствительных данных | `"token: " + t` |

## Поддерживаемые логгеры

- `log/slog` — все методы: `Info`, `Error`, `Warn`, `Debug`, `Log`, `LogAttrs`, `*Context`
- `go.uber.org/zap` — методы `*zap.Logger`: `Info`, `Error`, `Warn`, `Debug`, `Fatal`, `Panic`, `DPanic`
  > `*zap.SugaredLogger` не поддерживается

## Установка и запуск

### Через go vet (рекомендуется локально)

```bash
git clone https://github.com/Davidianol/loglint
cd loglint
make build
go vet -vettool=./loglint ./...
```

### Через golangci-lint Module Plugin System

Добавь в `.custom-gcl.yml` своего проекта:

```yaml
version: v1.64.6
plugins:
  - module: github.com/Davidianol/loglint
    path: github.com/Davidianol/loglint@latest
```

Добавь в `.golangci.yml`:

```yaml
version: "2"
linters:
  default: none
  enable:
    - loglint
```

Собери кастомный бинарник и запусти:

```bash
golangci-lint custom
./custom-gcl run ./...
```

## Тесты

```bash
# Все тесты
make test

# Unit-тесты правил
make test-unit

# Интеграционные тесты (analysistest)
make test-integration
```

## Проверка на реальных проектах

```bash
make projects-check
```

Запускает линтер на [pocketbase](https://github.com/pocketbase/pocketbase) (log/slog)
и [jaeger](https://github.com/jaegertracing/jaeger) (go.uber.org/zap).

Пример вывода на pocketbase:

```
core/base.go:438:21: log message must start with a lowercase letter: "OnBootstrap hook didn't fail..."
core/base.go:438:21: log message contains forbidden char '?': "OnBootstrap hook didn't fail..."
apis/record_auth_password_reset_request.go:58:24: log message may contain sensitive data (keyword "password" found)
```

## Структура проекта

```
loglint/
├── .github/workflows/ci.yml   # CI
├── cmd/loglint/main.go        # standalone-запуск
├── internal/rules/            # реализация правил
│   ├── lowercase.go
│   ├── english.go
│   ├── specialchars.go
│   └── sensitive.go
├── testdata/src/              # тестовые данные для analysistest
├── analyzer.go                # analysis.Analyzer
├── logcall.go                 # детектор log-вызовов
├── plugin.go                  # точка входа для golangci-lint
├── .custom-gcl.yml
├── .golangci.yml
└── Makefile
```

## Использование ИИ

При разработке использовался ChatGPT / Perplexity для:
- обсуждения архитектуры детектора log-вызовов (`logcall.go`)
- отладки граничных случаев в `extractStringArg` (рекурсивная конкатенация)
- подбора тестовых кейсов для `analysistest`
- Оформления этого README.md

Все архитектурные решения, отладка реальных ошибок и итоговый код писались и проверялись самостоятельно.
