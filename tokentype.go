package pyecon

type TokenType string

const (
	TokenTypePrincipalToken = TokenType("pye-principal-token")
	TokenTypeYieldToken     = TokenType("pye-yield-token")
)

func (receiver TokenType) String() string {
	return string(receiver)
}
