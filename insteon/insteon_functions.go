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
    return
}
func DeviceSearchID(dev_id int)(dev Device){
    for _, device := range DevList {
	if device.DeviceID == dev_id {
	    dev = device
	    return
	}
    }
    return
}
func DeviceSearch(s string)(dev Device){
    for _, device := range DevList {
	s1 := strings.ToLower(device.DeviceName) 
	if s1 == s {
	    dev = device
	    return
	}
    }
    return
}
func DeviceOn(device_id int)(success bool){
    var n Command
    device := DeviceSearchID(device_id)
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
