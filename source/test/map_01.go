package main

import "fmt"

type PersonInfo struct {
    ID string
    Name string
    Address string
}

func main() {
    var personDB map[string]PersonInfo
    personDB = make(map[string]PersonInfo)

    personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}
    personDB["1"] = PersonInfo{"1", "Jack", "Room 234"}
    personDB["2"] = PersonInfo{"2,", "Soon", "Room 456"}

    person, ok := personDB["1234"]

    if ok {
        fmt.Println("Found person", person.Name, " with ID 1234")
    } else {
        fmt.Println("Did not find person with ID 1234")
    }

    for key, value := range personDB {
        fmt.Println(key, "\t", value.Name, "\t", value.Address)
    }
}
