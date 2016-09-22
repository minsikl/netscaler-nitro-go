package datatypes

// Base
type BaseRes struct {
	Errorcode int    `json:"errorcode"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
}

// service
type Service struct {
	Name        string `json:"name"`
	Ip          string `json:"ip"`
	ServiceType string `json:"servicetype"`
	Port        int    `json:"port"`
}

type ServiceReq struct {
	Service Service `json:"service,omitempty"`
}

type ServiceRes struct {
	Service []Service `json:"service,omitempty"`
}

// lbvserver
type Lbvserver struct {
	Name        string `json:"name"`
	ServiceType string `json:"servicetype"`
	Port        int    `json:"port"`
	Lbmethod    string `json:"lbmethod"`
	Ipv46       string `json:"ipv46"`
}

type LbvserverReq struct {
	Lbvserver Lbvserver `json:"lbvserver,omitempty"`
}

type LbvserverRes struct {
	Lbvserver []Lbvserver `json:"lbvserver,omitempty"`
}

//lbvserver_service_binding
type LbvserverServiceBinding struct {
	Name        string `json:"name"`
	ServiceName string `json:"serviceName"`
}

type LbvserverServiceBindingReq struct {
	LbvserverServiceBinding LbvserverServiceBinding `json:"lbvserver_service_binding,omitempty"`
}

type LbvserverServiceBindingRes struct {
	LbvserverServiceBinding []LbvserverServiceBinding `json:"lbvserver_service_binding,omitempty"`
}

// systemfile
type Systemfile struct {
	Filename     string `json:"filename"`
	Filelocation string `json:"filelocation"`
	Filecontent  string `json:"filecontent"`
	Fileencoding string `json:"fileencoding"`
}

type SystemfileReq struct {
	Systemfile Systemfile `json:"systemfile,omitempty"`
}

type SystemfileRes struct {
	Systemfile []Systemfile `json:"systemfile,omitempty"`
}

// nsfeature
type Nsfeature struct {
	Feature []string `json:"feature"`
}

type NsfeatureReq struct {
	Nsfeature Nsfeature `json:"nsfeature,omitempty"`
}

type NsfeatureRes struct {
	Nsfeature []Nsfeature `json:"nsfeature,omitempty"`
}
