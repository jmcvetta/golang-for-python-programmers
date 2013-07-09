**************
Project Layout
**************

Just as ``$PYTHONPATH`` controls where the Python interpretter will look for
source files, ``$GOPATH`` controls where the Go compiler and tools will look for
source code.


Hello World
===========

Let us consider a modified version of the standard "Hello World" program.  Our
example will consist of a library package and an application package.  The
library exports one function, which returns the string "Hello world, the time
is: " plus the current time.  The application package calls the library
function, and prints the hello message to the console.


Python
------


In Python we might have a file layout like this::

   $PYTHONPATH/
      hello.py
      greeter.py

The library package is contained in ``greeter.py``:

.. literalinclude:: examples/python/hello/greeter.py
   :linenos:
   :lines: 5-

The application is contained in ``hello.py``:

.. literalinclude:: examples/python/hello/hello.py
   :linenos:
   :lines: 1,6-


We run the Python application like this:

.. code-block:: console

   $ python hello.py
   Hello world, the time is: Mon Jul  8 19:16:40 2013


Go
--

The "standard" layout for Go code is more complex than that for Python code.
The disadvantage is that setup is slightly more work.  The upside is a
well-structured codebase that encourages modular code and keeps things orderly
as a project grows in size.


Typically Go code is developed using `distributed version control systems`_ like
Git_ or Mercurial_.  Therefore import paths are conventionally named based on
the Github etc URL.  The Go toolchain is aware of this, and can gracefully
handle automatic dependency installation if your project conforms to the
convention.

Imagine for a moment we have created a Github repository named ``hello`` for our
Hello World example.  We would create a file layout like this::

   $GOPATH/
      src/
         github.com/
            jmcvetta/
               hello/
                  hello.go
                  greeter/
                     greeter.go

The library is located in ``greeter.go``.  Note the package name, ``greeter``,
is the same as the name of the containing folder.  This is mandatory.  The
package imports ``time`` from the standard library.


.. literalinclude:: examples/go/hello/src/github.com/jmcvetta/hello/greeter/greeter.go
   :language: go
   :linenos:
   :lines: 5-

The application is in ``hello.go``.  This file declares its package as ``main``,
and defines a ``main()`` function.  This tells the compiler to generate an
executable from the file.  The package imports ``fmt`` from the standard
library, and our ``greeter`` package specified by its path.

.. literalinclude:: examples/go/hello/src/github.com/jmcvetta/hello/hello.go
   :language: go
   :linenos:
   :lines: 5-

Functions imported from another package are *always* namespaced with the package
name.  In this case, we call ``greeter.Greeting()``.  Go intentionally has no
equivalent of Python's ``from some_package import *``.

We can build the application with the ``go build`` command:

.. code-block:: console

   $ go build -v
   github.com/jmcvetta/hello/greeter
   github.com/jmcvetta/hello

   $ ls
   greeter/  hello*  hello.go

   $ ./hello
   Hello world, the time is: 2013-07-08 19:49:39.946836748 -0700 PDT


.. _`distributed version control systems`: https://en.wikipedia.org/wiki/Distributed_revision_control
.. _Git: http://git-scm.com/
.. _Mercurial: http://mercurial.selenic.com/