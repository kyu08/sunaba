use ratatui::text::{Line, Span};
use syntect::easy::HighlightLines;
use syntect::highlighting::ThemeSet;
use syntect::parsing::SyntaxSet;
use syntect::util::LinesWithEndings;
use syntect_tui::into_span;

fn main() {
    const EXAMPLE: &str = "
pub struct Wow {
    hi: u64
}
fn blah() -> u64 {}
";

    let ps = SyntaxSet::load_defaults_newlines();
    let ts = ThemeSet::load_defaults();
    let syntax = ps.find_syntax_by_extension("rs").unwrap();
    let mut h = HighlightLines::new(syntax, &ts.themes["base16-ocean.dark"]);

    for line in LinesWithEndings::from(EXAMPLE) {
        // LinesWithEndings enables use of newlines mode
        let spans: Vec<Span> = h
            .highlight_line(line, &ps)
            .unwrap()
            .into_iter()
            .filter_map(|segment| into_span(segment).ok())
            .collect();

        let line = Line::from(spans);
        println!("{:?}", line);
    }
}
