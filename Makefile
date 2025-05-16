# 実行環境の構築&起動
up:
	@echo "🚀 Starting up the environment..."
	docker compose up -d
	make shell
.PHONY: up

# コンテナ停止
down:
	@echo "🛑 Stopping containers..."
	docker compose down
.PHONY: down

# コンテナ内シェル起動
shell:
	@echo "🔑 Entering the container..."
	docker compose exec yaylib bash
.PHONY: shell

# テスト実行
test:
	@echo "🧪 Running tests..."
	poetry run python -m unittest discover -s tests -p "test_*.py"
.PHONY: test

# ドキュメントの生成
doc:
	@echo "📖 Generating documentation..."
	poetry run sphinx-build -M html ./docs ./docs/_build
.PHONY: doc

# バージョンの更新 (useage: make update VERSION=1.6.0)
update: update-pyproject-version update-init-version
	@echo "✅ Updated to version $(VERSION) in pyproject.toml and yaylib/__init__.py"
.PHONY: update

# pyproject.tomlのバージョンを更新
update-pyproject-version:
	@if [ -z "$(VERSION)" ]; then \
		echo "❌ Version not specified. Usage: make update-version version=1.6.0"; \
		exit 1; \
	fi
	@awk -v v="$(VERSION)" '\
		BEGIN { in_poetry=0 } \
		/^\[tool.poetry\]/ { in_poetry=1; print; next } \
		/^\[/ { in_poetry=0 } \
		in_poetry && /^version *= *"/ { sub(/version *= *".*"/, "version = \"" v "\""); print; next } \
		{ print }' pyproject.toml > pyproject.tmp && mv pyproject.tmp pyproject.toml
.PHONY: update-pyproject-version

# yaylib/__init__.pyのバージョンを更新
update-init-version:
	@awk -v v="$(VERSION)" '\
		/^__version__ *= *"/ { sub(/__version__ *= *".*"/, "__version__ = \"" v "\""); print; next } \
		{ print }' yaylib/__init__.py > yaylib/__init__.tmp && mv yaylib/__init__.tmp yaylib/__init__.py
.PHONY: update-init-version
