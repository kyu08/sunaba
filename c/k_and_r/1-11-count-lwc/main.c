#include <stdio.h>
#include <stdlib.h>

#define IN 1
#define OUT 0

// Q. このプログラムのテストはどのようにするか。もしバグがあるとしてらそれをあばきだすにはどんな入力をするのがよいか。
// A. 下記のようなさまざまな値を入力する。
// - 空白文字だけを入力する
// - ncがオーバーフローするだけの文字数を入力する
// - 絵文字や全角空白などの考慮されていない文字を入力する(ncがおかしくなる)
int main(void) {
  int c, nl, nw, nc, state;

  state = OUT;
  nl = nw = nc = 0;
  while ((c = getchar()) != EOF) {
    ++nc;
    if (c == '\n') {
      ++nl;
    }
    if (c == ' ' || c == '\n' || c == '\t') {
      state = OUT;
    } else if (state == OUT) {
      state = IN;
      ++nw;
    }
  }

  printf("nl: %d, nw: %d, nc: %d\n", nl, nw, nc);
}
