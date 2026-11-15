package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func loadConfig(path string) (Config, error) {
	var config Config

	if path == "" {
		return config, fmt.Errorf("Use (--config or -config) to indicate configuration path")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("parse config %q: %w", path, err)
	}

	if config.InputFile == "" {
		return config, fmt.Errorf("config %q: input-file is required", path)
	}

	if config.OutputFile == "" {
		return config, fmt.Errorf("config %q: output-file is required", path)
	}

	return config, nil
}

func loadXML(path string) (ValCurs, error) {
	var valCurs ValCurs

	data, err := os.ReadFile(path)
	if err != nil {
		return valCurs, fmt.Errorf("read input-file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("parse XML %q: %w", path, err)
	}

	return valCurs, nil
}

func parseValue(s string) (float64, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, ",", ".")
	return strconv.ParseFloat(s, 64)
}

func sortValutes(valutes []Valute) error {
	for _, v := range valutes {
		if _, err := parseValue(v.Value); err != nil {
			return fmt.Errorf("parse value for %q: %w", v.CharCode, err)
		}
	}

	sort.Slice(valutes, func(i, j int) bool {
		vi, _ := parseValue(valutes[i].Value)
		vj, _ := parseValue(valutes[j].Value)
		return vi > vj
	})

	return nil
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := loadXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	if err := sortValutes(valCurs.Valutes); err != nil {
		panic(err)
	}

	fmt.Println("valutes:", len(valCurs.Valutes))
	for _, v := range valCurs.Valutes {
		fmt.Printf("%d %s %s\n", v.NumCode, v.CharCode, v.Value)
	}
}
