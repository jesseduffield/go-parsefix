package parsefix

import (
	"strings"
)

type fixer struct {
	match  func(*fixerContext) bool
	repair func(*fixerContext)
}

// fixerList is a list of all defined fixers.
// Initialized inside init() function.
var fixerList []fixer

func init() {
	fixerList = []fixer{
		missingByteFixer(
			`missing ',' before newline in composite literal`,
			',', '}'),

		// missingByteFixer(
		// 	`missing ',' in composite literal`,
		// 	','),

		missingByteFixer(
			`missing ',' in argument list`,
			',', ')'),

		missingByteFixer(
			`missing ',' in parameter list`,
			',', ')'),

		// missingByteFixer(
		// 	`expected ':', found newline`,
		// 	':'),

		// missingByteFixer(
		// 	`expected ';', found `, ';'),

		replacingFixer(
			`expected boolean or range expression, found assignment`,
			`:= `,
			`:= range `),
	}
}

func replacingFixer(errorPat, from, to string) fixer {
	return fixer{
		match: func(ctx *fixerContext) bool {
			return strings.Contains(ctx.issue, errorPat) &&
				ctx.contains(from)
		},
		repair: func(ctx *fixerContext) {
			ctx.replace(from, to)
		},
	}
}

func missingByteFixer(errorPat string, toInsert byte, nextNonWhitespace byte) fixer {
	return fixer{
		match: func(ctx *fixerContext) bool {
			return strings.Contains(ctx.issue, errorPat) && ctx.nextNonWhitespaceIs(nextNonWhitespace)
		},
		repair: func(ctx *fixerContext) {
			ctx.insertByte(toInsert)
		},
	}
}
