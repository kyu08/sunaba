# understanding-rust-from-concept
コンセプトから理解するRustの写経リポジトリ

## 第4章 Rustにおける値の型
- メモリは以下の4つの領域に大きく分けられる。
    - **テキスト領域** 機械語に翻訳されたプログラムがロードされる領域
    - **静的領域** グローバル変数や静的変数に束縛された値を格納する領域
    - **スタック領域** 関数のローカル変数に束縛された値（コンパイル時にサイズが決まるもの）が置かれる領域
    - **ヒープ領域** プログラムの実行中に動的に確保されてデータが置かれる領域
- データごとの配置場所
    - 関数の中で`let`で変数に束縛する値は**スタック領域**に置かれるのが基本
    - Vec本体のデータは**ヒープ領域**に置かれ、**スタック領域**にはヒープ領域を指すポインタとサイズが置かれる。
- `Box<T>`型を使うとヒープ領域にデータを配置することができる。
    - `let box = Box::new(1);`のようにすると`i32`型の`1`というデータがヒープ領域に置かれ、その値へのポインタと関連データがスタック領域に配置され、そのスタック領域の値が`box`に束縛される。

### `Rc<T>`型
次のようなコードを書くと...
```rust
fn setdata(data_a: &mut DataA, data_b: &mut DataB, value: i32) {
    let number = Box::new(value + 1);
    data_a.number_a = Some(&number);
    data_b.number_b = Some(&number);
}
```

- 「借用された値が十分に長く生存していない」という内容のエラーになる。（`data_a`, `data_b`は`setdata`から返った後も生き続けるが、`setdata`のスコープ内で束縛された`number`は関数の実行終了時に破棄されてしまうため。）
- これを回避するためには返却する値に所有権をもたせればいいが、1つの値に対して複数の変数が所有権を持つことはできない。
- このような場面でで`Rc<T>`型を使う。（リファレンスカウントポインタ）

```rust
fn setdata(data_a: &mut DataA, data_b: &mut DataB, value: i32) {
    let number = Rc::new(value + 1);
    data_a.number_a = Some(Rc::clone(&number));
    data_b.number_b = Some(Rc::clone(&number));
}
```

`let x = Rc::new(1);`とするとヒープ領域に`(値: 1, カウンタ: 1)`のようなデータが配置される。

`let y = Rc::clone(&x);`とすると↑で配置されたヒープ領域のデータが`(値: 1, カウンタ: 2)`のように書き換わる。このとき`y`は`x`と同じヒープ領域のアドレスを指す。

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

## 第8章 Rustによるスレッド・非同期プログラミング
```rust
use std::thread::spawn;

fn main() {
    let mut v_threads = Vec::new();
    for i in 0..10 {
        let thread = spawn(move || println!("{}", i));
        v_threads.push(thread);
    }

    let _x: Vec<()> = v_threads.into_iter().map(|th| th.join().unwrap()).collect();
}
```

- `std::thread::spawn()`でスレッドが作成される。
- `spawn()`はあくまでスレッドを起動するだけなので完了をブロックするには`spawn()`が返却した値(`JoinHandle型`)に対して`join()`メソッドを実行する。
- スレッド間でデータを共有するには`Arc<T>`型を利用する。(`"A"tomically "R"eference-"C"ountered`)
- スレッド間での共有データに書き込みをするためにはmutextを利用して都度ロックする。
- 冒頭の例だと`join()`を使って別スレッドのデータをメインスレッドで処理したが、以下の例のように`channel`を使うこともできる。
    - `tx`: 送信路
    - `rx`: 受信路

```rust
use std::{sync::mpsc::channel, thread::spawn};

fn main() {
    let data = vec![1, 2, 3, 4, 5];
    let (tx, rx) = channel();

    let data_len = data.len();

    for dd in data {
        let tx = tx.clone();
        spawn(move || tx.send(dd));
    }

    for _ in 0..data_len {
        println!("{}", rx.recv().unwrap());
    }
}
```
