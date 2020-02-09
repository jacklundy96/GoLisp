package main

import (
	"strings"
	"strconv"
	"fmt"
)

//Struct definitions
type TokenType int

type token struct {
	value string
	tokenType TokenType 
} 
//Constant definitions
const (
	LPAREN TokenType = 1
	RPAREN TokenType = 2
	NUMBER TokenType = 3
	ATOM TokenType = 4
	SYMBOL TokenType = 5
)



func main() {
	fmt.Print(tokenize("(+ (+ 6 7.5) 5)"))
}

//Parse token stream into token AST
func parse(tokens []token) string {
	return ""
}

//Convert input string into tokens
func tokenize(program string) []token {
	var out []token
	var t token 
	
	program = strings.Replace(program, "(", "( ", -1)
	program = strings.Replace(program, ")", " )", -1)
	rawTokens := strings.Fields(program)  

	for _, rawToken := range rawTokens {
		switch rawToken {
			case "(":
				t.tokenType = LPAREN
				t.value = rawToken
				out = append(out, t)
			case ")":		
				t.tokenType = RPAREN
				t.value = rawToken
				out = append(out, t)	
			default: 
				if _, err := strconv.ParseInt(rawToken,10,64); err == nil {
					t.tokenType = NUMBER
					t.value = rawToken
					out = append(out, t)
				} else if _, err := strconv.ParseFloat(rawToken,64); err == nil {
					t.tokenType = NUMBER
					t.value = rawToken
					out = append(out, t)
				} else {
					t.tokenType = SYMBOL
					t.value = rawToken
					out = append(out, t)
				}				
		}
	}
	return out
}