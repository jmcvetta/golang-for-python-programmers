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
   :lines: 5-17,28-


Curly Braces
============

The first thing one may notice is that whereas Python uses whitespace to denote
scope, Go uses curly braces.  The Go authors did, however, take a lesson from
Python's culture of readable code.  Although it is certainly possible to write
syntactically valid Go code without any indentation at all, nevertheless almost
all Go code one sees is formatted consistently and readably.  That's because,
Go includes a code formatter along with the compiler.  The formatter, ``gofmt``,
is considered canonically correct by the Go community - if you don't like how
``gofmt`` formats your code, *you are wrong*.  It is customary that ``gofmt``
always be run before checking in code.

.. code-block:: console

   $ gofmt -w hello.go  # -w option updates the file(s) in place


Package
=======

In any ``.go`` file, the first non-comment line is a ``package`` declaration. 
The package declaration is mandatory - *every* ``.go`` file must begin with one.
Every ``.go`` file in the same folder must have the same package name.

The ``package`` declaration is preceded by a comment, called the *package
comment*. Automatic documentation tools like ``godoc`` extract the package
commment as the description of your program.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 5-6


Import
======

Next comes the ``import`` declaration.  Although it is optional, most ``.go``
files will have one.  Each package to be imported is listed on a separate line,
inside quotation marks.  The packages in our example, ``fmt`` and ``time``, come
from the standard library.  By convention, 3rd-party packages are named after
their repository URL.  For example, my Neo4j library hosted on Github would be
imported as ``"github.com/jmcvetta/neo4j"``.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 7-11

Functions
=========

Our program has two functions, ``greeting()`` and ``main()``.

The function ``greeting()`` takes no arguments, and returns a string. Unlike
Java and C, in Go the type declaration *follows* the function name.  Like all
good function declarations, this one includes a doc comment describing what it
does.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 13-14

Return
======

Every function that declares a return type, must end with a ``return``
statement.  In this case we add a literal string to the output of
``time.Now().String()``  

Let's look at the documentatation for ``time``.  We can see that ``time.Now()``
returns an instance of type ``Time``.  That instance, in turn, exports a
``String()`` method that unsurprisingly returns a ``string``.  Using the
addition operator with two strings results in the strings being concatenated. 
If we had tried to add our greeting string to just ``time.Now()`` the code would
not compile, due to type mismatch.

.. literalinclude:: godoc.time.now.txt
   :language: console
   :emphasize-lines: 36-43

Variables
=========

Python is dynamically typed.  That means there is no difference between variable
assignment and declaration. For example, saying ``x = 3`` is equally valid if
``x`` was previously undeclared, or if it was already declared as an integer, or
even if it was already declared as e.g. a string.

Go, on the other hand, is statically typed.  A variable must be declared as a
specific type before a value can be assigned to it.  Once declared, a variable
may only be assigned values of its declared type.  Go also provides an operator,
``:=``, that combined declaration and assignment in the same statement.

Let's look at a needlessly complex version of our greeting function to see how
varable declaration and assignment work.

.. literalinclude:: examples/go/hello/hello.go
   :language: go
   :lines: 18-26
   :linenos:

On line 4 we used the ``var`` keyword to declare the variable ``now`` as a
string.  It is initially set to the `zero value`_ for ``string`` type, which is
"" the empty string. Line 5 assigns the (string) return value of
``time.Now().String()`` to ``now``.

On line 6 we use the `short variable declaration`_ syntax to  declare ``msg`` as
a ``string``, and immediately assign it the value of a literal string.  The
resulting variable is exactly the same as if we had declared it as ``var msg
string`` on a seperate line.

Line 7 appends the value of ``now`` to the value of ``msg``, working in this
case exactly like Python's ``+=`` operator.  However unlike Python, *only*
strings can be added to other strings, no automatic type coercion is ever
attempted.

``main()``
==========

Our Hello World program declares its package as ``main``, and contains a special
function ``main()``.  That tells Go to compile this code as an executable
program, with the entry point at ``main()``.  The function ``main()`` has no
arguments and no return value.  Command line arguments - called "argv" in many
languages - are available from the standard library, either raw from
``os.Args``, or with higher level abstraction from the package ``flag``.


.. _`zero value`: http://golang.org/ref/spec#The_zero_value
.. _`short variable declaration`: http://golang.org/ref/spec#Short_variable_declarations