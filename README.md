# Programming exercises

My solutions for Brian Hogan's exercises for programmers.

## Python

Just run the program with Python interpreter:
```shell
$ python python/<filename>.py
```

## C code

Compile the C code with this command:
```shell
$ cc C/<filename>.c -o a.out
```

Run the binary with this command:
```shell
$ ./a.out
```

## C nostdlib challange

This challange prohibits using the C standard library. Much fun!

Compile the C challange code with this command:
```shell
$ cc -nostdlib -e _start -fno-stack-protector <filename>.c -o a.out
```

Run the binary with this command:
```shell
$ ./a.out
```

## Assembly

Coming soon (;