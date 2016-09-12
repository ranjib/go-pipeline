## go-pipeline

Pipeline abstraction in go similar to GoCD


### Usage

Pipelines are consist of stages. Stages are composed of jobs.
Jobs are composed of tasks. Tasks are commands.

Stages execute sequentially. Within individual stage, jobs execute
parallely. Within individual job, tasks execute sequentially.

```go
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
```
