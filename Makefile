# Docker
.PHONY: up down shell

up: # 実行環境の構築&起動
	docker compose up -d
	make shell

down: # コンテナ停止
	docker compose down

shell:
	docker compose exec yaylib bash

# Distribution
.PHONY: build clean publish

build: # パッケージのビルド
	python setup.py sdist
	python setup.py bdist_wheel

clean: # ビルドファイル削除
	rm -rf build/
	rm -rf dist/
	rm -rf yaylib.egg-info/

publish: # PYPIにパッケージのアップロード
	make build
	twine upload --repository pypi dist/*
	make clean

# Documentation
.PHONY: doc clean-doc

doc: # ドキュメントの生成
	make clean-doc
	sphinx-apidoc -f -e -o ./docs . tests/* *_test.py setup.py
	sphinx-build -M html ./docs ./docs/_build

clean-doc: # temp
	rm -rf docs/yaylib.*rst docs/modules.rst docs/_build
