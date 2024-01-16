package main

import (
	"fmt"
	"net/http"
	"os"
)

var innerHtml = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Yearn</title>
    <style>html,body{margin:0px;background:#0e0e0e;display:flex;justify-content:center;align-items:center;height:100vh;width:100vw;}img{max-height:95vh;width:auto;-webkit-user-select:none;}</style>
</head>
<body>
    <img alt="美女" title="美女" src="%s">
</body>
</html>
`

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func index(w http.ResponseWriter, r *http.Request) {
	imageUrl := getEnv("IMAGE_URL", "https://jsd.nn.ci/gh/faniteimg/images@master/img/36d05a1add1d078a75438ba833cf5afd.jpg")
	fmt.Fprintf(w, innerHtml, imageUrl)
}

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
