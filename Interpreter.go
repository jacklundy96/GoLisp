package main

import (
	"strings"
	"fmt"
	"strconv"
	"errors"
)

//Struct definitions
type TokenType int
type symbol string  
type number float64

type Token struct {
	value string
	tokenType TokenType 
}
type expr interface{}

type Env map[symbol]interface{}

//If a value doesn't exist in the environment then insert it
func(e Env) update(sym symbol, expr interface{}) {
	if _, ok := e[sym]; ok {
		e[sym] = expr
	}
}

//Environment factory/constructor
func NewEnv() Env {
	Env := make(map[symbol]interface{})
	//Setyp global environment
	Env[symbol("null?")] = func(x interface{}) bool { return x == nil }

	return Env
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
	val,err := readTokens(tokenize("(defun area-circle(rad)(terpri)(format t rad)(format t  (* 3.141592 rad rad)))"))
	if err != nil {
		fmt.Println(val)
	}
}

//Parse token stream into token AST
func readTokens(tokens []Token) (expr, error) {	
	if len(tokens) == 0 {
		return nil, errors.New("Unexpected EOF")
	}

	token := tokens[0]
	if token.tokenType == LPAREN {
		val := parse(&tokens)
		return val, nil
	} else if token.tokenType == RPAREN {
		return nil, errors.New("Invalid Syntax Expected: ( ")
	}else {
		return nil, errors.New("Unknown Error Occoured")
	} 
}

//parse tokens in expression tree
func parse(tokens *[]Token) expr {	
	 //pop first element from tokens
	 token := (*tokens)[0]
	 *tokens = (*tokens)[1:]

	 switch token.tokenType {
		case LPAREN: 
			L := make([]expr, 0)
			for (*tokens)[0].tokenType != RPAREN {
				i := parse(tokens)
				L = append(L, i)
			}
			*tokens = (*tokens)[1:]
			return L
		default:
			if f, err := strconv.ParseFloat(token.value, 64); err == nil {
				return number(f)
			} else {
				return symbol(token.value)
			}
	 }
}

//Convert input string into tokens
func tokenize(program string) []Token {
	var out []Token
	var t Token 
	
	program = strings.Replace(program, "(", " ( ", -1)
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
				if _, err := strconv.ParseInt(rawToken,10,64); err == nil{
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
	//Convert token slice ([]Token) to []interface 
	slice := make([]Token, len(out))
	for i, v := range out {
		slice[i] = v
	}
	return slice
}