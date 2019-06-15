package msp

import "time"

/*Identity 接口 定义操作包含 证书。
提供签名认证。
*/
type Identity interface {
	ExpiresAt() time.Time
	GetIdentifier() *IdentityIdentifier
	GetMSPIdentifier() string

	Validate() error
	//匿名
	Anonymous() bool

	//使用作为引用的identify来给消息进行签名
	Verify(msg []byte, sig[]byte) error

	//将签名转化为字符串
	Serialize()([]byte,error)
}

// IdentityIdentifier is a holder for the identifier of a specific
// identity, naturally namespaced, by its provider identifier.
type IdentityIdentifier struct {

	// The identifier of the associated membership service provider
	Mspid string

	// The identifier for an identity within a provider
	Id string
}

type SigningIdentity interface {

	// Extends Identity
	Identity

	// Sign the message
	Sign(msg []byte) ([]byte, error)

	// GetPublicVersion returns the public parts of this identity
	GetPublicVersion() Identity
}

//证书提供者的类型
type ProviderType int


const (
	FABRIC ProviderType = iota // MSP is of FABRIC type
	IDEMIX                     // MSP is of IDEMIX type
	OTHER                      // MSP is of OTHER TYPE

	// NOTE: as new types are added to this set,
	// the mspTypes map below must be extended
)

var mspTypeStrings = map[ProviderType]string{
	FABRIC: "bccsp",
	IDEMIX: "idemix",
}

// 返回一个代表签名者类型的字符串
func ProviderTypeToString(id ProviderType) string{
	if res,found:=mspTypeStrings[id]; found {
		return res
	}

	return ""
}