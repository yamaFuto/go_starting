package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//型定義
var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {

	//tableの生成
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT
	)`
	//コマンドを起動
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// //追加
	// //valueに関しては?が使える
	// cmd = "INSERT INTO person (name, age) VALUES (?,?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// //更新
	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 15, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//マルチ取り出す
	cmd = "SELECT * FROM person"
		// 返り値があるためQueryを使う
	rows, _ := DbConnection.Query(cmd)
	//次のレコードを持ってくるために閉じる
	defer rows.Close()
	var pp []Person
		//rowsに格納されたものをNext()で回す
	for rows.Next() {
		var p Person
		//検出結果から一つだけ取り出し、pに格納
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	//errorをまとめて集めている
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	// //シングル取り出し
	// cmd = "SELECT * FROM person where age = ?"
	// row := DbConnection.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	//singleにはerrの中でもErrNoRowsが存在して判定できる
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// //ないときには初期値の０などが返ってくる
	// fmt.Println(p.Name, p.Age)

	// //削除
	// cmd = "DELETE FROM person WHERE name = ?"
	// //結果がいらないためExecを使っている
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// //sql injection
	// tableName := "person"
	// //table名を指定する際には?ではなく%sを使わなくてはならず、sqlInjectionされる可能性がある
	// cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	// rows,_ := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }
}