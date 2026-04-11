#include <stdio.h>
int main(void) {
  int i;
  char s[10];
  char c;
  int lim = 10;
  int len = 0;

  for (i = 0; i < lim - 1; ++i) {
    c = getchar();
    if (c == '\n') {
      break;
    }
    if (c == EOF) {
      break;
    }

    s[i] = c;
    len++;
  }

  for (int i = 0; i < len; i++) {
    printf("%c", s[i]);
  }
}
