package config

import (
	_ "embed"
	"fmt"
	"strings"

	"go.uber.org/multierr"
)

// GenerateDocs returns MarkDown documentation generated from doc.toml.
func GenerateDocs() (string, error) {
	items, err := parseTOMLDocs(docsTOML)
	var sb strings.Builder

	// Header.
	sb.WriteString(`[//]: # (Documentation generated from docs.toml - DO NOT EDIT.)

## Table of contents

`)
	// Link to each table group.
	for _, item := range items {
		switch t := item.(type) {
		case *table:
			indent := strings.Repeat("\t", strings.Count(t.name, "."))
			name := t.name
			if i := strings.LastIndex(name, "."); i > -1 {
				name = name[i+1:]
			}
			link := strings.ReplaceAll(t.name, ".", "-")
			fmt.Fprintf(&sb, "%s- [%s](#%s)\n", indent, name, link)
		}
	}
	fmt.Fprintln(&sb)

	for _, item := range items {
		fmt.Fprintln(&sb, item)
		fmt.Fprintln(&sb)
	}

	return sb.String(), err
}

//go:embed docs.toml
var docsTOML string

// lines holds a set of contiguous lines
type lines []string

func (d lines) String() string {
	return strings.Join(d, "\n")
}

type table struct {
	name  string
	codes lines
	desc  lines
}

// String prints a table as an H2, followed by a code block and description.
func (g *table) String() string {
	link := strings.ReplaceAll(g.name, ".", "-")
	return fmt.Sprint("## ", g.name, "<a id='", link, "'></a>",
		"\n```toml\n",
		g.codes,
		"\n```\n",
		g.desc)
}

type keyval struct {
	name string
	code string
	desc lines
}

// String prints a keyval as an H3, followed by a code block and description.
func (f keyval) String() string {
	name := f.name
	if i := strings.LastIndex(name, "."); i > -1 {
		name = name[i+1:]
	}
	link := strings.ReplaceAll(f.name, ".", "-")
	return fmt.Sprint("### ", name, "<a id='", link, "'></a>",
		"\n```toml\n",
		f.code,
		"\n```\n",
		f.desc)
}

func parseTOMLDocs(s string) (items []fmt.Stringer, err error) {
	defer func() {
		if err != nil {
			err = multiErrorList(multierr.Errors(err))
		}
	}()
	globalTable := &table{name: "Global"}
	items = append(items, globalTable)
	currentTable := globalTable
	var desc lines
	for _, line := range strings.Split(s, "\n") {
		if strings.HasPrefix(line, "#") {
			// comment
			desc = append(desc, strings.TrimSpace(line[1:]))
		} else if strings.TrimSpace(line) == "" {
			// empty
			currentTable = globalTable
			if len(desc) > 0 {
				items = append(items, desc)
				desc = nil
			}
		} else if strings.HasPrefix(line, "[") {
			currentTable = &table{
				name:  strings.Trim(line, "[]"),
				codes: []string{line},
				desc:  desc,
			}
			items = append(items, currentTable)
			desc = nil
		} else {
			line = strings.TrimSpace(line)
			kv := keyval{
				name: line[:strings.Index(line, " ")],
				code: line,
				desc: desc,
			}
			if currentTable != globalTable {
				kv.name = currentTable.name + "." + kv.name
			}
			if len(kv.desc) == 0 {
				err = multierr.Append(err, fmt.Errorf("%s: missing description", kv.name))
			}
			if !strings.HasSuffix(line, "# Default") && !strings.HasSuffix(line, "# Example") {
				err = multierr.Append(err, fmt.Errorf(`%s: is neither a "# Default" or "# Example"`, kv.name))
			}
			items = append(items, kv)
			if currentTable != nil {
				currentTable.codes = append(currentTable.codes, kv.code)
			}
			desc = nil
		}
	}
	if len(desc) > 0 {
		items = append(items, desc)
	}
	return
}

type multiErrorList []error

func (m multiErrorList) Error() string {
	l := len(m)
	if l == 1 {
		return m[0].Error()
	}
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d errors:", l)
	for _, e := range m {
		fmt.Fprintf(&sb, "\n\t- %v", e)
	}
	return sb.String()
}
