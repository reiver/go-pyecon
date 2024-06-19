package pyecon

type TokenType string

const (
	TokenTypePrincipalToken = TokenType("pye-principal-token")
	TokenTypeYieldToken     = TokenType("pye-yield-token")
)
