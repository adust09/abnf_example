package main

import (
	"fmt"
	"os"

	"github.com/elimity-com/abnf"
)

func main() {
	rawABNF := `rule1 = "foo" bar baz
	rule2 = 1*2rule1
	bar   = *ALPHA
	baz   = ("123" / "456") [SP] CRLF`

	corePkg := abnf.ExternalABNF{
		IsOperator:  true,
		PackageName: "github.com/elimity-com/abnf/core",
	}

	g := abnf.CodeGenerator{
		PackageName: "definition",
		RawABNF:     []byte(rawABNF),
		ExternalABNF: map[string]abnf.ExternalABNF{
			"ALPHA":  corePkg,
			"BIT":    corePkg,
			"CHAR":   corePkg,
			"CR":     corePkg,
			"CRLF":   corePkg,
			"CTL":    corePkg,
			"DIGIT":  corePkg,
			"DQUOTE": corePkg,
			"HEXDIG": corePkg,
			"HTAB":   corePkg,
			"LF":     corePkg,
			"LWSP":   corePkg,
			"OCTET":  corePkg,
			"SP":     corePkg,
			"VCHAR":  corePkg,
			"WSP":    corePkg,
		},
	}
	file, err := os.Create("./parser.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	g.GenerateABNFAsAlternatives(file)

	fmt.Println("Generated parser.go")
}
