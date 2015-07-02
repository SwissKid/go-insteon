package main

import (
    //"net/http"
    //"io/ioutil"
    "fmt"
    //"encoding/json"
    "../insteon"
    //"os"
    //"strconv"
    "log"
    "strings"
)

func checkErr(err error) bool{
    if err != nil {
	log.Println(err)
	return true
    }
    return false
}
//func saveDevices() {
//    dataFile, err := os.Create("save/devices.gob") 
//    if err != nil {
//	fmt.Println(err)
//	os.Exit(1)
//    }
//    dataEncoder := gob.NewEncoder(dataFile)
//    dataEncoder.Encode(data)
//
//    dataFile.Close()
//
//}
func PopulateAll(){
    insteon.DevList = insteon.GetDevices()
    insteon.SceneList = insteon.GetScenes()
    insteon.RoomList = insteon.GetRooms()
}
func searchString(s string)(dev_type string, id int, loc int){//Going to search Scene -> Device -> Room for now
    for num, scene := range insteon.SceneList {
	s1 := strings.ToLower(scene.SceneName) 
	if s1 == s {
	    dev_type = "scene"
	    id = scene.SceneID
	    loc = num
	    return
	}
    }
    for num, device := range insteon.DevList {
	s1 := strings.ToLower(device.DeviceName) 
	if s1 == s {
	    dev_type = "device"
	    id = device.DeviceID
	    loc = num
	    return
	}
    }
    for num, room := range insteon.RoomList {
	s1 := strings.ToLower(room.RoomName) 
	if s1 == s {
	    dev_type = "room"
	    id = room.RoomID
	    loc = num
	    return
	}
    }
    return
}

func main(){
    //mydev := getDevices()
    //fmt.Println(mydev)
    var n insteon.Command
    PopulateAll()
    search_dev := "media room lamp"
    direction := "on"
    res_type, id, loc := searchString(search_dev)
    switch res_type{
	case "device":
	    n.Device_Id = id
	    n.Level = insteon.DevList[loc].DimLevel/254 * 100
	case "scene":
	    n.Scene_Id = id
	case "room":
	    fmt.Println("Rooms are a pain")
	}

    n.Command = direction
    v := insteon.RunCommand(n)
    fmt.Println(v)


}
