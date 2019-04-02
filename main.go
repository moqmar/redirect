package main

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	"codeberg.org/momar/logg"
)

func main() {
	to, ok := os.LookupEnv("TO")
	if !ok {
		logg.Error("Environment variable \"TO\" is required.")
		os.Exit(2)
	}
	to = strings.TrimSuffix(to, "/")

	address := os.ExpandEnv("$HOST:")
	if port, ok := os.LookupEnv("PORT"); ok {
		address += port
	} else {
		address += "80"
	}

	prefixExpression, _ := os.LookupEnv("PREFIX")
	prefix, err := regexp.Compile("^" + prefixExpression)
	if err != nil {
		logg.Error("\"PREFIX\" contains an invalid regular expression: %s", err)
		os.Exit(2)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		path := strings.TrimPrefix(prefix.ReplaceAllLiteralString(req.URL.Path, ""), "/")
		if len(path) > 0 {
			res.Header().Set("Location", to+"/"+path)
		} else {
			res.Header().Set("Location", to)
		}

		if permanent, ok := os.LookupEnv("PERMANENT"); ok && strings.Contains("  0 false no ", " "+strings.ToLower(permanent)+" ") {
			res.WriteHeader(301)
			res.Write([]byte{})
		} else {
			res.WriteHeader(302)
			res.Write([]byte{})
		}
	})

	logg.Info("Starting server on %s.", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		logg.Error("%s", err)
		os.Exit(1)
	}
}
