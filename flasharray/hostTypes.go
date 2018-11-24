package flasharray

type Host struct{
	Name	string		`json:"name,omitempty"`
	wwn	[]string	`json:"wwn,omitempty"`
	iqn	[]string	`json:"iqn,omitempty"`
}

type ConnectedVolume struct {
	Vol	string		`json:"vol,omitempty"`
	Name	string		`json:"name,omitempty"`
	Lun	int		`json:"lun,omitempty"`
}

type HostPgroup struct {
	Name	string		`json:"name,omitempty"`
	Pgroup	string		`json:"pgroup,omitempty"`
}

type HostConnections struct {
        Name    string          `json:"name,omitempty"`
        wwn     []string        `json:"wwn,omitempty"`
        iqn     []string        `json:"iqn,omitempty"`
	hgroup	string		`json:"hgroup,omitempty"`
}
