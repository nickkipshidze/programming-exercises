#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
    char quote[64], author[64];

    printf("What is the quote? ");
    if (fgets(quote, sizeof(quote), stdin))
        quote[strlen(quote) - 1] = '\0';
    
    printf("Who said it? ");
    if (fgets(author, sizeof(author), stdin))
        author[strlen(author) - 1] = '\0';

    printf("%s says, \"%s\"\n", author, quote);

    return 0;
}