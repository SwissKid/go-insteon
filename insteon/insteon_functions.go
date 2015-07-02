package insteon
import (
    "strings"
    )
func PopulateAll(){
    DevList = GetDevices()
    SceneList = GetScenes()
    RoomList = GetRooms()
}
//I really need a better way to do this...
func SearchString(s string)(dev_type string, id int, loc int){//Going to search Scene -> Device -> Room for now
    for num, scene := range SceneList {
	s1 := strings.ToLower(scene.SceneName) 
	if s1 == s {
	    dev_type = "scene"
	    id = scene.SceneID
	    loc = num
	    return
	}
    }
    for num, device := range DevList {
	s1 := strings.ToLower(device.DeviceName) 
	if s1 == s {
	    dev_type = "device"
	    id = device.DeviceID
	    loc = num
	    return
	}
    }
    for num, room := range RoomList {
	s1 := strings.ToLower(room.RoomName) 
	if s1 == s {
	    dev_type = "room"
	    id = room.RoomID
	    loc = num
	    return
	}
    }
    dev_type = "not_found"
    id = 0
    loc = -1
    return
}
func DeviceSearchID(dev_id int)(dev Device, success bool){
    for _, device := range DevList {
	if device.DeviceID == dev_id {
	    dev = device
	    success = true
	    return
	}
    }
    success = false
    return
}
func SceneSearch(s string)(scene_return Scene, success bool){
    for _, scene := range SceneList {
	s1 := strings.ToLower(scene.SceneName) 
	if s1 == s {
	    scene_return = scene
	    success = true
	    return
	}
    }
    success = false
    return
}
func DeviceSearch(s string)(dev Device, success bool){
    for _, device := range DevList {
	s1 := strings.ToLower(device.DeviceName) 
	if s1 == s {
	    dev = device
	    success = true
	    return
	}
    }
    success = false
    return
}
func DeviceOn(device_id int)(success bool){
    var n Command
    device, success := DeviceSearchID(device_id)
    if success != true { return false}
    n.Command = "on"
    n.Level = device.DimLevel
    n.Device_Id = device_id
    result := RunCommand(n)
    if result.Status != "success" {
	success = false
    } else {
	success = true
    }
    return
}
