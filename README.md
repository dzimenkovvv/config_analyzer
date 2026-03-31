# Config Analyzer (Go)

Утилита для анализа конфигурационных файлов веб-приложений и выявления потенциально опасных настроек.

Поддерживает:
- CLI запуск
- HTTP REST API
- gRPC API
- анализ YAML и JSON конфигураций

---

# 📌 Возможности

Утилита выявляет следующие проблемы:

- включён debug режим
- пароль в открытом виде
- использование `0.0.0.0` без ограничений
- отключённый TLS
- использование устаревших/небезопасных алгоритмов
- слишком широкие настройки доступа

---

# ⚙️ Формат запуска (CLI)

## 📥 Базовый запуск

```bash
go run cmd/cli/main.go test_config.yaml
go run cmd/cli/main.go test_config.json
```
## 📥 Из stdin

```bash
cat test_config.yaml | go run cmd/cli/main.go --stdin
```

## 🚫 Поведение при ошибках

По умолчанию:
  если найдены проблемы → программа завершится с ошибкой (exit code ≠ 0)
Отключить:
```bash
-silent (go run cmd/cli/main.go -s test_config.yaml)
```
# 🌐 HTTP API

## ▶️ Запуск сервера
```bash
go run cmd/http/main.go
```

## 📡 Endpoint

POST /analyze

## 📥 Request

JSON  
```bash
{
  "config": {
    "log": {
      "level": "debug"
    },
    "storage": {
      "digest-algorithm": "MD5"
    }
  }
}
```

## 📥 Response

JSON
```bash
[
    {
        "Severity": "Low",
        "Message": "logging level is set to debug",
        "Recommendation": "use info level or higher in production"
    },
    {
        "Severity": "High",
        "Message": "weak cryptographic algorithm used: MD5",
        "Recommendation": "use stronger algorithms like SHA-256 or bcrypt"
    }
]
```

# 🔌 gRPC API

## ▶️ Запуск сервера
```bash
go run cmd/grpc/main.go
```

## 📡 Service 
```bash
service AnalyzeService
```

## 📥 Method
```bash
Analyze
```
# 📁 Поддерживаемые форматы
- YAML  
- JSON

# 🔎 Реализованные правила
- логирование в debug режиме
- пароль в открытом виде
- использование `0.0.0.0` без ограничений
- отключенный TLS
- использование устаревших или небезопасных алгоритмов
