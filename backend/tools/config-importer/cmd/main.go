package main

import "config-importer/internal"

func main() {
	i := internal.NewImporter()
	_ = i.Import()
}
