// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/etigerstudio/subr"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type script struct {
	key                 string
	command             string
	args                []string
	workingDir          string
	detectModifiedFiles bool
	ModifiedFilesKey    string
	watchingDir         string
}

func (s *script) Fetch(c *subr.Context) error {
	startTimestamp := time.Now()
	modified := []string{}
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(s.command, s.args...)
	cmd.Dir = s.workingDir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("%v:\nstdout:\n%sstderr:\n%s",
			err, stdout.String(), stderr.String()))
	}

	c.Buckets[s.key] = stdout.Bytes()
	duration := time.Now().Sub(startTimestamp)

	if s.detectModifiedFiles {
		// Recursively walking watching directory
		err = filepath.Walk(s.watchingDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.ModTime().After(startTimestamp) {
				modified = append(modified, path)
			}
			return nil
		})
		if err != nil {
			return err
		}

		c.Buckets[s.ModifiedFilesKey] = modified
	}

	prompt := "'" + s.command + "' is executed in " + fmt.Sprintf("%.1f", duration.Seconds()) + " seconds"
	if s.detectModifiedFiles {
		prompt += ": " + strconv.Itoa(len(modified)) + " file modifications:"
		for _, m := range modified {
			prompt += "\n" + m
		}
	}
	c.Logger.Infoln(prompt)

	return nil
}

func NewScript(key string, command string, args []string, workingDir string) *script {
	s := &script{
		key:        key,
		command:    command,
		args:       args,
		workingDir: workingDir,
	}
	return s
}

func NewScriptDetectingNewFiles(key string, command string, args []string,
	workingDir string, modifiedFilesKey string, watchingDir string) *script {
	s := &script{
		key:                 key,
		command:             command,
		args:                args,
		workingDir:          workingDir,
		detectModifiedFiles: true,
		ModifiedFilesKey:    modifiedFilesKey,
		watchingDir:         watchingDir,
	}
	return s
}
