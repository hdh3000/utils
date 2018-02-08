package main

import (
	"fmt"
	"time"
	"utils/matcher"
)

func main() {
	persons := []string{"Henry", "Max", "PJ", "Bill", "Milligan", "Frank", "FatherSchwartz"}

	tasks := []string{
		"Dinner-1",
		"Dinner-2",
		"Dinner-3",
		"Lunch-1",
		"Lunch-2",
		"Breakfast-1",
		"Breakfast-2",
		"Breakfast-3",
		"Dishes-1",
		"Dishes-2",
		"Dishes-3",
		"Dishes-4",
		"Dishes-5",
		"Dishes-6",
		"Dishes-7",
		"Dishes-8",
		"Shaman-1",
		"Shaman-2",
		"SpiritualLeader-1",
		"SpiritualLeader-2",
	}

	m := matcher.Match(time.Now().Unix(), persons, tasks)

	for _, p := range m {
		fmt.Println(p.Person)
		for _, t := range p.Task {
			fmt.Printf(" - %s\n", t)
		}
	}
}
