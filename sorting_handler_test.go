package tf_vars_sorter

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestSortingHandlerFixture(t *testing.T) {
	gunit.Run(new(SortingHandlerFixture), t)
}

type SortingHandlerFixture struct {
	*gunit.Fixture

	input   chan *Var
	output  chan *Var
	handler *SortingHandler
	sorter  *TFVarSorter
}

func (this *SortingHandlerFixture) Setup() {
	this.input = make(chan *Var, 10)
	this.output = make(chan *Var, 10)
	this.sorter = NewTFVarSorter()
	this.handler = NewSortingHandler(this.input, this.output, this.sorter)
}

func (this *SortingHandlerFixture) TestExpectedEnvelopeSentToOutput() {
	variable := &Var{Name: "aws_account_id"}
	this.input <- variable
	close(this.input)

	this.handler.Handle()

	this.AssertEqual(<-this.output, variable)
}

func (this *SortingHandlerFixture) TestUnorderedReceived_OrderedSent() {
	this.sendEnvelopesInSequence("tableau_cidr", "fdw_multi_az", "fdw_instance_class")

	this.handler.Handle()

	this.AssertDeepEqual([]string{"fdw_instance_class", "fdw_multi_az", "tableau_cidr"}, this.varNamesInSequence())
}

func (this *SortingHandlerFixture) sendEnvelopesInSequence(varNames ...string) {
	for _, varName := range varNames {
		variable := &Var{Name: varName}
		this.input <- variable
	}
	close(this.input)
}

func (this *SortingHandlerFixture) varNamesInSequence() (sequence []string) {
	for variable := range this.output {
		sequence = append(sequence, variable.Name)
	}
	return sequence
}
