package consuming_api

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	// "reflect"
	"encoding/json"
)
// A Response struct to map the Entire Response
type Response struct {
    Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}
func CallApi(url string) []uint8 {
	response, err := http.Get(url)
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	return responseData
}
func UnmarshalJson(responseData []uint8) *Response {
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	
	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))
	// fmt.Println("t1: %s\n", reflect.TypeOf(responseObject))
	return &responseObject

}
func ListAllPokeman(responseObject *Response)  {
    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }
}
func main() {
    responseData := CallApi("http://pokeapi.co/api/v2/pokedex/kanto/")
	fmt.Println("========================================================")
	fmt.Println("SHOW ALL DATA")
	fmt.Println(string(responseData))
	fmt.Println("======================Uint8==================================")
	a := UnmarshalJson(responseData)
	fmt.Println("=======================DATA FILTER WITH UNMARSHALL=================================")
	ListAllPokeman(a)
	// https://tutorialedge.net/golang/consuming-restful-api-with-go/

}