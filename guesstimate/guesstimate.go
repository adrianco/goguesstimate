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
	ID         string              `json:"id"`         // opaque ID to match to Guesstimate below
	ReadableID string              `json:"readableId"` // 2 character ID shown on user interface
	Name       string              `json:"name"`       // display string
	Location   GuessMetricLocation `json:"location"`   // unique cell location
}

type GuessMetricLocation struct {
	Row    int `json:"row"`
	Column int `json:"column"`
}

// Guesstimate refers to Metrics
type Guesstimate struct {
	Metric          string  `json:"metric"`          // must match one ID
	//THE "Input" field has been DEPRECATED in the Guesstimate API; use the Expression field instead
	Input           string  `json:"input"`           // [1,2]   3      null, =AB+AD
	Expression		string  `json:"expression"`           // [1,2]   3      null, =${metric:<metric ID1>} + ${metric:<metric ID2>}`
	GuesstimateType string  `json:"guesstimateType"` // NORMAL, POINT, DATA, FORMULA
	Data            []int64 `json:"data,omitempty"`  // [119958, 6066, 13914, 9595, 6773] samples from a distribution
}

// Save a guess to a file
func SaveGuess(g Guess, fn string) {
	dfile, err := os.Create(fn + ".guess")
	if err != nil {
		log.Fatal(err)
	}
	sj, _ := json.Marshal(g)
	dfile.WriteString(fmt.Sprintf("%v", string(sj)))
	dfile.Close()
}
