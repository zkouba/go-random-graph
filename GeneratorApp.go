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
	"io/ioutil"
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

type Config struct {
	NodeNumber int
	ProbOfClass0 float64
	ProbOfInterClassConnection float64
	ProbOfIntraClassConnection float64
	OutputNodesFile string
	OutputGraphFile string
	OutputFormat string
}

const JSON = "JSON"
const CSV = "CSV"

func main() {
	config, err := loadConfig()
	handleError(err)

	// Generate graph
	nodes, graph, err := GenerateRandomUndirectedGraph(
		config.NodeNumber,
		config.ProbOfClass0,
		config.ProbOfInterClassConnection,
		config.ProbOfIntraClassConnection)
	handleError(err)

	// Serialize graph
	graphSer, err := SerializeGraph(&graph, config.OutputFormat)
	handleError(err)
	nodesSer, err := SerializeNodes(&nodes, config.OutputFormat)
	handleError(err)

	// Print output to files
	graphFile, err := os.Create(config.OutputGraphFile)
	handleError(err)
	defer graphFile.Close()
	_, err = graphFile.WriteString(*graphSer) // The '_' means that the return value on that position is ignored
	handleError(err)

	nodesFile, err := os.Create(config.OutputNodesFile)
	handleError(err)
	defer nodesFile.Close()
	_, err = nodesFile.WriteString(*nodesSer)
	handleError(err)

	log.Println("Successfully written to output file")
}

func loadConfig() (*Config, error) {
	var argCount = len(os.Args)
	if argCount == 2 {
		// Read arguments from JSON file
		var configFileName = os.Args[1]
		confJson, err := ioutil.ReadFile(configFileName)
		if err != nil {
			return nil, err
		}

		var config = Config{}
		err = json.Unmarshal(confJson, &config)
		if err != nil {
			return nil, err
		}
		return &config, nil
	} else if argCount < 8 {
		err := errors.New("not enough command-line arguments")
		return nil, err
	}
	// Read arguments from CMD
	var arguments = os.Args[1:]
	nodeNum, err := strconv.Atoi(arguments[0])
	if err != nil {
		return nil, err
	}
	probCls, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		return nil, err
	}
	probClsInter, err := strconv.ParseFloat(arguments[2], 64)
	if err != nil {
		return nil, err
	}
	probClsIntra, err := strconv.ParseFloat(arguments[3], 64)
	if err != nil {
		return nil, err
	}

	nodesFileName := arguments[4]
	graphFileName := arguments[5]
	format := arguments[6]
	return &Config{
		nodeNum,
		probCls,
		probClsInter,
		probClsIntra,
		nodesFileName,
		graphFileName,
		format},
		nil
}

func SerializeNodes(nodes *[]Node, format string) (*string, error) {
	var buff bytes.Buffer
	switch format {
	case JSON: {
		s, err := json.Marshal(*nodes)
		if err != nil {
			return nil, err
		}
		buff.Write(s)
	}; break
	case CSV: {
		for _, node := range *nodes {
			buff.WriteString(strconv.Itoa(node.Id))
			buff.WriteString(",\"")
			buff.WriteString(node.Name)
			buff.WriteString("\",")
			buff.WriteString(strconv.Itoa(node.Class))
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
				buff.WriteString(",")
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