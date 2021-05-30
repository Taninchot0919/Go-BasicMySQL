package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("It works!")
	// mysqlLogin()
	db, err := sql.Open("mysql", `root:mysql@tcp(168.63.232.208:3306)/PostDB`)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO Post(title,body) VALUES('My post','My body')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	selectId(db)

}

func mysqlLogin() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("mysql username : ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Println("Username :", username)

	fmt.Printf("mysql password : ")
	scanner.Scan()
	password := scanner.Text()
	fmt.Println("password :", password)

	return username, password
}

type Post struct { // สร้าง struct ของ db ไว้
	id    int
	title string
	body  string
}

func selectId(db *sql.DB) { //รับ db มาจาก func อื่น
	posts, err := db.Query("SELECT id, title, body FROM Post") // ใช้ func Query เพื่อนำ id title body ออกมาจาก database Post
	if err != nil {
		panic(err.Error())
	}
	for posts.Next() { // วน for ทั้งหมดของ row ใน db
		var post Post                                        // ตัวแปรของ struct
		err := posts.Scan(&post.id, &post.title, &post.body) //นำค่าเก็บไว้ที่ post ถ้าไม่มี error ค่าของ err จะเป็น nil รึป่าว อันนี้งงๆ
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(post) //ปริ้นค่า post ออกมา
	}
}
