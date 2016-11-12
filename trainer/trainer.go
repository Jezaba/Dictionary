package trainer

import (
	D "Dictionary/dict"
	_ "fmt"
)

type Box struct {
	Number int
	Name int
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