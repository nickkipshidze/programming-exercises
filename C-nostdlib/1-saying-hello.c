void _start() {
    unsigned const char prompt[] = "What is your name? ";
    unsigned long int prompt_length = sizeof(prompt) - 1;

    __asm__(
        "movq $1, %%rax\n"
        "movq $1, %%rdi\n"
        "movq %0, %%rsi\n"
        "movq %1, %%rdx\n"
        "syscall\n"
        :
        : "r"(prompt), "r"(prompt_length)
        : "rax", "rdi", "rsi", "rdx"
    );

    char buffer[64];
    unsigned long read_length;

    __asm__(
        "movq $0, %%rax\n"
        "movq $0, %%rdi\n"
        "movq %0, %%rsi\n"
        "movq %1, %%rdx\n"
        "syscall\n"
        : 
        : "r"(buffer), "r"((unsigned long)sizeof(buffer))
        : "rax", "rdi", "rsi", "rdx"
    );

    for (int i = 0; i < 64; i++) {
        if (buffer[i] == '\n') {
            buffer[i] = '\0';
            read_length = i;
            break;
        }
    }
    
    char message[256];
    const char* greeting = "Hello, ";
    const char* closure = ", nice to meet you!";
    char* dest = message;

    while (*greeting) *dest++ = *greeting++;
    char* src = buffer;
    while (*src && *src != '\n') *dest++ = *src++;
    while (*closure) *dest++ = *closure++;

    *dest++ = '\n';
    *dest = '\0';

    __asm__(
        "movq $1, %%rax\n"
        "movq $1, %%rdi\n"
        "movq %0, %%rsi\n"
        "movq %1, %%rdx\n"
        "syscall\n"
        :
        : "r"(message), "r"((unsigned long)(dest - message))
        : "rax", "rdi", "rsi", "rdx"
    );

    __asm__(
        "movq $60, %%rax\n"
        "movq $0, %%rdi\n"
        "syscall\n"
        :
        :
        : "rax", "rdi"
    );
}