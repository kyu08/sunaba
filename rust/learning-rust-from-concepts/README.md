# understanding-rust-from-concept
コンセプトから理解するRustの写経リポジトリ


<!-- TODO: ここまでの学びを書く -->

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

## 第6章 ファイルやソケットの入出力
- `Read`をスコープに持ち込まないと`read()`メソッドを呼べないのは（Goだとそのシンボルやinterfaceに定義されている関数は何でも呼べるのでそれに比べると）若干面倒。
- Fileのreadとソケットの入出力が同じように書けるのはGoの`io.Reader`あたりの設計と似てるな、という感想を持った。

## 第7章 Rustの関数型プログラミング向けの機能
- Rustでは`Iterator`トレイトが実装されている型がイテレータになることができ、`next()`というメソッドを繰り返し呼ぶことで値の集まりから順番に値を取り出せる。
    - `Iterator`トレイトには70個ものメソッドがあるが、`next()`以外はデフォルトの実装が提供されているため、実質的には`next()`の実装さえあれば`Iterator`トレイトを満たすことができる。
- `next()`メソッドは返すものがなくなったら`None`を返す。panicしないのは型安全で良い。
    ```rust
    fn main() {
        let mut r = 1..3; // `1..3`は半開区間なので3は含まない。
        println!("{:?}", r.next()); // Some(1)
        println!("{:?}", r.next()); // Some(2)
        println!("{:?}", r.next()); // None
    }
    ```
- 3種類の`into_iter()`: イテレータが返すものが異なるので注意が必要。
    ||呼び方の例|イテレータが返すもの(`let vv= vec![1, 2, 3, 4];`のような定義がある前提)|
    |---|---|---|
    | ケース1 | 値 | `vv.into_iter()` | 
    | ケース2 | イミュータブルなリファレンス | `(&vv).into_iter()` | 
    | ケース3 | ミュータブルなリファレンス | `(&mut vv).into_iter()` | 
- `collect()`
    - `FromIterator`トレイトのメソッド`from_iter()`を使い、イテレータから値の列をつくることができる、
- 関数を返却する関数
    - `Fn`: 以下のように関数を返却する関数の返り値は`impl FN...`のように書く。
        ```rust
        fn func_of_func(b: i32) -> impl Fn(i32) -> i32 {
            move |x| x + b
        }
        ```
    - `FnOnce`: 所有権が最初の呼び出しで消費されてしまうクロージャ。以下のコードだと`func_of_func`を一度呼び出すとクロージャに閉じ込められたクロージャの外の変数である`p`の所有権がクロージャに移動してしまうため2回目以降の呼び出しではクロージャから`p`が見えなくなってしまう。（呼び出しが1回でもコンパイルエラーになる。）`impl Fn`の代わりに`impl FnOnce`を指定することで一度しか呼べないことが明示され、コンパイルエラーが解消できる。
        ```rust
        fn func_of_func(b: i32, p: Point) -> impl Fn(i32) -> i32 {
            move |x| {
                print_point(p);
                x + b
            }
        }
        ```
    - `FnMut`: クロージャに閉じ込められた環境の変数を変更する。
        ```rust
        fn func_of_func(b: i32) -> impl FnMut(i32) -> i32 {
            let mut count = 0;
            move |x| {
                count += 1;
                println!("count: {}", count);
                x + b
            }
        }
        ```
