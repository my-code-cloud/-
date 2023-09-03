//revive:disable
package pkg

//revive:enable:time-equal

import "time"

func t() bool {
	t := time.Now()
	u := t

	if !t.After(u) {
		return t.Equal(u) // MATCH /use t.Equal(u) instead of "==" operator/
	}

	return !t.Equal(u) // MATCH /use !t.Equal(u) instead of "!=" operator/
}
