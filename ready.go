package example

import "fmt"

// Ready implements the ready.Readiness interface, once this flips to true CoreDNS
// assumes this plugin is ready for queries; it is not checked again.
func (e Example) Ready() bool {
    fmt.Printf("ready\n")
    return true
}
