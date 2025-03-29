package lockfile

import (
	"bufio"
	"strings"

	"github.com/khulnasoft/tunnel/pkg/dependency"
	"github.com/khulnasoft/tunnel/pkg/dependency/parser/utils"
	ftypes "github.com/khulnasoft/tunnel/pkg/fanal/types"
	xio "github.com/khulnasoft/tunnel/pkg/x/io"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (Parser) Parse(r xio.ReadSeekerAt) ([]ftypes.Package, []ftypes.Dependency, error) {
	var pkgs []ftypes.Package
	scanner := bufio.NewScanner(r)
	var lineNum int
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") { // skip comments
			continue
		}

		// dependency format: group:artifact:version=classPaths
		dep := strings.Split(line, ":")
		if len(dep) != 3 { // skip the last line with lists of empty configurations
			continue
		}

		name := strings.Join(dep[:2], ":")
		version := strings.Split(dep[2], "=")[0] // remove classPaths
		pkgs = append(pkgs, ftypes.Package{
			ID:      dependency.ID(ftypes.Gradle, name, version),
			Name:    name,
			Version: version,
			Locations: []ftypes.Location{
				{
					StartLine: lineNum,
					EndLine:   lineNum,
				},
			},
			Relationship: ftypes.RelationshipUnknown,
		})

	}
	return utils.UniquePackages(pkgs), nil, nil
}
