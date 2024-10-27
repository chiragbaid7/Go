package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/*
we need to listen to multiple channels so 2 channels,
  - user input
  - timeout

Use select to wait for any 2 communication channels to send a value either user input or tick channel
*/
func startQuiz(records [][]string, tickChannel *time.Timer) (correct int) {

	for _, record := range records {
		userInput := make(chan int)
		go func() {
			var input int
			fmt.Println(record[0])
			fmt.Scan(&input)
			userInput <- input
		}()
		select {
		case i := <-userInput:
			if correctAns, _ := strconv.Atoi(record[1]); correctAns == i {
				correct += 1
			}
		case <-tickChannel.C:
			fmt.Println("Time out, quiz over")
			return
		}
	}
	return
}

func main() {
	//1. Parse flags and csv file
	filename := flag.String("filename", "problems.csv", "provide the filename")
	timer := flag.Int64("timer value", 5, "enter timer value")
	flag.Parse()
	f, err := os.Open(*filename) // access value to pointer
	if err != nil {
		log.Fatal("Unable to read file", err)
	}
	defer f.Close()
	parsedContent := csv.NewReader(f)
	if _, err := parsedContent.Read(); err != nil {
		panic(err)
	}
	records, _ := parsedContent.ReadAll()

	//2. Input from user
	// initialize the user input reader create standard input descriptor
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to start the quiz")
	_, errFile := reader.ReadString('\n')
	if errFile != nil {
		log.Fatal("Error", errFile)
	}

	//3. Start the timer and quiz
	// Create a timeOut channel -> when timer past duration the game should end

	tickChannel := time.NewTimer(time.Duration(*timer) * time.Second)
	ans := startQuiz(records, tickChannel)
	fmt.Println("Result::")
	fmt.Printf("Total questions: %v \n", len(records))
	fmt.Printf("Correct answers: %v \n", ans)
}

// func main() {
// 	fileName := flag.String("csvFilename", "problems.csv", "provide the filename")
// 	flag.Parse()
// 	fmt.Println(*fileName)
// 	f, err := os.Open(*fileName)
// 	if err != nil {
// 		log.Fatal("Unable to read file", err)
// 	}
// 	defer f.Close()
// 	content := csv.NewReader(f)
// 	var totalQuestions int
// 	var correct int
// 	if _, err := content.Read(); err != nil {
// 		panic(err)
// 	}
// 	records, _ := content.ReadAll()
// 	fmt.Println(records)
// 	totalQuestions = len(records)
// 	for _, record := range records {
// 		var ans int
// 		fmt.Println(record[0])
// 		fmt.Scan(&ans)
// 		correctAns, _ := strconv.Atoi(record[1])
// 		if ans == correctAns {
// 			correct += 1
// 		}
// 	}
// 	// Approach 2:
// 	// for {
// 	// 	rec, err := content.Read()
// 	// 	var ans int
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Println(rec[0])
// 	// 	fmt.Scan(&ans)
// 	// 	csvAns, _ := strconv.Atoi(rec[1])
// 	// 	if ans == csvAns {
// 	// 		correct += 1
// 	// 	}
// 	// 	totalQuestions += 1
// 	// }
// 	fmt.Println("total questions:", totalQuestions)
// 	fmt.Println("correct answers", correct)
// }
