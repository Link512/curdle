# curdle

Vanilla golang ASCII only command line version of the game Wordle

## Running

Either download the latest release from GitHub or run

```bash
    go install github.com/Link512/curdle@latest
```

and then run `curdle` to start playing

## Game description

The only supported mode is `easy mode` (where you don't have to use all revealed hints in subsequent guesses). Characters marked with `x` are the equivalent of the green hints (exact character at that position), while characters marked with `*` are the equivalent of yellow hints (exact character at a different position).

The words are taken from the original game word list (taken from the source code).

For example, if the word is `moist`, here is how a game of guessing it might look:

```
t r a s h
*     x

h e i s t
    x x x

h o i s t
  x x x x

m o i s t
x x x x x


CONGRATS!

1 ___*_
2 *__x_
3 __xxx
4 _xxxx
5 xxxxx
```
