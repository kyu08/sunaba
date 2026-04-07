#include <stdio.h>
#include <stdlib.h>

int const IN = 1;
int const OUT = 0;
int const MAX_WORD_LENGTH = 10;

// 値更新処理を関数に切り出すときに、返り値として更新後の値を返すのではなく、
// 更新したい値のポインタを引数として受け取るデザインパターンはよく見るかも。
void update_histogram(int *histogram, int word_length) {
  if (word_length == 0) {
    return;
  } else if (1 <= word_length && word_length < MAX_WORD_LENGTH) {
    histogram[word_length - 1]++;
  } else if (MAX_WORD_LENGTH <= word_length) {
    histogram[9]++;
  }
}

void print_bar(int len) {
  for (int i = 0; i < len; i++) {
    printf("■");
  }
}

void print_histogram(int *histogram) {
  printf("-------------\n");
  for (int i = 0; i < 10; i++) {
    // 階級ごとのlabel
    if (i < MAX_WORD_LENGTH - 1) {
      printf("len == %2d: ", i + 1);
    } else {
      // i == 9のとき
      printf("len >= %2d: ", MAX_WORD_LENGTH);
    }

    // 階級ごとのデータ
    print_bar(histogram[i]);
    printf("\n");
  }
}

//
/* 入力した単語の長さをヒストグラムにしてプリントするプログラム。以下の様な感じで動く

$ make run
hoge hoge make run
-------------
len ==  1:
len ==  2:
len ==  3: ■
len ==  4: ■■■
len ==  5:
len ==  6:
len ==  7:
len ==  8:
len ==  9:
len >= 10:

 */
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
