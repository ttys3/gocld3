module github.com/ttys3/gocld3

go 1.20

retract (
    // retract v0.1.0 through v0.1.2 because of memory leak issue.
    [v0.1.0, v0.1.2]
)
