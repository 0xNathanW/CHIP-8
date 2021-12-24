# **Chip-8** #
Chip-8 is a very basic interpreted programming language written for computers in the 1970's, mainly for writing games.

This is an immplementation, or emulator, of Chip-8 written in Go.  It loads and runs programs, and allows users to pause and step through individual opcodes.  Users can also set their own colour configs.

## Demo ##
[![asciicast](https://asciinema.org/a/GCeqUwVQ1ZU8Q4uppy2bi1AsE.svg)](https://asciinema.org/a/GCeqUwVQ1ZU8Q4uppy2bi1AsE)

## Installation ##
To install, clone the repository with git, and build with `go build`.

## Usage ##
Run the program, with optional flags -fg and -bg for foreground and background colours.  A list of available colours are in colours.txt.

`./chip8 -fg=green -bg=black`

## TODO ##

- [ ] Address performance issues PC vs Mac
- [ ] Instructions on menu screen
- [ ] Step backwards through opcodes
- [ ] Display opcode info on stepthrough