# curdle

Vanilla ASCII only command line version of the game Wordle written in go.

## Running

Either download the latest release from GitHub or run

```bash
    go install github.com/Link512/curdle@latest
```

and then run `curdle` to start playing

## Game description

The only supported mode is `easy mode` (where you don't have to use all revealed hints in subsequent guesses). Characters marked with `x` are the equivalent of the green hints (character match at that position), while characters marked with `*` are the equivalent of yellow hints (character match at a different position).

The words are taken from the original game word list (taken from the source code).

## Demo

If the word is `moist`, here is how a game of guessing it looks like:

[![asciicast](https://asciinema.org/a/467321.svg)](https://asciinema.org/a/467321)
