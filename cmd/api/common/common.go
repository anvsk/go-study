package common

import (
    "encoding/json"
    "go-ticket/pkg/store/db"
    "go-ticket/pkg/util"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

var IdentityKey = "ID"

type JwtPayload struct {
    User
}

type User struct {
    gorm.Model
    Username string `gorm:"index:idx_name,unique"`
    Password string
    Breif    Breif  // 一对一 hasone
    Tags     []Tag  `gorm:"many2many:user_tags;"`   // 多对多
    Friends  []User `gorm:"many2many:user_friends"` // 多对多[自引用]
    State    int
}

// many to many
type Tag struct {
    gorm.Model
    Name  string
    Users []User `gorm:"many2many:user_tags;"`
}

// hasone
type Breif struct {
    gorm.Model
    UserID    uint
    Introduce string
}

// 从jwt解析用户信息
func Uinfo(c *gin.Context) *User {
    claims := jwt.ExtractClaims(c)
    // payload := JwtPayload{}
    uinfo := &User{}
    util.Type2type(claims, &uinfo)
    // u.ID = claims["ID"].(uint)
    // u.Username = claims["Username"].(string)
    return uinfo
}

func InitUserTable() {
    db.Orm.AutoMigrate(&User{})
    db.Orm.AutoMigrate(&Breif{})
    issetUser := User{}
    db.Orm.Where("username='admin'").First(&issetUser)
    if issetUser.ID > 0 {
        return
    }
    // 在冲突时，更新除主键以外的所有列到新值。
    db.Orm.Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "name"}},
        UpdateAll: true,
    }).Create(&User{
        Username: "admin",
        Password: "admin",
        State:    1,
        Breif: Breif{
            Introduce: "管理员账号",
        },
        Tags: []Tag{
            {
                Name: "tag1",
            },
            {
                Name: "tag2",
            },
        },
    })
}

type login struct {
    Username string `form:"username" json:"username" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

func Authenticator(c *gin.Context) (interface{}, error) {
    var loginVals login
    if err := c.ShouldBind(&loginVals); err != nil {
        return "", jwt.ErrMissingLoginValues
    }
    var user User
    db.Orm.Model(&user).
        Where("username", loginVals.Username).
        Where("username", loginVals.Username).
        Preload("Tags").Preload("Breif").Preload("Friends").
        First(&user)
    if user.ID > 0 {
        return user, nil
    }
    return nil, jwt.ErrFailedAuthentication
}

func PayloadFunc(data interface{}) jwt.MapClaims {
    payload := jwt.MapClaims{}
    if v, ok := data.(User); ok {
        tmp, _ := json.Marshal(JwtPayload{User: v})
        json.Unmarshal(tmp, &payload)
        return payload
    }
    return jwt.MapClaims{}
}
