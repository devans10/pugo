package flasharray

type NetworkInterface struct{
	Name	string	`json:"name,omitempty"`
	Address	string	`json:"address,omitempty"`
	Gateway	string	`json:"gateway,omitempty"`
	Netmask	string	`json:"netmask,omitempty"`
	Enabled bool	`json:"enabled"`
}

type Subnet struct {
	Name	string	`json:"name"`
	Prefix	string	`json:"prefix"`
	Enabled bool	`json:"enabled"`
}

type VlanInterface struct {
        Name    string  `json:"name"`
	Address	string	`json:"address"`
	Gateway	string	`json:"gateway"`
	Netmask	string	`json:"netmask"`
	Enabled bool	`json:"enabled"`
        Subnet  string  `json:"subnet"`
}

type DNS struct {
	Servers string `json:"servers"`
}

type Port struct {
	Name	string	`json:"name"`
}
