## GET Requests
- Use `http.Get(url)` to make a GET request
- Always close the response body with `defer resp.Body.Close()`
- Read the response body using `io.ReadAll(resp.Body)`

Example:
```go
func PerformGetRequest() {
    const myurl = "http://localhost:8000/get"
    resp, err := http.Get(myurl)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    databytes, err := io.ReadAll(resp.Body)
    content := string(databytes)
    fmt.Println("Response Content: ", content)
}
```

## POST Requests
- JSON POST: Use `http.Post(url, "application/json", requestBody)`
- Form POST: Use `http.PostForm(url, formData)`
- Create form data using `url.Values{}`

Example:
```go
func PerformPostJsonRequest() {
    const myurl = "http://localhost:8000/post"
    requestBody := strings.NewReader(`{"courseName": "Golang","price": 100,"platform": "Udemy"}`)
    resp, err := http.Post(myurl, "application/json", requestBody)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    // Read and process response...
}

func PerformPostFormRequest() {
    const myurl = "http://localhost:8000/postform"
    data := url.Values{}
    data.Add("name", "John Doe")
    data.Add("age", "25")
    resp, err := http.PostForm(myurl, data)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    // Read and process response...
}
```

## Encoding JSON
- Define structs with JSON tags for field mapping
- Use `json.MarshalIndent(data, "", "\t")` for pretty-printed JSON

Example:
```go
type course struct {
    Name     string   `json:"CourseName"`
    Price    int      `json:"CoursePrice"`
    Platform string
    Password string   `json:"-"`
    Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
    courses := []course{
        {"ReactJS", 1000, "Udemy", "password", []string{"web-dev", "frontend"}},
        {"Django", 2000, "Udemy", "password", []string{"web-dev", "backend"}},
    }
    finalJson, err := json.MarshalIndent(courses, "", "\t")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", finalJson)
}
```

## Decoding JSON
- Use `json.Unmarshal(jsonData, &target)` to parse JSON into structs
- Check JSON validity with `json.Valid(jsonData)`
- Use `map[string]interface{}` for flexible JSON handling

Example:
```go
func DecodeJson() {
    jsonData := []byte(`[{"CourseName": "ReactJS","CoursePrice": 1000,"Platform": "Udemy"}]`)
    var courses []course
    if json.Valid(jsonData) {
        json.Unmarshal(jsonData, &courses)
        fmt.Printf("%#v\n", courses)
    }
}
```

## URL Handling

- Parse URLs using `url.Parse(urlString)`
- Access URL components like Scheme, Host, Path, etc.
- Get query parameters with `url.Query()`
- Create URLs using `&url.URL{}` struct

Example:
```go
func HandleURL() {
    myurl := "https://www.google.com/search?q=golang"
    result, _ := url.Parse(myurl)
    fmt.Println(result.Scheme)
    fmt.Println(result.Host)
    fmt.Println(result.Path)
    qparams := result.Query()
    fmt.Printf("Query Params: %v\n", qparams["q"])
}
```

## Web Server Creation

- Use `mux.NewRouter()` from the Gorilla Mux package for routing
- Define routes with `r.HandleFunc("/path", handlerFunc).Methods("GET")`
- Start the server with `http.ListenAndServe(":port", router)`

Example:
```go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", serverHome).Methods("GET")
    log.Fatal(http.ListenAndServe(":4000", r))
}

func serverHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the home page"))
}
```

## Go Modules

- Initialize a new module: `go mod init`
- Add dependencies: `go get packagename`
- Verify dependencies: `go mod verify`
- Remove unused dependencies: `go mod tidy`
- List all dependencies: `go list -m all`
- Create a vendor folder: `go mod vendor`

Example:
```shell
go mod init myproject
go get github.com/gorilla/mux
go mod tidy
go mod verify
go list -m all
go mod vendor
```
