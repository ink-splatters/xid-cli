package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rodaine/table"
	"github.com/rs/xid"
)

func main() {
	infoPtr := flag.String("info", "", "print info about supplied xid")
	timestampPtr := flag.String("timestamp", "", "create xid from specified timestamp (in Unix seconds)")

	flag.Parse()

	// Handle mutually exclusive flags
	if *infoPtr != "" && *timestampPtr != "" {
		fmt.Println("Error: -info and -timestamp flags are mutually exclusive.")
		os.Exit(1)
	}

	// If neither flag is set, print a new XID
	if *infoPtr == "" && *timestampPtr == "" {
		fmt.Println(xid.New())
		return
	}

	// Handle -info flag
	if *infoPtr != "" {
		id, err := xid.FromString(*infoPtr)
		if err != nil {
			fmt.Println("Invalid XID:", err)
			os.Exit(1)
		}

		t := table.New("Component", "Value")
		t.AddRow("Machine", id.Machine())
		t.AddRow("Pid", id.Pid())
		t.AddRow("Time", id.Time())
		t.AddRow("Counter", id.Counter())
		t.Print()
		return
	}

	// Handle -timestamp flag
	if *timestampPtr != "" {
		sec, err := strconv.ParseInt(*timestampPtr, 10, 64)
		if err != nil {
			fmt.Println("Invalid timestamp:", err)
			os.Exit(1)
		}

		t := time.Unix(sec, 0)
		id := xid.NewWithTime(t)
		fmt.Println(id)
		return
	}
}
