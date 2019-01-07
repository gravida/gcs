package models

import ()

const (
	MENU      = 1
	OPERATION = 2
)

// Role-Permission -
type RolePermission struct {
	Id      int64 `json:"id"`
	RId     int64
	TId     int64
	Type    int
	Created int64 `xorm:"created" json:"created"`
	Updated int64 `xorm:"updated" json:"updated"`
}

// Menu
type Menu struct {
	Id      int64 `json:"id"`
	PId     int64
	Name    string `xorm:"UNIQUE NOT NULL" json:"name"`
	Desc    string `xorm:"varchar(32)" json:"desc"`
	Key     string `json:"key"`
	Enable  bool   `json:"enable"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// Operation
type Operation struct {
	Id      int64  `json:"id"`
	Name    string `xorm:"UNIQUE NOT NULL" json:"name"`
	Desc    string `xorm:"varchar(32)" json:"desc"`
	Key     string `json:"key"`
	Enable  bool   `json:"enable"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

type PermissionOp struct {
	Id      int64 `json:"id"`
	TId     int64
	Type    int
	Name    string `xorm:"UNIQUE NOT NULL" json:"name"`
	Desc    string `xorm:"varchar(32)" json:"desc"`
	Key     string `json:"key"`
	Enable  bool   `json:"enable"`
	Created int64  `xorm:"created" json:"created"`
	Updated int64  `xorm:"updated" json:"updated"`
}

// Add
// func AddPermission(p *Permission) (err error) {
// 	sess := x.NewSession()
// 	defer sess.Close()
// 	if err = sess.Begin(); err != nil {
// 		return err
// 	}

// 	if _, err = sess.Insert(p); err != nil {
// 		return err
// 	}

// 	return sess.Commit()
// }

// Add Op and Permisson
func AddOperation(op *Operation) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(op); err != nil {
		return err
	}

	return sess.Commit()
}

// Del

// update

func UpdateOperation(op *Operation) (err error) {
	_, err = x.Id(op.Id).Update(op)
	return err
}

func UpdatePermission(op *Operation) (err error) {
	_, err = x.Id(op.Id).Update(op)
	return err
}

// // validate name -
func ExistOperationByName(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&Operation{Name: name})
}

// // validate key -
func ExistOperationByKey(uid int64, key string) (bool, error) {
	if len(key) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&Operation{Key: key})
}

// count
func CountOperations() (total int64, err error) {
	operation := new(Operation)
	total, err = x.Count(operation)
	return total, err
}

// query
func QueryOperationByID(id int64) (*Operation, bool, error) {
	operation := new(Operation)
	has, err := x.Id(id).Get(operation)
	return operation, has, err
}

// query all
func QueryAllOperations(page, pageSize int) ([]*Operation, error) {
	operations := make([]*Operation, 0, pageSize)
	return operations, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&operations)
	// return permissionOps, x.Sql("select * from Permission, Operation where Permission.t_id=Operation.id and Permission.type=2").Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&permissionOps)
}
