package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

/*
The Statement and Expression interfaces only contain dummy methods called statementNode and
expressionNode respectively. They are not strictly necessary but help us by guiding the Go
compiler and possibly causing it to throw errors when we use a Statement where an Expression
should’ve been used, and vice versa.
*/

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*
This Program node is going to be the root node of every AST our parser produces. Every valid Monkey program is a series of statements.
These statements are contained in the Program.Statements, which is just a slice of AST nodes that implement the Statement interface.
*/

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

/*
Identifiers in other parts of a Monkey program do produce values, thats why is it an Expression, even though the identifier in a let statement doesn’t produce
a value
*/
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
