package main

import (
	"bufio"
	"constant"
	"fmt"
	"os"
	"pipeline"
)

func main() {
	file, err := os.Create(constant.BigInFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(constant.BigFileSize / 8)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(constant.BigInFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	i := 1
	for v := range p {
		fmt.Println(i, v)
		i++
		if i > 100 {
			break
		}
	}

}

func chanDemo() {
	/*p := pipeline.ArraySource(3,4,8,1,2)
	for {
		if num,ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}

	}*/

	p := pipeline.ArraySource(3, 4, 8, 1, 2)
	for v := range p {
		fmt.Println(v)
	}
}

func memSort() {
	p := pipeline.InMemSort(pipeline.ArraySource(3, 4, 8, 0, 2))
	for v := range p {
		fmt.Println(v)
	}
}

func mergeDemo() {
	p1 := pipeline.InMemSort(pipeline.ArraySource(3, 5, 8, 0, 2))
	p2 := pipeline.InMemSort(pipeline.ArraySource(9, 3, 4, 7, 1))
	p := pipeline.Merge(p1, p2)
	for v := range p {
		fmt.Println(v)
	}
}
