package main

import (
	"strings"
	"fmt"
	"strconv"
	"errors"
)

//Struct definitions
type TokenType int

type Token struct {
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
	val, err := parse(tokenize("(+ (+ 6 7) (- 8 9) )"))
	if err == nil {
		fmt.Println(val)
	}	
}

//Parse token stream into token AST
func parse(tokens []interface{}) ([]interface{}, error) { 	 
	var L []interface{} 
	token := tokens[0].(Token)
	tokens = tokens[1:]

	if len(tokens) == 0 {
		return nil, errors.New("Syntax Error: Unexpected EOF)")
	}	

	switch token.tokenType {
		case LPAREN:
			var C []interface{} 
			offset := 0
			for tokens[0].(Token).tokenType != RPAREN {
				recur, err :=  parse(tokens)  
				if err == nil {
					C = append(C, recur)
				}		
				tokens = tokens[1:]		
			}	
			return C, nil
		case RPAREN:
			return nil, errors.New("Syntax Error: Unexpected )")
		default:
			L = append(L, token)
			return L, nil
	}
}

//Convert input string into tokens
func tokenize(program string) []interface{} {
	var out []Token
	var t Token 
	
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
	//Convert token slice ([]Tokne) to []interface 
	slice := make([]interface{}, len(out))
	for i, v := range out {
		slice[i] = v
	}
	return slice
}