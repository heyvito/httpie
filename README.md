# httpie

httpie is a poor-man's mock HTTP server with a minimal interface by wrapping
`net/http/httptest`.

## Usage

```go
// httpie.JSONKV is an alias to map[string]interface{}
commonResponse := httpie.JSONKV{
    "foo":  "bar",
    "test": true,
}

httpStatusCode := func(code int) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(code)
    }
}

server := httpie.New(
    httpie.WithJSON("/foo/a", commonResponse).AddHeader("Foo-Bar", "Baz"),
    httpie.WithJSON("/foo/b", commonResponse),
    httpie.WithCustom("/foo/c", httpStatusCode(204)),
    httpie.WithCustom("/foo/coffee", func(w http.ResponseWriter, _ *http.Request) {
        w.WriteHeader(418)
    }),
)
fmt.Println(server.URL) // => Random Port http://localhost:3666
defer server.Stop()

// Do your testing here.
```

## License

```
The MIT License (MIT)

Copyright (c) 2019-2023 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

```
