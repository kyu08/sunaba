## 第1章 本書で開発するもの
- 本書で開発するアプリケーションである本の貸し借りなどを管理する蔵書管理アプリケーションの概要について説明する章。主要なエンティティは以下。
    - ユーザー
    - 蔵書
    - 蔵書の貸し借り状態
- 登場するシステムコンポーネントは以下。
    - FE: React.js
    - BE: Rust(Axum)
    - DB: PostgreSQL
    - キャッシュ: Redis

## 第2章 開発環境の構築
- タイトルの通り環境構築の手順が書かれている。

## 第3章 最小構成アプリの実装
- `anyhow::Error`は`std::error::Error`を実装する任意のエラー型（ここでは`std::net::AddrParseError`や`hyper::error::Error`）を自動的に`anyhow::Error`型に変換する。
- axumにはハンドラ間で状態を共有できる`State`という機能がある。この機能を利用すると、たとえばコネクションプールのようなグローバルな情報を、ハンドラ間で効率よく共有することができるようになる。
- sqlxには専用のテストランナーが用意されている。`#[sqlx::test]`というマクロを使用すると、テスト用の接続設定を使ってデータベースに接続するようなテストを実行できる。

## 第4章 蔵書管理サーバーアプリケーションの設計
- 各レイヤーの役割
    - `api`: 画面からの入力情報を受け取るレイヤー
    - `kernel`: 受け取った入力情報をアプリケーションが扱いやすいデータ形式に変換しつつ必要な処理をかけるレイヤー
    - `adapter`: データベースを始めとしたいわゆる永続化層にアクセスし、データを保存するレイヤー
- 各レイヤー内のコンポーネント
    - `api`レイヤー
        - `handler`: いわゆるハンドラー
    - `kernel`レイヤー
        - `model`: ドメインモデル
        - `repository`: repositoryのinterface
    - `adapter`レイヤー
        - `repository`: repositoryの実装
        - `database`, `redis`: 各データストアへのアクセスに関する処理を担う

## 第5章 蔵書管理サーバーの実装


## GoユーザーがRustでサーバーを書くときのキャッチアップが必要そうなポイント
- マクロ
- アトリビュート Goに慣れていると↓こういうのとか若干黒魔術感を感じる。
    ```rust
    #[derive(Error, Debug)]
    pub enum AppError {
        #[error("{0}")]
        UnprocessableEntity(String),
        #[error("{0}")]
        EntityNotFound(String),
        #[error("{0}")]
        ValidationError(#[from] garde::Report),
        #[error("トランザクションを実行できませんでした。")]
        TransactionError(#[source] sqlx::Error),
        #[error("データベース処理実行中にエラーが発生しました。")]
        SpecificOperationError(#[source] sqlx::Error),
        #[error("No rows affected: {0}")]
        NoRowsAffectedError(String),
        #[error("{0}")]
        KeyValueStoreError(#[from] redis::RedisError),
        #[error("{0}")]
        BcryptError(#[from] bcrypt::BcryptError),
        #[error("{0}")]
        ConvertToUuidError(#[from] uuid::Error),
        #[error("ログインに失敗しました")]
        UnauthenticatedError,
        #[error("認可情報が誤っています")]
        UnauthorizedError,
        #[error("許可されていない操作です")]
        ForbiddenOperation,
        #[error("{0}")]
        ConversionEntityError(String),
    }
    ```
- 所有権
- 参照を扱おうとするとでてくるArcとか
- 非同期処理
