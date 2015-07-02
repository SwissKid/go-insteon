package insteon
import (
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "strconv"
    )
func checkErr(err error) bool{
    if err != nil {
	log.Println(err)
	return true
    }
    return false
}
func Get(endpoint string, all bool) (resp_body []byte){
    client := &http.Client{}
    if all {
	endpoint += "?properties=all"
    }
    req, _ := http.NewRequest("GET", Insteon_Url + endpoint, nil)
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authentication", "APIKey " + Client_Id)
    req.Header.Add("Authorization", "Bearer " + Access_Token)
    resp, err := client.Do(req)
    if err != nil {
	log.Println("Errored when sending request to the server")
	    return
    }
    defer resp.Body.Close()
    resp_body, _ = ioutil.ReadAll(resp.Body)
    return
}
func Post(endpoint string, data []byte) (resp_body []byte) {
    client := &http.Client{}
    req, _ := http.NewRequest("POST", Insteon_Url + endpoint, bytes.NewBuffer(data))
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authentication", "APIKey " + Client_Id)
    req.Header.Add("Authorization", "Bearer " + Access_Token)
    resp, err := client.Do(req)
    if err != nil {
	log.Println("Errored when sending request to the server")
	    return
    }
    defer resp.Body.Close()
    resp_body, _ = ioutil.ReadAll(resp.Body)
    return 
}
func CommandFollowup(id int) (response CommandResponse){
    endpoint := "commands/" + strconv.Itoa(id)
    body := Get(endpoint,false)
    err := json.Unmarshal(body, &response)
    if checkErr(err) {return}
    for response.Status == "pending" {
	body := Get(endpoint,false)
	err = json.Unmarshal(body, &response)
	if checkErr(err) {return}
    }
    return 
}
func RunCommand(data Command)(response CommandResponse){
    var resp CommandResponse
    cmd, err := json.Marshal(data)
    if checkErr(err) {return}
    nextString := Post("commands", cmd)
    err = json.Unmarshal(nextString, &resp)
    if checkErr(err) {return}
    response = CommandFollowup(resp.Id)
    return
}





func GetScenes() []Scene {
    var j SceneResponse
    var sceneList []Scene
    resp_body := Get("scenes", true)
    json.Unmarshal(resp_body, &j)
    sceneList = j.SceneList
    return sceneList
}
func GetRooms() []Room {
    var j RoomResponse
    var roomList []Room
    resp_body := Get("rooms", true)
    json.Unmarshal(resp_body, &j)
    roomList = j.RoomList
    return roomList
}
func GetDevices() []Device {
    var j DeviceResponse
    var devList []Device
    resp_body := Get("devices", true)
    json.Unmarshal(resp_body, &j)
    devList = j.DeviceList
    return devList
}
