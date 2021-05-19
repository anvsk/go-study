package common

import (
    "go-ticket/pkg/store/db"

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
    // claims := jwt.ExtractClaims(c)

    u := User{}
    // u.ID = claims["ID"].(uint)
    // u.Username = claims["Username"].(string)
    return &u
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
