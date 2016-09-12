package pipeline

import (
	"log"
	"os/exec"
	"sync"
)

type Task struct {
	Command *exec.Cmd
}

func (t *Task) Run() error {
	log.Println("    Running:", t.Command.Path, " with argument:", t.Command.Args)
	return t.Command.Run()
}

type Job struct {
	Tasks []Task
}

func (j *Job) Run(wg *sync.WaitGroup) error {
	defer wg.Done()
	for _, task := range j.Tasks {
		if err := task.Run(); err != nil {
			return err
		}
	}
	return nil
}

type Stage struct {
	Jobs []Job
}

func (s *Stage) Run() error {
	var wg sync.WaitGroup
	for i, job := range s.Jobs {
		wg.Add(1)
		log.Println("  Starting job:", i+1)
		go job.Run(&wg)
	}
	wg.Wait()
	return nil
}

type Pipeline struct {
	Stages []Stage
}

func (p *Pipeline) Run() error {
	for i, stage := range p.Stages {
		log.Println("Starting stage:", i+1)
		if err := stage.Run(); err != nil {
			return err
		}
	}
	return nil
}
