package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tm := time.Date(2023, 3, 17, 10, 15, 0, 0, time.UTC)
	hm := humanDate(tm)

	if hm != "17 Mar 2023 at 10:15" {
		t.Errorf("Got %q; wanted %q", hm, "17 Mar 2023 at 10:15")
	}
}
