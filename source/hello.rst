*************
Hello, World!
*************

Let's start out with a Hello World example.  Our program will greet the world
with a hello and the current time.

In Python it looks something like this:

.. literalinclude:: examples/python/hello/hello.py
   :linenos:
   :lines: 1,6-

   
It's not that different in Go:


.. literalinclude:: examples/go/hello/hello.go
   :linenos:
   :language: go
   :lines: 5-


The first non-comment line in the file is a ``package`` declaration.  The
package declaration is mandatory - every ``.go`` file must begin with one. 
Every ``.go`` file in the same folder must have the same package name.  

Our program declares the special package name ``main``.  Package ``main`` is
always compiled into an executable program, and must have a special method
``main()`` as the entry point into the application.

The ``package`` declaration is preceded by a comment, called the *package
comment*. Automatic documentation tools like ``godoc`` extract the package
commment as the description of your program.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 5-6

Next comes the ``import`` declaration.  Although it is optional, most ``.go``
files will have one.  Each package to be imported is listed on a separate line,
inside quotation marks.  The packages in our example, ``fmt`` and ``time``, come
from the standard library.  By convention, 3rd-party packages are named after
their repository URL.  For example, my Neo4j library hosted on Github would be
imported as ``"github.com/jmcvetta/neo4j"``.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 7-11

Our program has two functions, ``greeting()`` and ``main()``.  

The function ``greeting()`` takes no arguments, and returns a string containing
a pleasant hello message.  Unlike Java and C, in Go the type declaration
*follows* the function
name.  

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 14