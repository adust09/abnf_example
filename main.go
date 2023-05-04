package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elimity-com/abnf"
)

func main() {
	//read rawABNF from file
	// rawABNF, err := ioutil.ReadFile("./eip.abnf")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("rawABNF = ", rawABNF)

	//read sampleMessage from file
	// sampleMessage, err := ioutil.ReadFile("./sample_message.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }
	// fmt.Println("sampleMessage = ", sampleMessage)

	abnf_example := `rule1 = "foo" bar baz
	rule2 = 1*2rule1
	bar   = *ALPHA
	baz   = ("123" / "456") [SP] CRLF`
	fmt.Print("abnf_example = ", abnf_example)

	//generate functions from ParserGenerator
	g := abnf.ParserGenerator{RawABNF: []byte(abnf_example)}
	functions := g.GenerateABNFAsOperators()
	fmt.Println("functions = ", functions)

	parserCode := g.GenerateABNFAsOperators()
	code := ioutil.WriteFile("./parser/parser.go", []byte(parserCode), 0644)

	//execute generated function
	result := functions["sample_message"](sampleMessage)
	fmt.Println(result)

}
