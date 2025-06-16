package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/divan/num2words"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	command     = flag.String("command", "", "generate code for a specific command only")
	packageName = flag.String("package", "gitexec", "package name")
	output      = flag.String("output", ".", "output directory")
	buildTags   = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// CommandDoc represents the structure of a Git command documentation in JSON
type CommandDoc struct {
	CommandName string `json:"command_name"`
	Description string `json:"description"`
	Options     []struct {
		Argument    string `json:"argument"`
		Arguments   string `json:"arguments"`
		Description string `json:"description"`
	} `json:"options"`
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// accept path to docs directory as input
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"docs/git"}
	}

	docsDir := args[0]

	// Check if the provided path is a directory
	info, err := os.Stat(docsDir)
	if err != nil {
		log.Fatalf("Failed to access docs directory: %v", err)
	}

	if !info.IsDir() {
		log.Fatalf("Path '%s' is not a directory", docsDir)
	}

	// Read all JSON files in the directory
	entries, err := os.ReadDir(docsDir)
	if err != nil {
		log.Fatalf("Failed to read docs directory: %v", err)
	}

	for _, entry := range entries {
		// Skip if not a JSON file
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		jsonPath := filepath.Join(docsDir, entry.Name())
		data, err := os.ReadFile(jsonPath)
		if err != nil {
			log.Printf("Failed to read file %s: %v", jsonPath, err)
			continue
		}

		var cmd CommandDoc
		if err := json.Unmarshal(data, &cmd); err != nil {
			log.Printf("Failed to parse JSON file %s: %v", jsonPath, err)
			continue
		}

		// Skip if a specific command was requested and this isn't it
		if *command != "" && cmd.CommandName != *command {
			continue
		}

		// Generate code for this command
		g := &Generator{
			packageName: *packageName,
		}
		g.generateCode(cmd, tags)

		// Determine output file path
		outputPath := filepath.Join(*output, fmt.Sprintf("git-%s_gen.go", cmd.CommandName))

		// Write the generated code to file
		if err := os.WriteFile(outputPath, g.buf.Bytes(), 0o644); err != nil {
			log.Fatalf("Failed to write output file: %v", err)
		}

		fmt.Printf("Generated code for %s command in %s\n", cmd.CommandName, outputPath)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.

	packageName string
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// generateCode generates the Go code for a Git command
func (g *Generator) generateCode(cmd CommandDoc, tags []string) {
	// Generate the name of the execution function
	execFuncName := cases.Title(language.English).String(cmd.CommandName)
	execFuncName = strings.ReplaceAll(execFuncName, "-", "")

	// Create a new file
	f := jen.NewFile(g.packageName)

	// Add build tags if specified
	if len(tags) > 0 {
		f.Line()
		f.HeaderComment("//go:build " + strings.Join(tags, " && "))
		f.Line()
	}

	f.PackageComment(fmt.Sprintf("// Code generated for %s. DO NOT EDIT.", g.packageName))
	f.PackageComment("")

	// Add description
	f.PackageComment(fmt.Sprintf("git-%s - %s", cmd.CommandName, cmd.Description))
	f.PackageComment("")
	f.PackageComment(fmt.Sprintf("Reference: https://git-scm.com/docs/git-%s", cmd.CommandName))

	// Add imports
	f.ImportName("errors", "errors")
	f.ImportName("fmt", "fmt")
	f.ImportName("os/exec", "exec")

	// Generate options struct
	structName := execFuncName + "Options"
	structFields := []jen.Code{
		jen.Id("CmdDir").String(),
		jen.Line(),
	}

	// Add fields for command options
	for _, opt := range cmd.Options {
		fieldName := optionToFieldName(opt.Argument)
		fieldType := determineFieldType(opt.Argument)

		// Add comment with arguments
		if opt.Arguments != "" {
			args := []string{opt.Arguments}
			if strings.Contains(opt.Arguments, ",") {
				args = strings.Split(opt.Arguments, ",")
			}
			for _, arg := range args {
				structFields = append(structFields, jen.Comment(strings.TrimSpace(arg)))
			}
		}
		// Add comment with description
		structFields = append(structFields, jen.Comment(opt.Description))
		structFields = append(structFields, jen.Id(fieldName).Id(fieldType))
	}

	// Define the options struct
	f.Type().Id(structName).Struct(structFields...)
	f.Line()

	// Generate command function
	cmdFuncName := execFuncName + "Cmd"

	// Create the statements for the function body
	cmdStmts := []jen.Code{
		jen.Id("args").Op(":=").Index().String().Values(jen.Lit(cmd.CommandName)),
		jen.Line(),
	}

	// Add the option checks
	cmdStmts = append(cmdStmts, g.generateOptionChecks(cmd)...)

	// Add the remaining statements - updated to use execGit
	cmdStmts = append(cmdStmts,
		jen.Line(),
		jen.Return(jen.Id("execGit").Call(jen.Id("opts").Dot("CmdDir"), jen.Id("args"))),
	)

	// Generate the function
	f.Func().Id(cmdFuncName).Params(
		jen.Id("opts").Op("*").Id(structName),
	).Op("*").Qual("os/exec", "Cmd").Block(cmdStmts...)
	f.Line()

	// Generate execution function - updated to use run
	f.Func().Id(execFuncName).Params(
		jen.Id("opts").Op("*").Id(structName),
	).Params(
		jen.Index().Byte(),
		jen.Error(),
	).Block(
		jen.Return(jen.Id("run").Call(jen.Id(cmdFuncName).Call(jen.Id("opts")))),
	)

	// Render the file to the buffer
	if err := f.Render(&g.buf); err != nil {
		log.Fatalf("Failed to render code: %v", err)
	}
}

// generateOptionChecks generates the code to check each option and add it to args
func (g *Generator) generateOptionChecks(cmd CommandDoc) []jen.Code {
	var statements []jen.Code

	for _, opt := range cmd.Options {
		fieldName := optionToFieldName(opt.Argument)
		fieldType := determineFieldType(opt.Argument)

		switch {
		case !strings.HasPrefix(opt.Argument, "-"):
			// Handle command arguments
			switch fieldType {
			case "[]string":
				// Handle arguments with multiple options
				statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Nil()).Block(
					jen.Id("args").Op("=").Id("append").Call(jen.Id("args"), jen.Id("opts").Dot(fieldName).Op("...")),
				))
			case "string":
				fallthrough
			default:
				// Handle arguments with a single option
				statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
					jen.Id("args").Op("=").Id("append").Call(jen.Id("args"), jen.Id("opts").Dot(fieldName)),
				))
			}
		case strings.Contains(opt.Argument, "<n>"):
			optName := strings.Split(opt.Argument, "<n>")[0]
			optName = strings.ReplaceAll(optName, "[", "")
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op(">").Lit(0)).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Qual("fmt", "Sprintf").Call(
						jen.Lit(optName+"%d"),
						jen.Id("opts").Dot(fieldName),
					),
				),
			))
		case strings.Contains(opt.Argument, "<num>"):
			optName := strings.Split(opt.Argument, "<num>")[0]
			optName = strings.ReplaceAll(optName, "[", "")
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op(">").Lit(0)).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Qual("fmt", "Sprintf").Call(
						jen.Lit(optName+"%d"),
						jen.Id("opts").Dot(fieldName),
					),
				),
			))
		case strings.Contains(opt.Argument, "[="):
			// Handle options with optional values, like "--recurse-submodules[=(yes|on-demand|no)]"
			optName := strings.Split(opt.Argument, "[=")[0]
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Qual("fmt", "Sprintf").Call(
						jen.Lit(optName+"=%s"),
						jen.Id("opts").Dot(fieldName),
					),
				),
			))
		case strings.Contains(opt.Argument, "=<"), strings.Contains(opt.Argument, "=["):
			// Handle options with values using equals sign
			optName := strings.Split(opt.Argument, "=")[0]
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Qual("fmt", "Sprintf").Call(
						jen.Lit(optName+"=%s"),
						jen.Id("opts").Dot(fieldName),
					),
				),
			))
		case strings.Contains(opt.Argument, " <"):
			// Handle options with space-separated values like "--origin <name>"
			optName := strings.Split(opt.Argument, " ")[0]
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
				jen.Id("args").Op("=").Id("append").Call(jen.Id("args"), jen.Lit(optName)),
				jen.Id("args").Op("=").Id("append").Call(jen.Id("args"), jen.Id("opts").Dot(fieldName)),
			))
		case strings.HasPrefix(opt.Argument, "--no-"):
			// Handle negative boolean flags
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName)).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Lit(opt.Argument),
				),
			))
		case strings.Contains(opt.Argument, "[=<"):
			// Handle options with optional values like "--recurse-submodules[=<pathspec>]"
			optName := strings.Split(opt.Argument, "[")[0]
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
				jen.If(jen.Id("opts").Dot(fieldName).Op("==").Lit("true")).Block(
					jen.Id("args").Op("=").Id("append").Call(jen.Id("args"), jen.Lit(optName)),
				).Else().Block(
					jen.Id("args").Op("=").Id("append").Call(
						jen.Id("args"),
						jen.Qual("fmt", "Sprintf").Call(
							jen.Lit(optName+"=%s"),
							jen.Id("opts").Dot(fieldName),
						),
					),
				),
			))
		case strings.Contains(opt.Argument, "=("):
			// Handle options with parenthesis like "--path-format=(absolute|relative)"
			optName := strings.Split(opt.Argument, "=(")[0]
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName).Op("!=").Lit("")).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Qual("fmt", "Sprintf").Call(
						jen.Lit(optName+"=%s"),
						jen.Id("opts").Dot(fieldName),
					),
				),
			))
		default:
			// Handle regular boolean flags
			statements = append(statements, jen.If(jen.Id("opts").Dot(fieldName)).Block(
				jen.Id("args").Op("=").Id("append").Call(
					jen.Id("args"),
					jen.Lit(opt.Argument),
				),
			))
		}
	}

	return statements
}

// optionToFieldName converts a command line option to a struct field name
func optionToFieldName(option string) string {
	// Special handling for --
	if option == "--" {
		return "DoNotInterpretMoreArgumentsAsOptions"
	}

	if strings.HasPrefix(option, "-") {
		return optionFlagToFieldName(option)
	}

	return optionArgumentToFieldName(option)
}

// optionFlagToFieldName converts a command line option to a struct field name
func optionFlagToFieldName(option string) string {
	// Remove leading dashes
	name := strings.TrimLeft(option, "-")

	// Handle options with values
	if strings.Contains(name, "=") {
		name = strings.Split(name, "=")[0]
	}

	// Handle options with space-separated values like "--origin <name>"
	if strings.Contains(name, " ") {
		name = strings.Split(name, " ")[0]
	}

	// Handle options with square brackets like "--recurse-submodules[=<pathspec>]"
	if strings.Contains(name, "[") {
		name = strings.Split(name, "[")[0]
	}

	// Handle options with variable parameters like "--jobs=<n>"
	if strings.Contains(name, "<") {
		name = strings.Split(name, "<")[0]
	}

	// Special handling for negative options
	if strings.HasPrefix(name, "no-") {
		// For --no-prune, return NoPrune
		rest := strings.TrimPrefix(name, "no-")
		return "No" + kebabToCamelCase(rest)
	}

	if len(name) == 1 {
		if unicode.IsDigit(rune(name[0])) {
			num, _ := strconv.Atoi(name)
			name = num2words.Convert(num)
		}
	}

	return kebabToCamelCase(name)
}

func optionArgumentToFieldName(option string) string {
	name := option

	// Remove square brackets
	name = strings.Trim(name, "[]")

	// Remove trailing dots
	name = strings.TrimSuffix(name, "...")

	// Remove angle brackets
	name = strings.Trim(name, "<>")

	return kebabToCamelCase(name)
}

// determineFieldType determines the Go type for a command line option
func determineFieldType(option string) string {
	// handle arguments
	if !strings.HasPrefix(option, "-") {
		// Argument with multiple options
		if strings.Contains(option, "...") {
			return "[]string"
		}

		// Any other argument is treated as string
		return "string"
	}

	// Options with values are represented as numbers
	if strings.Contains(option, "<n>") || strings.Contains(option, "<num>") {
		return "uint64"
	}

	// Options with values are represented as strings
	if strings.Contains(option, "=") {
		return "string"
	}

	// Options with space-separated values like "--origin <name>"
	if strings.Contains(option, " <") {
		return "string"
	}

	// Options with optional values
	if strings.Contains(option, "[=<") {
		return "string"
	}

	// All other options are boolean flags
	return "bool"
}

// kebabToCamelCase converts kebab-case string to CamelCase
func kebabToCamelCase(name string) string {
	// Convert kebab-case to CamelCase
	parts := strings.Split(name, "-")
	for i := range parts {
		parts[i] = cases.Title(language.English).String(parts[i])
	}
	return strings.Join(parts, "")
}
