# jsonerr
This package works similarly to http.Error providing a quick way to respond to error conditions in handlers with json output

# Example usage 

```
const (
    MetricsErrorCode = "error.metrics"
)
/ only support posts
if r.Method != http.MethodGet {
    log.Printf("id %s: methd not allowed")
    resp := jsonerr.NewResponse(id, MetricsErrorCode, "method %s not allowed", r.Method)
    jsonerr.Error(w, resp, http.StatusMethodNotAllowed)
    return
}
```
NewResponse() method takes a trailing list of arguments to format the preceeding string with.
A naked return is required after Error() similar to http.Error() api. 
