package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	class string
	name  string
	ID    string
	age   int
	score float64
	note  string
}

func main() {

	file, err := os.Open("list.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	num, err := strconv.Atoi(text[0]) // so dong trong file text
	//fmt.Printf("NUM LA: %v\n", num)
	if err != nil {
		// ... handle error
		panic(err)
	}

	school := make([]Student, 0, num-1)    // tong hop danh sach hoc sinh trong truong
	classScore := make(map[string]float64) // day la map chua lop va tong diem lop do
	scoreEach := make(map[string]float64)  // day la map chua ten lop va sinh vien co diem cao nhat lop

	for i := 1; i < num; i++ {
		var tmp []string // slice
		tmp = strings.Split(text[i], "|")

		ageTmp, _ := strconv.Atoi(tmp[3])
		scoreTmp, _ := strconv.ParseFloat(tmp[4], 64)
		noteTmp := tmp[5]

		StudentTmp := Student{
			class: tmp[0],
			name:  tmp[1],
			ID:    tmp[2],
			age:   ageTmp,
			score: scoreTmp,
			note:  noteTmp,
		}
		school = append(school, StudentTmp)

		scoreEach[StudentTmp.class] = findThreeHighestStudents(StudentTmp.score, scoreEach[StudentTmp.class])
		classScore[StudentTmp.class] += scoreTmp // tinh tong diem cac lop

	}

	var classFirst, classSecond, classThird string
	classFirst, classSecond, classThird = threeHighestClasses(classScore)
	fmt.Printf("\nThe three highest classes are: %v, %v, %v \n", classFirst, classSecond, classThird)

	var firstStudent, secondStudent, thirdStudent float64
	firstStudent, secondStudent, thirdStudent = threeHighestScores(scoreEach)
	fmt.Printf("\nThe three highest scores of the top 3 classes are:\n%v: %v\n%v: %v\n%v: %v\n", classFirst, firstStudent, classSecond, secondStudent, classThird, thirdStudent)
}

func findThreeHighestStudents(newScore float64, oldScore float64) float64 {
	if newScore > oldScore {
		return newScore
	} else {
		return oldScore
	}
}

func threeHighestClasses(numbers map[string]float64) (string, string, string) {
	var classFirst, classSecond, classThird string
	first, third, second := 0.0, 0.0, 0.0
	for i, _ := range numbers {
		if numbers[i] > first {
			third = second
			second = first
			first = numbers[i]

			classThird = classSecond
			classSecond = classFirst
			classFirst = i
			//fmt.Printf("\ndinh dang cua num la %T %T\n", num, i)
		} else if numbers[i] > third && numbers[i] != second {
			third = second
			second = numbers[i]

			classThird = classSecond
			classSecond = i
		}
	}
	return classFirst, classSecond, classThird
}

func threeHighestScores(numbers map[string]float64) (float64, float64, float64) {
	first, third, second := 0.0, 0.0, 0.0
	for i, _ := range numbers {
		if numbers[i] > first {
			third = second
			second = first
			first = numbers[i]
		} else if numbers[i] > third && numbers[i] != second {
			third = second
			second = numbers[i]
		}
	}
	return first, second, third
}
