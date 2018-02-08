package matcher

import (
	"math/rand"
)

type Matched struct {
	Person string
	Task   []string
}

func Match(seed int64, persons, tasks []string) []Matched {
	r := rand.New(rand.NewSource(seed))

	// Shuffle persons and tasks
	persons = shuffle(persons, r.Perm(len(persons)))
	tasks = shuffle(tasks, r.Perm(len(tasks)))

	m := make([]Matched, len(persons))
	for i := range m {
		m[i].Person = persons[i]
	}

	for i := range tasks {
		p := i % len(persons)
		m[p].Task = append(m[p].Task, tasks[i])
	}

	return m
}

func shuffle(in []string, perm []int) []string {
	if len(in) != len(perm) {
		panic("wrong length perm")
	}

	out := make([]string, len(in))
	for i := range perm {
		out[i] = in[perm[i]]
	}

	return out
}
