package main

import (
	"log"
	"math"

	bipartitonlocalsearchlib "github.com/Rakiiii/goBipartitonLocalSearch"
)

func partiotion(gr *bipartitonlocalsearchlib.Graph, sol *bipartitonlocalsearchlib.Solution, groupSize int, number int64) *bipartitonlocalsearchlib.Solution {
	log.Println("Check number:", number)

	if float64(number) >= math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent())) {
		log.Println("finish:", math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent())))
		return sol
	}
	var newSol bipartitonlocalsearchlib.Solution

	log.Println("solution constructed")

	newSol.Init(gr)
	newSol.SetDependentAsBinnary(number)
	mark := newSol.CountMark()
	log.Println("mark:", mark)

	if sol == nil {
		log.Println("nil solution removed")
		if flag := newSol.PartIndependent(groupSize); flag {
			log.Println("better param:", newSol.CountParameter())
			return partiotion(gr, &newSol, groupSize, number+1)
		} else {
			log.Println("invalid disb for:", number)
			return partiotion(gr, nil, groupSize, number+1)
		}
	}
	if mark < sol.CountParameter() {
		log.Println("better mark for :", sol.Vector)
		if flag := newSol.PartIndependent(groupSize); flag {
			if newSol.CountParameter() < sol.CountParameter() {
				log.Println("better param:", newSol.Value)
				return partiotion(gr, &newSol, groupSize, number+1)
			} else {
				log.Println("low param for:", number, " new param:", newSol.Value, " old param:", sol.Value)
			}
		} else {
			log.Println("invalid disb for:", number)
		}
	} else {
		log.Println("low mark for:", number)
	}
	return partiotion(gr, sol, groupSize, number+1)

}
