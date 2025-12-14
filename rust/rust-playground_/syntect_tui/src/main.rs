use ratatui::text::{Line, Span};
use syntect::easy::HighlightLines;
use syntect::highlighting::ThemeSet;
use syntect::parsing::SyntaxSet;
use syntect::util::LinesWithEndings;
use syntect_tui::{self, translate_colour, translate_font_style}

fn main() {
    let ps = SyntaxSet::load_defaults_newlines();
    let ts = ThemeSet::load_defaults();
    let syntax = ps.find_syntax_by_extension("rs").unwrap();
    let mut h = HighlightLines::new(syntax, &ts.themes["base16-ocean.dark"]);
    let s = "pub struct Wow { hi: u64 }\nfn blah() -> u64 {}";
    for line in LinesWithEndings::from(s) {
        // LinesWithEndings enables use of newlines mode
        let line_spans: Vec<Span> = h
            .highlight_line(line, &ps)
            .unwrap()
            .into_iter()
            .filter_map(|(style, content)| {
                Ok(ratatui::text::Span::styled(
                    String::from(content),
                    translate_style(style),
                ))
            })
            .collect();
        let spans = Line::from(line_spans);
        print!("{:?}", spans);
    }
}

pub fn translate_style(
    syntect_style: syntect::highlighting::Style,
) -> ratatui::style::Style {
    ratatui::style::Style {
        fg: translate_colour(syntect_style.foreground),
        bg: translate_colour(syntect_style.background),
        underline_color: translate_colour(syntect_style.foreground),
        add_modifier: translate_font_style(syntect_style.font_style).unwrap(),
        sub_modifier: ratatui::style::Modifier::empty(),
    }
}

