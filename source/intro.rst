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
   The backend of cloud-based CAD application TinkerCAD_ is written entirely in
   Go.
   
Drone.io
   Continuous integration service Drone.io_ is written entirely in Go.


Background Assumptions
======================

* Familiarity with Python or a similar dynamic, interpretted language.
* Comfortable working on the UNIX command line.
* Basic understanding of internet protocols such as HTTP(S).


Student Software
================

Command line examples were written on `Ubuntu`_ 12.10 desktop environment.
Everything we do in this curriculum can also be done on other flavors of Linux,
on Mac OSX - even on Windows.


Resources
=========

The source code for this curriculum can be found on Github_.

The latest version of these notes is published at `Read the Docs`_. It is also
available in PDF_ and ePub_ formats.


Related Courses
===============

.. todo::

   Add related courses


Further Reading
===============

* `An Introduction to Programming in Go`_ by Caleb Doxsey
* `Go for Pythonists`_ by Aditya Mukurjee


Copying
=======


All source code used in this curriculum is Free Software, released under the
`terms of the GPL v3`_.

The remainder of the curriculum is released under a
`Creative Commons Attribution-ShareAlike 3.0 Unported License`_.

.. raw:: html

   <a rel="license" href="http://creativecommons.org/licenses/by-sa/3.0/deed.en_US"><img alt="Creative Commons License" style="border-width:0" src="http://i.creativecommons.org/l/by-sa/3.0/88x31.png" /></a>


`Sing along with Richard Stallman`_, founder of the Free Software movement!


.. _Vitess: https://code.google.com/p/vitess/
.. _pq: https://github.com/lib/pq
.. _Doozer: https://github.com/ha/doozerd
.. _Bazooka: http://backstage.soundcloud.com/2012/07/go-at-soundcloud/
.. _Bitly: http://word.bitly.com/post/29550171827/go-go-gadget
.. _nsq: https://github.com/bitly/nsq
.. _Railgun: http://blog.cloudflare.com/go-at-cloudflare
.. _Iron.io: http://blog.iron.io/2013/03/how-we-went-from-30-servers-to-2-go.html
.. _Drone.io: https://groups.google.com/forum/#!topic/golang-nuts/Lo7KP3rWP3o
.. _TinkerCAD:  http://www.youtube.com/watch?v=JE17r3n1kz4
.. _PDF: https://media.readthedocs.org/pdf/golang-for-python-programmers/latest/modern-api-development.pdf
.. _ePub: https://media.readthedocs.org/epub/golang-for-python-programmers/latest/golang-for-python-programmers.epub
.. _Github: http://github.com/jmcvetta/golang-for-python-programmers
.. _`Read the Docs`: http://golang-for-python-programmers.readthedocs.org/
.. _`Jason McVetta`: mailto:jason.mcvetta@gmail.com
.. _`terms of the GPL v3`: http://www.gnu.org/copyleft/gpl.html
.. _Ubuntu: http://www.ubuntu.com
.. _`Sing along with Richard Stallman`: https://upload.wikimedia.org/wikipedia/commons/9/9c/Stallman_free_software_song.ogv
.. _`Creative Commons Attribution-ShareAlike 3.0 Unported License`: http://creativecommons.org/licenses/by-sa/3.0/deed.en_US
.. _`An Introduction to Programming in Go`:  http://www.golang-book.com/
.. _`Go for Pythonists`: https://github.com/ChimeraCoder/go-for-pythonists
.. _`Programming in Go: An Introduction`: http://programming-in-go.readthedocs.org


.. rubric:: Citations

.. [#cit1] https://en.wikipedia.org/wiki/Go_(programming_language)
