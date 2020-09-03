package structbot

type Bot struct {

}
type Test struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

func (*Bot)MakeStruct(serstr string,out *Test)  {
	out.ID = 1
}

