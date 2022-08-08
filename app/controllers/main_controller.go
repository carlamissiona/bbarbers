package controllers

import (
	"carlamissiona/golang-barbers/pkg/database"
	"database/sql"
	_ "database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	m3o "go.m3o.com"
	"go.m3o.com/email"
)

var db *sql.DB

func Initcontroller() {
	log.Println("InitInitCon!")
	db = database.SetupDatabase()
}

func RenderHome(c *fiber.Ctx) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT   COUNT(title) FROM   bbr_articles;")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res);log.Println("Record nos");
		log.Printf(res);
		todos = append(todos, res)
	}

	// SELECT   COUNT(column) FROM   table_name WHERE  condition;
	rows, err = db.Query("Select * from public.bbr_articles")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}

	// strs := []string{"first", "second"}
	// names := make([]interface{}, len(strs))
	// for i, s := range strs {
	// 	names[i] = s
	// }
	
	rowPtr := make([]any,6)
	for rows.Next() {
		err := rows.Scan(rowPtr...)
	    log.Printf('\n v%', err)
		log.Println(rowPtr)
	}
	
	err = db.Ping();log.Println("Passed Ping & Error after OS env file");
	if err != nil {
		log.Fatalf("failed No DB connection %v", err)
	}



   log.Println(todos);log.Println("todos");log.Println("todos");log.Println("todos")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderPaid(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderPayment(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderAbout(c *fiber.Ctx) error {
	log.Println("Higshsss!")
	return c.Render("index", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}

func RenderContact(c *fiber.Ctx) error {
	log.Println("High")

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
func RenderContactSubmit(c *fiber.Ctx) error {
	log.Println("From", c.Params("from"))
	api_secret := os.Getenv("M3O_APIKEY")

	client := m3o.New(api_secret)
	rsp, err := client.Email.Send(&email.SendRequest{
		To:      "missiona.carla@gmail.com <missiona.carla@gmail.com>",
		From:    "Awesome Dot Com <codetuna@protonmail.com>",
		Subject: "Email verification",
		TextBody: `Hi there,

        Please verify your email by clicking this link: $micro_verification_link`,
	})
	log.Println(rsp, err)

	return c.Render("contact", fiber.Map{
		"FiberTitle": "Hello From Fiber Html Engine",
	}, "layouts/htm")
}
