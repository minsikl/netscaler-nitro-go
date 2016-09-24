package datatypes

// Base
type BaseRes struct {
	Errorcode *int    `json:"errorcode,omitempty"`
	Message   *string `json:"message,omitempty"`
	Severity  *string `json:"severity,omitempty"`
}

// service
type Service struct {
	Name        *string `json:"name,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	ServiceType *string `json:"servicetype,omitempty"`
	Port        *int    `json:"port,omitempty"`
}

type ServiceReq struct {
	Service *Service `json:"service,omitempty"`
}

type ServiceRes struct {
	BaseRes
	Service []Service `json:"service,omitempty"`
}

// lbvserver
type Lbvserver struct {
	Name        *string `json:"name,omitempty"`
	ServiceType *string `json:"servicetype,omitempty"`
	Port        *int    `json:"port,omitempty"`
	Lbmethod    *string `json:"lbmethod,omitempty"`
	Ipv46       *string `json:"ipv46,omitempty"`
}

type LbvserverReq struct {
	Lbvserver *Lbvserver `json:"lbvserver,omitempty"`
}

type LbvserverRes struct {
	BaseRes
	Lbvserver []Lbvserver `json:"lbvserver,omitempty"`
}

//lbvserver_service_binding
type LbvserverServiceBinding struct {
	Name        *string `json:"name,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
}

type LbvserverServiceBindingReq struct {
	LbvserverServiceBinding *LbvserverServiceBinding `json:"lbvserver_service_binding,omitempty"`
}

type LbvserverServiceBindingRes struct {
	BaseRes
	LbvserverServiceBinding []LbvserverServiceBinding `json:"lbvserver_service_binding,omitempty"`
}

// systemfile
type Systemfile struct {
	Filename     *string `json:"filename,omitempty"`
	Filelocation *string `json:"filelocation,omitempty"`
	Filecontent  *string `json:"filecontent,omitempty"`
	Fileencoding *string `json:"fileencoding,omitempty"`
}

type SystemfileReq struct {
	Systemfile *Systemfile `json:"systemfile,omitempty"`
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
	Nsfeature *Nsfeature `json:"nsfeature,omitempty"`
}

type NsfeatureRes struct {
	BaseRes
	Nsfeature []Nsfeature `json:"nsfeature,omitempty"`
}

// sslcertkey
type Sslcertkey struct {
	Certkey *string `json:"certkey,omitempty"`
	Cert    *string `json:"cert,omitempty"`
	Key     *string `json:"key,omitempty"`
}

type SslcertkeyReq struct {
	Sslcertkey *Sslcertkey `json:"sslcertkey,omitempty"`
}

type SslcertkeyRes struct {
	BaseRes
	Sslcertkey []Sslcertkey `json:"sslcertkey,omitempty"`
}

// sslvserver_sslcertkey_binding
type SslvserverSslcertkeyBinding struct {
	Vservername *string `json:"vservername,omitempty"`
	Certkeyname    *string `json:"certkeyname,omitempty"`
}

type SslvserverSslcertkeyBindingReq struct {
	SslvserverSslcertkeyBinding *SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding,omitempty"`
}

type SslvserverSslcertkeyBindingRes struct {
	BaseRes
	SslvserverSslcertkeyBinding []SslvserverSslcertkeyBinding `json:"sslvserver_sslcertkey_binding,omitempty"`
}
