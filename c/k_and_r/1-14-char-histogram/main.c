#include <stdio.h>
#include <stdlib.h>

const int HISTOGRAM_LENGTH = 256;

void update_count(int *count, int appeared_char) {
  count[appeared_char]++;
}

void print_bar(int len) {
  for (int i = 0; i < len; i++) {
    printf("#");
  }
}

void print_histogram(int *count) {
  printf("-------------\n");
  for (int i = 0; i < HISTOGRAM_LENGTH; i++) {
    if (count[i] == 0) {
      continue;
    }

    // 階級ごとのlabel
    printf("%2c: ", i);

    // 階級ごとのデータ
    print_bar(count[i]);
    printf("\n");
  }
}

//
/* 入力した文字ごとの出現頻度をヒストグラムにしてプリントするプログラム。以下の様な感じで動く。
 * 綺麗にやるなら改行文字の表示を\n:などにすべきだが演習問題の趣旨としては今の実装で問題ないのでこのままでOK。

$ make run
hoge hoge
-------------

: #
  : #
 e: ##
 g: ##
 h: ##
 o: ##

 */
int main(void) {
  int c;
  // 各ascii文字の出現回数
  int count[HISTOGRAM_LENGTH];
  for (int i = 0; i < HISTOGRAM_LENGTH; i++) {
    count[i] = 0;
  }

  while ((c = getchar()) != EOF) {
    update_count(count, c);
  }

  // プリントする
  print_histogram(count);
}
