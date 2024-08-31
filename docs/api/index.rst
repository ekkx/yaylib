:description: 好きでつながるバーチャルワールド - Yay!（イェイ）API ライブラリ 📚

API
======

.. rst-class:: lead

    好きでつながるバーチャルワールド - Yay!（イェイ）API ライブラリ 📚

----

**yaylib** で利用可能な API 関数一覧を紹介します。

API とは Yay! のサーバーにアクセスするための入り口となるインターフェースのことです。yaylib
では、その API に Python 関数を通して簡単にアクセスすることができます。

API 関数にアクセスするには、以下のように ``yaylib.Client`` クラスを使用します。

.. code-block:: python
    :caption: main.py

    import yaylib

    client = yaylib.Client()

    # タイムラインを取得する API にアクセス
    client.get_timeline()

また、``yaylib.Client.[APIの種類]`` のようにアクセスを行うことで非同期処理が可能です。

.. code-block:: python
    :caption: main.py

    import yaylib

    client = yaylib.Client()

    # API の種類を指定してタイムラインを取得（非同期）
    client.post.get_timeline()

非同期処理については `こちらを参照 <../async.html>`_ してください。


API 関数
---------

種類別に分類した yaylib で利用可能な API 関数一覧です。

新しく追加してほしい機能などがありましたら、 `GitHub <https://github.com/ekkx/yaylib/issues/new>`_ や
`Discord <https://discord.gg/MEuBfNtqRN>`_ にてお伝えください。

.. grid:: 1 1 1 2
    :gutter: 2
    :padding: 0

    .. grid-item-card:: :octicon:`shield-lock` 認証
        :link: ./auth.html

    .. grid-item-card:: :octicon:`person` ユーザー
        :link: ./user.html

    .. grid-item-card:: :octicon:`note` 投稿
        :link: ./post.html

    .. grid-item-card:: :octicon:`list-unordered` スレッド
        :link: ./thread.html

    .. grid-item-card:: :octicon:`bell` 通知
        :link: ./notification.html

    .. grid-item-card:: :octicon:`discussion-closed` チャット
        :link: ./chat.html

    .. grid-item-card:: :octicon:`read` レター
        :link: ./review.html

    .. grid-item-card:: :octicon:`apps` サークル
        :link: ./group.html

    .. grid-item-card:: :octicon:`unmute` 通話
        :link: ./call.html

    .. grid-item-card:: :octicon:`book` その他
        :link: ./misc.html

例外クラス
----------

yaylib では、実行中に発生しうる例外（エラー）を細かく定義しています。これにより利用者は特定の状況に応じて細かく処理を分岐することが可能です。

.. grid:: 1
    :gutter: 2
    :padding: 0

    .. grid-item-card:: クライアント例外
        :link: ./errors.html

        ``ClientError`` を継承するサーバーレスポンスの ``error_code`` に基づいたエラークラスです。

    .. grid-item-card:: HTTP 例外
        :link: ./errors.html

        ``HTTPError`` を継承するクライアント例外に該当しない、HTTP 通信上のエラークラスです。


.. toctree::
    :hidden:

    auth
    user
    post
    thread
    notification
    chat
    review
    group
    call
    misc
    errors
