package lib_test

import (
	"strconv"
	"testing"

	"github.com/hooph00p/secretary/lib"
)

func TestRun(t *testing.T) {

}

func loc(t *lib.Task) string {
	return "./test/" + strconv.Itoa(t.i) + ".log"
}
