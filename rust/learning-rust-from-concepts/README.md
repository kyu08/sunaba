# understanding-rust-from-concept
コンセプトから理解するRustの写経リポジトリ

## 第5章 Rustの抽象化プログラミング
- 型パラメータに加える「トレイトを満たしている」という条件をことを **「トレイト境界」** と呼ぶ。
    ```rust
    fn area<T: CalcArea>(x: &T) -> f64 {
        x.calc_area()
    }
    ```
- 複数のトレイト境界を指定する場合は、以下のように書く。
    ```rust
    <T: A + B + C>
    ```
- または、`where`というキーワードを使って以下のように書くこともできる。（`where`を使わない記法の方がが他の言語で見慣れてる感ある。）
    ```rust
    fn area<T>(x: &T) -> f64
        where T: CalcArea
    {
        x.calc_area()
    }
    ```
