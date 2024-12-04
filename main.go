package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var funcMap = template.FuncMap{
	"split":    strings.Split,
	"join":     interfaceJoiner,
	"datetime": time.Now,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
}

func main() {
	yamlFile := flag.String("yaml", "", "YAML file to substitute")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("yamlsubst %s (commit: %s) built on %s\n", version, commit, date)
		return
	}

	if *yamlFile == "" {
		log.Fatalln("YAML file not found:", *yamlFile)
	}

	if err := processTemplate(*yamlFile, os.Stdin, os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func processTemplate(yamlPath string, input io.Reader, output io.Writer) error {
	// Read input template
	yamlContent, err := os.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	// Parse YAML
	var config interface{}
	if err := yaml.Unmarshal(yamlContent, &config); err != nil {
		return err
	}

	// Read template content
	templateContent, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	// Parse and execute template
	tmpl, err := template.New("template").Funcs(funcMap).Parse(string(templateContent))
	if err != nil {
		return err
	}

	return tmpl.Execute(output, &config)
}

func interfaceJoiner(a []interface{}, sep string) string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.(string)
	}
	return strings.Join(s, sep)
}
