#include <stdio.h>
#include <stdlib.h>

#define IN 1
#define OUT 0

int main(void) {
  int c, state;

  state = OUT;
  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\n' || c == '\t') {
      // 単語が終わったので改行する
      if (state == IN) {
        printf("\n");
      }
      state = OUT;
    } else {
      putchar(c);
      state = IN;
    }
  }
}
