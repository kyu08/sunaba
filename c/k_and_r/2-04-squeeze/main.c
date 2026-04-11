#include <stdio.h>

void squeeze(char s1[], char s2[]) {
  int i, j;
  for (i = 0, j = 0; s1[i] != '\0'; i++) {
    bool found = false;
    for (int k = 0; s2[k] != '\0'; k++) {
      if (s1[i] == s2[k]) {
        found = true;
        break;
      }
    }
    if (!found) {
      s1[j] = s1[i];
      j++;
    }
  }
  s1[j] = '\0';
}

int main(void) {
  char s[] = "abcd";
  squeeze(s, "cde");
  for (int i = 0; s[i] != '\0'; i++) {
    printf("%c", s[i]);
  }
}
