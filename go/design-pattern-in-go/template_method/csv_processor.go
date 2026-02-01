package main

import "fmt"

// CSVProcessor はCSVファイル処理を担当する。
// BasePipelineを埋め込むことで、オプションのステップとフックのデフォルト実装を継承する。
type CSVProcessor struct {
	CommonProcessor
	SourceFile string
	OutputFile string
}

func (c *CSVProcessor) Extract() ([]map[string]any, error) {
	// CSVデータ読み込みをシミュレート
	fmt.Printf("  Reading from CSV file: %s\n", c.SourceFile)
	return []map[string]any{
		{"name": "Alice", "age": "30", "city": "Tokyo"},
		{"name": "Bob", "age": "25", "city": "Osaka"},
		{"name": "Charlie", "age": "35", "city": "Kyoto"},
	}, nil
}

func (c *CSVProcessor) Transform(data []map[string]any) []map[string]any {
	// CSV固有の変換: 年齢文字列をフォーマット済み文字列に変換
	for i := range data {
		data[i]["age_formatted"] = fmt.Sprintf("%s years old", data[i]["age"])
	}
	return data
}

func (c *CSVProcessor) Load(data []map[string]any) error {
	fmt.Printf("  Writing to CSV file: %s\n", c.OutputFile)
	for _, record := range data {
		fmt.Printf("    -> %v\n", record)
	}
	return nil
}

// CSVProcessorはBasePipelineを埋め込んでいるため、
// SanitizeData, Validate, BeforeExtractHook, AfterLoadHookは
// デフォルト実装が自動的に使用される。
