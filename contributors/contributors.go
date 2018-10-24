package contributors

type Person struct {
	FullName string
}

var all = [...]Person{
	Person{FullName: "Robert Griesemer"},
	Person{FullName: "Rob Pike"},
	Person{FullName: "Ken Thompson"},
	Person{FullName: "Russ Cox"},
	Person{FullName: "Ian Lance Taylor"},
}

func Details() []Person {
	res := all
	return res[:]
}
