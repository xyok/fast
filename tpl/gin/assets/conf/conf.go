package conf

//go:generate go-bindata -nomemcopy -pkg=conf -ignore="\\.DS_Store|README.md|TRANSLATORS|auth.d" -prefix=../../ -debug=false -o=conf_gen.go ../../sample/app.ini
