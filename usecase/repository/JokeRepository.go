package repository

import (
	"academy-go-q12021/domain/model"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)


func GetRandomJoke(joke model.ChuckJoke) (error, model.ChuckJoke) {
	return nil, joke
}

func GetAllCSVJokes(jokes map[int]model.ChuckJoke) (map[int]model.ChuckJoke, error) {
	csvfile, fileerr := os.Open("input.csv")
	if fileerr != nil {
		log.Fatalln("Couldn't open the csv file", fileerr)
		return nil, fileerr
	}
	r := csv.NewReader(csvfile)
	for i := 0; ; i++{
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		intId, error := strconv.ParseInt(record[0], 10, 64)
		if error != nil {
			log.Fatal(err)
		}else{
			jokes[int(intId)] = model.ChuckJoke{ID : int(intId), Joke : record[1]}
		}

		fmt.Printf("ID: %s Joke %s\n", record[0], record[1])
	}
	return jokes, nil
}