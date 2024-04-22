#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    char *months[] = {"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"};

    char strday[4];

    printf("Please enter the number of the month: ");
    scanf("%4s", strday);

    const int day = (atoi(strday)-1)%12;
    
    printf("The name of the month is %s.\n", months[day]);

    return 0;
}