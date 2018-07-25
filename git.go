package main

import (
	"fmt"
	"os/exec"
	"os"
)

func (p *Pro) Clone() bool {
	filepath := p.Path
	cmd := exec.Command("/bin/bash","-c", "cd "+filepath+" &&git clone "+p.GitPath+" "+p.Name)

	bytes, err := cmd.Output()

	if err != nil {
		fmt.Println("cmd.Output: ", err)
	}

	fmt.Printf("***********%s->%s************\n%s\n",p.Name, "Clone", string(bytes))
	return true
}

func (p *Pro) Pull() bool {
	filepath := p.Path+"/"+p.Name

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		p.Clone()
		return true
	}

	cmd := exec.Command("/bin/bash","-c", "cd "+filepath+" &&git pull")

	bytes, err := cmd.Output()

	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return false
	}
	fmt.Printf("***********%s->%s************\n%s\n",p.Name, "Pull", string(bytes))
	return true
}

func (p *Pro) Checkout() {
	panic("implement me")
}

func (p *Pro) Branch() {
	panic("implement me")
}




