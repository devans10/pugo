package flasharray

type Volume struct{
	Name	string	`json:"name,omitempty"`
	Source	string	`json:"source,omitempty"`
	Serial	string	`json:"serial,omitempty"`
	Size	int	`json:"size,omitempty"`
	Created	string	`json:"created,omitempty"`
}

type VolumePgroup struct {
	Name	string	`json:"name"`
	Pgroup	string	`json:"protection_group"`
}
