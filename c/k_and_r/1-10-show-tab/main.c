#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int current;

  while ((current = getchar()) != EOF) {
    if (current == '\t') {
      current = 't';
      printf("\\");
      // ただし、terminal側がbackspaceを処理してしまうのでここには入らない
    } else if (current == '\b') {
      current = 'b';
      printf("\\");
    } else if (current == '\\') {
      current = '\\';
      printf("\\");
    }
    putchar(current);
  }
}
