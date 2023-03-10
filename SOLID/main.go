package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	entries    []string
	entryCount int
}

func (j *Journal) AddEntry(text string) int {
	j.entryCount++
	entry := fmt.Sprintf("%d %s", j.entryCount, text)
	j.entries = append(j.entries, entry)

	return j.entryCount
}

func (j *Journal) RemoveEntry(index int) int {
	j.entryCount--
	j.entries = append(j.entries[:index], j.entries[index+1:]...)

	return j.entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concerns
// God Object: A Package da does everything

/// ❌ persistence should not be done by the Journal
/// I secondary function should be created for this purpose
func (j *Journal) SaveToFile(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) LoadFromFile(filename string) {
	//...funciton example
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...function example
}

/// ✅ a function created to handle persistence
/// we also separate other methods such as string() from the journal
func SaveToFile(j *Journal, filename string) {
	LineSeparator := "\n"
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("Drink water")
	j.AddEntry("Push harder")
	fmt.Println(j.String())

	SaveToFile(&j, "my_journal.txt")
}
