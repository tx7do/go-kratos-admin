// Code generated by ent, DO NOT EDIT.

package usercredential

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the usercredential type in the database.
	Label = "user_credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldIdentityType holds the string denoting the identity_type field in the database.
	FieldIdentityType = "identity_type"
	// FieldIdentifier holds the string denoting the identifier field in the database.
	FieldIdentifier = "identifier"
	// FieldCredentialType holds the string denoting the credential_type field in the database.
	FieldCredentialType = "credential_type"
	// FieldCredential holds the string denoting the credential field in the database.
	FieldCredential = "credential"
	// FieldIsPrimary holds the string denoting the is_primary field in the database.
	FieldIsPrimary = "is_primary"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExtraInfo holds the string denoting the extra_info field in the database.
	FieldExtraInfo = "extra_info"
	// FieldActivateToken holds the string denoting the activate_token field in the database.
	FieldActivateToken = "activate_token"
	// FieldResetToken holds the string denoting the reset_token field in the database.
	FieldResetToken = "reset_token"
	// Table holds the table name of the usercredential in the database.
	Table = "user_credentials"
)

// Columns holds all SQL columns for usercredential fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDeleteTime,
	FieldTenantID,
	FieldUserID,
	FieldIdentityType,
	FieldIdentifier,
	FieldCredentialType,
	FieldCredential,
	FieldIsPrimary,
	FieldStatus,
	FieldExtraInfo,
	FieldActivateToken,
	FieldResetToken,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// TenantIDValidator is a validator for the "tenant_id" field. It is called by the builders before save.
	TenantIDValidator func(uint32) error
	// IdentifierValidator is a validator for the "identifier" field. It is called by the builders before save.
	IdentifierValidator func(string) error
	// CredentialValidator is a validator for the "credential" field. It is called by the builders before save.
	CredentialValidator func(string) error
	// DefaultIsPrimary holds the default value on creation for the "is_primary" field.
	DefaultIsPrimary bool
	// ActivateTokenValidator is a validator for the "activate_token" field. It is called by the builders before save.
	ActivateTokenValidator func(string) error
	// ResetTokenValidator is a validator for the "reset_token" field. It is called by the builders before save.
	ResetTokenValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint32) error
)

// IdentityType defines the type for the "identity_type" enum field.
type IdentityType string

// IdentityTypeUsername is the default value of the IdentityType enum.
const DefaultIdentityType = IdentityTypeUsername

// IdentityType values.
const (
	IdentityTypeUsername    IdentityType = "USERNAME"
	IdentityTypeUserId      IdentityType = "USERID"
	IdentityTypeEmail       IdentityType = "EMAIL"
	IdentityTypePhone       IdentityType = "PHONE"
	IdentityTypeWechat      IdentityType = "WECHAT"
	IdentityTypeQQ          IdentityType = "QQ"
	IdentityTypeWeibo       IdentityType = "WEIBO"
	IdentityTypeDouYin      IdentityType = "DOUYIN"
	IdentityTypeKuaiShou    IdentityType = "KUAISHOU"
	IdentityTypeBaidu       IdentityType = "BAIDU"
	IdentityTypeAlipay      IdentityType = "ALIPAY"
	IdentityTypeTaoBao      IdentityType = "TAOBAO"
	IdentityTypeJD          IdentityType = "JD"
	IdentityTypeMeiTuan     IdentityType = "MEITUAN"
	IdentityTypeDingTalk    IdentityType = "DINGTALK"
	IdentityTypeBiliBili    IdentityType = "BILIBILI"
	IdentityTypeXiaohongshu IdentityType = "XIAOHONGSHU"
	IdentityTypeGoogle      IdentityType = "GOOGLE"
	IdentityTypeFacebook    IdentityType = "FACEBOOK"
	IdentityTypeApple       IdentityType = "APPLE"
	IdentityTypeTelegram    IdentityType = "TELEGRAM"
	IdentityTypeTwitter     IdentityType = "TWITTER"
	IdentityTypeLinkedIn    IdentityType = "LINKEDIN"
	IdentityTypeGitHub      IdentityType = "GITHUB"
	IdentityTypeMicrosoft   IdentityType = "MICROSOFT"
	IdentityTypeDiscord     IdentityType = "DISCORD"
	IdentityTypeSlack       IdentityType = "SLACK"
	IdentityTypeInstagram   IdentityType = "INSTAGRAM"
	IdentityTypeTikTok      IdentityType = "TIKTOK"
	IdentityTypeReddit      IdentityType = "REDDIT"
	IdentityTypeYouTube     IdentityType = "YOUTUBE"
	IdentityTypeSpotify     IdentityType = "SPOTIFY"
	IdentityTypePinterest   IdentityType = "PINTEREST"
	IdentityTypeSnapchat    IdentityType = "SNAPCHAT"
	IdentityTypeTumblr      IdentityType = "TUMBLR"
	IdentityTypeYahoo       IdentityType = "YAHOO"
	IdentityTypeWhatsApp    IdentityType = "WHATSAPP"
	IdentityTypeLINE        IdentityType = "LINE"
)

func (it IdentityType) String() string {
	return string(it)
}

// IdentityTypeValidator is a validator for the "identity_type" field enum values. It is called by the builders before save.
func IdentityTypeValidator(it IdentityType) error {
	switch it {
	case IdentityTypeUsername, IdentityTypeUserId, IdentityTypeEmail, IdentityTypePhone, IdentityTypeWechat, IdentityTypeQQ, IdentityTypeWeibo, IdentityTypeDouYin, IdentityTypeKuaiShou, IdentityTypeBaidu, IdentityTypeAlipay, IdentityTypeTaoBao, IdentityTypeJD, IdentityTypeMeiTuan, IdentityTypeDingTalk, IdentityTypeBiliBili, IdentityTypeXiaohongshu, IdentityTypeGoogle, IdentityTypeFacebook, IdentityTypeApple, IdentityTypeTelegram, IdentityTypeTwitter, IdentityTypeLinkedIn, IdentityTypeGitHub, IdentityTypeMicrosoft, IdentityTypeDiscord, IdentityTypeSlack, IdentityTypeInstagram, IdentityTypeTikTok, IdentityTypeReddit, IdentityTypeYouTube, IdentityTypeSpotify, IdentityTypePinterest, IdentityTypeSnapchat, IdentityTypeTumblr, IdentityTypeYahoo, IdentityTypeWhatsApp, IdentityTypeLINE:
		return nil
	default:
		return fmt.Errorf("usercredential: invalid enum value for identity_type field: %q", it)
	}
}

// CredentialType defines the type for the "credential_type" enum field.
type CredentialType string

// CredentialTypePasswordHash is the default value of the CredentialType enum.
const DefaultCredentialType = CredentialTypePasswordHash

// CredentialType values.
const (
	CredentialTypePasswordHash               CredentialType = "PASSWORD_HASH"
	CredentialTypeAccessToken                CredentialType = "ACCESS_TOKEN"
	CredentialTypeRefreshToken               CredentialType = "REFRESH_TOKEN"
	CredentialTypeEmailVerificationCode      CredentialType = "EMAIL_VERIFICATION_CODE"
	CredentialTypePhoneVerificationCode      CredentialType = "PHONE_VERIFICATION_CODE"
	CredentialTypeOauthToken                 CredentialType = "OAUTH_TOKEN"
	CredentialTypeApiKey                     CredentialType = "API_KEY"
	CredentialTypeSsoToken                   CredentialType = "SSO_TOKEN"
	CredentialTypeJWT                        CredentialType = "JWT"
	CredentialTypeSamlAssertion              CredentialType = "SAML_ASSERTION"
	CredentialTypeOpenidConnectIdToken       CredentialType = "OPENID_CONNECT_ID_TOKEN"
	CredentialTypeSessionCookie              CredentialType = "SESSION_COOKIE"
	CredentialTypeTemporaryCredential        CredentialType = "TEMPORARY_CREDENTIAL"
	CredentialTypeCustomCredential           CredentialType = "CUSTOM_CREDENTIAL"
	CredentialTypeBiometricData              CredentialType = "BIOMETRIC_DATA"
	CredentialTypeSecurityKey                CredentialType = "SECURITY_KEY"
	CredentialTypeOTP                        CredentialType = "OTP"
	CredentialTypeSmartCard                  CredentialType = "SMART_CARD"
	CredentialTypeCryptographicCertificate   CredentialType = "CRYPTOGRAPHIC_CERTIFICATE"
	CredentialTypeBiometricToken             CredentialType = "BIOMETRIC_TOKEN"
	CredentialTypeDeviceFingerprint          CredentialType = "DEVICE_FINGERPRINT"
	CredentialTypeHardwareToken              CredentialType = "HARDWARE_TOKEN"
	CredentialTypeSoftwareToken              CredentialType = "SOFTWARE_TOKEN"
	CredentialTypeSecurityQuestion           CredentialType = "SECURITY_QUESTION"
	CredentialTypeSecurityPin                CredentialType = "SECURITY_PIN"
	CredentialTypeTwoFactorAuthentication    CredentialType = "TWO_FACTOR_AUTHENTICATION"
	CredentialTypeMultiFactorAuthentication  CredentialType = "MULTI_FACTOR_AUTHENTICATION"
	CredentialTypePasswordlessAuthentication CredentialType = "PASSWORDLESS_AUTHENTICATION"
	CredentialTypeSocialLoginToken           CredentialType = "SOCIAL_LOGIN_TOKEN"
	CredentialTypeSsoSession                 CredentialType = "SSO_SESSION"
	CredentialTypeApiSecret                  CredentialType = "API_SECRET"
	CredentialTypeCustomToken                CredentialType = "CUSTOM_TOKEN"
	CredentialTypeOauth2ClientCredentials    CredentialType = "OAUTH2_CLIENT_CREDENTIALS"
	CredentialTypeOauth2AuthorizationCode    CredentialType = "OAUTH2_AUTHORIZATION_CODE"
	CredentialTypeOauth2ImplicitGrant        CredentialType = "OAUTH2_IMPLICIT_GRANT"
	CredentialTypeOauth2PasswordGrant        CredentialType = "OAUTH2_PASSWORD_GRANT"
	CredentialTypeOauth2RefreshGrant         CredentialType = "OAUTH2_REFRESH_GRANT"
)

func (ct CredentialType) String() string {
	return string(ct)
}

// CredentialTypeValidator is a validator for the "credential_type" field enum values. It is called by the builders before save.
func CredentialTypeValidator(ct CredentialType) error {
	switch ct {
	case CredentialTypePasswordHash, CredentialTypeAccessToken, CredentialTypeRefreshToken, CredentialTypeEmailVerificationCode, CredentialTypePhoneVerificationCode, CredentialTypeOauthToken, CredentialTypeApiKey, CredentialTypeSsoToken, CredentialTypeJWT, CredentialTypeSamlAssertion, CredentialTypeOpenidConnectIdToken, CredentialTypeSessionCookie, CredentialTypeTemporaryCredential, CredentialTypeCustomCredential, CredentialTypeBiometricData, CredentialTypeSecurityKey, CredentialTypeOTP, CredentialTypeSmartCard, CredentialTypeCryptographicCertificate, CredentialTypeBiometricToken, CredentialTypeDeviceFingerprint, CredentialTypeHardwareToken, CredentialTypeSoftwareToken, CredentialTypeSecurityQuestion, CredentialTypeSecurityPin, CredentialTypeTwoFactorAuthentication, CredentialTypeMultiFactorAuthentication, CredentialTypePasswordlessAuthentication, CredentialTypeSocialLoginToken, CredentialTypeSsoSession, CredentialTypeApiSecret, CredentialTypeCustomToken, CredentialTypeOauth2ClientCredentials, CredentialTypeOauth2AuthorizationCode, CredentialTypeOauth2ImplicitGrant, CredentialTypeOauth2PasswordGrant, CredentialTypeOauth2RefreshGrant:
		return nil
	default:
		return fmt.Errorf("usercredential: invalid enum value for credential_type field: %q", ct)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusEnabled is the default value of the Status enum.
const DefaultStatus = StatusEnabled

// Status values.
const (
	StatusDisabled   Status = "DISABLED"
	StatusEnabled    Status = "ENABLED"
	StatusExpired    Status = "EXPIRED"
	StatusUnverified Status = "UNVERIFIED"
	StatusRemoved    Status = "REMOVED"
	StatusBlocked    Status = "BLOCKED"
	StatusTemporary  Status = "TEMPORARY"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusDisabled, StatusEnabled, StatusExpired, StatusUnverified, StatusRemoved, StatusBlocked, StatusTemporary:
		return nil
	default:
		return fmt.Errorf("usercredential: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the UserCredential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByIdentityType orders the results by the identity_type field.
func ByIdentityType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentityType, opts...).ToFunc()
}

// ByIdentifier orders the results by the identifier field.
func ByIdentifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentifier, opts...).ToFunc()
}

// ByCredentialType orders the results by the credential_type field.
func ByCredentialType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredentialType, opts...).ToFunc()
}

// ByCredential orders the results by the credential field.
func ByCredential(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredential, opts...).ToFunc()
}

// ByIsPrimary orders the results by the is_primary field.
func ByIsPrimary(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPrimary, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByExtraInfo orders the results by the extra_info field.
func ByExtraInfo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtraInfo, opts...).ToFunc()
}

// ByActivateToken orders the results by the activate_token field.
func ByActivateToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActivateToken, opts...).ToFunc()
}

// ByResetToken orders the results by the reset_token field.
func ByResetToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResetToken, opts...).ToFunc()
}
