package main

import(
    "github.com/golang-jwt/jwt/v5"
    "time"
    "fmt"
    "log"
)


type Claims struct {
    User string
    Pass string
    jwt.RegisteredClaims
}



var secret string = "jaomoqhoi45frjih8r2q8"


func RunJWT(){
    tokenString := buildJWT()
    if(tokenString != ""){
        parseJWT(tokenString)
    }
}

func buildJWT() (string){

    claims := Claims{
            "manger",
            "weiKing",
            jwt.RegisteredClaims{
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
                IssuedAt:  jwt.NewNumericDate(time.Now()),
                NotBefore: jwt.NewNumericDate(time.Now()),
                Issuer: "scw",
                Subject: "mail",
            },
        }


    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString([]byte(secret))

    if err != nil {
        fmt.Println(err)
        return ""
    }else{
        return tokenString
    }
}

func parseJWT(tokenString string){

    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        log.Fatal(err)
    } else if claimBack, ok := token.Claims.(*Claims); ok {
        fmt.Println(claimBack.User, claimBack.RegisteredClaims.Issuer)
    } else {
        log.Fatal("unknown claims type, cannot proceed")
    }
}