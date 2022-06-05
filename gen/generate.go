package gen

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"io/fs"
)

func walk(folder string) []string {
	out := make([]string, 0)
	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	return out
}

func fsWalk(tplFS embed.FS, root string) []string {
	out := make([]string, 0)
	fs.WalkDir(tplFS, root, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	return out
}

func generate(rootpath, app, frame string, fs embed.FS) {

	dict := map[string]string{
		"AppName": app,
	}

	appTpl := fmt.Sprintf("tpl/%s", frame)

	tpls := fsWalk(fs, appTpl)
	for _, tpl := range tpls {
		relateFile := tpl[len(appTpl):]
		writeFsTplFile(fs, tpl, path.Join(rootpath, relateFile), dict)
	}

	tpls = walk(appTpl)

	for _, tpl := range tpls {
		relateFile := tpl[len(appTpl):]
		writeTplFile(tpl, path.Join(rootpath, relateFile), dict)
	}

}

func Execute(name, output, remote, frame string, fs embed.FS) (err error) {
	// root dir
	var rootpath = path.Join(output, name)
	if err = os.MkdirAll(rootpath, os.ModePerm); err != nil {
		return err
	}

	// git
	cmd := exec.Command("git", "init")
	cmd.Dir = rootpath
	if err = cmd.Run(); err != nil {
		return fmt.Errorf("git init failed: %v", err)
	}
	if remote != "" {
		cmd = exec.Command("git", "remote", "add", "origin", remote)
		cmd.Dir = rootpath
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("git add remote error: %v", err)
		}
	}

	generate(rootpath, name, frame, fs)

	cmds := [][]string{
		{"go", "mod", "init", name},
		{"go", "mod", "tidy"},
		{"swag", "init", "-g", "server/router.go"},
	}

	for _, cmdArg := range cmds {
		log.Printf("Running `%s`", strings.Join(cmdArg, " "))
		cmd = exec.Command(cmdArg[0], cmdArg[1:]...)
		cmd.Dir = rootpath
		if err = cmd.Run(); err != nil {
			return fmt.Errorf("%s error: %v", strings.Join(cmdArg, " "), err)
		}
	}

	log.Printf("Success!, cd %s && go run main.go", rootpath)
	return nil
}

func writeTplFile(tplFile string, filename string, data interface{}) error {
	tmpl, err := template.ParseFiles(tplFile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	log.Printf("load [%v] -> [%s]\n", tplFile, filename)

	return nil
}

func writeFsTplFile(res embed.FS, tplFile string, filename string, data interface{}) error {
	tmpl, err := template.ParseFS(res, tplFile)
	// tmpl, err := template.ParseFiles(tplFile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	log.Printf("load [%v] -> [%s]\n", tplFile, filename)

	return nil
}
