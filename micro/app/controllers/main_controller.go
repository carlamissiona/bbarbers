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


type Article struct {
 
	Title       string
	Content     string  
	Link        string  
	Changed_on interface{}
  }


var db *sql.DB 

func Initcontroller() {
	log.Println("InitInitCon!")
	db = database.SetupDatabase()
}

func RenderHome(c *fiber.Ctx) error {
	 
	 
	rows, err := db.Query("Select title, content, link, changed_on from public.bbr_articles")
	 
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
    
	log.Println("Struct")
	// database.OpeDatabase()
	var articles []Article
    for rows.Next() {
		var ar Article
		err := rows.Scan( &ar.Title, &ar.Content, &ar.Link, &ar.Changed_on )
        if err != nil {
			log.Printf("Err! %v", err)
        } 
        articles = append(articles, ar) 
    }
    if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)
    }
	log.Println("Row! %v", articles[0].Link)
	log.Println("Row! %v", articles)
	log.Println(articles);
	return c.Render("index", fiber.Map{
	   "Articles" : articles,
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



func GetApi_Articles(c *fiber.Ctx) error {
	log.Println("API ARTICLES")
	rows, err := db.Query("Select article_id, title, content, link, changed_on from public.bbr_articles")
	
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
    
	log.Println("Struct")
	// database.OpeDatabase()
	var articles []Article
    for rows.Next() {
		log.Printf("Print Next Row")
		var ar Article
		err := rows.Scan( &ar.Title, &ar.Content, &ar.Link, &ar.Changed_on )  
        if err != nil {
			log.Printf("Err! %v", err)
        } 
        articles = append(articles, ar)
    }
 

    if err := rows.Err(); err != nil {
		log.Printf("Err! %v", err)
    }
	// log.Println("Row! %v", articles[0].Article_id)
	log.Println("Row! %v", articles)
	log.Println(articles);
	type response struct {
		articles      []Article  
		code          int32  
	}
    
		var resp = new(response)
		resp.articles = articles
		resp.code = 200
		log.Println(resp)   

	return c.JSON( resp )
	 
}  

func GetApi_Users(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Users")
}

func GetApi_Maps(c *fiber.Ctx) error {
 
	return c.SendString("GetApi_Maps")
}
