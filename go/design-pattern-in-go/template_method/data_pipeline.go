package main

import (
	"fmt"
	"strings"
)

// DataProcessor はデータパイプライン操作のインターフェースを定義する。
type DataProcessor interface {
	// 抽象ステップ: 具象実装が必須で実装しなければならないメソッド群
	//
	// Extract はデータソースからデータを読み込む
	Extract() ([]map[string]any, error)
	// Transform は抽出したデータを加工する
	Transform(data []map[string]any) []map[string]any
	// Load は処理済みデータを保存する
	Load(data []map[string]any) error

	// オプションのステップ: デフォルト実装があり、必要な場合だけ上書き可能
	//
	// SanitizeData はデータをサニタイズする
	SanitizeData(data []map[string]any) []map[string]any
	// Validate はデータを検証する
	Validate(data []map[string]any) error

	// フック: 空のデフォルト実装を持つオプションのステップ
	// アルゴリズムの重要なステップの前後に配置され、サブクラスに拡張箇所を提供
	//
	// BeforeExtractHook は抽出前に呼ばれるフック
	BeforeExtractHook()
	// AfterLoadHook はロード後に呼ばれるフック
	AfterLoadHook()
}

// CommonProcessor は共通のパイプライン操作を提供する構造体。
// 具象クラスに埋め込むことで、デフォルト実装を継承できる。
type CommonProcessor struct{}

// オプションのステップ: デフォルト実装

// SanitizeData は全レコードの文字列フィールドをトリムし、空文字列をnilに変換する。
func (b *CommonProcessor) SanitizeData(data []map[string]any) []map[string]any {
	fmt.Println("[BasePipeline] Sanitizing data...")
	for _, record := range data {
		for key, value := range record {
			if str, ok := value.(string); ok {
				trimmed := strings.TrimSpace(str)
				if trimmed == "" {
					record[key] = nil
				} else {
					record[key] = trimmed
				}
			}
		}
	}
	return data
}

// Validate は空のデータセットでないことを検証する。
func (b *CommonProcessor) Validate(data []map[string]any) error {
	if len(data) == 0 {
		return fmt.Errorf("data is empty")
	}
	fmt.Printf("[BasePipeline] Validated: %d records\n", len(data))
	return nil
}

// フック: 空のデフォルト実装

// BeforeExtractHook は抽出前に呼ばれるフック（デフォルトは何もしない）
func (b *CommonProcessor) BeforeExtractHook() {}

// AfterLoadHook はロード後に呼ばれるフック（デフォルトは何もしない）
func (b *CommonProcessor) AfterLoadHook() {}

// RunPipeline はETLパイプラインのテンプレートを実行する。
// これがアルゴリズムの骨格を定義する「テンプレートメソッド」である。
func RunPipeline(p DataProcessor) error {
	// フック: 抽出前
	p.BeforeExtractHook()

	// 抽象ステップ: 抽出（具象クラスが実装）
	data, err := p.Extract()
	if err != nil {
		return err
	}

	// オプションのステップ: サニタイズ（デフォルト実装あり、オーバーライド可能）
	data = p.SanitizeData(data)

	// オプションのステップ: バリデーション（デフォルト実装あり、オーバーライド可能）
	if err := p.Validate(data); err != nil {
		return err
	}

	// 抽象ステップ: 変換（具象クラスが実装）
	transformed := p.Transform(data)

	// 抽象ステップ: ロード（具象クラスが実装）
	if err := p.Load(transformed); err != nil {
		return err
	}

	// フック: ロード後
	p.AfterLoadHook()

	return nil
}
