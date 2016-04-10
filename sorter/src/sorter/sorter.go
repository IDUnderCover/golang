package main

import "os"
import "flag"
import "fmt"
import "bufio"
import "strconv"
import "io"
import "time"
import "algorithms/bubblesort"

var infile *string = flag.String("i", "infile", "File contain values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted value")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		fmt.Println("read values", values)
		t1 := time.Now()
		switch *algorithm {
		case "bubblesort":
			bubblesort.Bubblesort(values)
		}
		t2 := time.Now()
		fmt.Println("sorting cost ", t2.Sub(t1))

		err1 := writeValues(values, outfile)
		if err1 != nil {
			fmt.Println("Failed to write values to file ", *outfile)
		}
		fmt.Println("write values successfully")
	} else {
		fmt.Println(err)
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexcepted.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile *string) error {
	file, err := os.Create(*outfile)
	if err != nil {
		fmt.Println("Failed to create file ", *outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil

}
