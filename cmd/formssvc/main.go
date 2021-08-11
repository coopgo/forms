package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/coopgo/forms"
)

func main() {

	ui := flag.Bool("ui", false, "ui flag")
	temp := flag.Bool("temp", false, "temporary flag (no persistent db)")
	flag.Parse()
	if *ui {
		go func() {
			fs := http.FileServer(http.Dir("./admin/public/"))
			http.Handle("/", http.StripPrefix("/admin/public", fs))
			fmt.Println("UI listing on port 8082")
			fmt.Println(http.ListenAndServe(":8082", fs))
		}()
	}

	fmt.Println("gRPC listening on port 8083")
	grpc := forms.NewGRPCTransport(":8083")

	fmt.Println("API listening on port 8080")
	rest := forms.NewRESTTransport(":8080")

	if *temp {
		// Temporary backend for testing/development
		s := forms.NewService(
			[]forms.Transport{rest},
			[]forms.Backend{forms.NewTempBackend()},
		)

		// Initialize forms for testing purpose
		s.AddForm(forms.NewUnstructuredForm("abc"))
		s.AddForm(forms.NewUnstructuredForm("def"))

		s.Run()
	} else {
		// Real database with PostgreSQL
		s := forms.NewService(
			[]forms.Transport{rest, grpc},
			[]forms.Backend{forms.NewPgsqlBackend("postgres://postgres:123456@localhost:5432/formssvc")},
		)

		s.Run()
	}
}
