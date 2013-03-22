logler
======

Golang's log library is cool, but I like automation of my log levels (function argument to determine at what log levels it gets logged). So I'm made this.

This package is identical to the builtin "log" package, except every method that would produce output has an added first argument of type 'int' to determine at what levels the output should actually be printed. This level is set with the SetLevel(level int) method of the Logger object.

All aspects are totally thread safe, it uses the builtin library "log" (also thread safe) to make an internal Logger object. and access to the internal level variable is controlled by channels (each Logger object uses up one goroutine, created when using SetLevel(level int) for the first time)