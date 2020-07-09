package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/adler32"
	"io/ioutil"
	"sort"
)

func main() {
	data, err := ioutil.ReadFile("types.json")
	if err != nil {
		panic(err)
	}
	var items map[string]string
	if err := json.Unmarshal(data, &items); err != nil {
		panic(err)
	}
	keys := []string{}

	for k := range items {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	generateTypes(keys)
	generateEmojiMappings(keys, items)
	generateFunctions(keys)
	generateTests(keys, items)
}

func generateTypes(keys []string) {
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, "package pine\n\nconst (")
	first := true
	for _, key := range keys {
		if first {
			fmt.Fprint(buffer, "\t")
			fmt.Fprint(buffer, key)
			fmt.Fprint(buffer, " msgType = iota\n")
			first = false
		} else {
			fmt.Fprint(buffer, "\t")
			fmt.Fprint(buffer, key)
			fmt.Fprint(buffer, "\n")
		}
	}
	fmt.Fprint(buffer, ")\n")
	if err := ioutil.WriteFile("generated_types.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func generateEmojiMappings(keys []string, items map[string]string) {
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, "package pine\n\nvar typeEmoji = map[msgType]string{\n")

	for _, key := range keys {
		fmt.Fprint(buffer, "\t")
		fmt.Fprint(buffer, key)
		fmt.Fprint(buffer, ": \"")
		fmt.Fprint(buffer, items[key])
		fmt.Fprint(buffer, "\",\n")
	}
	fmt.Fprint(buffer, "}\n\n")
	fmt.Fprint(buffer, "var typeMap = map[msgType]string{\n")
	for _, key := range keys {
		fmt.Fprint(buffer, "\t")
		fmt.Fprint(buffer, key)
		fmt.Fprint(buffer, ": \"")
		fmt.Fprint(buffer, key)
		fmt.Fprint(buffer, "\",\n")
	}
	fmt.Fprint(buffer, "}\n")
	if err := ioutil.WriteFile("generated_emoji_types.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func generateFunctions(keys []string) {
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, "package pine\n\n")
	fmt.Fprint(buffer, "import \"os\"\n\n")

	fmt.Fprintf(buffer, "type Writer interface{\n")
	for _, key := range keys {
		fmt.Fprintf(buffer, "\t%s(msg string, params ...interface{})\n", key)
	}
	fmt.Fprintf(buffer, "\tFatal(msg string, params ...interface{})\n")
	fmt.Fprintf(buffer, "}\n\n")

	for _, key := range keys {
		fmt.Fprintf(buffer, "func (w *PineWriter) %s(msg string, params ...interface{}) {\n", key)
		fmt.Fprintf(buffer, "\tw.parent.write(%s, w.name, nil, msg, params...)\n", key)
		fmt.Fprint(buffer, "}\n\n")

		fmt.Fprintf(buffer, "func (w *PineWriter) %sExtra(extra, msg string, params ...interface{}) {\n", key)
		fmt.Fprintf(buffer, "\tw.parent.write(%s, w.name, &extra, msg, params...)\n", key)
		fmt.Fprint(buffer, "}\n\n")

		fmt.Fprintf(buffer, "func (w *PineExtraWriter) %s(msg string, params ...interface{}) {\n", key)
		fmt.Fprintf(buffer, "\tw.parent.%sExtra(w.extra, msg, params...)\n", key)
		fmt.Fprint(buffer, "}\n\n")

	}

	fmt.Fprintf(buffer, "func (w *PineWriter) Fatal(msg string, params ...interface{}) {\n")
	fmt.Fprintf(buffer, "\tw.Error(msg, params...)\n")
	fmt.Fprintf(buffer, "\tos.Exit(1)\n")
	fmt.Fprint(buffer, "}\n\n")
	fmt.Fprintf(buffer, "func (w *PineWriter) FatalExtra(extra, msg string, params ...interface{}) {\n")
	fmt.Fprintf(buffer, "\tw.ErrorExtra(extra, msg, params...)\n")
	fmt.Fprintf(buffer, "\tos.Exit(1)\n")
	fmt.Fprint(buffer, "}\n\n")
	fmt.Fprintf(buffer, "func (w *PineExtraWriter) Fatal(msg string, params ...interface{}) {\n")
	fmt.Fprintf(buffer, "\tw.parent.FatalExtra(w.extra, msg, params...)\n")
	fmt.Fprint(buffer, "}\n\n")

	if err := ioutil.WriteFile("generated_functions.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func generateTests(keys []string, items map[string]string) {
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, `package pine

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/logrusorgru/aurora"
)

func testTimeProvider() time.Time {
	time, _ := time.Parse("20060102150405", "20081027102708")
	return time
}

const timeValue string = "10:27:08"

var inputChannel = make(chan string)
var lastOutput string

func testOutputProvider(msg string) {
	lastOutput = msg
}

func generateExpectedOutput(module, msg, emoji string, extra *string) string {
	prefix := fmt.Sprintf("%%s %%s  %%s ", aurora.Gray(timeValue), emoji, aurora.Magenta(module))
	ex := ""
	if extra != nil {
		ex = fmt.Sprintf("%%s ", aurora.Cyan(*extra))
	}
	return fmt.Sprintf("%%s%%s%%s\n", prefix, ex, msg)
}

func TestMain(m *testing.M) {
	pine.outputProvider = testOutputProvider
	pine.timeProvider = testTimeProvider
	pine.formatProvider = ttyFormatProvider
	os.Exit(m.Run())
}

`)

	writeBasicTest := func(inputA, inputB, method, emoji, mod string) {
		fmt.Fprintf(buffer, "\twriter := pine.NewWriter(\"%s\")\n", mod)
		fmt.Fprintf(buffer, "\twriter.%s(\"%s %%s\", \"%s\")\n", method, inputA, inputB)
		fmt.Fprintf(buffer, "\texp := generateExpectedOutput(\"%s\", \"%s %s\", \"%s\", nil)\n", mod, inputA, inputB, emoji)
		fmt.Fprint(buffer, "\tif lastOutput != exp {\n")
		fmt.Fprintf(buffer, "\t\tt.Errorf(\"Failed. Expectation was %%s, got instead %%s\", exp, lastOutput)\n")
		fmt.Fprint(buffer, "\t}\n")
	}

	writeExtraTest := func(inputA, inputB, method, emoji, mod string, extra string) {
		fmt.Fprintf(buffer, "\twriter := pine.NewWriter(\"%s\")\n", mod)
		fmt.Fprintf(buffer, "\twriter.%sExtra(\"%s\", \"%s %%s\", \"%s\")\n", method, extra, inputA, inputB)
		fmt.Fprintf(buffer, "\textraVal := \"%s\"\n", extra)
		fmt.Fprintf(buffer, "\texp := generateExpectedOutput(\"%s\", \"%s %s\", \"%s\", &extraVal)\n", mod, inputA, inputB, emoji)
		fmt.Fprint(buffer, "\tif lastOutput != exp {\n")
		fmt.Fprintf(buffer, "\t\tt.Errorf(\"Failed. Expectation was %%s, got instead %%s\", exp, lastOutput)\n")
		fmt.Fprint(buffer, "\t}\n")
	}

	writeExtraWriterTest := func(inputA, inputB, method, emoji, mod string, extra string) {
		fmt.Fprintf(buffer, "\twriter := pine.NewWriter(\"%s\")\n", mod)
		fmt.Fprintf(buffer, "\text := writer.WithExtra(\"%s\")\n", extra)
		fmt.Fprintf(buffer, "\text.%s(\"%s %%s\", \"%s\")\n", method, inputA, inputB)
		fmt.Fprintf(buffer, "\textraVal := \"%s\"\n", extra)
		fmt.Fprintf(buffer, "\texp := generateExpectedOutput(\"%s\", \"%s %s\", \"%s\", &extraVal)\n", mod, inputA, inputB, emoji)
		fmt.Fprint(buffer, "\tif lastOutput != exp {\n")
		fmt.Fprintf(buffer, "\t\tt.Errorf(\"Failed. Expectation was %%s, got instead %%s\", exp, lastOutput)\n")
		fmt.Fprint(buffer, "\t}\n")
	}

	for _, key := range keys {
		i := uint32(0)
		fmt.Fprintf(buffer, "func TestBasic%s(t *testing.T) {\n", key)
		kind := "Basic"
		writeBasicTest(testValue(items[key], kind, &i), testValue(items[key], kind, &i), key, items[key], testValue(items[key], kind, &i))
		fmt.Fprint(buffer, "}\n\n")

		fmt.Fprintf(buffer, "func TestExtra%s(t *testing.T) {\n", key)
		writeExtraTest(testValue(items[key], kind, &i), testValue(items[key], kind, &i), key, items[key], testValue(items[key], kind, &i), testValue(items[key], kind, &i))
		fmt.Fprint(buffer, "}\n\n")

		fmt.Fprintf(buffer, "func TestExtraWriter%s(t *testing.T) {\n", key)
		writeExtraWriterTest(testValue(items[key], kind, &i), testValue(items[key], kind, &i), key, items[key], testValue(items[key], kind, &i), testValue(items[key], kind, &i))
		fmt.Fprint(buffer, "}\n\n")
	}
	if err := ioutil.WriteFile("generated_test.go", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func testValue(emoji, kind string, i *uint32) string {
	iv := fmt.Sprintf("%d", *i)
	p := adler32.Checksum([]byte(emoji + kind + iv))
	*i = (*i) + 1
	a := p >> 24 & 0xFF
	b := p >> 16 & 0xFF
	c := p >> 8 & 0xFF
	return fmt.Sprintf("%02x%02x%02x", a, b, c)
}
