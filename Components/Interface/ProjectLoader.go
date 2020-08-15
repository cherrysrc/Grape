package Interface

import (
	"encoding/json"
	"github.com/cherrysrc/Grape/Components/Structures"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

//Load projects config file
func loadConfig(name string) Structures.GProjectConfig {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var projectConfig Structures.GProjectConfig
	err = json.Unmarshal(bytes, &projectConfig)
	if err != nil {
		panic(err)
	}

	return projectConfig
}

//load specific scene
func loadScene(name string) Structures.GScene {
	path, err := filepath.Abs(name)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var scene Structures.GScene
	err = json.Unmarshal(bytes, &scene)

	return scene
}

//Load a specific animation
func loadAnimations(name string, project Structures.GProject) []*Structures.GAnimation {
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

	var animations []*Structures.GAnimation

	for i := 0; i < len(animationBlocks)-1; i++ {
		framing := framingRegex.FindString(animationBlocks[i])
		body := bodyRegex.FindString(animationBlocks[i])

		var anim Structures.GAnimation

		anim.ParseFraming(framing)
		anim.ParseBody(body, project)

		animations = append(animations, &anim)
	}

	return animations
}

//load a given project
func LoadProject(name string) *Structures.GProject {
	projectConfig := loadConfig("./Projects/" + name + "/config.json")
	var project Structures.GProject

	project.Name = projectConfig.Name
	project.StageSize = projectConfig.StageSize
	project.Init()

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
		project.GenerateAnimationHooks(project.Scenes[i].Animations)
	}

	project.CalculateVertices()

	return &project
}
