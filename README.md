# go-random-graph

Generator of random undirected graphs.

## Usage: 
```bash
go run GeneratorApp.go \
    ${NUMBER_OF_NODES} \
    ${PROB_CLASS_0} \
    ${PROB_INTER_CLASS_EDGE} \
    ${PROB_INTRA_CLASS_EDGE} \
    ${NODES_FILE} \
    ${GRAPH_FILE} \
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
|`GRAPH_FILE`| Name of the file where the list of edges will be printed| 
|`FORMAT`| The format of the output files (`JSON`, `CSV`)| 

### Example:

```bash
go run GeneratorApp.go \
    10000 \ # number of nodes
    0.6 \ # probability of class 0
    0.3 \ # probability of inter-class connections
    0.85 \ # probability of connections within a class
    nodes.csv graph.csv CSV # outputs
```
