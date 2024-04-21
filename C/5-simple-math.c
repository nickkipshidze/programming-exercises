#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    char num1[8], num2[8];

    printf("What is the first number? ");
    scanf("%8s", num1);
    printf("What is the second number? ");
    scanf("%8s", num2);

    int dec1 = atoi(num1), dec2 = atoi(num2);

    printf("%d + %d = %d\n%d - %d = %d\n%d * %d = %d\n%d / %d = %d\n", dec1, dec2, dec1+dec2, dec1, dec2, dec1-dec2, dec1, dec2, dec1*dec2, dec1, dec2, dec1/dec2);
    
    return 0;
}