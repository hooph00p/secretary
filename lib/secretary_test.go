package lib_test

import (
	"testing"

	"github.com/hooph00p/secretary/lib"
)

func Test_Do(t *testing.T) {
	s := new(lib.Secretary)
	s.Do("run -command=\"go get github.com/gin-gonic/gin\"")
	s.Do("run --command='go get github.com/gin-gonic/gin'")
}
