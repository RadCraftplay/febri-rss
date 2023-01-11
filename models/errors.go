package models

type FebriRssError struct {
	InnerError error
	Message    string
	ReturnCode int
}

func (e FebriRssError) Error() string {
	return e.Message
}
