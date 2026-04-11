#include <stdio.h>

void print_delimiter() {
  printf("========================================\n");
}

int main(void) {
  float fahr, celsius;
  int lower, upper, step;
  lower = -20;
  upper = 150;
  step = 10;

  // print header
  print_delimiter();
  printf("🎉 MY celsius-to-fahr CALCULATOR 🎉\n");
  print_delimiter();
  printf("\n");

  // print results
  celsius = lower;
  while (celsius <= upper) {
    fahr = celsius * (9.0 / 5.0) + 32.0;
    /* celsius = (5.0 / 9.0) * (fahr - 32.0); */
    printf("%6.1f = %3.0f\n", celsius, fahr);
    /* printf("%3.0f = %6.1f\n", fahr, celsius); */
    celsius += step;
  }
}
