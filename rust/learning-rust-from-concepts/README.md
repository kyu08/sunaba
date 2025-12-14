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
- 次のように書くとtraitにデフォルト実装を持たせてその実装を利用したり、明示的に上書きしたりすることができたりする。
    ```rust
    trait PrintHello {
        fn print_hello(&self) {
            println!("Hello!!")
        }
    }

    struct Test1;
    struct Test2;

    impl PrintHello for Test1 {}

    impl PrintHello for Test2 {
        fn print_hello(&self) {
            println!("Hello from not default impl")
        }
    }

    fn main() {
        let test1 = Test1;
        test1.print_hello();

        let test2 = Test2;
        test2.print_hello();
    }
    ```
- 本筋ではないが、以下のようなTrait定義で出てくる`Rhs`はRight Hand Side(: 右側)の略らしい。ここでは`Add`オペレータの1つの目の(左側の)オペランドが`self`で、それに対して2つ目の(右側の)オペランドが第二仮引数なのでRight Sideなのだと理解した。
    ```rust
    pub trait Add<Rhs = Self> {
        type Output;
    
        // Required method
        fn add(self, rhs: Rhs) -> Self::Output;
    }
    ```
- トレイトを実装する型を関数の返り値として指定したい場合は2つの方法がある
    - `fn div4(x: i32) -> Result<(), Box<dyn MyError>>`のように`Box<dyn xxx>`と書くパターン
    - `fn generic_return_type() -> impl std::fmt::Display`のように`impl Trait`構文を利用するパターン
        - 関数を返却する関数を定義する際に便利らしい(後続の章で出てくるのでここでは深堀りしない)
- `Deref`トレイトが実装されている型が関数やメソッドの引数に指定された場合、メソッド`deref()`を使ってその関数やメソッドが要求する型の引数に変換できるのであれば、それが自動的になされる仕組みがある。
    - これは「デリファレンス型強制」(Deref coercion)と呼ばれる。
