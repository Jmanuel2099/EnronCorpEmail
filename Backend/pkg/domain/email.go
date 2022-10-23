package domain

//This is the model for the email
type Email struct {
	MessageId string
	From      string
	To        string
	Date      string
	Subject   string
	Content   string
}
