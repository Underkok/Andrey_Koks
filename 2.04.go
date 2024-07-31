package main

import "fmt"

func main() {
	IssueEditions := make(map[string]map[string][]string)
	IssueEditions["Толстой Л.Н."] = map[string][]string{
		"Война и мир":   []string{"Кокарев А,Н", "Казаков С.С."},
		"Анна Каренина": []string{"Джон Уик", "Путин В.В."},
	}
	totalReader := 0
	for _, books := range IssueEditions {
		for _, reader := range books {
			totalReader += len(reader)
		}
	}
	fmt.Println("Общее количество экземпляров на руках читателей", totalReader)

}
