package tournaments

type List struct {
	Caster string `json:"caster"`
	List   string `json:"list"`
	Theme  string `json:"theme"`
}

func (l List) String() string {
	return l.Caster
}
