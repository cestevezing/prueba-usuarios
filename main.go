package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AutoGenerated struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Login struct {
			UUID string `json:"uuid"`
		} `json:"login"`
		Location struct {
			City    string `json:"city"`
			Country string `json:"country"`
		} `json:"location"`
		Email string `json:"email"`
	} `json:"results"`
}

type NewData struct {
	ID      string `json:"id"`
	Gender  string `json:"gender"`
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	City    string `json:"city"`
	Country string `json:"country"`
}

var (
	mu   sync.Mutex
	stop = make(chan struct{}, 4)
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		newData := []NewData{}
		var wg sync.WaitGroup
		user := LoadUser()
		time.Sleep(100 * time.Millisecond)

		for k := 0; k < 15; k++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				stop <- struct{}{}
				defer func() { <-stop }()
				data, err := LoadData(&user)
				if err != nil {
					log.Println(err)
				}
				mu.Lock()
				newData = append(newData, data...)
				mu.Unlock()
			}()
		}
		wg.Wait()
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   newData,
		})
	})
	app.Listen(":8080")
}

func LoadUser() NewData {
	resp, err := http.Get("https://randomuser.me/api/?inc=gender,name,location,email,login&noinfo&results=1")
	if err != nil {
		log.Printf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()
	dataAux := new(AutoGenerated)
	err = json.NewDecoder(resp.Body).Decode(&dataAux)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}
	item := NewData{
		ID:      dataAux.Results[0].Login.UUID,
		Gender:  dataAux.Results[0].Gender,
		First:   dataAux.Results[0].Name.First,
		Last:    dataAux.Results[0].Name.Last,
		Email:   dataAux.Results[0].Email,
		City:    dataAux.Results[0].Location.City,
		Country: dataAux.Results[0].Location.Country,
	}
	return item
}

func LoadData(user *NewData) ([]NewData, error) {
	start := time.Now()
	var dataBatch []NewData
	resp, err := http.Get("https://randomuser.me/api/?inc=login&noinfo&results=1000")
	if err != nil {
		log.Printf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()
	dataAux := new(AutoGenerated)
	err = json.NewDecoder(resp.Body).Decode(&dataAux)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
	}
	fmt.Printf("Tiempo total: %s\n", time.Since(start))
	for i := range dataAux.Results {
		item := NewData{
			ID:      dataAux.Results[i].Login.UUID,
			Gender:  user.Gender,
			First:   user.First,
			Last:    user.Last,
			Email:   user.Email,
			City:    user.City,
			Country: user.Country,
		}
		dataBatch = append(dataBatch, item)
	}
	return dataBatch, nil
}
