# Kanren

An implementation of µKanren for Go. Work in progress.

See [the miniKanren page](http://minikanren.org/).

## TODO

* Stringize everything (maybe unneeded with reify)
* Num data type
* Vars as numbers instead of strings
* Infinite streams
* Interleaving streams
* User level functionality (reify, run, ...)
* Can streams be implemented with channels?
* Correct package name in go.mod
* Maybe remove IsVar, etc. and leave only Equal
* Use reflect and Ifz data type to support Go types?