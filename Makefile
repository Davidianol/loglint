.PHONY: test-unit test-integration test build vet-run \
        lint-install lint-build-plugin lint-run-plugin \
        projects-clone projects-check clean

BINARY      := loglint
CUSTOM_GCL  := ./custom-gcl
GOLANGCI    := golangci-lint

# -------- Проекты для проверки --------
# pocketbase | log/slog (Go 1.21+)
SLOG_URL := https://github.com/pocketbase/pocketbase
SLOG_DIR := /tmp/loglint-check/pocketbase

# jaeger | go.uber.org/zap
ZAP_URL  := https://github.com/jaegertracing/jaeger
ZAP_DIR  := /tmp/loglint-check/jaeger


# ---------------  Тесты ---------------

test-unit:
	@echo ">>> Unit-тесты правил"
	go test -v ./internal/rules/... ./internal/config/...

test-integration:
	@echo ">>> Интеграционные тесты (analysistest)"
	go test -v .

test: test-unit test-integration
	@echo ">>> Все тесты прошли успешно"

# ----- Standalone (основной способ) -----

build:
	@echo ">>> Собираем standalone-бинарник"
	go build -o $(BINARY) ./cmd/loglint/

vet-run: build
	@echo ">>> Запускаем loglint через go vet"
	go vet -vettool=./$(BINARY) ./...

# -------- golangci-lint custom --------

lint-install:
	@if ! command -v $(GOLANGCI) > /dev/null 2>&1; then \
		echo ">>> Устанавливаем golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo ">>> golangci-lint: $$(golangci-lint version)"; \
	fi

lint-build-plugin: lint-install
	@echo ">>> Собираем custom golangci-lint (best effort)..."
	$(GOLANGCI) custom -v || \
		( echo ">>> custom build не удался — используй make vet-run"; exit 0 )

lint-run-plugin: lint-build-plugin
	@if [ -x "$(CUSTOM_GCL)" ]; then \
		$(CUSTOM_GCL) run ./...; \
	else \
		echo ">>> custom binary не собран — используй make vet-run"; \
	fi

# --- Проверка на реальных проектах ---

projects-clone:
	@mkdir -p /tmp/loglint-check

	@if [ ! -d $(SLOG_DIR) ]; then \
		echo ">>> Клонируем pocketbase (log/slog)..."; \
		git clone --depth=1 $(SLOG_URL) $(SLOG_DIR); \
	else \
		echo ">>> pocketbase уже клонирован, пропускаем"; \
	fi

	@if [ ! -d $(ZAP_DIR) ]; then \
		echo ">>> Клонируем jaeger (go.uber.org/zap)..."; \
		git clone --depth=1 $(ZAP_URL) $(ZAP_DIR); \
	else \
		echo ">>> jaeger уже клонирован, пропускаем"; \
	fi

projects-check: build projects-clone
	@echo ""
	@echo ">>> pocketbase — log/slog"
	cd $(SLOG_DIR) && \
		go vet -vettool=$(CURDIR)/$(BINARY) ./... 2>&1 ; true

	@echo ""
	@echo ">>> jaeger — go.uber.org/zap"
	cd $(ZAP_DIR) && \
		go vet -vettool=$(CURDIR)/$(BINARY) ./... 2>&1 ; true

	@echo ""
	@echo ">>> Проверка завершена"

# -------------- Утилиты --------------

clean:
	@echo ">>> Очистка артефактов"
	rm -f $(BINARY) $(CUSTOM_GCL)
	rm -rf /tmp/loglint-check
