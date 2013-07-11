#!/usr/bin/env python
#
# Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
# terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
# Resist intellectual serfdom - the ownership of ideas is akin to slavery.

import time


def greeting():
    '''Returns a pleasant, semi-useful greeting.'''
    return "Hello world, the time is: " + time.ctime()


def main():
    print greeting()


if __name__ == '__main__':
    main()

