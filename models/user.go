package models

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"github.com/freelifer/goapi/pkg/utils"
	"golang.org/x/crypto/pbkdf2"
)

// User -
// Status 1: 未激活 2: 已激活 3: 冻结 4: 注销
type User struct {
	Id      int64  `json:"id"`
	Name    string `xorm:"varchar(20) notnull" json:"name"`
	Passwd  string `xorm:"varchar(32) notnull" json:"-"`
	Salt    string `xorm:"VARCHAR(10)" json:"-"`
	Email   string `xorm:"varchar(32) notnull" json:"email"`
	Phone   string `xorm:"varchar(11)" json:"phone"`
	Status  int    `xorm:"notnull default(1)" json:"status"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// Add
func AddUser(u *User) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(u); err != nil {
		return err
	}

	return sess.Commit()
}

// Del

// update
func UpdateUser(u *User) (err error) {
	_, err = x.Id(u.Id).Update(u)
	return err
}

// query
func QueryUserByID(id int64) (*User, bool, error) {
	user := new(User)
	has, err := x.Id(id).Get(user)
	return user, has, err
}

func QueryAllUsers(page, pageSize int) ([]*User, error) {
	users := make([]*User, 0, pageSize)
	return users, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&users)
}

// validate name -
func ExistUserByName(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&User{Name: name})
}

// validate email -
func ExistUserByEmail(uid int64, email string) (bool, error) {
	if len(email) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&User{Email: email})
}

// validate phone -
func ExistUserByPhone(uid int64, phone string) (bool, error) {
	if len(phone) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&User{Phone: phone})
}

//----------
// GetUserSalt returns a ramdom user salt token.
func (u *User) GetUserSaltAndEncodePasswd() error {
	value, err := utils.RandomString(10)
	if err != nil {
		return err
	}
	u.Salt = value

	u.EncodePasswd()
	return nil
}

// EncodePasswd encodes password to safe format.
func (u *User) EncodePasswd() {
	newPasswd := pbkdf2.Key([]byte(u.Passwd), []byte(u.Salt), 10000, 50, sha256.New)
	u.Passwd = fmt.Sprintf("%x", newPasswd)
}

// ValidatePassword checks if given password matches the one belongs to the user.
func (u *User) ValidatePassword(passwd string) bool {
	newUser := &User{Passwd: passwd, Salt: u.Salt}
	newUser.EncodePasswd()
	return subtle.ConstantTimeCompare([]byte(u.Passwd), []byte(newUser.Passwd)) == 1
}
