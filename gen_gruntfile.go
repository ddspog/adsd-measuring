package main

import (
	"html/template"
	"os"
)

func main() {
	// Create a new template by wrapping the JavaScript contents
	// in an HTML <script> tag.
	t, err := template.New("").Parse(`'use strict';
module.exports = function (grunt) {

	grunt.loadNpmTasks('grunt-contrib-compress');

	// Project configuration.
	grunt.initConfig({
		compress: {
			main: {
				options: {
					archive: function () {
						return 'temp/output.{{ .Compression}}';
					},
					mode: "{{ .Compression}}",
					quality: {{ .Quality}}
				},
				files: [{
					cwd: '.',
					src: 'temp/{{ .Size }}.{{ .Type }}'
				}]
			}
		}
	});

	grunt.registerTask('default', ['compress']);

};`)
	if err != nil {
		panic(err)
	}

	var variables = map[string]string{
		"Type":        os.Args[1],
		"Size":        os.Args[2],
		"Compression": os.Args[3],
		"Quality":     os.Args[4],
	}

	f, _ := os.Create("Gruntfile.js")
	_ = f.Truncate(0)

	defer f.Close()

	if err := t.Execute(f, variables); err != nil {
		panic(err)
	}
}
