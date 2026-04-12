#include <stdio.h>

void print_delimiter() { printf("========================================\n"); }
float fahr_to_celsius(float fahr) {
  return (5.0 / 9.0) * (fahr - 32.0);
  ;
}

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
  fahr = lower;
  while (fahr <= upper) {
    celsius = fahr_to_celsius(fahr);
    printf("%3.0f = %6.1f\n", fahr, celsius);
    fahr += step;
  }
}
