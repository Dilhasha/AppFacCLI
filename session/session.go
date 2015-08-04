package session

type Session struct {
	UserName string
	Cookie string
}

func NewSession()(Session){
	return Session{
		UserName:"",
		Cookie : "",
	}

}
