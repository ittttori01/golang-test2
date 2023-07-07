// package main

// import (
// 	"archive/zip"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"path/filepath"
// 	"runtime"
// 	"sync"
// )

// func failOnError(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// //file download
// func download(url string) (string, error) {
// 	response, err := http.Get(url)
// 	failOnError(err)

// 	//밑에 함수 호출
// 	filename, err := urlToFileName(url)
// 	failOnError(err)

// 	f, err := os.Create(filename)
// 	failOnError(err)

// 	defer f.Close()
// 	_, err = io.Copy(f,response.Body)

// 	return filename, err
// }

// //url을 file 이름으로
// func urlToFileName (rawurl string) (string, error) {
// 	url, err := url.Parse(rawurl)
// 	failOnError(err)
// 	return filepath.Base(url.Path), nil
// }

// //압축
// func writeZip(outFilename string, filenames []string) error {
// 	outFile, err := os.Create(outFilename)
// 	failOnError(err)

// 	zipWriter := zip.NewWriter(outFile)

// 	for _, filename := range filenames{
// 		w , err := zipWriter.Create(filename)
// 		failOnError(err)

// 		f, err := os.Open(filename)
// 		failOnError(err)

// 		defer f.Close()

// 		_,err = io.Copy(w,f)
// 		failOnError(err)

// 	}
// 	return zipWriter.Close()
// }

// func main(){

// 	//최대 사용 프로세서를 4개로 설정
// 	runtime.GOMAXPROCS(4)

// 	var wait sync.WaitGroup
	
// 	//filenames[]
// 	var urlArray = []string{
// 		"http://xkxqjlzvieat874751.gcdn.ntruss.com/2/2019/d265/2d2651001bb575d64812b398661b39589500a9084c38a772f4b409035f74bf4e5_o_st.jpg",
// 		"https://file.mk.co.kr/meet/neds/2019/06/image_readtop_2019_441884_15610734753796599.jpg",
// 		"https://t1.daumcdn.net/news/201806/15/seouleconomy/20180615151617982vtvw.jpg",
// 	}

// 	//for문이 도는동안
// 	for _, url := range urlArray{

// 		wait.Add(1)
// 		//go 예약어를 통해 익명함수로 goroutine 들어갈 함수  들어온 url로 download 실행
// 		go func(url string){
// 			defer wait.Done()

// 			if _, err := download(url); err!=nil {
// 				failOnError(err)
// 			}
// 		}(url)
// 	}

// 	//go func가 끝나기 전까지 이후의 함수가 실행되지 않게.
// 	wait.Wait()

// 	filenames, err := filepath.Glob("*.jpg")
// 	failOnError(err)

// 	err = writeZip("kei_img.zip",filenames)
// 	failOnError(err)

// 	fmt.Println("++++++++++++++++DONE+++++++++++++")

// }