# A BASIC Interpreter in Go

[![Go CI/CD](https://github.com/cander/gobasic/actions/workflows/commit-actions.yaml/badge.svg)](https://github.com/cander/gobasic/actions/workflows/commit-actions.yaml)

This is a simple BASIC interpreter implement in Go. It's not intended for anything resembling real use.
It's not intended to implement most/all of any BASIC dialect. It's not intended to be the best/fastest/optimal
implementation. It's just a little sandbox.

## Why?
Way back when, I reimplemented the original Pascal compiler in C as an
exercise to learn C and to work with a real compiler (after university).
Later, I implemented a BASIC interpreter in Perl as a way to learn Perl. So, this is a bit of a learning kata.
As such, this isn't the first time I've implemented a BASIC interpreter or
an interpreter in Go  - I did it a while ago, and it got lost when I lost access
to the computer I implemented it on. (Shame on me for not pushing it to
GitHub.)  Most recently, I'm using this to play around with Visual Code.

Why BASIC? Because back in the day, before the public internet, when we
still used [UUCP](https://en.wikipedia.org/wiki/UUCP) to access
[Usenet](https://en.wikipedia.org/wiki/Usenet), someone posted a
[BASIC interpreter implemented entirely in Bourne Shell](https://gist.github.com/cander/2785819)
(`sh` - long before `bash`, `zsh`,
etc.).  At the time I was amazed by the painful lengths he went through to
parse and execute a BASIC program. I figured there had to be a better way
