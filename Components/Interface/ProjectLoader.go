package Interface

import (
	"encoding/json"
	"fmt"
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
//Uses regex to parse .anim files
//Framing refers to the start and end frame of the animation
//Body refers to the actions to be performed in this time frame
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

		reducedBody := body[1 : len(body)-1]
		lines := strings.Split(reducedBody, "\n")

		for i := range lines {
			var anim Structures.GAnimation

			anim.ParseFraming(framing)

			anim.ParseLine(lines[i], project)

			if anim.Target == nil {
				continue
			}

			animations = append(animations, &anim)
		}
	}

	return animations
}

//Loads a project specified by its name
//Displays warnings in case of partly incompletely specified objects
//Generates animation hooks
func LoadProject(name string) *Structures.GProject {
	projectConfig := loadConfig("./Projects/" + name + "/config.json")
	var project Structures.GProject

	project.Name = projectConfig.Name
	project.StageSize = projectConfig.StageSize
	project.Init()
	project.PreProcess()

	for i := range projectConfig.Scenes {
		scene := loadScene("./Projects/" + name + "/" + projectConfig.Scenes[i] + ".json")

		for i := range scene.Objects {
			//Generate ID if necessary
			if scene.Objects[i].ID == "" {
				fmt.Printf("[Warning] Object doesn't have an ID, generating a random one.\n")
				scene.Objects[i].GenerateID(5)
			}

			//Central position is missing, just assuming zero
			if len(scene.Objects[i].GeometricCenter) == 0 {
				fmt.Printf("[Warning] Object %s doesn't have a center specified, assuming zero.\n", scene.Objects[i].ID)
				scene.Objects[i].GeometricCenter = []float64{0, 0}
			}

			//No vertices found, spit out warning
			if len(scene.Objects[i].Vertices) == 0 {
				fmt.Printf("[Warning] Object %s has no vertices.", scene.Objects[i].ID)
			}

			//Objects having no color specified default to white
			if len(scene.Objects[i].Colors) == 0 {
				fmt.Printf("[Warning] Object %s has no colors, assuming white.", scene.Objects[i].ID)
				scene.Objects[i].Colors[0] = []float64{1, 1, 1, 1}
			}
		}

		project.Scenes = append(project.Scenes, scene)
	}

	sceneOffset := 0.0
	for i := range project.Scenes {

		project.Scenes[i].Animations = loadAnimations("./Projects/"+name+"/"+projectConfig.Scenes[i]+".anim", project)
		project.GenerateAnimationHooks(project.Scenes[i].Animations, sceneOffset)

		sceneOffset += project.Scenes[i].Frames
	}

	project.PostProcess()
	project.CalculateVertices()

	return &project
}
