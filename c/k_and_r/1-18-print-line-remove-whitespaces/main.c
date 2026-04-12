#include <stdio.h>

int my_getline(char s[], int lim);
void copy(char to[], char from[]);

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
    // 空白行だったらスキップ
    if (len == 1 && line[0] == '\n') {
      continue;
    }

    // 行の後のblank, tabをtrim
    // len - 1は改行文字なのでその1つ前から走査したい
    for (int i = len - 2; i >= 0; i--) {
      if (line[i] == ' ' || line[i] == '\t') {
        line[i] = '\0';
        len--;
      } else {
        // 空白でもタブ文字でもない文字に出会ったらそこで終了
        break;
      }
    }

    copy(lines[line_index], line);
    lens[line_index] = len;
    line_index++;
  }

  for (int i = 0; i < line_index; i++) {
    printf("[%2d]%s\n", lens[i], lines[i]);
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
