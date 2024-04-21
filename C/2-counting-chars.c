#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
    char string[64];

    printf("What is the input string? ");
    scanf("%63s", string);

    printf("%s has %d characters.\n", string, strlen(string));

    return 0;
}