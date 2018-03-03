package main

import (
	"github.com/zkouba/go-random-graph/graph"
	"log"
	"os"
)

func main() {
	config, err := (&graph.Config{}).Load(os.Args)
	handleError(err)

	// Generate graph
	nodes, edges, err := graph.GenerateRandomUndirectedGraph(
		config.NodeNumber,
		config.ProbOfClass0,
		config.ProbOfInterClassConnection,
		config.ProbOfIntraClassConnection)
	handleError(err)

	// Serialize graph
	edgesSer, err := graph.SerializeEdgeList(&edges, config.OutputFormat)
	handleError(err)
	nodesSer, err := graph.SerializeNodeList(&nodes, config.OutputFormat)
	handleError(err)

	// Print output to files
	edgesFile, err := os.Create(config.OutputGraphFile)
	handleError(err)
	defer edgesFile.Close()
	_, err = edgesFile.WriteString(*edgesSer) // The '_' means that the return value on that position is ignored
	handleError(err)

	nodesFile, err := os.Create(config.OutputNodesFile)
	handleError(err)
	defer nodesFile.Close()
	_, err = nodesFile.WriteString(*nodesSer)
	handleError(err)

	log.Println("Successfully written to output file")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
