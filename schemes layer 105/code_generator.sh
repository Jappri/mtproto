#!/bin/sh

go run ./tl_scheme_builder.go < ./layer105.json > api.go 
gofmt -w api.go

# menjalankan tl_scheme_builder dengan input json, hasilnya letakkan di tl_schema.go