#include <stdio.h>

int my_getline(char s[], int lim);
void copy(char to[], char from[]);
void detab(char arr[], int len);
const int TAB_STOP_UNIT = 4;

// detabする
int main() {
  const int MAX_LINE = 1000;
  const int MAX_LINE_COUNT = 20;

  int len = 0;                           /* 現在行の長さ */
  int line_index = 0;                    /* いままで入力されてきた行数 */
  char line[MAX_LINE];                   /* 現在の入力行 */
  int lens[MAX_LINE_COUNT];              // 各行の長さ
  char lines[MAX_LINE_COUNT][MAX_LINE];  // 入力された各行

  while ((len = my_getline(line, MAX_LINE)) > 0) {
    detab(line, len);
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

// tabを次のタブストップまでの空白に置換する
void detab(char str[], int len) {
  for (int i = 0; i < len; i++) {
    if (str[i] == '\t') {
      // タブをいくつのスペースに置換するか
      int num_spaces = TAB_STOP_UNIT - (i % TAB_STOP_UNIT);
      int shift = num_spaces - 1;  // タブ1文字分の枠は既にある

      // '\0'を含めてタブ以降の要素をうしろにずらす
      for (int j = len; j > i; j--) {
        str[j + shift] = str[j];
      }

      // タブをnum_spaces個のスペースに置換
      for (int k = 0; k < num_spaces; k++) {
        str[i + k] = ' ';
      }

      len += shift;
      i += num_spaces - 1;  // forの++で次の文字に進む
    }
  }
}
