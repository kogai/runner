package runner

import (
	"bufio"
	"log"
	"os/exec"
	"sync"
)

const (
	MaxRetry int    = 5
	testDir  string = "./test/"
)

type TestContext struct {
	Path      string
	IsSuccess bool
	Retried   int
	TestCmds  TestCmds
	Wg        *sync.WaitGroup
}

type TestContexts []*TestContext

func (t *TestContext) ExecTest() {
	if t.Retried > MaxRetry {
		t.IsSuccess = false
		t.Wg.Done()
		return
	}

	err := t.showResult()

	if err != nil {
		t.IsSuccess = false
	}
	t.IsSuccess = err == nil

	if !t.IsSuccess {
		t.Retried++
		t.ExecTest()
		return
	}
	t.Wg.Done()
	return
}

func (t *TestContext) showResult() error {
	log.Println("Starting test in goroutine:", t.Retried, "-", t.Path)

	cmds := append(t.TestCmds, t.Path)
	cmd := exec.Command(cmds[0], cmds[1:]...)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	cmd.Wait()
	return nil
}
