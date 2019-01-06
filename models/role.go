package models

import ()

// Role -
type Role struct {
	Id      int64  `json:"id"`
	Name    string `xorm:"UNIQUE NOT NULL" json:"name"`
	Enable  bool   `json:"enable"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// Add
func AddRole(r *Role) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(r); err != nil {
		return err
	}

	return sess.Commit()
}

// Del

// update
func UpdateRole(r *Role) (err error) {
	_, err = x.Id(r.Id).Update(r)
	return err
}

// validate name -
func ExistRoleByName(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&Role{Name: name})
}

// query
func QueryRoleByID(id int64) (*Role, bool, error) {
	role := new(Role)
	has, err := x.Id(id).Get(role)
	return role, has, err
}

func QueryAllRoles(page, pageSize int) ([]*Role, error) {
	roles := make([]*Role, 0, pageSize)
	return roles, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&roles)
}
