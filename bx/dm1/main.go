package main 

import (
	"fmt"
	"os"
	"bufio"
	"bingxing/pipeline"
)

func main() {
	const filename = "small.in"
	const n = 64
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file,err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}


func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3,6,2,1,9,8)),
		pipeline.InMemSort(
			pipeline.ArraySource(7,4,0,10,12)))
	for v := range p {
		fmt.Println(v)
	}
}