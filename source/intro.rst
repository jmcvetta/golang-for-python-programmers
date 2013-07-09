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

* Familiarity with Python or a similar dynamic, interpretted language.
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


Resources
=========

The Sphinx_ `source code`_ for this book is available on on Github.

The latest version of these notes is published at `Read the Docs`_. It is also
available in PDF_ and ePub_ formats.


Related Courses
===============

.. todo::

   Add related courses.


Further Reading
===============

* `Effective Go`_ by the creators of the language.  Perhaps the single most
  valuable documentation resource for Go programmers.
* `An Introduction to Programming in Go`_ by Caleb Doxsey
* `Go for Pythonists`_ by Aditya Mukurjee


About the Author
================

`Jason McVetta`_ is an independent consultant, teacher, and Free Software
activist based in beautiful-but-too-cold San Francisco.  He studied government
and philosophy in college, but was unwilling to sell his soul to the
Demopublicratican Party, so he went to work in software. After 12 years as a
professional Pythonist he met his new love, Go.   He is able to fly without
mechanical assistance.



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
.. _PDF: https://media.readthedocs.org/pdf/golang-for-python-programmers/latest/golang-for-python-programmers.pdf
.. _ePub: https://media.readthedocs.org/epub/golang-for-python-programmers/latest/golang-for-python-programmers.epub
.. _Sphinx: http://sphinx-doc.org
.. _`source code`: http://github.com/jmcvetta/golang-for-python-programmers
.. _`Read the Docs`: http://golang-for-python-programmers.readthedocs.org/
.. _`Jason McVetta`: mailto:jason.mcvetta@gmail.com
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
.. _`Jason McVetta`: mailto:jason.mcvetta@gmail.com


.. rubric:: Citations

.. [#cit1] https://en.wikipedia.org/wiki/Go_(programming_language)
