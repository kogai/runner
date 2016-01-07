package runner

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

type TestCmds []string

type Runner struct {
	TestDir  string
	TestCmds TestCmds
}

func (r *Runner) Run() {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	var tests TestContexts

	filepath.Walk(r.TestDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if f.IsDir() == false {
			t := &TestContext{
				Path:      path,
				IsSuccess: false,
				Retried:   0,
				TestCmds:  r.TestCmds,
				Wg:        wg,
			}
			tests = append(tests, t)
		}
		return nil
	})

	for _, t := range tests {
		wg.Add(1)
		go t.ExecTest()
	}
	wg.Wait()

	for _, t := range tests {
		log.Println(t.IsSuccess, ": ", t.Path)
	}
}

func New(testDir string, testCmds ...string) *Runner {
	return &Runner{
		TestDir:  testDir,
		TestCmds: testCmds,
	}
}
