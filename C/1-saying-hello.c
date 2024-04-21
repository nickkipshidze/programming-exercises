#include <stdio.h>

int main(int argc, char *argv[]) {
    char name[64];

    printf("What is your name? ");
    scanf("%63s", name);

    printf("Hello, %s, nice to meet you!\n", name);

    return 0;
}