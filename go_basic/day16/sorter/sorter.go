package sorter

import (
	"bufio"
	"flag"
	"fmt"
	"go_basic/day16/sorter/algorithms/bubblesorter"
	quickSorter "go_basic/day16/sorter/algorithms/qsorter"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	infile    *string = flag.String("i", "infile", "file contains values for sorting")
	outfile   *string = flag.String("o", "outfile", "file to receive sorted values ")
	algorithm *string = flag.String("a", "qSorter", "sort algorithm")
)

func readValues(infile string) ([]int, error) {
	var values []int
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("failed to open infile", err)
		return nil, err
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err != io.EOF {
			err = err1
			return nil, err
		}
		if isPrefix {
			fmt.Println("a too long line,seems unexpected")
			return nil, err
		}
		str := string(line)
		value, err := strconv.Atoi(str)
		if err != nil {
			err = err1
			return nil, err
		}
		values = append(values, value)
	}
	return values, err
}
func writeValues(outfile string, values []int) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create outfile")
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			quickSorter.QuickSort(values)
		case "bubblesort":
			bubblesorter.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete")
		writeValues(*outfile, values)
	} else {
		fmt.Println(err)
	}

}
