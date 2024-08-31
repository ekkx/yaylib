:description: yaylib における非同期処理について解説します。

非同期処理
===========

.. rst-class:: lead

   ここでは同期処理、非同期処理を行う方法について解説します 💁‍♀️

----

yaylib では、直接 ``yaylib.Client()`` のメソッドにアクセスすることで同期処理を行います。

.. code-block:: python
    :caption: main.py

    import yaylib

    client = yaylib.Client()

    # 同期処理
    client.login('your_email', 'your_password')
    client.create_post('Hello with yaylib!')
    client.get_user(93)

非同期処理を行うには、単純に各 API メンバ経由で同様のメソッドを実行するだけです。

.. code-block:: python
    :caption: main.py

    import yaylib

    client = yaylib.Client()

    # 非同期処理
    client.auth.login('your_email', 'your_password')
    client.post.create_post('Hello with yaylib!')
    client.user.get_user(93)

.. caution::

    上記の非同期コードは動作しません。

    Python で非同期関数を実行するためには、直前に ``await`` を付与し、 さらに ``async`` で囲まれた関数内で実行する必要があります。

    .. code-block:: python
        :caption: main.py

        import asyncio
        import yaylib

        client = yaylib.Client()

        async def do_something():
            await client.auth.login('your_email', 'your_password')
            await client.post.create_post('Hello with yaylib!')
            await client.user.get_user(93)

        asyncio.run(do_something())

    `参照 <https://docs.python.org/ja/3.12/library/asyncio-task.html>`_
