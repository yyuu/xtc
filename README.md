# xtc - an eXtremely Tentative Compiler

[![Build Status](https://travis-ci.org/yyuu/xtc.svg?branch=master)](https://travis-ci.org/yyuu/xtc)

An experimental pseudo-C compiler implementation written in [Go](http://golang.org/).

For now, most of the concept are brought from [Cflat Compiler](http://i.loveruby.net/archive/cbc/cbc-1.0.tar.gz) (licensed under 3-clause BSD license) by [Minero Aoki](http://i.loveruby.net/).


## Requirements

All development is progressed on Debian GNU/Linux sid (amd64).

Build dependencies:

(any golang non-standard libraries aren't used.)

  * golang 1.3.2-1

Runtime dependencies:

  * binutils 2.24.51.20140918-1
  * libc6-i386 2.19-11
  * libc6-dev-i386 2.19-11


## Build

```
$ ./build
```


## Test

```
$ ./test
```


## Run

Create a pseudo-C file like following.

```
$ cat hello.xtc
extern int printf(char* format, ...);
int main(int argc, char*[] argv) {
  printf("hello, world\n");
  return 0;
}
$ ./bin/xtc hello.xtc
$ ./hello
hello, world
```


## License

(The MIT license)

Copyright (c) 2014 Yamashita, Yuu <<peek824545201@gmail.com>>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
