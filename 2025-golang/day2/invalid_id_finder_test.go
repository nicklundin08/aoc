package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	A "github.com/IBM/fp-go/array"
	"github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/option"
)

func Test__PrintAnswer(t *testing.T) {
	answer, err := getAnswer("input.txt")
	if err != nil {
		fmt.Println("An error occured", err)
	}

	fmt.Println(answer)
}

func Test__SuppliedCaseWithFile(t *testing.T) {
	answer, err := getAnswer("sample.txt")
	if err != nil {
		fmt.Println("An error occured", err)
	}

	want := 1227775554
	if answer != want {
		t.Errorf(`Default Case = %d, want %d`, answer, want)
	}
}

func TestIsInvalid(t *testing.T) {
	if isInvalid(100){
		t.Errorf("100 should be valid")
	}

	if isInvalid(101){
		t.Errorf("101 should be valid")
	}

	if !isInvalid(1010){
		t.Errorf("1001 should be invalid")
	}
}


func getAnswer(path string) (int, error) {
	lines, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, err
	}
	ranges := A.Map(BuildRangeFromString)(lines)
	invalidSums := A.FilterMap(func(e either.Either[error, Range]) option.Option[int] {
		mapr := option.Map(func(r Range)int{return r.InvalidIdSum()})
		result := mapr(either.ToOption(e))
		return result
	})(ranges)
	return sum(invalidSums), nil;
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}

	split := func(s string) []string { return strings.Split(s, ",") }
	result := A.Chain(split)(lines)
	return result, nil
}
