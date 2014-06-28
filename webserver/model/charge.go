package model

type Charge struct {
	Id         	int			`json:"id"`
	Content_id 	int			`json:"content_id"`
	Username	string		`json:"username"`
	Charged     bool		`json:"charged"`
}
