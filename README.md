#Effective-go CLI

Having read [effecive go](https://golang.org/doc/effective_go.html) use this CLI
to run the code snippets (if possible).

A somewhat superfluous CLI, but it helped me learn and may be of use to you.

##Build
```
git clone https://github.com/jacec/effective-go.git
cd effective-go
make
```

##Execute CLI
```
./bin/effective-go <chapter> --code-snipet=[1..]
```

for example
```
./bin/effective-go channels -code-snipet=4
calling CodeSnipet4


--------------------------------------------------------------------------------
  func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below.
            <-sem
        }()
    }
  }
--------------------------------------------------------------------------------

now running it...


putting 8 requests on the queue, but we have a limit of 4...


press enter to exit

doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


that's that done!


that's that done!


that's that done!


that's that done!


doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


doing something for a while, this is gonna take at least 5secs...


that's that done!


that's that done!


that's that done!


that's that done!
```
