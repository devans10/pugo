package flasharray

type Hostgroup struct {
	Name	string		`json:"name,omitempty"`
	Hosts	[]string	`json:"hosts"`
}

type HostgroupPgroup struct {
	Name	string		`json:"name,omitempty"`
	Pgroup	string		`json:"protection_group"`
}

type HostgroupConnection struct {
        Name    string          `json:"name,omitempty"`
        Vol     string		`json:"vol,omitempty"`
        Lun     string	        `json:"lun,omitempty"`
}
