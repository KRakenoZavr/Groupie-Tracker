package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations`
}

type Relations struct {
	Index []Relation `json:"index`
}

type Group struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Concerts     []string
}

type Groups struct {
	Arr []Group
}

func GetData(Url1, Url2 string) (Groups, error) {
	var group Groups
	resp, err := http.Get(Url1)
	if err != nil {
		fmt.Println("Url1 Json file not found")
		return group, errors.New("500")
	}
	Art, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("read data error")
		return group, errors.New("500")
	}
	resp, err = http.Get(Url2)
	if err != nil {
		fmt.Println("Url2 Json file not found")
		return group, errors.New("500")
	}
	Rel, err1 := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	var (
		artist   []Artist
		relation Relations
	)

	err = json.Unmarshal(Art, &artist)
	if err != nil {
		fmt.Println("unmarshal error")
		return group, errors.New("500")
	}

	err = json.Unmarshal(Rel, &relation)
	if err != nil {
		fmt.Println("unmarshal error")
		return group, errors.New("500")
	}
	return GetGroups(artist, relation.Index)
}

func formatLocation(str string) string {
	str = strings.Replace(str, "-", ", ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Title(str)
	return str
}
func formatDate(str string) string {
	str = strings.Replace(str, "-", ".", -1)
	return str
}

func GetGroups(artist []Artist, relation []Relation) (Groups, error) {
	var group Groups
	s := ""
	for i := range artist {
		g := Group{artist[i].ID, artist[i].Image, artist[i].Name, artist[i].Members, artist[i].CreationDate, artist[i].FirstAlbum, nil}
		for j, k := range relation[i].DatesLocations {
			s += formatLocation(j)
			s += ":"
			for _, date := range k {
				s += " "
				s += formatDate(date)
			}
			g.Concerts = append(g.Concerts, s)
			s = ""
		}
		g.Concerts = append(g.Concerts, s)
		group.Arr = append(group.Arr, g)
	}
	MakeJson(group.Arr)
	return group, nil
}

func MakeJson(Groups []Group) {

	file, err := os.Create("js/Json.json")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	var all []byte

	f, err := json.Marshal(Groups)
	if err != nil {
		fmt.Println("Cannot Marshal")
		os.Exit(1)
	}
	text := []byte("var Artists =")
	all = append(all, text...)
	all = append(all, f...)

	_, err = file.Write(all)
	if err != nil {
		fmt.Println("Cannot Write")
		os.Exit(1)
	}

	defer file.Close()
}
