package tf_vars_sorter

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestReaderHandler(t *testing.T) {
    gunit.Run(new(ReaderHandler), t)
}

type ReaderHandler struct {
    *gunit.Fixture
}

func (this *ReaderHandler) Setup() {
}

func (this *ReaderHandler) Test() {
}