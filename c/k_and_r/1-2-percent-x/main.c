#include <stdio.h>
int main(void) {
  // %xはstring literalらしい
  printf("%x\n", 'a');
  printf("%s\n", "foo");
  printf("%6d,%6d\n", 10, 200);
}
