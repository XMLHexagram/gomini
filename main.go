package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

func main() {
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	err = w.Pull(&git.PullOptions{
		RemoteName:        "origin",
		ReferenceName:     "",
		SingleBranch:      false,
		Depth:             0,
		Auth:              nil,
		RecurseSubmodules: 0,
		Progress:          nil,
		Force:             false,
		InsecureSkipTLS:   false,
		CABundle:          nil,
	})
	if err != nil {
		panic(err)
	}
	ref, err := r.Head()
	if err != nil {
		panic(err)
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println(commit)
}
