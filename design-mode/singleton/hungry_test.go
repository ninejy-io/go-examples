package main

import "testing"

func Test_hungry(t *testing.T) {
	de := GetDemo()
	putInstance(de)
	t.Error(de.Work())
}

func putInstance(d Instance) {
	d.Work()
}

func Test_lazy(t *testing.T) {
	la := GetLazy()
	t.Error(la.Work2())
}
