#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
    char strpeople[8], strpizzas[8], strslices[8];
    int people, pizzas, slices;

    printf("How many people? ");
    scanf("%8s", strpeople);

    printf("How many pizzas do you have? ");
    scanf("%8s", strpizzas);

    printf("How may slices per pizza? ");
    scanf("%8s", strslices);

    printf("%s people with %s pizzas\n", strpeople, strpizzas);

    people = atoi(strpeople);
    pizzas = atoi(strpizzas);
    slices = atoi(strslices);

    printf("Each person gets %d pieces of pizza.\n", (pizzas * slices) / people);
    printf("There are %d leftover pieces.\n", (pizzas * slices) % people);

    return 0;
}