package main

import (
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/rotoku/go-webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// lista de produtos
	// produtos := []Produto{
	// 	{Codigo: 1, Nome: "Arroz", Descricao: "Arroz Tipo 1 Bom Prato", Preco: 25.99, Quantidade: 1},
	// 	{2, "Nuggets", "Nuggets de Frango Crocante Sadia 300g", 7.99, 1},
	// 	{3, "Papel Higiênico", "Papel Higiênico Folha Dupla, Branco 24 unidades, Personal Vip", 49.90, 1},
	// 	{4, "Leite Condensado", "Leite Condensado Semidesnatado Piracanjuba 395g", 5.99, 1},
	// 	{5, "Carne Moída", "Paleta Bovina Moída Congelada Swift 500 g", 19.90, 1},
	// 	{6, "Creme de Leite", "Creme de Leite Nestlé 200g", 4.59, 1},
	// 	{7, "Pão de Forma Integral", "Pão de Forma Integral Visconti Pacote 400 g", 6.99, 1},
	// 	{8, "Salsicha", "Salsicha Tradicional Sadia 500g", 9.50, 1},
	// }
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
