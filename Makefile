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
	make clean-doc
	poetry run sphinx-apidoc -f -e -o ./docs . tests/* *_test.py setup.py
	poetry run sphinx-build -M html ./docs ./docs/_build

clean-doc: # temp
	rm -rf docs/yaylib.*rst docs/modules.rst docs/_build
