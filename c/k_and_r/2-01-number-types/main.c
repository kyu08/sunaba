#include <limits.h>
#include <stdint.h>
#include <stdio.h>

int main(void) {
  printf("signed char: [%d] ~ [%d]\n", SCHAR_MIN, SCHAR_MAX);
  printf("unsigned char: [%d] ~ [%d]\n", 0, UCHAR_MAX);
  // 以下略
  //  SCHAR_MIN
  //  SCHAR_MAX
  //  UCHAR_MAX
  //  SHRT_MIN
  //  SHRT_MAX
  //  USHRT_MAX
  //  INT_MIN
  //  INT_MAX
  //  UINT_MAX
  //  LONG_MIN
  //  LONG_MAX
  //  ULONG_MAX
}
