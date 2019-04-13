// Code generated by goyacc - DO NOT EDIT.

package bsonparse

import __yyfmt__ "fmt"

type pair struct {
	key string
	val interface{}
}

func setResult(l yyLexer, v []interface{}) {
	l.(*lex).result = v
}

type yySymType struct {
	yys  int
	obj  map[string]interface{}
	list []interface{}
	pair pair
	val  interface{}
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault     = 57359
	yyEofCode     = 57344
	BsonError     = 57347
	DBRef         = 57358
	ISODate       = 57352
	LexError      = 57346
	Literal       = 57350
	MaxKey        = 57357
	MinKey        = 57356
	Number        = 57349
	NumberDecimal = 57354
	NumberLong    = 57353
	ObjectID      = 57351
	String        = 57348
	Undefined     = 57355
	yyErrCode     = 57345

	yyMaxDepth = 200
	yyTabOfs   = -24
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		44:    0,  // ',' (26x)
		125:   1,  // '}' (21x)
		41:    2,  // ')' (20x)
		93:    3,  // ']' (19x)
		34:    4,  // '"' (9x)
		91:    5,  // '[' (6x)
		57360: 6,  // array (6x)
		123:   7,  // '{' (5x)
		57358: 8,  // DBRef (5x)
		57352: 9,  // ISODate (5x)
		57350: 10, // Literal (5x)
		57357: 11, // MaxKey (5x)
		57356: 12, // MinKey (5x)
		57349: 13, // Number (5x)
		57354: 14, // NumberDecimal (5x)
		57353: 15, // NumberLong (5x)
		57363: 16, // object (5x)
		57351: 17, // ObjectID (5x)
		57348: 18, // String (5x)
		57355: 19, // Undefined (5x)
		57365: 20, // value (5x)
		57344: 21, // $end (2x)
		58:    22, // ':' (2x)
		57364: 23, // pair (2x)
		40:    24, // '(' (1x)
		57361: 25, // elements (1x)
		57362: 26, // members (1x)
		57359: 27, // $default (0x)
		57347: 28, // BsonError (0x)
		57345: 29, // error (0x)
		57346: 30, // LexError (0x)
	}

	yySymNames = []string{
		"','",
		"'}'",
		"')'",
		"']'",
		"'\"'",
		"'['",
		"array",
		"'{'",
		"DBRef",
		"ISODate",
		"Literal",
		"MaxKey",
		"MinKey",
		"Number",
		"NumberDecimal",
		"NumberLong",
		"object",
		"ObjectID",
		"String",
		"Undefined",
		"value",
		"$end",
		"':'",
		"pair",
		"'('",
		"elements",
		"members",
		"$default",
		"BsonError",
		"error",
		"LexError",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {6, 3},
		2:  {16, 3},
		3:  {26, 0},
		4:  {26, 1},
		5:  {26, 3},
		6:  {23, 3},
		7:  {23, 5},
		8:  {25, 0},
		9:  {25, 1},
		10: {25, 3},
		11: {20, 1},
		12: {20, 3},
		13: {20, 1},
		14: {20, 1},
		15: {20, 1},
		16: {20, 4},
		17: {20, 2},
		18: {20, 2},
		19: {20, 2},
		20: {20, 1},
		21: {20, 1},
		22: {20, 1},
		23: {20, 4},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [46][]uint8{
		// 0
		{5: 26, 25},
		{21: 24},
		{16, 3: 16, 31, 26, 34, 28, 42, 36, 32, 41, 40, 30, 38, 37, 33, 35, 19: 39, 29, 25: 27},
		{68, 3: 67},
		{21, 21, 4: 57, 18: 56, 23: 55, 26: 54},
		// 5
		{15, 3: 15},
		{13, 13, 13, 13},
		{18: 52},
		{11, 11, 11, 11},
		{10, 10, 10, 10},
		// 10
		{9, 9, 9, 9},
		{24: 49},
		{2: 48},
		{2: 47},
		{2: 46},
		// 15
		{4, 4, 4, 4},
		{3, 3, 3, 3},
		{2, 2, 2, 2},
		{43},
		{18: 44},
		// 20
		{2: 45},
		{1, 1, 1, 1},
		{5, 5, 5, 5},
		{6, 6, 6, 6},
		{7, 7, 7, 7},
		// 25
		{4: 31, 26, 34, 28, 42, 36, 32, 41, 40, 30, 38, 37, 33, 35, 19: 39, 50},
		{2: 51},
		{8, 8, 8, 8},
		{4: 53},
		{12, 12, 12, 12},
		// 30
		{65, 64},
		{20, 20},
		{22: 62},
		{18: 58},
		{4: 59},
		// 35
		{22: 60},
		{4: 31, 26, 34, 28, 42, 36, 32, 41, 40, 30, 38, 37, 33, 35, 19: 39, 61},
		{17, 17},
		{4: 31, 26, 34, 28, 42, 36, 32, 41, 40, 30, 38, 37, 33, 35, 19: 39, 63},
		{18, 18},
		// 40
		{22, 22, 22, 22},
		{4: 57, 18: 56, 23: 66},
		{19, 19},
		{23, 23, 23, 23, 21: 23},
		{4: 31, 26, 34, 28, 42, 36, 32, 41, 40, 30, 38, 37, 33, 35, 19: 39, 69},
		// 45
		{14, 3: 14},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 29

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
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
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.val = yyS[yypt-1].list
			setResult(yylex, yyS[yypt-1].list)
		}
	case 2:
		{
			yyVAL.obj = yyS[yypt-1].obj
		}
	case 3:
		{
			yyVAL.obj = map[string]interface{}{}
		}
	case 4:
		{
			yyVAL.obj = map[string]interface{}{
				yyS[yypt-0].pair.key: yyS[yypt-0].pair.val,
			}
		}
	case 5:
		{
			yyS[yypt-2].obj[yyS[yypt-0].pair.key] = yyS[yypt-0].pair.val
			yyVAL.obj = yyS[yypt-2].obj
		}
	case 6:
		{
			yyVAL.pair = pair{key: yyS[yypt-2].val.(string), val: yyS[yypt-0].val}
		}
	case 7:
		{
			yyVAL.pair = pair{key: yyS[yypt-3].val.(string), val: yyS[yypt-0].val}
		}
	case 8:
		{
			yyVAL.list = []interface{}{}
		}
	case 9:
		{
			yyVAL.list = []interface{}{yyS[yypt-0].val}
		}
	case 10:
		{
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].val)
		}
	case 12:
		{
			yyVAL.val = yyS[yypt-1].val
		}
	case 14:
		{
			yyVAL.val = yyS[yypt-0].obj
		}
	case 16:
		{
			yyVAL.val = map[string]interface{}{"$oid": yyS[yypt-1].val}
		}
	case 17:
		{
			yyVAL.val = map[string]interface{}{"$date": yyS[yypt-1].val}
		}
	case 18:
		{
			yyVAL.val = map[string]interface{}{"$numberLong": yyS[yypt-1].val}
		}
	case 19:
		{
			yyVAL.val = map[string]interface{}{"$numberDecimal": yyS[yypt-1].val}
		}
	case 20:
		{
			yyVAL.val = map[string]interface{}{"$undefined": true}
		}
	case 21:
		{
			yyVAL.val = map[string]interface{}{"$minKey": true}
		}
	case 22:
		{
			yyVAL.val = map[string]interface{}{"$maxKey": true}
		}
	case 23:
		{
			yyVAL.val = map[string]interface{}{"$ref": yyS[yypt-3].val, "$id": yyS[yypt-1].val}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
