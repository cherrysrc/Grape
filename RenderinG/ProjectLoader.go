package RenderinG

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

//
//Load projects config file
func loadConfig(name string) GProjectConfig {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var projectConfig GProjectConfig
	err = json.Unmarshal(bytes, &projectConfig)
	if err != nil {
		panic(err)
	}

	return projectConfig
}

//load specific scene
func loadScene(name string) GScene {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var scene GScene
	err = json.Unmarshal(bytes, &scene)

	return scene
}

//load a given project
func LoadProject(name string) *GProject {
	projectConfig := loadConfig("./Projects/" + name + "/config.json")
	var project GProject

	project.Name = projectConfig.Name
	project.StageSize = projectConfig.StageSize

	for i := range projectConfig.Scenes {
		scene := loadScene("./Projects/" + name + "/" + projectConfig.Scenes[i] + ".json")

		for i := range scene.Objects {
			//Calculate centers for every object of every scene
			scene.Objects[i].CalculateCenter()

			//Generate ID if necessary
			scene.Objects[i].GenerateID(5)
		}

		project.Scenes = append(project.Scenes, scene)
	}

	return &project
}
