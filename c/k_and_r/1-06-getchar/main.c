#include <stdio.h>

int main(void) {
  int c;
  while (1) {
    c = getchar();
    int ret = c;

    printf("ret != EOF: %d\n\n", ret != EOF);
    if (ret == EOF) {
      break;
    }
    printf("\n");
  }
}
