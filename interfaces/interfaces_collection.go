package main

type named interface {
	getName() string
}

type Sity struct {
	Name string
}

func (s Sity) getName() string {
	return s.Name
}

type Star struct {
	Name string
}

func (s Star) getName() string {
	return s.Name
}

func main() {
	var rows = []named{
		Sity{"Rome"},
		Star{"Sirius"},
	}

	var list = ""
	for _, row := range rows {
		if len(list) > 0 {
			list += ", "
		}
		list += row.getName()
	}
}
