package graph

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

type Config struct {
	NodeNumber                 int     `json:"nodeNumber"`
	ProbOfClass0               float64 `json:"probOfClass0"`
	ProbOfInterClassConnection float64 `json:"probOfInterClassConnection"`
	ProbOfIntraClassConnection float64 `json:"probOfIntraClassConnection"`
	OutputNodesFile            string  `json:"outputNodesFile"`
	OutputEdgesFile            string  `json:"outputEdgesFile"`
	OutputFormat               string  `json:"outputFormat"`
}

func (c *Config) Load(args []string) (*Config, error) {
	var argCount = len(args)
	if argCount == 2 {
		// Read arguments from JSON file
		var configFileName = args[1]
		confJson, err := ioutil.ReadFile(configFileName)
		if err != nil {
			return nil, err
		}

		// var config = Config{}
		err = json.Unmarshal(confJson, &c)
		if err != nil {
			return nil, err
		}
		return c, nil
	} else if argCount < 8 {
		err := errors.New("not enough command-line arguments")
		return nil, err
	}
	// Read arguments from CMD
	var arguments = args[1:]
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
	edgesFileName := arguments[5]
	format := arguments[6]
	c.NodeNumber = nodeNum
	c.ProbOfClass0 = probCls
	c.ProbOfInterClassConnection = probClsInter
	c.ProbOfIntraClassConnection = probClsIntra
	c.OutputNodesFile = nodesFileName
	c.OutputEdgesFile = edgesFileName
	c.OutputFormat = format
	return c, nil
}
