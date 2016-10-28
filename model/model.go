package model

type Request struct{
	Statement string `json:"statement"`
}

type Response struct{
	Input Request `json:"request"`
	Response string `json:"response"`
}



