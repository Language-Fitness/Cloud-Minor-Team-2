package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type Question struct {
	Type                 string
	QuestionText         string
	OpenQuestionAnswer   string
	MultipleChoiceA      string
	MultipleChoiceB      string
	MultipleChoiceC      string
	MultipleChoiceD      string
	MultipleChoiceAnswer string
}

func main() {
	// Open the Excel file.
	f, err := excelize.OpenFile("C:\\Users\\jaimy\\GolandProjects\\Cloud-Minor-Team-2\\Excel_Data\\src\\TestData.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize the list to store questions.
	var questions []Question

	// Loop through the sheets in the Excel file.
	for _, sheetName := range f.GetSheetList() {
		rows, err := f.GetRows(sheetName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Iterate through the rows and create Question structs.
		for rowIndex, row := range rows {
			if rowIndex == 0 { // Skip the first row (index 0).
				continue
			}

			// Depending on the type of question, handle it accordingly.
			if row[0] == "Multiple choice" {
				// Multiple choice question
				if len(row) != 8 {
					fmt.Printf("Invalid row in sheet %s: %v\n", sheetName, row)
					continue
				}
				question := Question{
					Type:                 row[0],
					QuestionText:         row[1],
					MultipleChoiceA:      row[3],
					MultipleChoiceB:      row[4],
					MultipleChoiceC:      row[5],
					MultipleChoiceD:      row[6],
					MultipleChoiceAnswer: row[7],
				}
				questions = append(questions, question)
			} else if row[0] == "Open" {
				// Open question
				if len(row) != 3 {
					fmt.Printf("Invalid row in sheet %s: %v\n", sheetName, row)
					continue
				}
				question := Question{
					Type:               row[0],
					QuestionText:       row[1],
					OpenQuestionAnswer: row[2],
				}
				questions = append(questions, question)
			}
		}
	}

	// Print or process the list of questions.
	for _, q := range questions {
		if q.Type == "Multiple choice" {
			fmt.Printf("Type: %s\n", q.Type)
			fmt.Printf("Question: %s\n", q.QuestionText)
			fmt.Printf("Multiple Choice A: %s\n", q.MultipleChoiceA)
			fmt.Printf("Multiple Choice B: %s\n", q.MultipleChoiceB)
			fmt.Printf("Multiple Choice C: %s\n", q.MultipleChoiceC)
			fmt.Printf("Multiple Choice D: %s\n", q.MultipleChoiceD)
			fmt.Printf("Multiple Choice Answer: %s\n", q.MultipleChoiceAnswer)
			fmt.Println()
		}
		if q.Type == "Open" {
			fmt.Printf("Type: %s\n", q.Type)
			fmt.Printf("Question: %s\n", q.QuestionText)
			fmt.Printf("Open Answer: %s\n", q.OpenQuestionAnswer)
			fmt.Println()
		}
	}
}
