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

   hello/
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

Typically Go code is developed using distributed version control systems like
Git or Mercurial.  Therefore import paths are conventionally named based on the
Github etc URL.  The Go toolchain is aware of this, and can gracefully handle
automatic dependency installation if your project conforms to the convention.

