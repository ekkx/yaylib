# Docker
.PHONY: up down shell

up: # 実行環境の構築&起動
	docker compose up -d
	make shell

down: # コンテナ停止
	docker compose down

shell: # コンテナのシェルを起動
	docker compose exec yaylib bash

# テスト
.PHONY: test

test:
	poetry run python -m unittest discover -s tests -p "test_*.py"

# ドキュメント
.PHONY: doc clean-doc

doc: # ドキュメントの生成
	poetry run sphinx-build -M html ./docs ./docs/_build
