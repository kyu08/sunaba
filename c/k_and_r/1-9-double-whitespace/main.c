#include <stdio.h>

int main(void) {
  int current, previous;

  while ((current = getchar()) != EOF) {
    if (current == ' ' && previous == ' ') {
      continue;
    }
    putchar(current);
    previous = current;
  }
}
