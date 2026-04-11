#include <stdio.h>

int my_getline(char s[], int lim);
void copy(char to[], char from[]);
void reverse(char arr[], int len);

/* 最も長い入力行を印字する */
int main() {
  const int MAX_LINE = 1000;
  const int MAX_LINE_COUNT = 20;

  int len = 0;                           /* 現在行の長さ */
  int line_index = 0;                    /* いままで入力されてきた行数 */
  char line[MAX_LINE];                   /* 現在の入力行 */
  int lens[MAX_LINE_COUNT];              // 各行の長さ
  char lines[MAX_LINE_COUNT][MAX_LINE];  // 入力された各行

  while ((len = my_getline(line, MAX_LINE)) > 0) {
    reverse(line, len);
    copy(lines[line_index], line);
    lens[line_index] = len;
    line_index++;
  }

  for (int i = 0; i < line_index; i++) {
    printf("%s", lines[i]);
  }

  return 0;
}

/* getline: s に行を入れ、長さを返す */
int my_getline(char s[], int lim) {
  int c, i;

  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i) s[i] = c;
  if (c == '\n') {
    s[i] = c;
    ++i;
  }
  s[i] = '\0';
  return i;
}

/* copy: from を to にコピー ; to は十分大きいと仮定 */
void copy(char to[], char from[]) {
  int i;

  i = 0;
  while ((to[i] = from[i]) != '\0') ++i;
}

void reverse(char str[], int len) {
  // lenが偶数の時
  int max = (len / 2) - 1;
  // lenが奇数の時
  // len7:3
  // len9:4
  // len13:6
  if ((len % 2) == 1) {
    max = (len - 1) / 2;
  }

  for (int i = 0; i <= max; i++) {
    // iとlen-1-iをswapする
    char temp = str[len - 1 - i];
    str[len - 1 - i] = str[i];
    str[i] = temp;
  }
}
