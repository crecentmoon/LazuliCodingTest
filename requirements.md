# Assignment for prospective Server-Side Engineers

# Index

---

---

# 提出課題

「要件」に従って、下記の二点の対応と提出をお願いします。

---

1. 「コーディングテスト対象」に従って、処理を実装してください
    - ソースコードを提出してください (提出物1)
        - Dockerを利用してください
        - DBとアプリケーションの定義は1つのdocker-compose内に記述してください (Dockerfileは複数で問題ございません)
        - アプリケーションの実装言語は問いません
        - 提出フォーマットは問いません (e.g. Githubのリンク, ファイル etc.)
2. 「サンプルデータ」をRDBに格納するにあたって最適と思われるER図を書いてください
    - ER図を提出してください (提出物2)
        - RDBの種類は問いません
        - 提出フォーマットは問いません (e.g. パワーポイント, Google Slide etc.)

---

二次面接ではこれらに関しての説明/ディスカッションの場とさせて頂きます。

# 要件

### コーディングテスト対象:

「サンプルデータ」を用いたデータ参照APIを作成してください

---

- DBに格納したサンプルデータを元に、データ参照のAPIを2種類作成してください
    - 商品検索API
        - endpoint: /products?jan=xxx
            - method: GET
        - APIへの入力値:
            - `jan`
        - APIからの出力値 (key: value):
            - `jan` : janの値
            - `product_name` : product_nameの値
            - `attributes` : attributesの値
            - `maker` : makerの値
            - `brand` : brandの値
            - `tags_from_description` : tags_from_descriptionの値
            - `tags_from_review` : tags_from_reviewの値
    - 格納されたデータの集計値参照API
        - endpoint: /reports
            - method: GET
        - APIへの入力値:
            - 無し
        - APIからの出力値 (key: value):
            - `product_name` : 充足率 (null or 空以外の件数値/全レコード数)
            - `attributes`: 充足率 (要素が1以上の件数値/全レコード数)
            - `maker` : 充足率 (null or 空以外の件数値/全レコード数)
            - `brand` : 充足率 (null or 空以外の件数値/全レコード数)
            - `tags_from_description` : 充足率 (要素が1以上の件数値/全レコード数)
            - `tags_from_review` : 充足率 (要素が1以上の件数値/全レコード数)

---

### サンプルデータ:

- サンプルデータ (jsonl)
    - https://drive.google.com/file/d/1tVI-VSUHLvdIV1aXuFc9_bzlIpdngOkZ/view?usp=sharing
- 項目
    1. jan (JANコード, string): 商品を特定するためのキー情報
    2. product_name (商品名, string): 商品の名称
    3. attributes (属性, json): 商品の構成要素
    4. maker (メーカー, string): 商品のメーカー
    5. brand (ブランド, string): 商品のブランド
    6. tags_from_description (説明文タグ, array[string]): 商品説明文から生成されたタグ
    7. tags_from_review (レビュータグ, array[string]): レビュー/口コミから生成されたタグ
- 補足
    - 公開されているデータをクローリングして得られた商品のデータです
    - サンプルデータ内の情報は `jan` で一意になっています

### 環境:

- 環境構築にはDockerを利用してください
- DBの種類は問いませんがRDBを利用してください
- アプリケーションを実装する言語は問いません

以上

