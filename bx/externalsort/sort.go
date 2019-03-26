package main 

import (
	"os"
	"fmt"
	"bufio"
	"bingxing/pipeline"
)

func main() {
	pipeline.Init()
	p := createPipeline("large.in", 80000000, 4)
	writeTofile(p, "large.out")
	printFile("large.out")
}

func writeTofile(p <-chan int, filename string) {
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func printFile(filename string) {
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	p := pipeline.ReaderSource(file, -1)
	count := 0	
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func createPipeline(filename string, filesize, chunkcount int ) <-chan int {
	chunkSize := filesize/chunkcount
	pipeline.Init()
	sortResults := []<-chan int{}
	for i:=0; i<chunkcount; i++ {
		file,err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))

	}
	return pipeline.MergeN(sortResults...)
}