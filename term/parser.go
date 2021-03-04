package term

import (
	"errors"
	"strconv"

	// "strconv"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}

// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	//panic("TODO: implement NewParser")
	return &ParserImpl{
		lex: nil,
		peekTok: nil,
		terms: make(map[string]*Term),
		termID: make(map[*Term]int),
		termCounter: 0,
	}
}

type ParserImpl struct {
	lex     *lexer
	peekTok *Token
	// Map from a string representing a Term to a Term
	terms map[string] *Term
	// Map from Term to its ID
	termID map[*Term] int
	termCounter int
}

// nextToken gets the next token either by reading peekTok or from the lexer
func (p *ParserImpl) nextToken() (*Token, error){
	// Get the next token from peekTok, and set peekTok = nil
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}
	// Get the next token from the lexer
	return p.lex.next()
}

// backToken puts back tok into peekTok
func (p *ParserImpl) backToken(tok *Token){
	p.peekTok = tok
}

// Parse a term
func (p *ParserImpl) Parse (input string) (*Term, error){
	//Initialize p.lex using the input and initialize p.peekTok
	p.lex = newLexer(input)
	p.peekTok = nil

	// If the input is empty string
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	if tok.typ == tokenEOF{
		return nil, nil
	}

	// If input not empty, then parse the input by calling parseNextTerm
	p.backToken(tok)
	term, err := p.parseNextTerm()
	if err != nil{
		return nil, ErrParser
	}
	// Error if we have not consumed all the input
	if tok, err := p.nextToken(); err != nil || tok.typ != tokenEOF {
		return nil, ErrParser
	}
	return term, nil
}

// parseNextTerm parses a prefix of the string(via the lexer) into a Term, or return an error
func (p *ParserImpl) parseNextTerm() (*Term, error){
	// Get the next token
	tok, err := p.nextToken()
	if err != nil{
		return nil, err
	}

	switch tok.typ {
	case tokenEOF:
		return nil, nil
	case tokenNumber:
		return p.mkSimpleTerm(TermNumber, tok.literal), nil
	case tokenVariable:
		return p.mkSimpleTerm(TermVariable, tok.literal), nil
	case tokenAtom:
		// Create a term of type atom for this current token(type tokenAtom)
		a := p.mkSimpleTerm(TermAtom, tok.literal)
		nxt, err := p.nextToken()
		if err != nil {
			return nil, err
		}
		if nxt.typ != tokenLpar{
			// If the next token is not of type tokenLpar, then this atom is not the functor for a compound term
			p.backToken(nxt)
			return a, nil
		}

		// nxt.typ == tokenLpar, so Atom might be the functor of a compound term
		arg, err := p.parseNextTerm()
		if err != nil {
			return nil, err
		}

		// Args of a compound term must contain at least one Term
		args := []*Term{arg}
		nxt, err = p.nextToken()
		if err != nil {
			return nil, err
		}
		// Parse the rest of the arguments, if any
		for ;nxt.typ == tokenComma; nxt, err = p.nextToken(){
			arg, err = p.parseNextTerm()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
		}
		if nxt.typ != tokenRpar {
			return nil, ErrParser
		}
		return p.mkCompoundTerm(a, args), nil
	default:
		return nil, ErrParser
	}
}


// mkSimpleTerm makes a simple term (atom, num, var)
func (p *ParserImpl) mkSimpleTerm(typ TermType, lit string) *Term {
	// convert the lit into a key for the map
	key := lit
	// Try to access the map "p.terms" with the key to see if a term of this form has already been created
	// => DAG (improved time/space complexity)
	term, ok := p.terms[key]
	if !ok {
		// If the term of this form has not exist in the mapping, then create a new one and store in the map
		term = &Term{Typ: typ, Literal: lit}
		p.insertTerm(term, key)
	}
	return term
}

// mkCompoundTerm makes a compound term
func (p *ParserImpl) mkCompoundTerm (functor *Term, args[]*Term) *Term{
	// Look up the mapping termID using *Term functor to find corresponding ID for this term
	// if it exists then convert it to string
	key := strconv.Itoa(p.termID[functor])
	for _, arg := range args{
		key += ", " + strconv.Itoa(p.termID[arg])
	}
	term, ok := p.terms[key]
	if !ok {
		// If the compound term has not exist, then create new one and store in the dictionary
		term = &Term{
			Typ: TermCompound,
			Functor: functor,
			Args: args,
		}
		p.insertTerm(term, key)
	}
	return term
}

// insertTerm inserts term with given key into the mapping "p.terms" and "p.termsID"
func (p *ParserImpl) insertTerm(term *Term, key string){
	// terms map from string key to a term pointer
	// termIsD map from Term pointer to a unique integer(ID) for each unique term
	// so each term can be accessed with a key and also has a corresponding ID
	p.terms[key] = term
	p.termID[term] = p.termCounter
	p.termCounter ++
}