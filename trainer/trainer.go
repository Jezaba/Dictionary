package trainer

import (
	D "Dictionary/dict"
	_ "fmt"
)

type Box struct {
	Name string
	Sections []Section
}

type Section struct{
	Number int  // zB 0-9 Sektoren
        Expressions []Expression
}


type Expression struct{
	Dict *D.Dictionary
	Term D.Vocable
	falses int
	Rights int
}

func main()  {

}
