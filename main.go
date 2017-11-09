package main

import(
  "flag"
  "fmt"
  "os"
  "sort"
  "strings"
)

func main() {

  show := flag.Bool("s", false, "show current PATH, ordered and \\n separated")
  dedup := flag.Bool("d", false, "deduplicates PATH and show")

  flag.Parse()

  if(! *show && ! *dedup) {
    fmt.Printf("Usage: %s -h\n", os.Args[0])
    os.Exit(1)
  } else {
    if(*show) { doShow() }
    if(*dedup) { doDedup() }
  }

}

func doShow() {

  path := strings.Split(os.Getenv("PATH"), ":")
  sort.Strings(path)

  for _,item := range(path) {
    fmt.Println(item)
  }
}

func doDedup() {
  items := make(map[string]bool)

  for _,item := range(strings.Split(os.Getenv("PATH"), ":")) {
    items[item] = true
  }

  var path []string

  for item,_ := range(items) {
    path = append(path,item)
  }

  fmt.Print(strings.Join(path, ":"))
}
