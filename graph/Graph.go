package graph

import (
	"errors"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
)

type Node struct {
	Id    int
	Class int
	Name  string
}

type Edge struct {
	X Node
	Y Node
}

func GenerateRandomUndirectedGraph(
	nodeNum int,
	clsProb float64,
	interClsProb float64,
	intraClsProb float64) ([]Node, []Edge, error) {
	if clsProb < 0 || clsProb > 1 || interClsProb < 0 || interClsProb > 1 || intraClsProb < 0 || intraClsProb > 1 {
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
	for n := int(0); n < nodeNum-1; n++ {
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
