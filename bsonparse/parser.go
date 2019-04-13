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
	yyDefault     = 57360
	yyEofCode     = 57344
	BinData       = 57359
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
	yyTabOfs   = -25
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		44:    0,  // ',' (28x)
		125:   1,  // '}' (22x)
		34:    2,  // '"' (20x)
		93:    3,  // ']' (20x)
		57348: 4,  // String (12x)
		40:    5,  // '(' (6x)
		41:    6,  // ')' (6x)
		91:    7,  // '[' (5x)
		57361: 8,  // array (5x)
		123:   9,  // '{' (4x)
		57359: 10, // BinData (4x)
		57358: 11, // DBRef (4x)
		57352: 12, // ISODate (4x)
		57350: 13, // Literal (4x)
		57357: 14, // MaxKey (4x)
		57356: 15, // MinKey (4x)
		57349: 16, // Number (4x)
		57354: 17, // NumberDecimal (4x)
		57353: 18, // NumberLong (4x)
		57364: 19, // object (4x)
		57351: 20, // ObjectID (4x)
		57355: 21, // Undefined (4x)
		57366: 22, // value (4x)
		57344: 23, // $end (2x)
		58:    24, // ':' (2x)
		57365: 25, // pair (2x)
		57362: 26, // elements (1x)
		57363: 27, // members (1x)
		57360: 28, // $default (0x)
		57347: 29, // BsonError (0x)
		57345: 30, // error (0x)
		57346: 31, // LexError (0x)
	}

	yySymNames = []string{
		"','",
		"'}'",
		"'\"'",
		"']'",
		"String",
		"'('",
		"')'",
		"'['",
		"array",
		"'{'",
		"BinData",
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
		"Undefined",
		"value",
		"$end",
		"':'",
		"pair",
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
		1:  {8, 3},
		2:  {19, 3},
		3:  {27, 0},
		4:  {27, 1},
		5:  {27, 3},
		6:  {25, 3},
		7:  {25, 5},
		8:  {26, 0},
		9:  {26, 1},
		10: {26, 3},
		11: {22, 1},
		12: {22, 3},
		13: {22, 1},
		14: {22, 1},
		15: {22, 1},
		16: {22, 6},
		17: {22, 6},
		18: {22, 6},
		19: {22, 6},
		20: {22, 1},
		21: {22, 1},
		22: {22, 1},
		23: {22, 10},
		24: {22, 6},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [72][]uint8{
		// 0
		{7: 27, 26},
		{23: 25},
		{17, 2: 32, 17, 7: 27, 35, 29, 44, 43, 37, 33, 42, 41, 31, 39, 38, 34, 36, 40, 30, 26: 28},
		{95, 3: 94},
		{22, 22, 84, 4: 83, 25: 82, 27: 81},
		// 5
		{16, 3: 16},
		{14, 14, 3: 14},
		{4: 79},
		{12, 12, 3: 12},
		{11, 11, 3: 11},
		// 10
		{10, 10, 3: 10},
		{5: 74},
		{5: 69},
		{5: 64},
		{5: 59},
		// 15
		{5, 5, 3: 5},
		{4, 4, 3: 4},
		{3, 3, 3: 3},
		{5: 50},
		{5: 45},
		// 20
		{4: 46},
		{47},
		{4: 48},
		{6: 49},
		{1, 1, 3: 1},
		// 25
		{2: 51},
		{4: 52},
		{2: 53},
		{54},
		{2: 55},
		// 30
		{4: 56},
		{2: 57},
		{6: 58},
		{2, 2, 3: 2},
		{2: 60},
		// 35
		{4: 61},
		{2: 62},
		{6: 63},
		{6, 6, 3: 6},
		{2: 65},
		// 40
		{4: 66},
		{2: 67},
		{6: 68},
		{7, 7, 3: 7},
		{2: 70},
		// 45
		{4: 71},
		{2: 72},
		{6: 73},
		{8, 8, 3: 8},
		{2: 75},
		// 50
		{4: 76},
		{2: 77},
		{6: 78},
		{9, 9, 3: 9},
		{2: 80},
		// 55
		{13, 13, 3: 13},
		{92, 91},
		{21, 21},
		{24: 89},
		{4: 85},
		// 60
		{2: 86},
		{24: 87},
		{2: 32, 7: 27, 35, 29, 44, 43, 37, 33, 42, 41, 31, 39, 38, 34, 36, 40, 88},
		{18, 18},
		{2: 32, 7: 27, 35, 29, 44, 43, 37, 33, 42, 41, 31, 39, 38, 34, 36, 40, 90},
		// 65
		{19, 19},
		{23, 23, 3: 23},
		{2: 84, 4: 83, 25: 93},
		{20, 20},
		{24, 24, 3: 24, 23: 24},
		// 70
		{2: 32, 7: 27, 35, 29, 44, 43, 37, 33, 42, 41, 31, 39, 38, 34, 36, 40, 96},
		{15, 3: 15},
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
	const yyError = 30

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
			yyVAL.val = map[string]interface{}{"$oid": yyS[yypt-2].val}
		}
	case 17:
		{
			yyVAL.val = map[string]interface{}{"$date": yyS[yypt-2].val}
		}
	case 18:
		{
			yyVAL.val = map[string]interface{}{"$numberLong": yyS[yypt-2].val}
		}
	case 19:
		{
			yyVAL.val = map[string]interface{}{"$numberDecimal": yyS[yypt-2].val}
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
			yyVAL.val = map[string]interface{}{"$ref": yyS[yypt-6].val, "$id": yyS[yypt-2].val}
		}
	case 24:
		{
			yyVAL.val = map[string]interface{}{"$binary": yyS[yypt-1].val, "$type": yyS[yypt-3].val}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
