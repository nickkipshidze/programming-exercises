#include <stdio.h>
#include <stdlib.h>

#define CONVFACTOR 0.09290304

int main(int argc, char *argv[]) {
    char lengthft[4], widthft[4];

    printf("What is the length of the room in feet? ");
    scanf("%4s", lengthft);

    printf("What is the width of the room in feet? ");
    scanf("%4s", widthft);

    printf("You entered dimensions of %s feet by %s feet.\n", lengthft, widthft);

    int squareft = atoi(lengthft) * atoi(widthft);
    float squarem = (float)squareft * CONVFACTOR;

    printf("The area is\n%d square feet\n%f square meters\n", squareft, squarem);
    
    return 0;
}