************
Introduction
************

.. pull-quote::

   Go is like C and Python had a kid, who inherited Python's good looks and
   pleasant demeanor, along with C's confidence and athletic ability.


What is Go?
===========

Go is an open source, compiled, garbage-collected, concurrent system programming
language. It was first designed and developed at Google Inc. beginning in
September 2007 by Robert Griesemer, Rob Pike, and Ken Thompson. [#cit1]_

Minimalist
----------

The creators of Go advocate a minimalist approach to language design.  This
minimal elegance is often contrasted with the "kitchen sink" approach of C++.

Opinionated
-----------

Like Python, in Go there is often "one right way to do things", although that
phrase is not so commonly used as in the Pythonist community.  The compiler is
renowed for its strictness.  Things that would at most be warnings in other
languages - e.g. unused imports, unused variables - are hard compiler errors in
Go.  At first that may seem like a draconian buzzkill - but soon you will begin
to appreciate how clean and cruft-free it keeps your code.

Fast
----

No question about it - Go is blazing fast.  Despite a still young compiler with
minimal speed optimizations, Go regularly scores at or near the top of
inter-language benchmark comparisons.  That's because Go is statically typed,
compiled - and although garbage-collected, allows the programmer some degree of
control over memory usage.  

Batteries Included
------------------

To a large degree Go follows Python's "batteries included" philosophy.  Although
there are not yet as many 3rd party libraries as there are for Python,
nevertheless Go's standard library provides a solid foundation for many modern
programming tasks.  Built-in support for serving HTTP, de/serializing JSON,
template rendering, and strong cryptography make Go ideal for modern web
services development.

Productive
----------

Go is frequently - and favorably - compared with Python and Ruby for programmer
productivity.  Its clean and elegant syntax, unobtrusive static typing, optional
duck typing, batteries-included standard library, lightning-fast compilation,
and support for a variety of modern continuous integration tools make Go a
productivity champion.


Searching
=========

Since "go" is such a common English word, it is completely useless as a search
term when googling.  Instead search for "golang".  It is standard practice in
the community that all Go-related blog posts, Github repos, tweets, job
postings, etc be tagged with "golang".


Who's Using Go?
===============

Google
   Home of the Go authors and core team, signifiant parts of the Google
   infrastructure are thought to be written in Go.  The company however does not
   publicly disclose which Google services are written in which languages.
   Youtube is known to be powered by Vitess_, an open source, massively scalable
   database multiplexing layer written in Go.

Heroku
   Maintainers of the pq_ driver for PostgreSQL, significant parts of
   Heroku's infrastructure are said to be written in Go.  One component that has
   been open sourced is Doozer_, a consistent distributed data store.

Canonical
   The company behind Ubuntu Linux recently rewrote_ its Juju_ devops management
   application from Python to Go, with pleasing results.

Soundcloud
   Internet music thought leaders SoundCloud use a bespoke build and deployment
   system called Bazooka_, designed to be a platform for managing the deployment
   of internal services. It is promised that Bazooka will be open sourced soon.

Bitly
   Significant parts of Bitly_ are written in Go, including nsq_, a realtime
   distributed message processing system.

Cloudflare
   Large volumes of internet traffic are served by Cloudflare's Railgun_ web
   optimization software, designed to speed up the delivery of content that
   cannot be cached.

Iron.io
   Cloud infrastructure company and core Go community supporters Iron.io_
   replaced parts of their Ruby codebase with Go, allowing them to handle a
   greater load with 2 workers than with 30 Ruby workers.

TinkerCAD
   Cloud-based CAD application TinkerCAD_ is written in Go.

Drone.io
   Continuous integration service Drone.io_ is written in Go.


Background Assumptions
======================

* Familiarity with Python or a similar dynamic, interpreted language.
* Comfortable working on the UNIX command line.
* Basic understanding of internet protocols such as HTTP(S).


Requirements
============

Command line examples were written on an `Ubuntu`_ 12.10 desktop environment.
Everything we do in this book can also be done on other flavors of Linux, on Mac
OSX - probably even on Windows.

In addition to Go language basics, this book also covers cloud_ deployment and
`continuous integration`_ tools & techniques popular in the Go community.  You
will need free accounts on the following services:

* Github_
* Drone_
* Coveralls_
* Heroku_


Further Reading
===============

* `Effective Go`_ by the creators of the language.  Perhaps the single most
  valuable documentation resource for Go programmers.
* `An Introduction to Programming in Go`_ by Caleb Doxsey
* `Go for Pythonists`_ by Aditya Mukurjee




.. _Vitess: https://code.google.com/p/vitess/
.. _pq: https://github.com/lib/pq
.. _Doozer: https://github.com/ha/doozerd
.. _rewrote: https://www.youtube.com/watch?v=USr0Bvg1ZOo
.. _Juju: https://juju.ubuntu.com/
.. _Bazooka: http://backstage.soundcloud.com/2012/07/go-at-soundcloud/
.. _Bitly: http://word.bitly.com/post/29550171827/go-go-gadget
.. _nsq: https://github.com/bitly/nsq
.. _Railgun: http://blog.cloudflare.com/go-at-cloudflare
.. _Iron.io: http://blog.iron.io/2013/03/how-we-went-from-30-servers-to-2-go.html
.. _Drone.io: https://groups.google.com/forum/#!topic/golang-nuts/Lo7KP3rWP3o
.. _TinkerCAD:  http://www.youtube.com/watch?v=JE17r3n1kz4
.. _Ubuntu: http://www.ubuntu.com
.. _`Effective Go`: http://golang.org/doc/effective_go.html
.. _`An Introduction to Programming in Go`:  http://www.golang-book.com/
.. _`Go for Pythonists`: https://github.com/ChimeraCoder/go-for-pythonists
.. _`Programming in Go: An Introduction`: http://programming-in-go.readthedocs.org
.. _`continuous integration`: https://en.wikipedia.org/wiki/Continuous_integration
.. _cloud: http://en.wikipedia.org/wiki/Cloud_computing
.. _Github:  http://github.com
.. _Drone: http://drone.io
.. _Coveralls: http://coveralls.io
.. _Heroku:  http://heroku.com


.. rubric:: Citations

.. [#cit1] https://en.wikipedia.org/wiki/Go_(programming_language)
