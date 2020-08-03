package RenderinG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

//JSON Parsing helper struct
type JSONProjectInfo struct {
	Name      string
	StageSize []float64
	Scenes    []string
}

//
//Struct holding project information
//
type GProject struct {
	Name      string
	StageSize []float64
	Scenes    []GScene
}

//Get a given scenes configuration and objects
//param project: the project to extract from
//param sceneIdx: index of the scene to retrieve information about
func GetSceneProperties(project GProject, sceneIdx int) (GConfig, []GObject) {
	return project.Scenes[sceneIdx].Config, project.Scenes[sceneIdx].Objects
}

//Loads a project by its file name
//param name: name of the project to load
//returns: The project structure
func LoadProject(name string) GProject {
	absPath, err := filepath.Abs("./Projects/" + name + "/config.json")
	content, err := ioutil.ReadFile(absPath)

	if err != nil {
		log.Fatal(err)
	}

	var projectInfo JSONProjectInfo

	err = json.Unmarshal(content, &projectInfo)

	if err != nil {
		log.Fatal(err)
	}

	var scenes []GScene
	for i := range projectInfo.Scenes {
		absPath, err := filepath.Abs("Projects/" + name + "/" + projectInfo.Scenes[i] + ".json")

		if err != nil {
			log.Fatal(err)
		}

		scenes = append(scenes, LoadScene(absPath))
	}

	return GProject{
		projectInfo.Name,
		projectInfo.StageSize,
		scenes,
	}
}

//
//GPrintable interface (GProject) implementation
//
func (g GProject) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GProject")
	for i := range g.Scenes {
		printSpacer(depth)
		g.Scenes[i].Print(depth + 1)
	}
}
