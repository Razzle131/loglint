package logcheck

import (
	"bytes"
	"errors"
	"go/ast"
	"go/printer"
	"go/token"
	"slices"
	"strings"
	"unicode"

	"github.com/Razzle131/loglint/config"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/types/typeutil"
)

var (
	analyzer     *analysis.Analyzer
	funcs        []config.Func
	avoidedData  []string
	enabledRules int
)

func NewAnalyzer(cfg config.Config) *analysis.Analyzer {
	funcs = cfg.EnabledFuncs
	avoidedData = cfg.AvoidedData
	enabledRules = cfg.EnabledRules

	analyzer = &analysis.Analyzer{
		Name:     "loglint",
		Doc:      "checks logging calls",
		Run:      Run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return analyzer
}

func Run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			fn := typeutil.StaticCallee(pass.TypesInfo, call)
			if fn == nil {
				return true
			}

			idx := slices.IndexFunc(funcs, func(f config.Func) bool {
				return f.Name == fn.FullName()
			})
			if idx == -1 {
				return true
			}

			msg := call.Args[funcs[idx].MsgPos]
			params := call.Args[funcs[idx].ArgPos:]

			errs := checkArg(msg)
			for _, param := range params {
				errs = append(errs, checkArg(param)...)
			}

			for _, err := range errs {
				pass.Reportf(call.Pos(), err.Error()+": %q",
					render(pass.Fset, call))
			}

			return true
		})
	}

	return nil, nil
}

func checkArg(expr ast.Expr) []error {
	res := []error{}

	res = append(res, withTypeCheck(expr, checkLiteral)...)
	res = append(res, withTypeCheck(expr, checkIdent)...)
	res = append(res, withTypeCheck(expr, checkBinaryExpression)...)

	return res
}

const stringMinLen = 2

func checkLiteral(literal *ast.BasicLit) []error {
	if len(literal.Value) < stringMinLen {
		return []error{}
	}

	msg := literal.Value[1 : len(literal.Value)-1] // remove quotes
	checks := []error{checkFirstLetterCase(msg), checkEnglish(msg), checkSpecialSymbols(msg)}

	res := []error{}
	for _, err := range checks {
		if err == nil {
			continue
		}
		res = append(res, err)
	}

	return res
}

func checkIdent(ident *ast.Ident) []error {
	if err := checkSensetive(ident.Name); err != nil {
		return []error{err}
	}

	return []error{}
}

func checkBinaryExpression(binaryExpr *ast.BinaryExpr) []error {
	res := []error{}

	expr := binaryExpr
	for expr != nil {
		res = append(res, withTypeCheck(expr.Y, checkLiteral)...)
		res = append(res, withTypeCheck(expr.Y, checkIdent)...)

		res = append(res, withTypeCheck(expr.X, checkLiteral)...)
		res = append(res, withTypeCheck(expr.X, checkIdent)...)

		expr, _ = expr.X.(*ast.BinaryExpr)
	}

	return res
}

func checkFirstLetterCase(s string) error {
	if enabledRules>>0%2 == 0 {
		return nil
	}

	if len(s) < 1 {
		return nil
	}

	firstLetter := []rune(s)[0]
	if unicode.IsLetter(firstLetter) && !unicode.IsLower(firstLetter) {
		return errors.New("first letter must be in lower case")
	}

	return nil
}

func checkEnglish(s string) error {
	if enabledRules>>1%2 == 0 {
		return nil
	}

	if strings.ContainsFunc(s, func(r rune) bool {
		return r > unicode.MaxLatin1 && unicode.IsLetter(r)
	}) {
		return errors.New("message must contain only english letters")
	}

	return nil
}

func checkSpecialSymbols(s string) error {
	if enabledRules>>2%2 == 0 {
		return nil
	}

	if strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' '
	}) {
		return errors.New("message must not contain special symbols")
	}

	return nil
}

func checkSensetive(s string) error {
	if enabledRules>>3%2 == 0 {
		return nil
	}

	if slices.Contains(avoidedData, strings.ToLower(s)) {
		return errors.New("must not log sensitive data")
	}

	return nil
}

func render(fset *token.FileSet, x any) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
