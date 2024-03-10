setup: # 仮想環境構築
	pip install pipenv
	pipenv install --dev

active: # 仮想環境の起動
	pipenv shell

clean-venv: # 仮想環境削除
	pipenv --rm

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
