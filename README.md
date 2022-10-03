diet - Discrete interval encoding trees for Go
==============================================

[![build and test](https://github.com/mrogaski/diet/actions/workflows/ci.yml/badge.svg)](https://github.com/mrogaski/diet/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/mrogaski/diet/branch/main/graph/badge.svg?token=xbOzAjNDCr)](https://codecov.io/gh/mrogaski/diet)

A Go implementation of discrete interval encoding trees (DIETs), as described by Martin Erwig in
[Diets for Fat Sets](https://web.engr.oregonstate.edu/~erwig/papers/Diet_JFP98.pdf).

DIETs store subsets of types having total order, a predecessor function, and an operator function (e.g. integers).
The superset is represented as a binary search tree in which maximal adjacent subsets are encoded as an interval.

---

## Installation

To install the diet package, use `go get`:

    go get github.com/mrogaski/diet
