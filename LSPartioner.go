package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	bipartitonlocalsearchlib "github.com/Rakiiii/goBipartitonLocalSearch"
)

var proc int = 1

func main() {

	graphPath := os.Args[1]

	disbalance, er := strconv.ParseFloat(os.Args[2], 64)
	if er != nil {
		log.Println(er)
		return
	}

	var graph bipartitonlocalsearchlib.Graph

	if err := graph.ParseGraph(graphPath); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex()/2 - int(float64(graph.AmountOfVertex())*disbalance)

	fmt.Println("GroupSize:", groupSize)

	log.Println("amount of independent:", graph.GetAmountOfIndependent(), "|amount of vertex:", graph.AmountOfVertex())
	var ord []int

	if os.Args[1] == "-h" || os.Args[1] == "-dh" {
		ord = graph.HungryNumIndependent()
	} else {
		ord = graph.NumIndependent()
	}

	if os.Args[1] == "-d" || os.Args[1] == "-dh" {
		log.Println(ord)

	}else{

		log.Println("amount of independent:", graph.GetAmountOfIndependent(), "|amount of vertex:", graph.AmountOfVertex())

		timeStart := time.Now()

		res := bipartitonlocalsearchlib.LSPartiotionAlgorithm(&graph, nil, groupSize, 0)

		timeEnd := time.Now()
		elapced := timeEnd.Sub(timeStart)
		timeFile, err := os.Create("time")
		defer timeFile.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			timeFile.WriteString(strconv.FormatInt(elapced.Milliseconds(), 10) + "ms")
		}

		formatedRes := make([]int, len(ord))
		strRes := ""
		for i, num := range ord {
			if res.Vector[i] {
				formatedRes[num] = 1
			} else {
				formatedRes[num] = 0
			}
		}

		for _,v := range formatedRes{
			strRes += strconv.Itoa(v)
		}

		fmt.Println(ord)
		fmt.Println(res.Vector)

		f, err := os.Create("result_" + graphPath)
		if err != nil {
			fmt.Println("res:", formatedRes)
			fmt.Println("value", res.CountParameter())
			log.Panic(err)
		}
		defer f.Close()

		f.WriteString(strconv.FormatInt(res.Value, 10) + "\n")
		f.WriteString(strRes)
}
}
