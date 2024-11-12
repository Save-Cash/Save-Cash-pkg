package config

import (
    "fmt"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/google/uuid"
    "github.com/joho/godotenv"
)

var SecretKey []byte

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found, using environment variables instead")
    }

    secret := os.Getenv("SECRET_KEY")
    if secret == "" {
        fmt.Println("SECRET_KEY is not set in environment variables.")
    }
    SecretKey = []byte(secret)
}

type Claims struct {
    UserID uuid.UUID `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID uuid.UUID) (string, error) {
    claims := Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
            Issuer:    "save-cash",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(SecretKey)
    if err != nil {
        return "", fmt.Errorf("could not create token: %v", err)
    }

    return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return SecretKey, nil
    })

    if err != nil {
        return nil, fmt.Errorf("could not parse token: %v", err)
    }

    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}
