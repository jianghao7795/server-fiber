package utils

import (
	errors "errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	ErrNotECPublicKey  = errors.New("密钥不是有效的 ECDSA 公钥")
	ErrNotECPrivateKey = errors.New("密钥不是有效的 ECDSA 私钥")
)

var (
	ErrNotEdPrivateKey = errors.New("密钥不是有效的 Ed25519私钥")
	ErrNotEdPublicKey  = errors.New("密钥不是有效的 Ed25519公钥")
)

var (
	ErrInvalidKey                = errors.New("token密钥无效")
	ErrInvalidKeyType            = errors.New("token密钥的类型无效")
	ErrHashUnavailable           = errors.New("请求的哈希函数不可用")
	ErrTokenMalformed            = errors.New("token无效")
	ErrTokenUnverifiable         = errors.New("令牌无法验证")
	ErrTokenSignatureInvalid     = errors.New("令牌签名无效")
	ErrTokenRequiredClaimMissing = errors.New("令牌丢失")
	ErrTokenInvalidAudience      = errors.New("令牌的访问者无效")
	ErrTokenExpired              = errors.New("令牌过期了")
	ErrTokenUsedBeforeIssued     = errors.New("发出前使用的标记")
	ErrTokenInvalidIssuer        = errors.New("令牌的发行者无效")
	ErrTokenInvalidSubject       = errors.New("令牌的主题无效")
	ErrTokenNotValidYet          = errors.New("令牌尚未有效")
	ErrTokenInvalidId            = errors.New("令牌的 id 无效")
	ErrTokenInvalidClaims        = errors.New("令牌的索赔无效")
	ErrInvalidType               = errors.New("无效索赔类型")
)

var (
	ErrKeyMustBePEMEncoded = errors.New("无效密钥: 密钥必须是 PEM 编码的 PKCS1或 PKCS8密钥")
	ErrNotRSAPrivateKey    = errors.New("密钥不是有效的 RSA 私钥")
	ErrNotRSAPublicKey     = errors.New("密钥不是有效的 RSA 公钥")
)

var (
	ErrECDSAVerification = errors.New("crypto/ecdsa: 验证错误")
)

var (
	ErrEd25519Verification = errors.New("Ed25519: 验证错误")
)

func ReportError(err error) error {
	switch {
	case errors.Is(err, jwt.ErrNotECPublicKey):
		return ErrNotECPublicKey
	case errors.Is(err, jwt.ErrNotECPrivateKey):
		return ErrNotECPrivateKey
	case errors.Is(err, jwt.ErrNotEdPrivateKey):
		return ErrNotEdPrivateKey
	case errors.Is(err, jwt.ErrNotEdPublicKey):
		return ErrNotEdPublicKey
	case errors.Is(err, jwt.ErrInvalidKey):
		return ErrInvalidKey
	case errors.Is(err, jwt.ErrInvalidKeyType):
		return ErrInvalidKeyType
	case errors.Is(err, jwt.ErrHashUnavailable):
		return ErrHashUnavailable
	case errors.Is(err, jwt.ErrTokenMalformed):
		return ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenUnverifiable):
		return ErrTokenUnverifiable
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return ErrTokenSignatureInvalid
	case errors.Is(err, jwt.ErrTokenRequiredClaimMissing):
		return ErrTokenRequiredClaimMissing
	case errors.Is(err, jwt.ErrTokenInvalidAudience):
		return ErrTokenInvalidAudience
	case errors.Is(err, jwt.ErrTokenExpired):
		return ErrTokenExpired
	case errors.Is(err, jwt.ErrTokenUsedBeforeIssued):
		return ErrTokenUsedBeforeIssued
	case errors.Is(err, jwt.ErrTokenInvalidIssuer):
		return ErrTokenInvalidIssuer
	case errors.Is(err, jwt.ErrTokenInvalidSubject):
		return ErrTokenInvalidSubject
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return ErrTokenNotValidYet
	case errors.Is(err, jwt.ErrTokenInvalidId):
		return ErrTokenInvalidId
	case errors.Is(err, jwt.ErrTokenInvalidClaims):
		return ErrTokenInvalidClaims
	case errors.Is(err, jwt.ErrInvalidType):
		return ErrInvalidType
	case errors.Is(err, jwt.ErrKeyMustBePEMEncoded):
		return ErrKeyMustBePEMEncoded
	case errors.Is(err, jwt.ErrNotRSAPrivateKey):
		return ErrNotRSAPrivateKey
	case errors.Is(err, jwt.ErrNotRSAPublicKey):
		return ErrNotRSAPublicKey
	case errors.Is(err, jwt.ErrECDSAVerification):
		return ErrECDSAVerification
	case errors.Is(err, jwt.ErrNotRSAPublicKey):
		return ErrEd25519Verification
	default:
		return ErrInvalidKey
	}

}
