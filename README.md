# go-random-graph

Generator of random undirected graphs.

## Usage: 
### With a config file:

```bash
go run GeneratorApp.go ${PATH_TO_CONFIG_FILE}
```

**Where:**

|Parameter |Description|
|---|---|
|`PATH_TO_CONFIG_FILE` | Path to the text file that contains the JSON-serialization of the configuration|

#### Configuration JSON structure:

```json
{
        "OutputFormat": "CSV",
        "NodeNumber": 1000,
        "ProbOfClass0": 0.3,
        "ProbOfInterClassConnection": 0.3,
        "ProbOfIntraClassConnection": 0.8,
        "OutputNodesFile": "nodes.csv",
        "OutputEdgesFile": "edges.csv"
}
```

**Where:**

|Parameter |Type |Description |
|---|---|---|
|`NodeNumber` |int | Number of nodes in the graph | 
|`ProbOfClass0` |float64 | Probability that a node belongs to class 0 (there are only 2 classes)|
|`ProbOfInterClassConnection` |float64 | Probability that any two nodes of different classes are connected| 
|`ProbOfIntraClassConnection` |float64 | Probability that any two nodes of the same class are connected| 
|`OutputNodesFile` |string | Name of the file where the list of nodes, their names and classes will be printed| 
|`OutputEdgesFile` |string | Name of the file where the list of edges will be printed| 
|`OutputFormat` |string | The format of the output files (`JSON`, `CSV`)| 

#### Example:

```bash
go run GeneratorApp.go config.json
```

### Spelling out the parameters:

```bash
go run GeneratorApp.go \
    ${NUMBER_OF_NODES} \
    ${PROB_CLASS_0} \
    ${PROB_INTER_CLASS_EDGE} \
    ${PROB_INTRA_CLASS_EDGE} \
    ${NODES_FILE} \
    ${EDGES_FILE} \
    ${FORMAT}
```

**Where:**

|Parameter |Description |
|---|---|
|`NUMBER_OF_NODES` | Number of nodes in the graph | 
|`PROB_CLASS_0`| Probability that a node belongs to class 0 (there are only 2 classes)|
|`PROB_INTER_CLASS_EDGE`| Probability that any two nodes of different classes are connected| 
|`PROB_INTRA_CLASS_EDGE`| Probability that any two nodes of the same class are connected| 
|`NODES_FILE`| Name of the file where the list of nodes, their names and classes will be printed| 
|`EDGES_FILE`| Name of the file where the list of edges will be printed| 
|`FORMAT`| The format of the output files (`JSON`, `CSV`)| 

#### Example:

```bash
go run GeneratorApp.go \
    10000 \ # number of nodes
    0.6 \ # probability of class 0
    0.3 \ # probability of inter-class connections
    0.85 \ # probability of connections within a class
    nodes.csv edges.csv CSV # outputs
```
