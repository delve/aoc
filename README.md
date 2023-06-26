# Generalization
exercise sets are called gopherholes

using AOC as an example
gopherhole: AOC
gopherset: year
gophercise: day (blatantly stolen from gophercises.com)

```
cmd init (--gopherhole) \<gopherhole name> (--gopherset) (\<what to call gophersets EG year>) (--gophercise) (\<what to call gohpercises EG day>)
```
Initialize a base gopherhole with syntax for subpackages utilizing the given semantics. If gopherset and gophercize are left out they default to set and exercise

Create at general 'gohperhole' package with common code (EG retrieving problem inputs for AOC). Each set is a (sub??)package containing common code for the set (EG to execute all available days for AOC for the year). Each exercise is a distinct package, providing name isolation for functions (EG paresLine()) 


# AOCgen

AOCgen is a tool to assist in solving Advent of Code in Go.

## Setup

Run AOCgen via executable: ```./aocgen```

### Commands

- **bench**: run benchmarks for a given puzzle or year of puzzles
- **build**: run code generation suite, useful for when you've had to remove any code
- **gen**: generate a puzzle
- **input**: display input for a puzzle in the console
- **list**: list all years or puzzles in a year
- **rm**: delete a puzzle and its input
- **run**: run a puzzle

## Generating Code

Use ```aocgen``` via the ```gen``` subcommand to generate code: ```./aocgen gen -y <year> -d <day>```

This will generate two files: the puzzle (```pkg/year<year>/<day>.go```) and its input (```pkg/year<year>/inputs/<day>.txt```)

Open up the puzzle and remove the DO NOT EDIT line to begin working.

Run the puzzle through the ```aocgen``` command as well: ```./aocgen run -y <year> -d <day>```

### Automatically pulling puzzle input from website

Export the environment variable ```AOC_SESSION``` with your adventofcode.com session cookie value.  Otherwise, you'll need to manually copy the input into your generated input file.

NOTE: This is set in gitpod variables https://gitpod.io/variables and loaded into gitpod automatically.

### Benchmarking

Again, use ```aocgen``` to run benchmarks for a specific day's puzzle or the entire year:

Day: ```./aocgen bench -y <year> -d <day>```

Year: ```./aocgen bench -y <year>```