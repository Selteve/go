package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	Config "gitee.com/under-my-umbrella/cloud/config"
)

// 生成token
func GenerateToken(ID int, username string) (string, error) {
    // 设置过期时间
    expirationTime := time.Now().Add(1 * time.Hour)

    // 创建自定义声明
    claims := Config.CustomClaims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
            Issuer:    "XiongHaiyin", // 发行者
            // 你可以设置更多的标准声明
        },
        ID:    ID,
        Username:  username,
    }
    // 使用HS256算法创建token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // 使用密钥签名token
    signedToken, err := token.SignedString([]byte("secretAtXiongHaiyin"))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}