package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Person struct {
	Id           int
	Name         string `json:"name" form:"name"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	City         string `json:"city" form:"city"`
	State        string `json:"state" form:"state"`
	Street1      string `json:"street1" form:"street1"`
	Street2      string `json:"street2" form:"street2"`
	Zip_code     string `json:"zip_code" form:"zip_code"`
}

func (p Person) get() (person Person, err error) {

	row := db.QueryRow("SELECT per.id, per.name, ph.number as phone_number, ad.city, ad.state, ad.street1, ad.street2, ad.zip_code FROM PERSON per left outer join phone ph on per.id=ph.person_id left outer join address_join adj on per.id=adj.person_id left outer join address ad on adj.address_id=ad.id where per.id=?", p.Id)
	err = row.Scan(&person.Id, &person.Name, &person.Phone_number, &person.City, &person.State, &person.Street1, &person.Street2, &person.Zip_code)
	if err != nil {
		return
	}
	return
}

func (p Person) getAll() (persons []Person, err error) {
	rows, err := db.Query("SELECT per.name, ph.number as phone_number, ad.city, ad.state, ad.street1, ad.street2, ad.zip_code FROM PERSON per left outer join phone ph on per.id=ph.person_id left outer join address_join adj on per.id=adj.person_id left outer join address ad on adj.address_id=ad.id")
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Name, &person.Phone_number, &person.City, &person.State, &person.Street1, &person.Street2, &person.Zip_code)
		persons = append(persons, person)
	}
	defer rows.Close()
	return
}

func (p Person) add() (Id int, Id2 int, Id3 int, Id4 int, err error) {
	stmt, err := db.Prepare("INSERT INTO person(name) VALUES (?)")
	stmt2, err := db.Prepare("INSERT INTO phone(number, person_id) VALUES (?, ?)")
	stmt3, err := db.Prepare("INSERT INTO address(city, state, street1, street2, zip_code) VALUES (?,?,?,?,?)")
	stmt4, err := db.Prepare("Insert INTO address_join(person_id, address_id) VALUES (?,?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Name)

	rs3, err := stmt3.Exec(p.City, p.State, p.Street1, p.Street2, p.Zip_code)

	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	id3, err := rs3.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	Id = int(id)
	rs2, err := stmt2.Exec(p.Phone_number, Id)
	id2, err := rs2.LastInsertId()
	Id2 = int(id2)
	Id3 = int(id3)
	rs4, err := stmt4.Exec(Id, Id3)
	id4, err := rs4.LastInsertId()
	Id4 = int(id4)
	defer stmt.Close()
	return
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:1900)/infilon")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.GET("/persons", func(c *gin.Context) {
		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
		})

	})

	router.GET("/person/:id/info", func(c *gin.Context) {
		var result gin.H
		id := c.Param("id")

		Id, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln(err)
		}

		p := Person{
			Id: Id,
		}
		person, err := p.get()
		if err != nil {
			result = gin.H{
				"result": nil,
			}
		} else {
			result = gin.H{
				"result": person,
			}

		}
		c.JSON(http.StatusOK, result)
	})

	router.POST("/person/add", func(c *gin.Context) {

		var p Person
		err := c.Bind(&p)
		if err != nil {
			log.Fatalln(err)
		}

		Id, Id2, Id3, Id4, err := p.add()

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(Id)
		fmt.Println(Id2)
		fmt.Println(Id3)
		fmt.Println(Id4)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully created"),
		})

	})

	router.Run(":8000")

}
