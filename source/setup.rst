**********
Tool Setup
**********

Installing Go
=============

The first thing we need is a working Go installation on our system.  The
following instructions were derived from the official `Getting Started`_ documentation.

Binary Packages
---------------

Binary distributions of Go are available for the FreeBSD, Linux, Mac OS X,
NetBSD, and Windows operating systems from the `official downloads page`_.


Building from Source
--------------------

Many developers prefer to install Go directly from the source code.  Building is
easy and quick - the Go compiler and tool chain can be compiled in a matter of
minutes on even a modest workstation or laptop.


Install Build Tools
^^^^^^^^^^^^^^^^^^^

The Go tool chain is written in C. To build it, you need a C compiler installed.
You will also need the Mercurial DVCS tool to download the source. On Ubuntu run
the following command to install the necessary packages:

.. code-block:: console

   $ sudo apt-get install gcc libc6-dev mercurial

See `Installing Go from source`_ and `Install C tools`_ for build instructions
on other platforms.

Fetch the repository
^^^^^^^^^^^^^^^^^^^^

Go will install to a directory named ``go``. Change to the directory that will
be its parent and make sure the ``go`` directory does not exist. Then check out
the repository:

.. code-block:: console

   $ hg clone -u release https://code.google.com/p/go

Build Go
^^^^^^^^

To build the Go distribution, run

.. code-block:: console

   $ cd go/src
   $ ./all.bash

If all goes well, it will finish by printing output like::

   ALL TESTS PASSED

   ---
   Installed Go for linux/amd64 in /home/you/go.
   Installed commands in /home/you/go/bin.
   *** You need to add /home/you/go/bin to your $PATH. ***

where the details on the last few lines reflect the operating system,
architecture, and root directory used during the install.


Environment Variables
=====================

These instructions assume Go is installed at ``~/go``.  If you installed
elsewhere, change the commands to reflect your actual installation path.

The Go tools look for several environment variables.  Typically these are set in
your shell profile.  If you are using Bash, add the following lines to your
``~/.bashrc``::

   export GOBIN=~/go/bin
   export GOARCH=amd64
   export GOOS=linux
   export GOROOT=~/go


PATH
----

You probably also want to add the path to your ``go/bin`` directory to your ``PATH``::

   export PATH=$PATH:/home/you/go/bin

Reload your ``~/.bashrc`` file to active the new environment variables.  Once
the ``go`` binary is on your ``PATH``, you can confirm your install is working
correctly by running the ``go version`` command.

.. code-block:: console

   $ go version
   go version go1.1.1 linux/amd64


GOPATH
------





.. _`Getting Started`: http://golang.org/doc/install
.. _`official downloads page`: https://code.google.com/p/go/downloads/list
.. _`Installing Go from source`: http://golang.org/doc/install/source
.. _`Install C tools`: https://code.google.com/p/go-wiki/wiki/InstallFromSource#Install_C_tools