fn main() {
    let tag_name = "keyword";
    let key = Keyword::Do;
    // let i = Token::Key(key);
    println!("{}", to_xml_tag(key))
        std::any::
}

fn to_xml_tag<T: std::fmt::Debug>(value: T) -> String {
    let tag_name = std::any::type_name_of_val(&value)
        .split("::")
        .collect::<Vec<&str>>()[1]
        .to_lowercase();
    format!("<{}> {:?} </{}>", &tag_name, value, &tag_name)
}

#[derive(Debug, PartialEq, Eq)]
enum Token {
    Key(Keyword),
    IntegerConstant(u32),
    StringConstant(String),
    Identifier(String),
}

#[derive(Debug, PartialEq, Eq)]
enum Keyword {
    Class,
    Constructor,
    Function,
    Method,
    Field,
    Static,
    Var,
    Int,
    Char,
    Boolean,
    Void,
    True,
    False,
    Null,
    This,
    Let,
    Do,
    If,
    Else,
    While,
    Return,
}
