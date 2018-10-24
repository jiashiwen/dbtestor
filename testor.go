package main

import (
	"database/sql"
	"fmt"
	"os"
	// randomdata "github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icrowley/fake"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var looptimes int = 100
	var databaseurl string = "root:Git785230@tcp(rm-m5e48r0qq973fjf5jjo.mysql.rds.aliyuncs.com:3306)/testdb?charset=utf8"
	db, err := sql.Open("mysql", databaseurl)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(200)
	checkErr(err)

	//查询数据
	// rows, err := db.Query("SELECT id,name,email,domain FROM company")
	// checkErr(err)

	// for rows.Next() {
	// 	var id string
	// 	var name string
	// 	var email string
	// 	var domain string
	// 	err = rows.Scan(&id, &name, &email, &domain)
	// 	fmt.Println(id, name, email, domain)
	// }
	//插入数据
	stmt, err := db.Prepare("INSERT INTO testtable SET number=?,name=?,email=?,image=?,updatetime=?")
	// checkErr(err)
	number := fake.DigitsN(10)
	name := fake.FullName()
	email := fake.EmailAddress()
	image := fake.CharactersN(4096)
	imagebytes := []byte(image)
	updatetime := randate()

	fmt.Println(number, name, email, image, updatetime)

	res, err := stmt.Exec(number, name, email, imagebytes, updatetime)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	var wg sync.WaitGroup

	if len(os.Args) == 2 {
		// looptimes, err := strconv.ParseInt(os.Args[1], 10, 64)
		looptimes, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	for i := 0; i < looptimes; i++ {
		wg.Add(1)
		go func() {
			defer wg.Add(-1)
			number := fake.DigitsN(10)
			name := fake.CharactersN(20)
			email := fake.EmailAddress()
			image := fake.CharactersN(20)
			imagebytes := []byte(image)
			updatetime := randate()

			res, err := stmt.Exec(number, name, email, imagebytes, updatetime)
			checkErr(err)
			affect, err := res.RowsAffected()
			checkErr(err)
			fmt.Println(affect)
			time.Sleep(500)

		}()
	}
	wg.Wait()
	// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09", "abc@bcd.com", "abc@bcd.com", "West")
	// checkErr(err)

	// fmt.Println(res)
	// err := fake.SetLang("zh")
	// if err != nil {
	// 	panic(err)
	// }

	// id := fake.CharactersN(18)
	// name := fake.FirstName()
	// sdate := randate()
	// email := randomdata.Email()
	// domain := randomdata.Email()
	// city := randomdata.City()

	// year := fake.Year(1918, 2030)
	// month := fake.MonthNum()
	// day := fake.Day()

	// for i := 0; i < 100; i++ {

	// 	id := fake.CharactersN(18)
	// 	name := fake.FirstName()
	// 	sdate := randate()
	// 	email := randomdata.Email()
	// 	domain := randomdata.Email()
	// 	city := randomdata.City()

	// 	res, err := stmt.Exec(id, name, sdate, email, domain, city)
	// 	checkErr(err)
	// 	affect, err := res.RowsAffected()
	// 	checkErr(err)

	// 	fmt.Println(affect)
	// 	time.Sleep(500)

	// }

	// fmt.Println(id)
	// fmt.Println(name)
	// fmt.Println(sdate.Format("2006-01-02"))
	// fmt.Println(email)
	// fmt.Println(domain)
	// fmt.Println(city)
	// fmt.Println(year)
	// fmt.Println(month)
	// fmt.Println(day)

	// fmt.Println(fake.GetLangs())
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func randate() time.Time {
	min := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
