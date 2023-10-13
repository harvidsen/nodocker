package main

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

type Pokemon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CaptureRate int    `json:"capture_rate"`
	IsLegendary bool   `json:"is_legendary"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Generation struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

func inspect(client *resty.Client, pokemonSpecies []NamedAPIResource) {
	var (
		res     *resty.Response
		err     error
		pokemon Pokemon
	)

	for _, apiResource := range pokemonSpecies {
		log.Printf("Inspect pokemon %s", apiResource.Name)
		res, err = client.R().Get(apiResource.URL)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(res.Body(), &pokemon)
		log.Printf(
			"Pokemon %s has ID %d, capture rate %d and is legendary: %t",
			pokemon.Name, pokemon.ID, pokemon.CaptureRate, pokemon.IsLegendary,
		)
	}
}

func main() {
	log.Println("Go Start!")

	client := resty.New()

	res, err := client.R().Get("https://pokeapi.co/api/v2/generation/1")
	if err != nil {
		log.Fatal(err)
	}

	var generation Generation

	err = json.Unmarshal(res.Body(), &generation)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got info about generation %s, with %d species", generation.Name, len(generation.PokemonSpecies))

	inspect(client, generation.PokemonSpecies)

	//// Go routine and channel
	// FastInspect(client, generation.PokemonSpecies)

	//// Context with timeout
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// defer cancel()
	// _, err = HardThing(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
