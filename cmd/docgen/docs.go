package main

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/nyaruka/goflow/flows"
	"github.com/nyaruka/goflow/test"
	"github.com/nyaruka/goflow/utils"
)

var tagLineRegex = regexp.MustCompile(`@\w+\s+(?P<value>\w+)(?P<extra>.+)?`)

var docSets = []struct {
	contextKey string
	searchDirs []string
	tag        string
	handler    renderFunc
}{
	{"contextDocs", []string{"flows"}, "context", renderContextDoc},
	{"functionDocs", []string{"excellent/functions"}, "function", renderFunctionDoc},
	{"testDocs", []string{"flows/routers/tests"}, "test", renderFunctionDoc},
	{"actionDocs", []string{"flows/actions"}, "action", renderActionDoc},
	{"eventDocs", []string{"flows/events"}, "event", renderEventDoc},
	{"triggerDocs", []string{"flows/triggers"}, "trigger", renderTriggerDoc},
}

type documentedItem struct {
	typeName    string   // actual go type name
	tagName     string   // tag used to make this as a documented item
	tagValue    string   // identifier value after @tag
	tagExtra    string   // additional text after tag value
	examples    []string // any indented line
	description []string // any other line
}

type renderFunc func(output *strings.Builder, item *documentedItem, session flows.Session) error

// builds the documentation generation context from the given base directory
func buildDocsContext(baseDir string) (map[string]string, map[string]bool, error) {
	fmt.Println("Generating docs...")

	server, err := test.NewTestHTTPServer(49998)
	if err != nil {
		return nil, nil, fmt.Errorf("error starting mock HTTP server: %s", err)
	}
	defer server.Close()

	defer utils.SetRand(utils.DefaultRand)
	defer utils.SetUUIDGenerator(utils.DefaultUUIDGenerator)
	defer utils.SetTimeSource(utils.DefaultTimeSource)

	utils.SetRand(utils.NewSeededRand(123456))
	utils.SetUUIDGenerator(utils.NewSeededUUID4Generator(123456))
	utils.SetTimeSource(utils.NewFixedTimeSource(time.Date(2018, 4, 11, 18, 24, 30, 123456000, time.UTC)))

	session, err := test.CreateTestSession(server.URL, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating example session: %s", err)
	}

	context := make(map[string]string, len(docSets))
	linkTargets := make(map[string]bool)

	for _, ds := range docSets {
		var tagValues []string
		if context[ds.contextKey], tagValues, err = buildDocSet(baseDir, ds.searchDirs, ds.tag, ds.handler, session); err != nil {
			return nil, nil, err
		}
		for _, tagValue := range tagValues {
			linkTargets[ds.tag+":"+tagValue] = true
		}
	}

	return context, linkTargets, nil
}

func buildDocSet(baseDir string, searchDirs []string, tag string, renderer renderFunc, session flows.Session) (string, []string, error) {
	items := make([]*documentedItem, 0)
	for _, searchDir := range searchDirs {
		fromDir, err := findDocumentedItems(baseDir, searchDir, tag)
		if err != nil {
			return "", nil, err
		}
		items = append(items, fromDir...)
	}

	// sort documented items by their tag value
	sort.SliceStable(items, func(i, j int) bool { return items[i].tagValue < items[j].tagValue })

	fmt.Printf(" > Found %d documented items with tag %s\n", len(items), tag)

	buffer := &strings.Builder{}
	tagValues := make([]string, len(items))

	for i, item := range items {
		tagValues[i] = item.tagValue

		if err := renderer(buffer, item, session); err != nil {
			return "", nil, fmt.Errorf("error rendering %s:%s: %s", item.tagName, item.tagValue, err)
		}
	}

	return buffer.String(), tagValues, nil
}

// finds all documented items in go files in the given directory
func findDocumentedItems(baseDir string, searchDir string, tag string) ([]*documentedItem, error) {
	documentedItems := make([]*documentedItem, 0)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path.Join(baseDir, searchDir), nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	tag = "@" + tag

	for _, f := range pkgs {
		p := doc.New(f, "./", 0)
		for _, t := range p.Types {
			if strings.Contains(t.Doc, tag) {
				documentedItems = append(documentedItems, parseDocString(tag, t.Doc, t.Name))
			}
		}
		for _, t := range p.Funcs {
			if strings.Contains(t.Doc, tag) {
				documentedItems = append(documentedItems, parseDocString(tag, t.Doc, t.Name))
			}
		}
	}

	return documentedItems, nil
}

func parseDocString(tag string, docString string, typeName string) *documentedItem {
	var tagValue, tagExtra string
	examples := make([]string, 0)
	description := make([]string, 0)

	docString = removeTypeNamePrefix(docString, typeName)

	for _, l := range strings.Split(docString, "\n") {
		trimmed := strings.TrimSpace(l)

		if strings.HasPrefix(l, tag) {
			parts := tagLineRegex.FindStringSubmatch(l)
			tagValue = parts[1]
			tagExtra = parts[2]
		} else if strings.HasPrefix(l, "  ") { // examples are indented by at least two spaces
			examples = append(examples, trimmed)
		} else {
			description = append(description, l)
		}
	}

	return &documentedItem{typeName: typeName, tagName: tag[1:], tagValue: tagValue, tagExtra: tagExtra, examples: examples, description: description}
}

// Golang convention is to start all docstrings with the type name, but the actual type name can differ from how the type is
// referenced in the flow spec, so we remove it.
func removeTypeNamePrefix(docString string, typeName string) string {
	// remove type name from start of description and capitalize the next word
	if strings.HasPrefix(docString, typeName) {
		docString = strings.Replace(docString, typeName, "", 1)
		docString = strings.TrimSpace(docString)
		docString = strings.ToUpper(docString[0:1]) + docString[1:]
	}
	return docString
}
