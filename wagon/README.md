tiny-wasm
=====

[![Build Status](https://travis-ci.org/go-interpreter/wagon.svg?branch=master)](https://travis-ci.org/go-interpreter/wagon)
[![codecov](https://codecov.io/gh/go-interpreter/wagon/branch/master/graph/badge.svg)](https://codecov.io/gh/go-interpreter/wagon)
[![GoDoc](https://godoc.org/github.com/iuscript/wasmvm?status.svg)](https://godoc.org/github.com/iuscript/wasmvm)

`wagon` is a [WebAssembly](http://webassembly.org)-based interpreter in [Go](https://golang.org), for [Go](https://golang.org).

**NOTE:** `wagon` requires `Go >= 1.9.x`.

## Purpose
WIP. 
Thie project is forked from `wagon`.`wagon` aims to provide tools (executables+libraries) to:

- decode `wasm` binary files
- load and execute `wasm` modules' bytecode.

`wagon` doesn't concern itself with the production of the `wasm` binary files;
these files should be produced with another tool (such as [wabt](https://github.com/WebAssembly/wabt) or [binaryen](https://github.com/WebAssembly/binaryen).)
`wagon` *may* provide a utility to produce `wasm` files from `wast` or `wat` files (and vice versa.)

The primary goal of `wagon` is to provide the building blocks to be able to build an interpreter for Go code, that could be embedded in Jupyter or any Go program.
