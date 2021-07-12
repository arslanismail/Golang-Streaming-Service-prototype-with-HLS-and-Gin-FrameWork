package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	mediaPath = filepath.Join("../", "assets/media/")
)

func main() {
	if len(os.Args) <= 1 {
		panic("Need to pass the file name in arguments")
	}
	fileName := os.Args[1]
	splitedName := strings.Split(fileName, ".")
	if len(splitedName) < 2 {
		panic("input file is not valid")
	}
	setDirName := splitedName[0]
	ffmpegFileSegmentation := fmt.Sprintf("ffmpeg -i %s -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls %s/hls/index.m3u8", fileName, setDirName)

	dirToCreate := mediaPath + "/" + setDirName + "/hls"

	err := os.MkdirAll(dirToCreate, os.ModePerm)
	if err != nil {
		panic("Paniced")
	}
	println(ffmpegFileSegmentation)

	// ------------------------ Working on File Parser till this point
	// cmd := exec.Command("/usr/bin/ffmpeg", "-i", fileName, "-profile:v", "baseline", "-level", "3.0", "-s", "640x360", "start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", "tseries/hls/index.m3u8")
	// cmd.Dir = mediaPath
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
