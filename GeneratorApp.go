package main

import (
	"math/rand"
	"errors"
	"os"
	"strconv"
	"log"
	"bytes"
	"encoding/json"
)

type Edge struct {
	X int
	Y int
}

const JSON = "JSON"
const CSV = "CSV"

func main() {
	if len(os.Args) < 5 {
		err := errors.New("not enough command-line arguments")
		handleError(err)
	}

	// Read arguments
	var arguments = os.Args[1:]
	nodeNum, err := strconv.Atoi(arguments[0])
	handleError(err)
	prob, err := strconv.ParseFloat(arguments[1], 64)
	handleError(err)
	outputFileName := arguments[2]
	format := arguments[3]

	// Generate graph
	graph, err := GenerateRandomUndirectedGraph(nodeNum, prob)
	handleError(err)

	// Serialize graph
	graphSer, err := SerializeGraph(&graph, format)
	handleError(err)

	// Print output to file
	outputFile, err := os.Create(outputFileName)
	handleError(err)
	defer outputFile.Close()
	_, err = outputFile.WriteString(*graphSer) // The '_' means that the return value on that possition is ignored
	handleError(err)

	log.Println("Successfully written to output file")
}

func SerializeGraph(graph *[]Edge, format string) (*string /* returning pointer enables us to return nil if desired */, error) {
	var buff bytes.Buffer
	switch format {
		case JSON: {
			s, err := json.Marshal(*graph)
			if err != nil {
				return nil, err
			}
			buff.Write(s)
		}; break
		case CSV: {
			for _, edge := range *graph {
				buff.WriteString(strconv.Itoa(edge.X))
				buff.WriteString(", ")
				buff.WriteString(strconv.Itoa(edge.Y))
				buff.WriteString("\n")
			}
		}; break
		default: {
			return nil, errors.New("unknown serialization format '" + format + "'")
		}
	}
	var retVal = buff.String()
	return &retVal, nil
}

func GenerateRandomUndirectedGraph(nodeNum int, prob float64) ([]Edge, error) {
	if prob < 0 || prob > 1 {
		return nil,
			errors.New("probability that there's an edge between any two nodes has to be between 0 and 1")
	}
	var edges = make([]Edge, 0, int(float64(nodeNum^2) * prob))
	for n := int(0); n < nodeNum - 1; n++ {
		for m := int(n + 1); m < nodeNum; m++ {
			var x = rand.Float64()
			if x < prob {
				e := Edge{n, m}
				edges = append(edges, e)
			}
		}
	}
	return edges, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}