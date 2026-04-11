#include <math.h>
#include <stdio.h>

int htoi_single(char s) {
  // int
  if ('0' <= s && s <= '9') {
    return s - '0';
  }

  // hex小文字
  if ('a' <= s && s <= 'f') {
    return s - 'a' + 10;  // aは10なので10足している
  }
  // hex大文字
  if ('A' <= s && s <= 'F') {
    return s - 'A' + 10;  // aは10なので10足している
  }

  return -1;
}

// hex to int
// 0-9,a-f,A-Fが入力されうる。
int htoi(char s[]) {
  int len = 0;
  for (int i = 0; s[i] != '\0'; i++) {
    len++;
  }
  /* printf("%d\n", len); */

  int total = 0;
  for (int i = len - 1; i >= 0; i--) {
    // 下から何桁目か。iは文字列の先頭 == 上位桁から数えているが、計算上
    // 下から何桁目かが必要なため変換している。
    // なお、計算上便利なため0-origin。
    int digit_number = len - i - 1;
    int ii = htoi_single(s[i]);
    float p = pow(16, digit_number);
    total += ii * p;
  }

  return total;
}

int main(void) {
  char in[] = "1f";
  printf("%s -> %d\n", in, htoi(in));
}
