package main
import (
	"github.com/muhsufyan/dasar-golang2/consuming_api"
	"fmt"
)
func main()  {
	// consuming_api
	responseData := consuming_api.CallApi("http://pokeapi.co/api/v2/pokedex/kanto/")
	fmt.Println("========================================================")
	fmt.Println("SHOW ALL DATA")
	fmt.Println(string(responseData))
	fmt.Println("======================Uint8==================================")
	a := consuming_api.UnmarshalJson(responseData)
	fmt.Println("=======================DATA FILTER WITH UNMARSHALL=================================")
	consuming_api.ListAllPokeman(a)
}