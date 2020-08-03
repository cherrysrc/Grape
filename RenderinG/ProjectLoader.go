package RenderinG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

//
//Struct holding project information
//
type GProject struct {
	scenes []GScene
}

func GetSceneProperties(project GProject, sceneIdx int) (GConfig, []GObject) {
	return project.scenes[sceneIdx].Config, project.scenes[sceneIdx].Objects
}

//
//Loads a project by its file name
//
func LoadProject(name string) GProject {
	absPath, err := filepath.Abs("./Projects/" + name + "/config.json")
	content, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal(err)
	}

	var sceneStrings []string

	err = json.Unmarshal(content, &sceneStrings)

	if err != nil {
		log.Fatal(err)
	}

	var scenes []GScene
	for i := 0; i < len(sceneStrings); i++ {
		absPath, err := filepath.Abs("Projects/" + name + "/" + sceneStrings[i] + ".json")

		if err != nil {
			log.Fatal(err)
		}

		scenes = append(scenes, LoadScene(absPath))
	}

	return GProject{
		scenes,
	}
}

//
//GPrintable interface (GProject) implementation
//
func (g GProject) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GProject")
	for i := range g.scenes {
		printSpacer(depth)
		g.scenes[i].Print(depth + 1)
	}
}
