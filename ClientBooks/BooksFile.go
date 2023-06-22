package ClientBooks

import (
	"encoding/json"
	"fmt"
	Client "github.com/jordinislic/utilities/utilities/ClientJordi"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type BooksClient struct {
	cl Client.ClientJordi
}

func New(url string) BooksClient {
	return BooksClient{
		cl: Client.New(url),
	}
}

func (b BooksClient) List() {
	resp := b.cl.Get("/books")
	jsonData := b.cl.GetBodyResp(resp)
	BookResp := []Book{}
	err := json.Unmarshal(jsonData, &BookResp)
	if err != nil {
		return
	}
	fmt.Println(BookResp)
}

func (b BooksClient) Get(id string) {
	variables := []string{}
	variables = append(variables, id)
	resp := b.cl.GetWithVariables("/books", variables)
	jsonData := b.cl.GetBodyResp(resp)
	BookResp := Book{}
	err := json.Unmarshal(jsonData, &BookResp)
	if err != nil {
		return
	}
	fmt.Println(BookResp)
}

func (b BooksClient) Add(body string) {
	resp := b.cl.Add("/books", body)
	jsonData := b.cl.GetBodyResp(resp)
	BookResp := Book{}
	err := json.Unmarshal(jsonData, &BookResp)
	if err != nil {
		return
	}
	fmt.Println(BookResp)
}

func (b BooksClient) Delete(id string) {
	variables := []string{}
	variables = append(variables, id)
	resp := b.cl.Delete("/books", variables)
	jsonData := b.cl.GetBodyResp(resp)
	BookResp := []Book{}
	err := json.Unmarshal(jsonData, &BookResp)
	if err != nil {
		return
	}
	fmt.Println(BookResp)
}

func (b BooksClient) Sort() {
	resp := b.cl.Add("/sorts/books", "")
	jsonData := b.cl.GetBodyResp(resp)
	BookResp := []Book{}
	err := json.Unmarshal(jsonData, &BookResp)
	if err != nil {
		return
	}
	fmt.Println(BookResp)
}
