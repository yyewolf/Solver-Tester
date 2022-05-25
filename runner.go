package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

// Juste pour simplifier
func startExecutable(executable string, size int) (*exec.Cmd, io.WriteCloser, *bufio.Scanner) {
	c := exec.Command("stdbuf", "-oL", fmt.Sprintf("./%s", executable), fmt.Sprintf("%d", size))

	si, err := c.StdinPipe()
	if err != nil {
		panic(err)
	}
	c.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	so, err := c.StdoutPipe()
	if err != nil {
		panic(err)
	}

	c.Start()

	scanner := bufio.NewScanner(so)

	return c, si, scanner
}

// n is the amount of games to play
func runGame(executable string, n int) {
	counter := &Counters{
		Name: executable,
	}
	statCounters[executable] = counter

	var wg sync.WaitGroup

	bar := progressbar.New(n)
	for i := 0; i < n; i++ {
		bar.Add(1)
		if i%1000 == 0 {
			wg.Wait()
		}
		wg.Add(1)
		go func(i int) {
			game, err := CreateWordle(counter)
			if err != nil {
				panic(err)
			}
			cmd, stdin, stdout := startExecutable(executable, len(game.Word))
			//fmt.Printf("		Starting game %d (%s) : ", i+1, game.Word)
			win, err := game.GameLoop(stdin, stdout)
			if err != nil {
				fmt.Println(color.RedString("	Erreur [%s]", err))
			} else if win {
				//fmt.Print(color.GreenString("Victoire\n"))
				counter.WinCount++
				counter.Total++
			} else {
				//fmt.Print(color.RedString("Défaite\n"))
				counter.Total++
			}
			counter.UpdateAverageTries(len(game.Tries))
			wg.Done()
			cmd.Wait()
		}(i)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("\n		All games are running")
	wg.Wait()
	fmt.Printf("	%dW / %dL (%fT | %v/guess)\n", counter.WinCount, counter.Total-counter.WinCount, counter.AverageTry, time.Duration(counter.AverageTime))
	fmt.Println()
}
