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


godoc.org
---------

The website godoc.org_ extracts and displays automatic documentation, much like
the ``godoc`` command, for any open source Go module specified by its import URL. 

See :doc:`GoDoc.org <godoc>` for more detail.



.. _godoc.org: http://godoc.org
.. _gowalker.org: http://gowalker.org