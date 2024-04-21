#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
    char noun[32], verb[32], adjective[32], adverb[32];

    printf("Enter a noun: ");
    fgets(noun, 32, stdin);

    printf("Enter a verb: ");
    fgets(verb, 32, stdin);

    printf("Enter a adjective: ");
    fgets(adjective, 32, stdin);

    printf("Enter a adverb: ");
    fgets(adverb, 32, stdin);

    noun[strcspn(noun, "\n")] = 0;
    verb[strcspn(verb, "\n")] = 0;
    adjective[strcspn(adjective, "\n")] = 0;
    adverb[strcspn(adverb, "\n")] = 0;

    printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb);
    
    return 0;
}