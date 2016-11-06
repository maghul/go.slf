# slf

This is a LGPL licensed library for facading logs.

The primary purpose is to allow leveled logging (Debug and Info) and
fine grained loggers in libraries which can be controlled by the
main application.

It doesn't do any logging output by itself, instead relying on
other logging implementations to do that. The most trivial being
io.Writer.

License
-------
The raopd package is licensed under the LGPL with an exception that allows it to be linked statically. Please see the LICENSE file for details.

Documentation
-------------
Reference API documentation can be generated using godoc. The documentation can
alse be found at https://godoc.org/github.com/maghul/go.slf


