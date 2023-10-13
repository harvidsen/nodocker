package main

import (
	"encoding/json"
	"log"

	"github.com/go-resty/resty/v2"
)

func inspectSpecies(client *resty.Client, pokemonSpecies NamedAPIResource, pokechan chan<- Pokemon) {
	var (
		res     *resty.Response
		err     error
		pokemon Pokemon
	)

	res, err = client.R().Get(pokemonSpecies.URL)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(res.Body(), &pokemon)
	pokechan <- pokemon
}

func FastInspect(client *resty.Client, pokemonSpecies []NamedAPIResource) {
	var (
		pokemon     Pokemon
		bestPokemon Pokemon
	)
	pokechan := make(chan Pokemon)

	for _, apiResource := range pokemonSpecies {
		log.Printf("Inspect pokemon %s", apiResource.Name)
		go inspectSpecies(client, apiResource, pokechan)
	}

	for range pokemonSpecies {
		pokemon = <-pokechan
		log.Printf(
			"Pokemon %s has ID %d, capture rate %d and is legendary: %t",
			pokemon.Name, pokemon.ID, pokemon.CaptureRate, pokemon.IsLegendary,
		)
		if (pokemon.CaptureRate < bestPokemon.CaptureRate) || (bestPokemon.CaptureRate == 0) {
			bestPokemon = pokemon
		}
	}
	log.Printf(
		"Best pokemon is %s(legendary: %t) with capture rate %d",
		bestPokemon.Name, bestPokemon.IsLegendary, bestPokemon.CaptureRate,
	)

	//// Will wait forever since no new messages are sent to the channel
	// msg := <-results
	// log.Printf(msg)

}
