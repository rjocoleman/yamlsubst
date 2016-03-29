// Copyright © 2016 Robert Coleman <github@robert.net.nz>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

var config interface{}
var yamlFile = flag.String("yaml", "", "YAML file to substitute")

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
	flag.Parse()

	if *yamlFile == "" {
		log.Fatalln("YAML file not found:", *yamlFile)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	yamlFileContent, err := ioutil.ReadFile(*yamlFile)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFileContent, &config)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(string(input))
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, &config)
	if err != nil {
		log.Fatal(err)
	}
}

func interfaceJoiner(a []interface{}, sep string) string {
	s := make([]string, len(a))
	for i, v := range a {
		s[i] = v.(string)
	}

	return strings.Join(s, sep)
}
