package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/db-tech-uk/csync/waitgroup"
	"github.com/db-tech-uk/uiprogress"
	"github.com/db-tech-uk/uiprogress/util/strutil"
)

var steps = []string{
	"downloading source",
	"installing deps",
	"compiling",
	"packaging",
	"seeding database",
	"deploying",
	"staring servers",
}

func main() {
	fmt.Println("apps: deployment started: app1, app2")
	uiprogress.Start()

	wg := waitgroup.New()
	wg.Add(1)
	go deploy("app1", wg)
	wg.Add(1)
	go deploy("app2", wg)
	wg.Wait()

	fmt.Println("apps: successfully deployed: app1, app2")
}

func deploy(app string, wg *waitgroup.WaitGroup) {
	defer wg.Done()
	bar := uiprogress.AddBar(len(steps)).AppendCompleted().PrependElapsed()
	bar.Width = 50

	// prepend the deploy step to the bar
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return strutil.Resize(app+": "+steps[b.Current()-1], 22)
	})

	rng := rand.New(rand.NewSource(500))
	for bar.Incr() {
		time.Sleep(time.Millisecond * time.Duration(rng.Intn(2000)))
	}
}
