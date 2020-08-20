<h1 align="center">Grape</h1>
<p align="center">
Go Rendering and Animation Program Experiment
</p>

<br><br><br>

<h2 align="center">
Demonstration:
</h2>

[![Animation Example](https://img.youtube.com/vi/U7uCXSVZa1s/0.jpg)](https://www.youtube.com/watch?v=U7uCXSVZa1s)

<h2 align="center">How to use</h2>

<h3 align="center">Run the example project</h3>

```
sh CompileAndRun.sh TestP
```

This will render the example project into "(ProjectName).mp4"

<h3 align="center">Make your own</h3>

0. Check the example project, in case you're lost

1. Go into the "Projects" folder, and create a new folder with your projects name
1. Create a config.json. It needs to contain the keys "Name", "StageSize" and "Scenes". 
    1. Name is simply a string, should match the folder name.
    1. StageSize is an array of length 2. It represents the width and height of your animation.
    1. Scenes is an array of strings. It contains the names of your projects scenes.
1. Create your scenes. Each scene consists of a json and an anim file. The json file is used to declare the objects of your scene. The anim file on the other hand is used for describing change aka movement/rotation etc. The filenames should match with your config.

<h3 align="center">JSON</h3>
<p align="center">How to JSON a scene</p>

A scene essentially consists of two keys:
1. Frames
1. Objects

The former is simply a number describing how many frames the current scene will have.  
The latter is more complicated. It contains all the objects used inside your scene. JSON speaking, it's an array of objects. These objects have these properties:
1. ID
    * Its the 'name' of the object. Its used to reference this particular object in the animation files.
1. geometriccenter
    * This is the position of the object inside the scene. Note that the origin is in the bottom left corner
1. transparency
    * Sets the transparency of the object. Ranges from 0.0 to 1.0
1. scale
    * Scales the object by x. Range can be whatever. Note however that a scale too large will result in the object not being visible anymore, due to its vertices being outside the srceen
1. vertices
    * Array of x-y-touples. Corner points of your geometric shape, so to speak.
1. colors
    * Defines a color for each vertex. If there are less colors than vertices, the last given color will be used for the remaining vertices. Values for RGB range from 0.0 to 1.0

Example:
```
{
  "Frames": 600,
  "Objects": [
    {
      "id": "RectangleA",
      "geometriccenter": [100, 100],
      "transparency": 1.0,
      "scale": 1.0,
      "vertices": [
        [-50, -50], [50, -50], [50, 50], [-50, 50]
      ],
      "colors": [
        [0.9, 0.9, 0.9]
      ]
    }
  ]
}
```

<h3 align="center">Anim</h3>
<p align="center">How to bring a scene to life</p>

A .anim file should always have the same name as its json counterpart.  
Animations are made out of blocks. Blocks are seperated by semicolons. Each block contains two parts:
1. Time frame
    * Specifies on which frame the actions specified in the Action set start and end.
1. Action set
    * Contains all actions to be done in the corresponding time frame

Actions always follow this pattern:  

```
[target object id][action][action parameter 1][action parameter 2] ... [action parameter n]
```
Currently available actions are:
1. move
    * moves an object to a specified position
1. rotate
    * rotates an object around its geometric center by an angle specified in degrees
1. scale
    * scales an object by a given factor
1. fade
    * sets the transparency of a given object

Comments can be specified by using #,  
Examples:  
```
#This is a comment
RectA move 100 100
RectA rotate 45
RectA scale 1.5
RectA fade 0.25
```