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
	BaseRes
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
	BaseRes
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
	BaseRes
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
	BaseRes
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
	BaseRes
	Nsfeature []Nsfeature `json:"nsfeature,omitempty"`
}

// sslcertkey
type Sslcertkey struct {
	Certkey string `json:"certkey"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
}

type SslcertkeyReq struct {
	Sslcertkey Sslcertkey `json:"sslcertkey,omitempty"`
}

type SslcertkeyRes struct {
	BaseRes
	Sslcertkey []Sslcertkey `json:"sslcertkey,omitempty"`
}

// sslvserver_sslcertkey_binding
type SslvserverSslcertkeyBinding struct {
	Vservername string `json:"vservername"`
	Certkeyname    string `json:"certkeyname"`
}

type SslvserverSslcertkeyBindingReq struct {
	SslvserverSslcertkeyBinding SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding,omitempty"`
}

type SslvserverSslcertkeyBindingRes struct {
	BaseRes
	SslvserverSslcertkeyBinding []SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding,omitempty"`
}
