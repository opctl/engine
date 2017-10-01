package dirs

import (
	"fmt"
	"github.com/opspec-io/sdk-golang/model"
	"path/filepath"
	"strings"
)

func (d _Dirs) Interpret(
	pkgHandle model.PkgHandle,
	scope map[string]*model.Value,
	scgContainerCallDirs map[string]string,
	scratchDirPath string,
) (map[string]string, error) {
	dcgContainerCallDirs := map[string]string{}
dirLoop:
	for scgContainerDirPath, dirExpression := range scgContainerCallDirs {

		if "" == dirExpression {
			// bound implicitly
			dirExpression = fmt.Sprintf("$(%v)", scgContainerDirPath)
		}

		if !strings.HasPrefix(dirExpression, "$(") && !strings.HasSuffix(dirExpression, ")") {
			// handle deprecated ref
			dirExpression = fmt.Sprintf("$(%v)", dirExpression)
		}

		dirValue, err := d.expression.EvalToDir(
			scope,
			dirExpression,
			pkgHandle,
		)
		if nil != err {
			return nil, fmt.Errorf(
				"unable to bind %v to %v; error was %v",
				scgContainerDirPath,
				dirExpression,
				err,
			)
		}

		if "" != *dirValue.Dir && !strings.HasPrefix(*dirValue.Dir, d.rootFSPath) {
			// bound to non rootFS dir
			dcgContainerCallDirs[scgContainerDirPath] = *dirValue.Dir
			continue dirLoop
		}
		dcgContainerCallDirs[scgContainerDirPath] = filepath.Join(scratchDirPath, scgContainerDirPath)

		if "" == *dirValue.Dir {

			if err := d.os.MkdirAll(
				dcgContainerCallDirs[scgContainerDirPath],
				0700,
			); nil != err {
				return nil, err
			}

		} else {

			if err := d.dirCopier.OS(
				*dirValue.Dir,
				dcgContainerCallDirs[scgContainerDirPath],
			); nil != err {
				return nil, fmt.Errorf(
					"unable to bind %v to %v; error was %v",
					scgContainerDirPath,
					dirExpression,
					err,
				)
			}

		}
	}
	return dcgContainerCallDirs, nil
}