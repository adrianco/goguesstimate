package guesstimate

import (
	"fmt"
	"testing"
)

// Example of a valid Guesstimate object
var example = `{ 
      "space": {
      "name": "gotest",
      "description": "Testing",
      "is_private": "true",
      "graph":  {
          "metrics": [
              {"id": "AB", "readableId": "AB", "name": "memcached", "location": {"row": 2, "column":4}},
              {"id": "AC", "readableId": "AC", "name": "memcached percent", "location": {"row": 2, "column":3}},
              {"id": "AD", "readableId": "AD", "name": "staash cpu", "location": {"row": 3, "column":3}},
              {"id": "AE", "readableId": "AE", "name": "staash", "location": {"row": 3, "column":2}}
              ], 
          "guesstimates": [
	      {"metric": "AB", "expression": null, "guesstimateType": "DATA", "data": [119958,6066,13914,9595,6773,5867,2347,1333,9900,9404,13518,9021,7915,3733,10244,5461,12243,7931,9044,11706,5706,22861,9022,48661,15158,28995,16885,9564,17915,6610,7080,7065,12992,35431,11910,11465,14455,25790,8339,9991]},
              {"metric": "AC", "expression": "40", "guesstimateType": "POINT"},
              {"metric": "AD", "expression": "[1000,4000]", "guesstimateType": "LOGNORMAL"},
              {"metric": "AE", "expression": "=100+((randomInt(0,100)>AC)?${metric:AB}:${metric:AD})", "guesstimateType": "FUNCTION"}
          ]
      }
  } 
}
`

func TestGuess(t *testing.T) {
	var g Guess
	g = Guess{
		Space: GuessModel{
			Name:        "gotest",
			Description: "Testing",
			IsPrivate:   "true",
			Graph: GuessGraph{
				Metrics: []GuessMetric{
					GuessMetric{
						ID:         "AB",
						ReadableID: "AB",
						Name:       "memcached",
						Location:   GuessMetricLocation{2, 4},
					},
					GuessMetric{"AC", "AC", "memcached percent", GuessMetricLocation{2, 3}},
					GuessMetric{"AD", "AD", "staash cpu", GuessMetricLocation{3, 3}},
					GuessMetric{"AE", "AE", "staash", GuessMetricLocation{3,2}},
				},
				Guesstimates: []Guesstimate{
					Guesstimate{
						Metric:          "AB",
						GuesstimateType: "DATA",
						Data:            []int64{119958, 6066, 13914, 9595, 6773, 5867, 2347, 1333, 9900, 9404, 13518, 9021, 7915, 3733, 10244, 5461, 12243, 7931, 9044, 11706, 5706, 22861, 9022, 48661, 15158, 28995, 16885, 9564, 17915, 6610, 7080, 7065, 12992, 35431, 11910, 11465, 14455, 25790, 8339, 9991},
					},
					Guesstimate{
						Metric:          "AC",
						Expression:      "40",
						GuesstimateType: "POINT",
					},
					Guesstimate{"AD", "[1000,4000]", "LOGNORMAL", nil},
					Guesstimate{"AE", "=100+((randomInt(0,100)>AC)?AB:AD)", "FUNCTION", nil},
				},
			},
		},
	}
	SaveGuess(g, "gotest")
	fmt.Println("\n$ sh pretty.sh to pretty print the output file in standard json and compare want.guess with got.guess")
}
