#include <stdio.h>

void print_delimiter() { printf("========================================\n"); }

int main(void) {
  float fahr, celsius;
  int lower, upper, step;
  lower = 0;
  upper = 300;
  step = 20;

  // print header
  print_delimiter();
  printf("🎉 MY fahr-to-celsius CALCULATOR 🎉\n");
  print_delimiter();
  printf("\n");

  // print results
  for (fahr = upper; lower <= fahr; fahr -= step) {
    celsius = (5.0 / 9.0) * (fahr - 32.0);
    printf("%3.0f = %6.1f\n", fahr, celsius);
  }
}
