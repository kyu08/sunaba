package main

import "fmt"

func main() {
	// 例1: CSVファイルを処理
	fmt.Println("\n########## CSV Data Pipeline ##########")
	csvProcessor := &CSVProcessor{
		SourceFile: "users.csv",
		OutputFile: "users_processed.csv",
	}
	if err := RunPipeline(csvProcessor); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 例2: APIからJSONを処理
	fmt.Println("\n########## JSON API Data Pipeline ##########")
	jsonProcessor := &JSONProcessor{
		APIEndpoint: "https://api.example.com/users",
		OutputTable: "processed_users",
	}
	if err := RunPipeline(jsonProcessor); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
