****************
The Go Toolchain
****************

Go is a compiled language - the source code you write must be converted to a
binary format understandable to a computer before it is executed.  The compiler
is accessed through the ``go`` command.  In addition, the ``go`` command
provides several other tools

The sections below describe several of the more popular ``go`` commands, but do
not cover the complete set.  For full information run:

.. code-block:: console

   $ go help



``go build``
============

Build compiles the packages named by the import paths, along with their
dependencies, but it does not install the results.

If the arguments are a list of .go files, build treats them as a list of source
files specifying a single package.

When the command line specifies a single main package, build writes the
resulting executable to output. Otherwise build compiles the packages but
discards the results, serving only as a check that the packages can be built.

.. code-block:: console

   $ go build somefile.go  # Build just this one file
   $ go build              # Build package in current folder


``godoc``
=========

Godoc extracts and generates documentation for Go programs.

It has two modes.

Without the ``-http`` flag, it runs in command-line mode and prints plain text
documentation to standard output and exits. If both a library package and a
command with the same name exists, using the prefix cmd/ will force
documentation on the command rather than the library package. If the ``-src``
flag is specified, godoc prints the exported interface of a package in Go source
form, or the implementation of a specific exported language entity::

   godoc fmt                # documentation for package fmt
   godoc fmt Printf         # documentation for fmt.Printf
   godoc cmd/go             # force documentation for the go command
   godoc -src fmt           # fmt package interface in Go source form
   godoc -src fmt Printf    # implementation of fmt.Printf

With the ``-http`` flag, it runs as a web server and presents the documentation
as a web page.

::

   godoc -http=:6060


godoc.org
=========

The website godoc.org_ extracts and displays automatic documentation, much like
the ``godoc`` command, for any Go module specified by its import URL. You may
also want to check out gowalker.org_, an enhanced version of godoc.org that
displays code snippets alongside the documentation.


.. _godoc.org: http://godoc.org
.. _gowalker.org: http://gowalker.org