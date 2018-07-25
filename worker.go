package main

func Worker(jobs chan Pro) {

	for j := range jobs {
		if j.GitOp == "pull" {
			j.Pull()
		}
	}
}

