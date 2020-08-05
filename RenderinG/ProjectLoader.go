package RenderinG

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func LoadProject(name string) *GProject {
	path, err := filepath.Abs("./Projects/" + name + "/config.json")
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var projectConfig GProjectConfig
	var project GProject

	err = json.Unmarshal(bytes, &projectConfig)
	if err != nil {
		panic(err)
	}

	project.Name = projectConfig.Name
	project.StageSize = projectConfig.StageSize

	for i := range projectConfig.Scenes {
		path, err = filepath.Abs("./Projects/" + name + "/" + projectConfig.Scenes[i] + ".json")
		if err != nil {
			panic(err)
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var scene GScene
		err = json.Unmarshal(bytes, &scene)

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
