package tokenizer

func isWhitespaceChar(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func isIdentifierChar(r rune) bool {
	return isAlphaChar(r) || isNumberChar(r) || isSeparatorChar(r)
}

func isQuoteChar(r rune) bool {
	return r == '"'
}

func isLeftParenChar(r rune) bool {
	return r == '('
}

func isRightParenChar(r rune) bool {
	return r == ')'
}

/*
// TODO: Implement these as dict literals
func isLeftBraceChar(r rune) bool {
	return r == '{'
}

func isRightBraceChar(r rune) bool {
	return r == '}'
}

// TODO: Implement these as list literals
func isLeftBracketChar(r rune) bool {
	return r == '['
}

func isRightBracketChar(r rune) bool {
	return r == ']'
}
*/

func isFloatSeparatorChar(r rune) bool {
	return r == '.'
}

func isSeparatorChar(r rune) bool {
	return r == '_'
}

func isAlphaChar(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isNumberChar(r rune) bool {
	return r >= '0' && r <= '9'
}

func isOperatorChar(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '%' || r == '=' || r == '<' || r == '>' || r == '&' || r == '|' || r == '^' || r == '~' || r == '!'
}
