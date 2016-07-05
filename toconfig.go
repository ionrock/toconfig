package main

import (
	"os"
	"os/exec"
	"text/template"

	"github.com/codegangsta/cli"
)

type TmplData struct {
	Filename string
}

func ApplyConfig(c *cli.Context) error {

	funcMap := template.FuncMap{
		"Get": os.Getenv,
	}

	t := c.String("template")
	tmpl, err := template.New(t).Funcs(funcMap).ParseFiles(t)
	if err != nil {
		panic(err)
	}

	d := TmplData{Filename: t}

	fh := os.Stdout
	if c.String("config") != "" {
		fh, err = os.Create(c.String("config"))
		if err != nil {
			panic(err)
		}
		defer fh.Close()
	}

	err = tmpl.Execute(fh, d)
	if err != nil {
		panic(err)
	}
	return nil
}

func RunWrapped(parts ...string) error {
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func CommandAction(c *cli.Context) error {
	return RunWrapped(c.Args()...)
}

func main() {
	app := cli.NewApp()
	app.Name = "toconfig"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "template, t",
			Usage: "The template(s) to fill in from env vars",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "",
			Usage: "Where to write the template output",
		},
	}
	app.Before = ApplyConfig
	app.Action = CommandAction
	app.Run(os.Args)
}
