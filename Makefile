up: # 実行環境の構築&起動
	docker compose up -d
	make shell
.PHONY: up

down: # コンテナ停止
	docker compose down
.PHONY: down

shell: # コンテナ内シェル起動
	docker compose exec yaylib bash
.PHONY: shell

test: # テスト実行
	poetry run python -m unittest discover -s tests -p "test_*.py"
.PHONY: test

doc: # ドキュメントの生成
	poetry run sphinx-build -M html ./docs ./docs/_build
.PHONY: doc
