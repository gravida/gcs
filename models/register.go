package models

import ()

// Register Code -
type Code struct {
	Id      int64  `json:"id"`
	Code    string `xorm:"varchar(20) notnull"`
	UserId  int64
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}

// Add
func AddCode(c *Code) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(c); err != nil {
		return err
	}

	return sess.Commit()
}

// query
func QueryCodeByCode(code string) (*Code, bool, error) {
	c := new(Code)
	c.Code = code
	has, err := x.Get(c)
	return c, has, err
}
