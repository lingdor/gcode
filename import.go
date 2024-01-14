package gcode

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
)

func AddImport(code []byte, name, path string) ([]byte, error) {

	path = fmt.Sprintf("%q", path)
	//bs := file.Bytes()
	var fset = &token.FileSet{}
	var importStart, importEnd int
	var importBuf = &bytes.Buffer{}
	importBuf.WriteString("import(")

	if f, err := parser.ParseFile(fset, "", code, parser.ImportsOnly); err == nil {
		if len(f.Imports) < 1 {
			//create
			if f.Name == nil || f.Name.End() == 0 {
				return nil, ErrNoPackage
			}
			importStart = int(f.Name.End())
			importEnd = int(f.Name.End())
		} else {
			importStart = int(f.Decls[0].Pos())
			importEnd = int(f.Decls[0].End())
		}
		for _, item := range f.Imports {
			itemName := ""
			if item.Name != nil {
				itemName = item.Name.String()
			}
			itemPath := item.Path.Value
			if itemName == name && itemPath == path {

				return code, nil // no need change
			}
			if itemName == "" {
				importBuf.WriteString(fmt.Sprintf("\n    %s", itemPath))
			} else {
				importBuf.WriteString(fmt.Sprintf("\n    %s %s", itemName, itemPath))
			}
		}
	}

	if name == "" {
		importBuf.WriteString(fmt.Sprintf("\n    %s", path))
	} else {
		importBuf.WriteString(fmt.Sprintf("\n    %s %s", name, path))
	}
	importBuf.WriteString("\n)")

	w := &bytes.Buffer{}
	w.Write(code[0:importStart])
	if importStart == importEnd {
		w.WriteString("\n")
	}
	w.Write(importBuf.Bytes())
	if importStart == importEnd {
		w.WriteString("\n")
	}
	w.Write(code[importEnd:])
	return w.Bytes(), nil
}
