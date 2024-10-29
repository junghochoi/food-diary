### Steps to Reading JSON


```go

var res interface {}
err := json.NewDecoder(req.Body).Decode(res)

if err == nil {
    return nil
}

```


### Steps to Writing JSON

```go
js, err := json.Marshal(data)
if err != nil {
    return err
}

for key, value := range headers {
    w.Header()[key] = value
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(status)
w.Write(js)

return nil
```
