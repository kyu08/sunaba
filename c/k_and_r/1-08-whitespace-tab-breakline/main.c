#include <stdio.h>

int main(void) {
  int c;

  // それぞれの登場回数
  int whitespace = 0, tab = 0, line_break = 0;
  while ((c = getchar()) != EOF) {
    switch (c) {
      case ' ':
        whitespace++;
        // NOTE: Cではデフォルトでfollow-throughされてしまうので注意
        break;
      case '\t':
        tab++;
        break;
      case '\n':
        line_break++;
        break;
    }
  }

  printf("whitespace: %d, tab: %d, line_break: %d\n", whitespace, tab, line_break);
}
