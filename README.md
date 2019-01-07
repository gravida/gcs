# gcs


```

权限操作关键字格式
	添加：A
	修改：U
	删除：D
	查找：Q

```

```
列表返回格式
{
	“pager”: {
		"page":1,
		"pageSize": 10,
		"total": 11,
	},
	"data":数组
}
```

CREATE TABLE Permission
(
id int primary key,
t_id int,
type int
);

CREATE TABLE Operation
(
id int primary key,
name varchar(255),
desc1 varchar(255),
key1 varchar(255)
);


INSERT INTO Permission VALUES (1, 1, 1);
INSERT INTO Permission VALUES (2, 2, 1);
INSERT INTO Permission VALUES (3, 1, 2);


INSERT INTO Operation VALUES (1, "添加用户", "add user ....", "add-user");
INSERT INTO Operation VALUES (2, "删除用户", "del user ....", "del-user");


select *,Permission.id from Operation,Permission where Operation.id 
in (select Permission.t_id from Permission where Permission.type=1);
select Permission.*,Operation.* from Permission
left
join 
Operation 
ON Permission.t_id=Operation.id
where Permission.type=1 
order by Permission.id asc limit 2 ;

select * from Permission, Operation where Permission.t_id = Operation.id
and Permission.type = 1;


select rowid,t.* from ( select *,(select count(1) from Permission) count from Permission ) t;