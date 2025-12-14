package app

import (
	"fmt"
	"testing"
)

func TestWireAnalyzer_AnalyzeWireFile(t *testing.T) {
	// サンプルディレクトリのwire.goを解析
	workDir := "../../sample/basic"
	wireFilePath := "../../sample/basic/wire.go"
	searchPattern := "./..."

	analyzer := NewWireAnalyzer(workDir, searchPattern)
	results, err := analyzer.AnalyzeWireFile(wireFilePath)
	if err != nil {
		t.Fatalf("AnalyzeWireFile failed: %v", err)
	}

	// 結果を表示
	for _, result := range results {
		printStructAnalysis(t, &result, 0)
	}
}

// printStructAnalysis は構造体の解析結果を階層的に表示する
func printStructAnalysis(t *testing.T, result *StructAnalysisResult, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}

	if result.Skipped {
		t.Logf("%s[SKIPPED] %s: %s", prefix, result.StructName, result.SkipReason)
		return
	}

	t.Logf("%s%s (Package: %s)", prefix, result.StructName, result.PackagePath)

	// 初期化関数を表示
	if len(result.InitFunctions) > 0 {
		for _, initFunc := range result.InitFunctions {
			t.Logf("%s  [Init] %s (Package: %s)", prefix, initFunc.Name, initFunc.PackagePath)
		}
	}

	for _, field := range result.Fields {
		pointer := ""
		if field.IsPointer {
			pointer = "*"
		}

		if field.IsInterface {
			// インターフェースの場合
			if field.InterfaceSkipped {
				t.Logf("%s>%s -> %s%s -> [SKIPPED] %s",
					prefix, field.Name, pointer, field.TypeName, field.InterfaceSkipReason)
			} else if field.ResolvedStruct != nil {
				t.Logf("%s>%s -> %s%s ->",
					prefix, field.Name, pointer, field.TypeName)
				printStructAnalysis(t, field.ResolvedStruct, indent+1)
			} else {
				t.Logf("%s>%s -> %s%s",
					prefix, field.Name, pointer, field.TypeName)
			}
		} else if field.ResolvedStruct != nil {
			// 構造体フィールドで解決できた場合
			t.Logf("%s>%s ->",
				prefix, field.Name)
			printStructAnalysis(t, field.ResolvedStruct, indent+1)
		} else {
			// 基本型など
			t.Logf("%s>%s -> %s%s",
				prefix, field.Name, pointer, field.TypeName)
		}
	}
}

func TestWireAnalyzer_AnalyzeStruct(t *testing.T) {
	workDir := "../../sample/basic"
	searchPattern := "./..."

	analyzer := NewWireAnalyzer(workDir, searchPattern)

	// ControllerSetを直接解析
	result, err := analyzer.analyzeStruct("", "ControllerSet")
	if err != nil {
		t.Fatalf("analyzeStruct failed: %v", err)
	}

	printStructAnalysis(t, result, 0)
}

func ExampleWireAnalyzer_AnalyzeWireFile() {
	workDir := "../../sample/basic"
	wireFilePath := "../../sample/basic/wire.go"
	searchPattern := "./..."

	analyzer := NewWireAnalyzer(workDir, searchPattern)
	results, err := analyzer.AnalyzeWireFile(wireFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, result := range results {
		printStructAnalysisExample(&result, 0)
	}
}

func printStructAnalysisExample(result *StructAnalysisResult, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}

	if result.Skipped {
		fmt.Printf("%s[SKIPPED] %s: %s\n", prefix, result.StructName, result.SkipReason)
		return
	}

	fmt.Printf("%s%s (Package: %s)\n", prefix, result.StructName, result.PackagePath)

	// 初期化関数を表示
	if len(result.InitFunctions) > 0 {
		for _, initFunc := range result.InitFunctions {
			fmt.Printf("%s  [Init] %s (Package: %s)\n", prefix, initFunc.Name, initFunc.PackagePath)
		}
	}

	for _, field := range result.Fields {
		pointer := ""
		if field.IsPointer {
			pointer = "*"
		}

		if field.IsInterface {
			// インターフェースの場合
			if field.InterfaceSkipped {
				fmt.Printf("%s>%s -> %s%s -> [SKIPPED] %s\n",
					prefix, field.Name, pointer, field.TypeName, field.InterfaceSkipReason)
			} else if field.ResolvedStruct != nil {
				fmt.Printf("%s>%s -> %s%s ->\n",
					prefix, field.Name, pointer, field.TypeName)
				printStructAnalysisExample(field.ResolvedStruct, indent+1)
			} else {
				fmt.Printf("%s>%s -> %s%s\n",
					prefix, field.Name, pointer, field.TypeName)
			}
		} else if field.ResolvedStruct != nil {
			// 構造体フィールドで解決できた場合
			fmt.Printf("%s>%s ->\n",
				prefix, field.Name)
			printStructAnalysisExample(field.ResolvedStruct, indent+1)
		} else {
			// 基本型など
			fmt.Printf("%s>%s -> %s%s\n",
				prefix, field.Name, pointer, field.TypeName)
		}
	}
}
