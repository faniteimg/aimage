package main

import (
	"fmt"
	"net/http"
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
    <img src="https://s2.loli.net/2022/07/29/kRVJgN795QFrxjO.jpg">
</body>
</html>
`

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, innerHtml)
}

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}