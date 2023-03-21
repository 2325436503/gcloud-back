package define

// JwtKey
type UserClaim struct {
	Id       int
	Identity string
	Name     string
	Email    string
	jwt.StandardClaims
}

var JwtKey = "JwtKey"
var MailPassword = "MailPassword"
var RedisPassword = ""
var RedisAddr = "redis"
var MySQLAddr = "mysql"
var MySQLPassword = "root"

// CodeLength 验证码长度
var CodeLength = 6

// CodeExpire 验证码过期时间（s）
var CodeExpire = 300

// TencentSecretKey 腾讯云对象存储

var TencentSecretID = "TencentSecretID"
var TencentSecretKey = "TencentSecretKey"
var CosBucket = "CosBucket"
var CosFolderName = ""
var AvatarBaseUrl = CosBucket + "/avatars/"

// PageSize 分页的默认参数
var PageSize = 20

var Datetime = "2000-01-01 00:00:01"

var TokenExpire = 60 * 60 * 24 * 3        // 3 days
var RefreshTokenExpire = 60 * 60 * 24 * 7 // 7 days

var UserRepositoryMaxSize = 1000 * 1024 * 1024  // 1GB
var PublicRepositoryMaxSize = 500 * 1024 * 1024 // 0.5GB
var UserRepositoryMinSize = 200 * 1024 * 1024   // 200MB
