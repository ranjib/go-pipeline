package main

import (
	"github.com/ranjib/go-pipeline"
	"log"
	"os/exec"
)

func main() {
	task := pipeline.Task{
		Command: exec.Command("sleep", "3"),
	}
	job := pipeline.Job{
		Tasks: []pipeline.Task{task},
	}
	stage := pipeline.Stage{
		Jobs: []pipeline.Job{job},
	}
	p := pipeline.Pipeline{
		Stages: []pipeline.Stage{stage},
	}
	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
