package main

import (
    "database/sql"
    "fmt"
    "runtime"
)

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
    if err != nil {
        fmt.Println(err)
    }
    //
    defer db.Close()
    
    // 用户模型, 表结构, 需要一个结构来接收查询结果集
    type User struct {
        Id int32
        Name string
        Age int8
    }
    
    // 保存用户信息列表
    var user User
	user.Id,user.Name,user.Age,err := findUserInfo(2)
	switch {
    case err == sql.ErrNoRows:
		fmt.Println("result is: ", err)
    case err != nil:
		fmt.Println("error is: ", err)
    }
	fmt.Println("result is",user.Id,user.Name,user.Age)
}	
	
	
	
func findUserInfo(id int) (Id int32, Name String,Age int8,err String){
	err := db.QueryRow(`SELECT id,name,age WHERE id = ?`, id).Scan(&user.Id, &user.Name, &user.Age,)
	return userTmp,err
}
