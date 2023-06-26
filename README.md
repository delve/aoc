# Generalization
## Organization
gopherHole: This may be a site, a class, as curriculum, etc. Some examples are below.

gopherTunnel: This is an optional subgrouping for another layer (or more) of shared code (EG to execute all available days for AOC for the year).

gophercise: This is a single, atomic exercise that can be executed in isolation or as part of a set. It is a distinct package and thus avoiding name collisions.  (blatantly stolen from gophercises.com)

### Some examples

A hole with a single tunnel, using [AoC](https://adventofcode.com/) as an example
```
gopherHole: AOC
gopherTunnel: year
gophercise: day
```
A multi-tunnel example, with fake books:
```
gopherHole: bookstudies
gopherTunnel: goOnCoding
gopherTunnel: chapter5
gophercise: exercise1
```
A no-tunnel example with Project Euler:
```
gopherHole: projectEuler
gophercise: multiplesOf3Or5
```
(or exercise1 to conform to [style guidelines](https://go.dev/blog/package-names))

## commands
```
cmd init (--gopherHole) \<gopherHole name> (--gopherTunnel) (\<what to call gopherTunnelss EG year>) (--gophercise) (\<what to call gohpercises EG day>)
```
Initialize a base gopherHole with syntax for subpackages utilizing the given semantics. If gopherset and gophercize are left out they default to set and exercise

Create a 'gohperhole' package with common code (EG codegen stubs for tunnels &&|| exercises, retrieving problem inputs for AOC). Codegen files must be fully boilerplate. Write inputs to a config file in the package dir for re-generation && tweaking. gopherhole is a required var, but manual config file can be used for orther params

if custom code files DNE then create them and erase the codegen protection line. if they already exist then warn and do nothing. create override param to force rewrite of custom files.

Each gopherTunnel is a subpackage containing common code for the tunnel. 


init 



## Layout
Abusing the "standard" layout somewhat.
```
.
├── cmd          # entry point, CLI code. may contain gopherhole names in codegen files
│   └── aocgen
├── gopherholes  # code for specific gopherholes including custom codegen logic
└── pkg          # code for the framework only.
    └── common   # common code leveraged all over the place (generics like the Check() shorthand for testing 'e')
```

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