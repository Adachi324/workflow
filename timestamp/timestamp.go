package main

// Package is called aw
import (
	"fmt"
	"github.com/deanishe/awgo"
	"strconv"
	"strings"
	"time"
)

// Workflow is the main API
var wf *aw.Workflow

const splitArr = ".-:"

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	// Add a "Script Filter" result
	args := wf.Args()
	result := time.Time{}
	if len(args) == 0 {

		result = time.Now()
	} else if len(args) == 1 {
		//t := args[0]
		//
		//index := strings.IndexAny(t, splitArr)
		var err error
		result, err = time.Parse(args[0], args[0])
		if err != nil {
			split := strings.Split(args[0], " ")
			if len(split) == 1 {

				ymd, err := getSplitedArr(split[0])
				if len(ymd) == 1 {
					if err != nil {
						wf.NewItem("Invalid timestamp, should be nums only")
						wf.SendFeedback()
						return
					}
					result = time.Unix(int64(ymd[0]), 0)
				} else if len(ymd) == 3 {
					result = time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, time.Local)
				}
			} else {

				ymd, err := getSplitedArr(split[0])
				if len(ymd) == 1 {

					if err != nil {
						wf.NewItem("Invalid timestamp, should be nums only")
						wf.SendFeedback()
						return
					}
					result = time.Unix(int64(ymd[0]), 0)
				} else if len(ymd) == 3 {

					hms, err := getSplitedArr(split[1])
					if err != nil {
						wf.NewItem("Invalid timestamp, should be nums only")
						wf.SendFeedback()
						return
					}
					result = time.Date(ymd[0], time.Month(ymd[1]), ymd[2], hms[0], hms[1], hms[2], 0, time.Local)
				}
			}

		}
	}
	//wf.NewItem("First result!!!")
	// Send results to Alfred
	wf.NewItem(fmt.Sprintf("timestamp: %d", result.Unix())).Arg(strconv.FormatInt(result.Unix(), 10)).Valid(true)
	wf.NewItem(fmt.Sprintf("format: %s", result.String())).Var("format", result.String()).Arg(result.String()).Valid(true)
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

func getSplitedArr(str string) ([]int, error) {
	index := strings.IndexAny(str, splitArr)
	if index == -1 {
		tInt, err := strconv.Atoi(str)
		return []int{tInt}, err
	}
	ft := strings.Split(str, string(str[index]))
	arr := make([]int, len(ft))
	for i, v := range ft {
		tInt, _ := strconv.Atoi(v)
		arr[i] = tInt
	}
	if len(arr) < 3 {
		arr = append(arr, 0)
	}
	if len(arr) > 3 {
		arr = arr[:3]
	}
	return arr, nil
}
