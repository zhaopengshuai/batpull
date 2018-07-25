package main

import (
	"runtime"
)

type Pro struct {
	Path string
	Name string
	GitPath string

	GitOp string
}

type Op interface {
	Clone()
	Pull()
	Checkout()
	Branch()
}

const WorkerNum  = 3


func InitWorker (jobs chan Pro) {
	for w := 1; w <= WorkerNum; w++ {
		go Worker(jobs)
	}
}

func run() {
	//利用计算机的多核执行代码
	runtime.GOMAXPROCS(runtime.NumCPU())

	f := "pull"
	p := ParseCnf("batpull/.cnf")

	jobs := make(chan Pro)

	InitWorker(jobs)

	for _,v := range p {
		v.GitOp  = f
		jobs <- v
	}

	<-jobs
	close(jobs)

}

