package flasharray

type User struct {
	Name	string	`json:"name"`
	Role	string	`json:"role"`
	Api_token	string	`json:"api_token"`
	Publickey	string	`json:"publickey"`
}

type ApiToken struct {
	Api_token	string	`json:"api_token"`
}

type GlobalAdmin struct {
}

type LockoutInfo struct {
}
