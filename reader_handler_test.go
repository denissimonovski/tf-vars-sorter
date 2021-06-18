package tf_vars_sorter

import (
	"testing"

	"github.com/smartystreets/gunit"
)

func TestReaderHandlerFixture(t *testing.T) {
	gunit.Run(new(ReaderHandlerFixture), t)
}

type ReaderHandlerFixture struct {
	*gunit.Fixture

	dir    string
	output chan *Var
	reader *ReaderHandler
}

func (this *ReaderHandlerFixture) Setup() {
	this.dir = "testdata"
	this.output = make(chan *Var, 10)
	this.reader = NewReaderHandler(this.dir, this.output)
}

func (this *ReaderHandlerFixture) TestAllItemsParsedAndSentToOutput() {
	this.reader.Handle()
	for variable := range this.output {
		this.AssertDeepEqual(expectedVariables[variable.Name], *variable)
	}
}

var expectedVariables = map[string]Var{
	"access_config": {
		Name:        "access_config",
		Type:        "list(object({\n    nat_ip       = string\n    network_tier = string\n  }))",
		Description: "Access configurations, i.e. IPs via which the VM instance can be accessed via the Internet.",
		Default:     []interface{}{},
		Required:    false,
	},
	"additional_disks": {
		Name: "additional_disks",
		Type: "list(object({\n    disk_name    = string\n    device_name  = string\n    auto_delete  = bool\n" +
			"    boot         = bool\n    disk_size_gb = number\n    disk_type    = string\n    " +
			"disk_labels  = map(string)\n  }))",
		Description: "List of maps of additional disks. See https://www.terraform.io/docs/providers/google/r/compute_instance_template.html#disk_name",
		Default:     []interface{}{},
		Required:    false,
	},
	"allowed_ips": {
		Name:        "allowed_ips",
		Type:        "list(string)",
		Description: "The IP address ranges which can access the load balancer.",
		Default:     []interface{}{"0.0.0.0/0"},
		Required:    false,
	},
	"auto_delete": {
		Name:        "auto_delete",
		Type:        "string",
		Description: "Whether or not the disk should be auto-deleted",
		Default:     "true",
		Required:    false,
	},
	"backends": {
		Name: "backends",
		Type: "map(object({\n    " +
			"protocol  = string\n    " +
			"port      = number\n    " +
			"port_name = string\n\n    " +
			"description            = string\n    " +
			"enable_cdn             = bool\n    " +
			"security_policy        = string\n    " +
			"custom_request_headers = list(string)\n\n    " +
			"timeout_sec                     = number\n    " +
			"connection_draining_timeout_sec = number\n    " +
			"session_affinity                = string\n    " +
			"affinity_cookie_ttl_sec         = number\n\n    " +
			"health_check = object({\n      " +
			"check_interval_sec  = number\n      " +
			"timeout_sec         = number\n      " +
			"healthy_threshold   = number\n      " +
			"unhealthy_threshold = number\n      " +
			"request_path        = string\n      " +
			"port                = number\n      " +
			"host                = string\n      " +
			"logging             = bool\n    " +
			"})\n\n    " +
			"log_config = object({\n      " +
			"enable      = bool\n      " +
			"sample_rate = number\n    " +
			"})\n\n    " +
			"groups = list(object({\n      " +
			"group = string\n\n      " +
			"balancing_mode               = string\n      " +
			"capacity_scaler              = number\n      " +
			"description                  = string\n      " +
			"max_connections              = number\n      " +
			"max_connections_per_instance = number\n      " +
			"max_connections_per_endpoint = number\n      " +
			"max_rate                     = number\n      " +
			"max_rate_per_instance        = number\n      " +
			"max_rate_per_endpoint        = number\n      " +
			"max_utilization              = number\n    " +
			"}))\n    " +
			"iap_config = object({\n      " +
			"enable               = bool\n      " +
			"oauth2_client_id     = string\n      " +
			"oauth2_client_secret = string\n    " +
			"})\n  " +
			"}))",
		Description: "Append more backends than the one that is created by default.",
		Default:     map[string]interface{}{},
		Required:    false,
	},
}
