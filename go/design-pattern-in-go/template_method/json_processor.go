package main

import (
	"fmt"
	"strings"
)

// JSONProcessor はJSON APIデータ処理を担当する。
// BasePipelineを埋め込むことで、オプションのステップとフックのデフォルト実装を継承する。
type JSONProcessor struct {
	CommonProcessor
	APIEndpoint string
	OutputTable string
}

func (j *JSONProcessor) Extract() ([]map[string]any, error) {
	// APIからJSONを取得するシミュレート
	fmt.Printf("  Fetching from API: %s\n", j.APIEndpoint)
	return []map[string]any{
		{"user_id": 1, "email": "alice@example.com", "status": "active"},
		{"user_id": 2, "email": "bob@example.com", "status": "inactive"},
		{"user_id": 3, "email": "charlie@example.com", "status": "active"},
	}, nil
}

func (j *JSONProcessor) Transform(data []map[string]any) []map[string]any {
	// JSON固有の変換: アクティブユーザーのフィルタリングとメール正規化
	var result []map[string]any
	for _, record := range data {
		if record["status"] == "active" {
			record["email"] = strings.ToLower(record["email"].(string))
			record["processed"] = true
			result = append(result, record)
		}
	}
	return result
}

func (j *JSONProcessor) Load(data []map[string]any) error {
	fmt.Printf("  Inserting into database table: %s\n", j.OutputTable)
	for _, record := range data {
		fmt.Printf("    -> INSERT: %v\n", record)
	}
	return nil
}

// デフォルト実装を利用したいので何も書かない（構造体を埋め込んでいるので問題なく要求されたI/Fを満たすことができる）
// func (j *JSONProcessor) SanitizeData(data []map[string]any) []map[string]any {}

// フックをオーバーライド: 抽出前にAPI接続を検証
func (j *JSONProcessor) BeforeExtractHook() {
	fmt.Println("[Hook] Validating API connection...")
}

// フックをオーバーライド: ロード成功後に通知を送信
func (j *JSONProcessor) AfterLoadHook() {
	fmt.Println("[Hook] Sending completion notification to monitoring system...")
}
