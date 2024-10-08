// Code generated by goyacc -o scheme_parser.go scheme_parser.y. DO NOT EDIT.

//line scheme_parser.y:2
package main

import __yyfmt__ "fmt"

//line scheme_parser.y:2
import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type Expression interface{}

type IntLiteral struct {
	Value int
}

type Var struct {
	Name string
}

type Token struct {
	token   int
	literal string
}

type IfExpr struct {
	Cond Expression
	Then Expression
	Else Expression
}

type LetExpr struct {
	Bindings []Binding
	Body     Expression
}

type Binding struct {
	Name  string
	Value Expression
}

type Application struct {
	Func Expression
	Args []Expression
}

type DefineExpr struct {
	Name  string
	Value Expression
}

type BinaryOp struct {
	Operator string
	Left     Expression
	Right    Expression
}

type WhileExpr struct {
	Cnd  Expression
	Body Expression
}
type SetExpr struct {
	Name  string
	Value Expression
}

type BeginExpr struct {
	Exprs []Expression
}

//line scheme_parser.y:72
type yySymType struct {
	yys    int
	token  Token
	expr   Expression
	str    string
	intval int
}

const NAME = 57346
const INTEGER = 57347
const LPAREN = 57348
const RPAREN = 57349
const PLUS = 57350
const LT = 57351
const GT = 57352
const LET = 57353
const IF = 57354
const DEFINE = 57355
const LAMBDA = 57356
const EQ = 57357
const WHILE = 57358
const SET = 57359
const BEGIN = 57360

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NAME",
	"INTEGER",
	"LPAREN",
	"RPAREN",
	"PLUS",
	"LT",
	"GT",
	"LET",
	"IF",
	"DEFINE",
	"LAMBDA",
	"EQ",
	"WHILE",
	"SET",
	"BEGIN",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line scheme_parser.y:153

type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.Scan()
	lit := l.TokenText()

	switch tok {
	case scanner.Int:
		tokVal, _ := strconv.Atoi(lit)
		lval.intval = tokVal
		return INTEGER
	case '(':
		return LPAREN
	case ')':
		return RPAREN
	case scanner.Ident:
		switch lit {
		case "if":
			return IF
		case "let":
			return LET
		case "define":
			return DEFINE
		case "while":
			return WHILE
		case "set":
			return SET
		case "begin":
			return BEGIN
		default:
			lval.str = lit
			return NAME
		}
	case '+':
		return PLUS
	case '<':
		return LT
	case '>':
		return GT
	case '=':
		return EQ
	}

	return 0
}

func (l *Lexer) Error(e string) {
	fmt.Printf("Lex error: %s\n", e)
}

func Parse(input string) (Expression, error) {
	lexer := &Lexer{}
	lexer.Init(strings.NewReader(input))

	if yyParse(lexer) == 0 {
		return lexer.result, nil
	} else {
		return nil, fmt.Errorf("Parsing error")
	}
}

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 57

var yyAct = [...]int8{
	20, 2, 4, 3, 5, 28, 9, 52, 17, 51,
	50, 21, 22, 23, 24, 19, 47, 46, 29, 30,
	45, 44, 33, 34, 35, 36, 37, 43, 42, 39,
	41, 26, 38, 31, 16, 40, 32, 25, 18, 27,
	48, 49, 4, 3, 5, 1, 10, 11, 12, 6,
	7, 8, 0, 0, 13, 14, 15,
}

var yyPact = [...]int16{
	-2, -1000, -1000, -1000, -1000, 38, 28, -2, 34, -2,
	-2, -2, -2, -2, 33, -2, -1, -2, -2, 26,
	-2, -2, -2, -2, -2, -2, 25, 22, 31, -2,
	21, -1000, -1000, 20, 14, 13, 10, 9, -1000, -2,
	-2, 3, -1000, -1000, -1000, -1000, -1000, -1000, 2, 0,
	-1000, -1000, -1000,
}

var yyPgo = [...]int8{
	0, 45, 0, 39, 15,
}

var yyR1 = [...]int8{
	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 4, 4,
}

var yyR2 = [...]int8{
	0, 1, 1, 1, 7, 6, 5, 4, 5, 5,
	5, 5, 5, 4, 4, 0, 2,
}

var yyChk = [...]int16{
	-1000, -1, -2, 5, 4, 6, 11, 12, 13, -2,
	8, 9, 10, 16, 17, 18, 6, -2, 4, -4,
	-2, -2, -2, -2, -2, 4, -4, -3, 6, -2,
	-2, 7, -4, -2, -2, -2, -2, -2, 7, 7,
	4, -2, 7, 7, 7, 7, 7, 7, -2, -2,
	7, 7, 7,
}

var yyDef = [...]int8{
	0, -2, 1, 2, 3, 0, 0, 0, 0, 15,
	0, 0, 0, 0, 0, 15, 0, 0, 0, 0,
	15, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 7, 16, 0, 0, 0, 0, 0, 13, 0,
	0, 0, 6, 8, 9, 10, 11, 12, 0, 0,
	5, 4, 14,
}

var yyTok1 = [...]int8{
	1,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line scheme_parser.y:96
		{
			yyVAL.expr = yyDollar[1].expr
			yylex.(*Lexer).result = yyVAL.expr
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line scheme_parser.y:102
		{
			yyVAL.expr = IntLiteral{Value: yyDollar[1].intval}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line scheme_parser.y:105
		{
			yyVAL.expr = Var{Name: yyDollar[1].str}
		}
	case 4:
		yyDollar = yyS[yypt-7 : yypt+1]
//line scheme_parser.y:108
		{
			yyVAL.expr = LetExpr{Bindings: []Binding{yyDollar[4].expr.(Binding)}, Body: yyDollar[6].expr}
		}
	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
//line scheme_parser.y:111
		{
			yyVAL.expr = IfExpr{Cond: yyDollar[3].expr, Then: yyDollar[4].expr, Else: yyDollar[5].expr}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:114
		{
			yyVAL.expr = DefineExpr{Name: yyDollar[3].str, Value: yyDollar[4].expr}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line scheme_parser.y:117
		{
			yyVAL.expr = Application{Func: yyDollar[2].expr, Args: yyDollar[3].expr.([]Expression)}
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:120
		{
			yyVAL.expr = BinaryOp{Operator: "+", Left: yyDollar[3].expr, Right: yyDollar[4].expr}
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:123
		{
			yyVAL.expr = BinaryOp{Operator: "<", Left: yyDollar[3].expr, Right: yyDollar[4].expr}
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:126
		{
			yyVAL.expr = BinaryOp{Operator: ">", Left: yyDollar[3].expr, Right: yyDollar[4].expr}
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:129
		{
			yyVAL.expr = WhileExpr{Cnd: yyDollar[3].expr, Body: yyDollar[4].expr}
		}
	case 12:
		yyDollar = yyS[yypt-5 : yypt+1]
//line scheme_parser.y:132
		{
			yyVAL.expr = SetExpr{Name: yyDollar[3].str, Value: yyDollar[4].expr}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
//line scheme_parser.y:135
		{
			yyVAL.expr = BeginExpr{Exprs: yyDollar[3].expr.([]Expression)}
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
//line scheme_parser.y:141
		{
			yyVAL.expr = Binding{Name: yyDollar[2].str, Value: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line scheme_parser.y:146
		{
			yyVAL.expr = []Expression{}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
//line scheme_parser.y:149
		{
			yyVAL.expr = append([]Expression{yyDollar[1].expr}, yyDollar[2].expr.([]Expression)...)
		}
	}
	goto yystack /* stack new state and value */
}
