package main

import (
	"fmt"
)

type Config struct{ jobs int }

type Modifier interface {
	Modify(*Config) error
}

type numJob int

func (nj numJob) Modify(conf *Config) error {
	conf.jobs = int(nj)
	return nil
}

func NumberOfJobs(n int) numJob { return numJob(n) }

func Create(args ...Modifier) (conf Config, err error) {
	for _, arg := range args {
		if err = arg.Modify(&conf); err != nil {
			return
		}
	}
	return
}

func main() {
	fmt.Println(Create(NumberOfJobs(1), NumberOfJobs(2)))
}
