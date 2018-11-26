package flasharray

// type console_lock describes the console_lock status of the array.
type console_lock struct {
	console_lock	string
}

// type Array gives information about the array
type Array struct {
	Id		string	`json:"id,omitempty"`
	Array_name	string	`json:"array_name,omitempty"`
	Version		string	`json:"version,omitempty"`
	Revision	string	`json:"revision,omitempty"`
}

