fn my_clear(msg: &mut String) {
    msg.clear();
}

fn my_append(msg: &mut String) {
    *msg += "!";
}

fn main() {
    let mut m = "Hello".to_string();
    println!("m:{}\np: {:p}\n", m, &m);

    let mut_ref1 = &mut m;
    my_clear(mut_ref1);
    println!("m:{}\np: {:p}\n", mut_ref1, mut_ref1);
    // 参照のスコープはそのリファレンスが最後に使われたところで終わる。
    // 可変参照は同時に複数存在できないが、このケースでは、L15でmut_ref1のスコープが終わっているため新たな可変参照mut_ref2を初期化することができる。

    let mut_ref2 = &mut m;
    my_append(mut_ref2);
    println!("m:{}\np: {:p}\n", mut_ref2, mut_ref2);
}
