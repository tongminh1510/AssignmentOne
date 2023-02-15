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

	// and then a loop iterates through
	// and prints each of the slice values.
	//for _, each_ln := range text {
	//	fmt.Println(each_ln)
	//}

	num, err := strconv.Atoi(text[0]) // so dong trong file text
	//fmt.Printf("NUM LA: %v\n", num)
	if err != nil {
		// ... handle error
		panic(err)
	}

	school := make([]Student, 0, num-1) // tong hop danh sach hoc sinh trong truong

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
		//fmt.Println(StudentTmp)
	}

	//fmt.Println(school)
	rs := map[string]float64{}

	var first, second, third float64 = 0.0, 0.0, 0.0
	var classFirst, classSecond, classThird string

	for i := 0; i < len(school); i++ {
		rs[school[i].class] += school[i].score
		//fmt.Printf("\nday la school[i]class %v va school[i]score %v \n", school[i].class, school[i].score)
	}

	for i, num := range rs {
		//fmt.Printf("\nday la class %v va NUM %v + %v\n", i, rs[i], num)
		rs[i] += num
	}

	classFirst, classSecond, classThird = threeHighestScore(rs, first, second, third)
	fmt.Printf("\nThe three highest classes are: %v, %v, %v \n", classFirst, classSecond, classThird)

	var firstStudent, secondStudent, thirdStudent float64
	firstStudent = findThreeHighestStudents(school, classFirst)
	secondStudent = findThreeHighestStudents(school, classSecond)
	thirdStudent = findThreeHighestStudents(school, classThird)
	fmt.Printf("\nThe three highest scores of the top 3 classes are:\n%v: %v\n%v: %v\n%v: %v\n", classFirst, firstStudent, classSecond, secondStudent, classThird, thirdStudent)
}

func findThreeHighestStudents(school []Student, check string) float64 {
	tmp := 0.0
	for i := 0; i < len(school); i++ {
		if school[i].class == check && school[i].score >= tmp {
			tmp = school[i].score
		}
	}
	return tmp
}

func threeHighestScore(numbers map[string]float64, first float64, second float64, third float64) (string, string, string) {
	var classFirst, classSecond, classThird string
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
