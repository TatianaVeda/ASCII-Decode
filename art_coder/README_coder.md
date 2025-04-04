# ART-DECODER/ENCODER

A command line tool, which takes a string as an input andconverts art data into text-based art and back.

*ASCII (American Standard Code for Information Interchange)
      7-bit as defined in ISO-646 is a basic set of 128 numbered symbols
      which almost all kinds of computer can display.

## Usage 

The tool accepts a single command line argument, which will be an one or multiple strings of characters describing the art to be generated, using the provided flags.
```
Usage if strings: 
go run . '<decode-content>'.
go run . -e '<encode-content>'.

Usage if .txt file: 
go run . [-m|-e] <strings-or-art.txt>

```
## Decoder

Decoder function enables to convert consecutive sets of ASCII printable characters with square brackets into the picture or diagram as follows:

```
$ go run . '[5 #][5 -_]-[5 #]'
Decoded content:
#####-_-_-_-_-_-#####
>>>>>>> 9b79acb (changes in project description about -e flag usage)

$ go run . -m txt_files/lion.encoded.txt 
Decoded content:
        @|\@@
       -  @@@@
      /7   @@@@
     /    @@@@@@
     \-' @@@@@@@@`-_______________
      -@@@@@@@@@             /    \
 _______/    /_       ______/      |__________-
/,__________/  `-.___/,_____________----------_)

```
## Flags

Multi-line files decoder is toggled on with an extra command line argument: -m flag. 

The program can convert text-based art.txt or part of art back into a decoded string by adding -e flag in command line. 

-m: Use for multi-line txt file(default behavior).

-e: Use for encoding art.txt file or multiple lines(part of art) as a command line argument.

```
$ go run . -e '     | 
____________    __ -+-  ____________ 
\_____     /   /_ \ |   \     _____/
 \_____    \____/  \____/    _____/
  \_____                    _____/
     \___________  ___________/
               /____\'

Encoded content:

[20  ]| 
[12 _][4  ][2 _] -+-[2  ][12 _] 
\[5 _][5  ]/[3  ]/_ \ |[3  ]\[5  ][5 _]/
 \[5 _][4  ]\[4 _]/[2  ]\[4 _]/[4  ][5 _]/
[2  ]\[5 _][20  ][5 _]/
[5  ]\[11 _][2  ][11 _]/
[15  ]/[4 _]\

```
## Error handling

The program detects and reports malformed inputs (displays "Error") if:

- square brackets marks are unbalanced.

- the second argument is an empty string.

- the first argument of a string to be expanded is not a number.

- the arguments of a string to be expanded are not separated by a space.

 ## Bonus Features
  + __Multi-line + Encode mode.__ Encoder converts text-images into the string format, working with multi-line files and toggling on with a command line argument (-e flag).

