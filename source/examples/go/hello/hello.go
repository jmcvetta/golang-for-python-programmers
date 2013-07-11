// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

// A "Hello World" program that prints a greeting with the current time.
package main

import (
	"fmt"
	"time"
)

// greeting returns a pleasant, semi-useful greeting.
func greeting() string {
	msg := "Hello world, the time is: " + time.Now().String()
	return msg
}

func main() {
	fmt.Println(greeting())
}
