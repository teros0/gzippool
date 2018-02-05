gzippool
=======

Description
-----------

Pool wrapper for encoding/gzip

Installation
------------

This package can be installed with the go get command:

    go get github.com/teros0/gzippool

Usage
------------

Simply import this package as:

import gzip "github.com/teros0/gzippol"

and you're good to go.

Documentation
-------------

Original API left (almost) untouched: https://golang.org/pkg/compress/gzip/

Benchmarks
-------------

Here is the comparison of naive, simple pool and this package ways to create Writers and Readers:

* BenchmarkReadNaive-4     	    3000	      7283 ns/op
* BenchmarkReadPool-4      	    3000	      1137 ns/op
* BenchmarkReadMyPool-4    	    3000	      1234 ns/op

* BenchmarkWritePool-4     	    3000	     41414 ns/op
* BenchmarkWriteMyPool-4   	    3000	     39237 ns/op
* BenchmarkWriteNaive-4    	    3000	    345736 ns/op


ToDo
----

* Add comments for clarity
* Change names of variables
* Think about adding NewWriterWithLevel

