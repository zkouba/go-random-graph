package graph

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

const JSON = "JSON"
const CSV = "CSV"

func SerializeNodeList(nodes *[]Node, format string) (*string, error) {
	var buff bytes.Buffer
	switch format {
	case JSON:
		{
			s, err := json.Marshal(*nodes)
			if err != nil {
				return nil, err
			}
			buff.Write(s)
		}
		break
	case CSV:
		{
			for _, node := range *nodes {
				buff.WriteString(strconv.Itoa(node.Id))
				buff.WriteString(",\"")
				buff.WriteString(node.Name)
				buff.WriteString("\",")
				buff.WriteString(strconv.Itoa(node.Class))
				buff.WriteString("\n")
			}
		}
		break
	default:
		{
			return nil, errors.New("unknown serialization format '" + format + "'")
		}
	}
	var retVal = buff.String()
	return &retVal, nil
}

func SerializeEdgeList(edges *[]Edge, format string) (*string /* returning pointer enables us to return nil if desired */, error) {
	var buff bytes.Buffer
	switch format {
	case JSON:
		{
			s, err := json.Marshal(*edges)
			if err != nil {
				return nil, err
			}
			buff.Write(s)
		}
		break
	case CSV:
		{
			for _, edge := range *edges {
				buff.WriteString(strconv.Itoa(edge.X.Id))
				buff.WriteString(",")
				buff.WriteString(strconv.Itoa(edge.Y.Id))
				buff.WriteString("\n")
			}
		}
		break
	default:
		{
			return nil, errors.New("unknown serialization format '" + format + "'")
		}
	}
	var retVal = buff.String()
	return &retVal, nil
}
