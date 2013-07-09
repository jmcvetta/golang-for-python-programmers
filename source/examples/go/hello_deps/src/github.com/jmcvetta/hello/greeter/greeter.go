// Copyright (c) 2013 Jason McVetta.  This is Free Software, released under the
// terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package greeter

import (
	"github.com/darkhelmet/env"
	"time"
)

// Greeting returns a pleasant, semi-useful greeting.
func Greeting() string {
	msg := "Hello world, the time is "
	msg += time.Now().String()
	msg += " and your PATH is "
	msg += env.String("PATH")
	return msg
}
