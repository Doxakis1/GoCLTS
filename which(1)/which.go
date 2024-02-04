package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func main(){
    if len(os.Args) <= 1{
        fmt.Println("Use the program as followed:")
        fmt.Println(os.Args[0] + " followed by the arguments you wish to search :)");
        os.Exit(1);
    }
    pathENV := os.Getenv("PATH");
    pathSPLIT := filepath.SplitList(pathENV);
    for _, arg := range os.Args{
    for _, directory := range pathSPLIT {
        fullpath := filepath.Join(directory, arg);
        fileinfo, err := os.Stat(fullpath);
        if err != nil {
            continue;
        }
        mode := fileinfo.Mode()
        if mode.IsRegular() && mode&0111 != 0{ //we use mode AND 0111 because we want to chdck is executable
            fmt.Println("Match found for " + arg +": " + fullpath); 
        }
    }
    }
}
