:description: 好きでつながるバーチャルワールド - Yay!（イェイ）API ライブラリ 📚

モデル一覧
==========

.. rst-class:: lead

    好きでつながるバーチャルワールド - Yay!（イェイ）API ライブラリ 📚

----

**yaylib** では、サーバーからのレスポンスを Python オブジェクトとして扱えるようよう「モデル」を定義しています。これらのモデルは、レスポンス内のデータを明確化し操作を簡単にします。

.. grid:: 1
    :gutter: 2
    :padding: 0

    .. grid-item-card:: :octicon:`package` 純粋なモデル
        :link: ./models.html

        単一のデータを表した純粋なモデルです。一つのモデルは複数のモデルを保持する場合があり、互いに参照し合うことがあります。

    .. grid-item-card:: :octicon:`archive` レスポンスモデル
        :link: ./responses.html

        複数の純粋なモデルを統括するサーバーからのレスポンスモデルです。一つの処理に対して一つのレスポンスという関係性から、互いに参照し合うことはありません。


.. toctree::
    :hidden:

    models
    responses
