package main

import (
	"bufio"
	"constant"
	"fmt"
	"os"
	"pipeline"
	"strconv"
)

func main() {
	//p,files:= createPipeline(constant.SMALL_IN_FILE_NAME,constant.SMALL_FILE_SIZE,4)
	//for _,file := range files {
	//	//TODO FIX defer may leak,no defer cannot read completely,how to do?
	//	defer file.Close()
	//}

	//p, _ := createNetworkPipeline(constant.SMALL_IN_FILE_NAME, constant.SMALL_FILE_SIZE, 4)
	//writeToFile(p, constant.SMALL_OUT_FILE_NAME)
	//printFile(constant.SMALL_OUT_FILE_NAME)

	p, _ := createPipeline(constant.BigInFileName, constant.BigFileSize, 1)
	writeToFile(p, constant.BigOutFileName)
	printFile(constant.BigOutFileName)

	//createNetworkPipeline(constant.SMALL_IN_FILE_NAME,constant.SMALL_FILE_SIZE,4)
	//time.Sleep(time.Hour)
}
func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 1
	for v := range p {
		fmt.Println(count, v)
		count++
		if count >= 1 {
			break
		}
	}
}
func writeToFile(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeline(fileName string, fileSize, chunkCount int) (<-chan int, []*os.File) {
	chunkSize := (fileSize / (8 * chunkCount)) * 8
	chunkLastSize := chunkSize

	chunkLeftSize := fileSize % (chunkCount * 8)
	if chunkLeftSize != 0 {
		chunkLastSize = chunkLeftSize + chunkLastSize
	}

	pipeline.Init()
	var sortResults []<-chan int
	var files []*os.File

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)

		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		if i == chunkCount-1 {
			chunkSize = chunkLastSize
		}

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
		files = append(files, file)
	}
	return pipeline.MergeN(sortResults...), files
}

func createNetworkPipeline(fileName string, fileSize, chunkCount int) (<-chan int, []*os.File) {
	chunkSize := (fileSize / (8 * chunkCount)) * 8
	chunkLastSize := chunkSize

	chunkLeftSize := fileSize % (chunkCount * 8)
	if chunkLeftSize != 0 {
		chunkLastSize = chunkLeftSize + chunkLastSize
	}

	pipeline.Init()
	var sortResults []<-chan int
	var files []*os.File

	var sortAddrs []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)

		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		if i == chunkCount-1 {
			chunkSize = chunkLastSize
		}

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddrs = append(sortAddrs, addr)
		files = append(files, file)
	}

	for _, addr := range sortAddrs {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...), files
}