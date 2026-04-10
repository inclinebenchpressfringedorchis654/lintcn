// lintcn:name no-single-use-top-level-type
// lintcn:severity warn
// lintcn:description Warn on root-level interfaces and type aliases that are only referenced once in the program.
//
// Package no_single_use_top_level_type warns on root-level type declarations
// that only add indirection for a single use site.
package no_single_use_top_level_type

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/microsoft/typescript-go/shim/ast"
	"github.com/microsoft/typescript-go/shim/checker"
	"github.com/typescript-eslint/tsgolint/internal/rule"
	"github.com/typescript-eslint/tsgolint/lintcn-rules/program_refs"
)

const maxInlineableTopLevelTypeChars = 120

func buildSingleUseTopLevelTypeMessage(name string, kind string) rule.RuleMessage {
	return rule.RuleMessage{
		Id: "singleUseTopLevelType",
		Description: fmt.Sprintf(
			"%s `%s` is used once. Inline it at the use site.",
			kind,
			name,
		),
	}
}

func declarationCharacterCount(sourceFile *ast.SourceFile, node *ast.Node) int {
	if sourceFile == nil || node == nil {
		return 0
	}

	text := sourceFile.Text()
	if node.Pos() < 0 || node.End() > len(text) || node.Pos() >= node.End() {
		return 0
	}

	return utf8.RuneCountInString(strings.TrimSpace(text[node.Pos():node.End()]))
}

func isTopLevelTypeDeclaration(node *ast.Node) bool {
	if node == nil || node.Parent == nil {
		return false
	}
	if node.Parent.Kind != ast.KindSourceFile {
		return false
	}
	if node.Name() == nil {
		return false
	}

	switch {
	case ast.IsTypeAliasDeclaration(node), ast.IsInterfaceDeclaration(node):
		return !ast.HasSyntacticModifier(node, ast.ModifierFlagsAmbient|ast.ModifierFlagsDefault|ast.ModifierFlagsExport)
	default:
		return false
	}
}

func declarationKind(node *ast.Node) string {
	if ast.IsInterfaceDeclaration(node) {
		return "interface"
	}
	return "type alias"
}

func isExportedViaNamedExports(typeChecker *checker.Checker, node *ast.Node, symbol *ast.Symbol) bool {
	if typeChecker == nil || node == nil || symbol == nil {
		return false
	}

	sourceFile := ast.GetSourceFileOfNode(node)
	if sourceFile == nil {
		return false
	}

	for _, statement := range sourceFile.Statements.Nodes {
		if !ast.IsExportDeclaration(statement) {
			continue
		}

		exportDecl := statement.AsExportDeclaration()
		if exportDecl.ModuleSpecifier != nil || exportDecl.ExportClause == nil || !ast.IsNamedExports(exportDecl.ExportClause.AsNode()) {
			continue
		}

		for _, specifierNode := range exportDecl.ExportClause.AsNamedExports().Elements.Nodes {
			specifier := specifierNode.AsExportSpecifier()
			localName := specifier.Name().AsNode()
			if specifier.PropertyName != nil {
				localName = specifier.PropertyName.AsNode()
			}

			if program_refs.SymbolAtLocation(typeChecker, localName) == symbol {
				return true
			}
		}
	}

	return false
}

func isWithinNode(node *ast.Node, ancestor *ast.Node) bool {
	for current := node; current != nil; current = current.Parent {
		if current == ancestor {
			return true
		}
	}
	return false
}

func hasRecursiveReference(references []program_refs.Reference, node *ast.Node) bool {
	for _, reference := range references {
		if reference.Node == nil || !isWithinNode(reference.Node, node) || ast.IsDeclarationName(reference.Node) {
			continue
		}
		return true
	}
	return false
}

var NoSingleUseTopLevelTypeRule = rule.Rule{
	Name: "no-single-use-top-level-type",
	Run: func(ctx rule.RuleContext, options any) rule.RuleListeners {
		checkTypeDeclaration := func(node *ast.Node) {
			if !isTopLevelTypeDeclaration(node) {
				return
			}

			name := node.Name()
			symbol := program_refs.SymbolAtLocation(ctx.TypeChecker, name)
			if symbol == nil {
				return
			}
			if isExportedViaNamedExports(ctx.TypeChecker, node, symbol) {
				return
			}

			allReferences := program_refs.FindSymbolReferences(ctx.Program, ctx.TypeChecker, symbol, program_refs.FindOptions{})
			if hasRecursiveReference(allReferences, node) {
				return
			}
			if declarationCharacterCount(ctx.SourceFile, node) > maxInlineableTopLevelTypeChars {
				return
			}

			externalReferences := program_refs.FindSymbolReferences(ctx.Program, ctx.TypeChecker, symbol, program_refs.FindOptions{ExcludeWithin: node})
			if len(externalReferences) != 1 {
				return
			}

			ctx.ReportNode(name, buildSingleUseTopLevelTypeMessage(name.Text(), declarationKind(node)))
		}

		return rule.RuleListeners{
			ast.KindTypeAliasDeclaration: checkTypeDeclaration,
			ast.KindInterfaceDeclaration: checkTypeDeclaration,
		}
	},
}
