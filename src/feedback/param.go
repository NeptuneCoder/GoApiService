package feedback


type LogFileParam struct {
	AccountId string
	Date      string
	Path      string
	FileName  string
}

type ProblemInfoParam struct {
	Id          int
	AccountId   string
	VpnId       string
	Problem     string
	ContactInfo string
	Date        string
}
