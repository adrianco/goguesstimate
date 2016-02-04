// Interface to getguesstimate.com Montecarlo spreadsheet service
package guesstimate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Guesstimate top level
type Guess struct {
	Space GuessModel `json:"space"`
}

// Guesstimate model
type GuessModel struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsPrivate   string     `json:"is_private"`
	Graph       GuessGraph `json:"graph"`
}

// Guesstimate Graph
type GuessGraph struct {
	Metrics      []GuessMetric `json:"metrics"`
	Guesstimates []Guesstimate `json:"guesstimates"`
}

// Guesstimate Metric
type GuessMetric struct {
	ID         string              `json:"id"`
	ReadableID string              `json:"readableId"`
	Name       string              `json:"name"`
	Location   GuessMetricLocation `json:"location"`
}

type GuessMetricLocation struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

type Guesstimate struct {
	Metric          string  `json:"metric"`
	Input           string  `json:"input"`
	GuesstimateType string  `json:"guesstimateType"`
	Data            []int64 `json:"data,omitempty"`
}

// Save a guess to a file
func SaveGuess(g Guess, fn string) {
	dfile, err := os.Create(fn + ".guess")
	if err != nil {
		log.Fatal("%v: %v\n", fn, err)
	}
	sj, _ := json.Marshal(g)
	dfile.WriteString(fmt.Sprintf("%v", string(sj)))
	dfile.Close()
}
