package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	bipartitonlocalsearchlib "github.com/Rakiiii/goBipartitonLocalSearch"
)

func main() {

	disbalance, er := strconv.ParseFloat(os.Args[2], 64)
	if er != nil {
		log.Println(er)
		return
	}

	var graph bipartitonlocalsearchlib.Graph

	if err := graph.ParseGraph(os.Args[1]); err != nil {
		log.Println(err)
		return
	}

	/*for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(graph.GetEdges(i))
	}*/

	gropuSize := graph.AmountOfVertex()/2 - int(float64(graph.AmountOfVertex())*disbalance)

	fmt.Println("GroupSize:", gropuSize)

	log.Println("amount of independent:", graph.GetAmountOfIndependent(), "|amount of vertex:", graph.AmountOfVertex())

	ord := graph.NumIndependent()

	log.Println("amount of independent:", graph.GetAmountOfIndependent(), "|amount of vertex:", graph.AmountOfVertex())

	res := partiotion(&graph, nil, gropuSize, 0)

	formatedRes := make([]int, len(ord))
	for i, num := range ord {
		if res.Vector[i] {
			formatedRes[num] = 1
		} else {
			formatedRes[num] = 0
		}
	}
	fmt.Println("res:", formatedRes)
	fmt.Println("value", res.CountParameter())

}
