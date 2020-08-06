package RenderinG

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
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

func loadAnimations(name string, project GProject) []GAnimation {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	animationBlocks := strings.Split(content, ";")

	framingRegex, _ := regexp.Compile("\\((.*?)\\)")
	bodyRegex, _ := regexp.Compile("{([^}]*)}")

	var animations []GAnimation

	for i := 0; i < len(animationBlocks)-1; i++ {
		framing := framingRegex.FindString(animationBlocks[i])
		body := bodyRegex.FindString(animationBlocks[i])

		var anim GAnimation

		anim.ParseFraming(framing)
		anim.ParseBody(body, project)

		animations = append(animations, anim)
	}

	return animations
}

//load a given project
func LoadProject(name string) *GProject {
	projectConfig := loadConfig("./Projects/" + name + "/config.json")
	var project GProject

	project.Name = projectConfig.Name
	project.StageSize = projectConfig.StageSize
	project.sceneIdx = 0

	for i := range projectConfig.Scenes {
		scene := loadScene("./Projects/" + name + "/" + projectConfig.Scenes[i] + ".json")

		for i := range scene.Objects {
			//Calculate centers for every object of every scene
			scene.Objects[i].CalculateCenter()

			//Generate ID if necessary
			if scene.Objects[i].ID == "" {
				scene.Objects[i].GenerateID(5)
			}
		}

		project.Scenes = append(project.Scenes, scene)
	}

	for i := range project.Scenes {
		project.Scenes[i].Animations = loadAnimations("./Projects/"+name+"/"+projectConfig.Scenes[i]+".anim", project)
	}

	return &project
}
