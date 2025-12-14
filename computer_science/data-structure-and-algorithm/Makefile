# `make new ARG=4-3`みたいな感じで呼ぶ
.PHONY: new
new:
	mkdir ${ARG}
	cd ${ARG} \
		&& touch go.mod \
		&& echo "module kyu08/data-structure-and-algorithm/${ARG}\n\nrequire lib v0.0.0\n\nreplace lib => ../lib\n\ngo 1.21" > go.mod \
		&& touch main.go \
		&& echo "package main \n\nimport \"lib\"\n\n func main() {\n_, nums := lib.GetIntAndIntSliceFromArg()\n}" > main.go
