# Design Notes
* gopherholes will codegen a new binary that executes the gohpercises
* gopherholes will rely on the YAML file as the map of truth for the gopherfield. CLI params will be expressed into the YAML file first, then codegen functions will turn the YAML into concrete output.
* gopherholes will require shortnames suitable for Go code var names. can be provided or generated, generation will need to check a list of shortnames for uniqueness, and can generate names as [0:5]\<index# between 01..99>
* During codegen customizable files will be stubbed out only if they do not currently exist. post generation the codegen prefix must be removed to prevent Go checkers complaining about customizations
* after 1.0 changes to customizable file templates will be considered breaking, as it's likely to require manual corrections to customizations

user interaction becomes:
1. create a repo for the gopherfield
1. build a yaml map of the desired gopherfield, or use gopherholes params to let the CLI build it
1. 1. gopherholes codegen writes the concrete gopherfield to disk
   1. at the same time gopherholes codegen writes the Go code for a binary named for the gopherfield. codegen tags are left in these files to reduce accidental breakage
1. optional edits to custom code files
1. 1. `go run` is used to test & debug? how does this work?
   1. how does IDE F5 debugging function?
1. `go build` is used to compile the `gopherfield` binary and included custom code and gophercises
1. `gopherfield` is executed to run the gophercises or any custom code such as perf tests, cumulative perf tests, etc


CLI construction of the YAML to be done first, to wet feet with YAML handling

# Gopherholes
A framework for housing GoLang coding exercises. Becaue gophers have to live somewhere.

# Generalization from the original AOC framework
## Organization
gopherHole: This may be a site, a class, a curriculum, etc. Some examples are below.

gopherTunnel: This is an optional subgrouping for another layer (or more) of shared code (EG to execute all available days for AOC for the year).

gophercise: This is a single, atomic exercise that can be executed in isolation or as part of a set. It is a distinct package and thus avoiding name collisions.  (blatantly stolen from gophercises.com)

### Some examples
#### Single tunnel
A hole with a single tunnel, using [AoC](https://adventofcode.com/) as an example
```
gopherHole: AOC
    gopherTunnel: y2022
        gophercise: d01
        gophercise: d02
    gopherTunnel: y2023
        gophercise: d01
        gophercise: d02
```
#### Multi-tunnel
A multi-tunnel example, with fake books:
```
gopherHole: bookstudies
    gopherTunnel: goOnCoding
        gopherTunnel: chapter5
            gophercise: exercise1
            gophercise: exercise2
    gopherTunnel: codeForDummies
        gopherTunnel: chapterTheFirst
            gophercise: obligatoryHelloWorld
```
#### No tunnels
A no-tunnel example with Project Euler:
```
gopherHole: projectEuler
    gophercise: multiplesOf3Or5
```
(or exercise1 to conform to [style guidelines](https://go.dev/blog/package-names) if you prefer, this framework is not dogmatic here)
#### Mixed depth
untested, not on roadmap, possibly functioning, or may be added later


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
│   └── gopherholes
├── gopherfield  # code for specific gopherholes including custom codegen logic
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