package tf_vars_sorter

import (
	"log"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type ReaderHandler struct {
	dir    string
	output chan *Var
}

func NewReaderHandler(dir string, output chan *Var) *ReaderHandler {
	return &ReaderHandler{
		dir:    dir,
		output: output,
	}
}

func (this *ReaderHandler) Handle() {
	module, diagnostics := tfconfig.LoadModule(this.dir)
	if diagnostics.HasErrors() {
		log.Fatalf("Failed to read config for %s, %#v", this.dir, diagnostics)
	}
	for _, variable := range module.Variables {
		this.output <- &Var{
			Name:        variable.Name,
			Type:        variable.Type,
			Description: variable.Description,
			Default:     variable.Default,
			Required:    variable.Required,
		}
	}
	close(this.output)
}
