package insteon
import (
    "net/http"
    "net/url"
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
func Refresh_Bearer(refresh_token string)(access_token string, success bool) {
    endpoint := "oauth2/token"
    client := &http.Client{}
    v := url.Values{}
    v.Set("grant_type", "refresh_token") 
    v.Set("refresh_token", refresh_token)
    v.Set("client_id", Client_Id)
    data := v.Encode()
    //log.Println(data)
    //log.Println(endpoint)
    req, _ := http.NewRequest("POST", Insteon_Url + endpoint, bytes.NewBuffer([]byte(data)))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    if err != nil { success = false; return}
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    if err != nil { success = false; return}
    var b BearerResponse
    err = json.Unmarshal(resp_body, &b)
    if err != nil { success = false; return}
    access_token = b.Access_Token
    return 
}

