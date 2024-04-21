#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    char age[8], retire[8];
    int intage, intretire;

    printf("What is your current age? ");
    scanf("%8s", age);

    printf("At what age would you like to retire? ");
    scanf("%8s", retire);
    
    char curMonth[8], curDay[8], curYear[8];
    sscanf(__DATE__, "%s %s %s", curMonth, curDay, curYear);

    intage = atoi(age);
    intretire = atoi(retire);

    printf("You have %d years left until you can retire.\n", (intretire - intage));
    printf("It's %s, so you can retire in %d.\n", curYear, atoi(curYear) + (intretire - intage));
    
    return 0;
}