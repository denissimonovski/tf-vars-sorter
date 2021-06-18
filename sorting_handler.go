package tf_vars_sorter

type Sorter interface {
	Sort([]*Var) []*Var
}

type SortingHandler struct {
	input  chan *Var
	output chan *Var
	sorter Sorter
}

func NewSortingHandler(input chan *Var, output chan *Var, sorter Sorter) *SortingHandler {
	return &SortingHandler{
		input:  input,
		output: output,
		sorter: sorter,
	}
}

func (this *SortingHandler) Handle() {
	for _, v := range this.sorter.Sort(this.composeVariablesList()) {
		this.output <- v
	}
	close(this.output)
}

func (this *SortingHandler) composeVariablesList() []*Var {
	var unsorted []*Var
	for variable := range this.input {
		unsorted = append(unsorted, variable)
	}
	return unsorted
}
