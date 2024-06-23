.PHONY:
up:
	docker compose up -d
	docker compose exec -it yaylib bash

clean-doc:
	rm -rf docs/yaylib.* docs/modules.rst docs/_build/

doc: # ドキュメントの生成
	make clean-doc
	sphinx-apidoc -f -e -o ./docs . tests/* *_test.py setup.py
	sphinx-build -b singlehtml ./docs ./docs/_build

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
