package main

import (
	"math/rand"
	"errors"
	"os"
	"strconv"
	"log"
	"bytes"
	"encoding/json"
	"github.com/Pallinder/go-randomdata"
)

type Node struct {
	Id    int
	Class int
	Name string
}

type Edge struct {
	X Node
	Y Node
}

const JSON = "JSON"
const CSV = "CSV"

func main() {
	if len(os.Args) < 8 {
		err := errors.New("not enough command-line arguments")
		handleError(err)
	}

	// Read arguments
	var arguments = os.Args[1:]
	nodeNum, err := strconv.Atoi(arguments[0])
	handleError(err)
	probCls, err := strconv.ParseFloat(arguments[1], 64)
	handleError(err)
	probClsInter, err := strconv.ParseFloat(arguments[2], 64)
	handleError(err)
	probClsIntra, err := strconv.ParseFloat(arguments[3], 64)
	handleError(err)
	nodesFileName := arguments[4]
	graphFileName := arguments[5]
	format := arguments[6]

	// Generate graph
	nodes, graph, err := GenerateRandomUndirectedGraph(nodeNum, probCls, probClsInter, probClsIntra)
	handleError(err)

	// Serialize graph
	graphSer, err := SerializeGraph(&graph, format)
	handleError(err)
	nodesSer, err := json.Marshal(nodes)
	handleError(err)

	// Print output to files
	graphFile, err := os.Create(graphFileName)
	handleError(err)
	defer graphFile.Close()
	_, err = graphFile.WriteString(*graphSer) // The '_' means that the return value on that possition is ignored
	handleError(err)
	nodesFile, err := os.Create(nodesFileName)
	handleError(err)
	defer nodesFile.Close()
	_, err = nodesFile.Write(nodesSer)
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
				buff.WriteString(strconv.Itoa(edge.X.Id))
				buff.WriteString(", ")
				buff.WriteString(strconv.Itoa(edge.Y.Id))
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

func GenerateRandomUndirectedGraph(
	nodeNum int,
	clsProb float64,
	interClsProb float64,
	intraClsProb float64) ([]Node, []Edge, error) {
	if clsProb < 0 || clsProb > 1 || interClsProb < 0 || interClsProb > 1 || intraClsProb < 0 || intraClsProb > 1{
		return nil, nil, errors.New("probability has to be between 0 and 1")
	}
	var nodes = make([]Node, nodeNum)
	for i := int(0); i < nodeNum; i++ {
		var x = rand.Float64()
		var cls int
		if x < clsProb {
			cls = 0
		} else {
			cls = 1
		}
		nodes[i] = Node{i, cls, randomdata.FullName(randomdata.RandomGender)}
	}

	var edges = make([]Edge, 0)
	for n := int(0); n < nodeNum - 1; n++ {
		for m := int(n + 1); m < nodeNum; m++ {
			var x = rand.Float64()
			var prob float64
			if nodes[n].Class == nodes[m].Class {
				prob = intraClsProb
			} else {
				prob = interClsProb
			}
			if x < prob {
				e := Edge{nodes[n], nodes[m]}
				edges = append(edges, e)
			}
		}
	}
	return nodes, edges, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}