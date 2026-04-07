#include <stdio.h>
#include <stdlib.h>

int const IN = 1;
int const OUT = 0;
int const MAX_WORD_LENGTH = 10;

void update_histogram(int *histogram, int word_length) {
  if (word_length == 0) {
    return;
  } else if (1 <= word_length && word_length < MAX_WORD_LENGTH) {
    histogram[word_length - 1]++;
  } else if (MAX_WORD_LENGTH <= word_length) {
    histogram[9]++;
  }
}

void print_histogram(int *histogram) {
  for (int i = 0; i < 10; i++) {
    if (i < MAX_WORD_LENGTH) {
      printf("%d文字: %d\n", i + 1, histogram[i]);
      continue;
    }
    // i == 9のとき
    printf("%d文字以上: %d\n", MAX_WORD_LENGTH, histogram[i]);
  }
}

// 入力した単語の長さをヒストグラムにしてプリントするプログラム
int main(void) {
  // 入力した単語の長さを階級ごとに分ける
  int c, state = OUT, current_word_length = 0;
  // 0: 1文字の単語の数, 1: 2文字の単語の数, ..., 8: 9文字の単語の数, 9: 10文字以上の単語の数
  int histogram[10];
  for (int i = 0; i < 10; i++) {
    histogram[i] = 0;
  }

  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\n' || c == '\t') {
      // 単語が終わったので配列に追加し、カウントを0に戻す
      if (state == IN) {
        update_histogram(histogram, current_word_length);
        current_word_length = 0;
      }
      state = OUT;
    } else {
      current_word_length++;
      state = IN;
    }
  }

  // プリントする
  print_histogram(histogram);
}
