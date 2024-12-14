package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"aocgen/pkg/aoc"
	"aocgen/pkg/common"
	"aocgen/pkg/gen"
	"aocgen/pkg/years"

	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var year, day int
var file string
var sample bool

var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Run benchmarks for a given puzzle or whole year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if year == 0 {
			year = time.Now().Year()
		}

		benchArgRegex := fmt.Sprintf("^Benchmark%d", year)
		if day > 0 {
			benchArgRegex += fmt.Sprintf("Day%s", gen.FormatDay(day))
		}

		cmdArgs := fmt.Sprintf("go test -benchmem -run=^$ -bench %s aocgen/pkg/year%d", benchArgRegex, year)
		c := exec.Command("bash", "-c", cmdArgs)
		out, err := c.Output()
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println(string(out))
	},
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build generated code",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		years.RegisterYears()

		gen.InitializeYearsPackages()

		years := aoc.Years()
		for _, y := range years {
			gen.InitializePackage(y)
			gen.NewBenchmarks(y)
		}
	},
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a puzzle",
	Long:  "Generate a puzzle from year and day inputs",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			year = time.Now().Year()
		}
		if day <= 0 {
			if time.Now().Month() == 12 {
				day = time.Now().Day()
			} else if time.Now().Month() == 11 && time.Now().Day() == 30 {
				day = 1
			} else {
				return errors.New("invalid day")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gen.InitializePackage(year)
		gen.NewSampleFile(year, day)
		gen.NewInputFile(year, day)
		gen.NewPuzzleFile(year, day)
		gen.InitializePackage(year)
		gen.InitializeYearsPackages()
		gen.NewBenchmarks(year)
	},
}

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Display puzzle input for a given puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(string(gen.WebInput(year, day)))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all years or list all puzzles in a year",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		years.RegisterYears()

		if year != 0 {
			puzzles := aoc.Puzzles(year)
			keys := make([]int, 0)
			keysStrings := make([]string, 0)

			for k := range puzzles {
				keys = append(keys, k)
			}
			sort.Ints(keys)
			for k := range keys {
				keysStrings = append(keysStrings, strconv.Itoa(keys[k]))
			}

			fmt.Printf("%d puzzles completed or in progress:\n", len(keys))
			fmt.Println(strings.Join(keysStrings, ", "))
			return
		}

		years := aoc.Years()
		var yearsStrings []string
		for y := range years {
			yearsStrings = append(yearsStrings, strconv.Itoa(years[y]))
		}

		fmt.Printf("%d years completed or in progress:\n", len(years))
		fmt.Println(strings.Join(yearsStrings, ", "))
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a puzzle and its input",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if year <= 0 {
			return errors.New("invalid year")
		}
		if day <= 0 {
			return errors.New("invalid day")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gen.RemovePuzzle(year, day)
		gen.RemovePuzzleInput(year, day)
		gen.RemovePuzzleSample(year, day)
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a puzzle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(file) > 0 {
			logrus.Infof("debugging file: %s", file)
			parser := regexp.MustCompile(`.*[\\/]+year([0-9]{4})[\\/]+day([0-9]{2})`)
			parsedDay := parser.FindAllStringSubmatch(file, -1)
			if parsedDay == nil {
				logrus.Fatalf("could not parse %s for year and day", file)
			}
			parsed, err := strconv.Atoi(parsedDay[0][1])
			common.Check(err)
			year = parsed

			parsed, err = strconv.Atoi(parsedDay[0][2])
			common.Check(err)
			day = parsed

			logrus.Infof("Year %d Day %d", year, day)
		}

		if year <= 0 {
			logrus.Fatal("invalid year")
		}

		years.RegisterYears()

		if day > 0 {
			runDay(year, day, sample)
			return
		}

		if day < 0 {
			logrus.Fatal("invalid day")
		}

		runYear(year, sample)
	},
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "AOC is a tool to support completing Advent of Code puzzles",
	Long:  "AOC supports generating puzzle data, including inputs directly from the website, and benchmarking answers",
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&year, "year", "y", 0, "year input")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "day input")
	rootCmd.PersistentFlags().StringVarP(&file, "filename", "f", "", "filename to run for vscode debug integration")
	rootCmd.PersistentFlags().BoolVarP(&sample, "sample", "s", false, "run on sample data when true")

	rootCmd.AddCommand(benchCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(inputCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runYear(year int, sample bool) {
	puzzles := aoc.Puzzles(year)
	for i := 1; i <= len(puzzles); i++ {
		runDay(year, i, sample)
	}
}

func runDay(year, day int, sample bool) {
	aoc.Run(year, day, aoc.NewPuzzle(year, day), aoc.Input(year, day, sample))
}

func run() {
	Execute()
}

func main() {
	opengl.Run(run)
}
