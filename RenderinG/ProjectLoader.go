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

type GProjectInstance interface {
	GetCurrentScene() GScene
	NextScene()
	GetObjectById(string) GObject
}

//
//Struct holding project information
//
type GProject struct {
	Name      string
	StageSize []float64
	Scenes    []GScene
	sceneIdx  int
}

func (project GProject) GetCurrentScene() GScene {
	return project.Scenes[project.sceneIdx]
}

func (project *GProject) NextScene() {
	project.sceneIdx++
}

func (project *GProject) GetObjectById(id string) *GObject {
	scene := project.GetCurrentScene()

	for i := range scene.Objects {
		if scene.Objects[i].ID == id {
			return &scene.Objects[i]
		}
	}

	return nil
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
		0,
	}
}

//
//GPrintable interface (GProject) implementation
//
func (project GProject) Print(depth int) {
	printSpacer(depth)
	fmt.Println("GProject")
	for i := range project.Scenes {
		printSpacer(depth)
		project.Scenes[i].Print(depth + 1)
	}
}
