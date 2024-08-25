:description: yaylib のインストール方法を解説します。

インストール
============

.. rst-class:: lead

   ここでは **yaylib** のインストール方法を解説します 💁‍♀️

----

yaylib は `好きでつながるバーチャルワールド - Yay! (イェイ) <https://yay.space>`_ の API クライアントです。

あらゆる操作の自動化や、ボットの開発が可能になります。

pip install
-----------

yaylib は以下のコマンドから簡単にインストールすることが可能です。

.. code-block:: shell

    pip install yaylib

.. hint::
    yaylib の動作条件は Python 3.10 以上からになります。

Getting Started
---------------

インストールが完了したら、さっそくログインしてみましょう！

以下のコードをコピーして、``your_email`` ``your_password`` をご自身のものに変更し実行してみましょう。

.. code-block:: python
    :caption: main.py

    import yaylib

    client = yaylib.Client()
    client.login('your_email', 'your_password')

    client.create_post('Hello with yaylib!')

アプリを開き、「**Hello with yaylib!**」と投稿されていれば準備完了です！
