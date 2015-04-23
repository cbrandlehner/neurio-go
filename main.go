package main 

import (
	"fmt"
	"os"
	"code.google.com/p/gcfg" // https://code.google.com/p/gcfg/
)

type Config struct {
        Neurio struct {
                NeurioClientID string
                NeurioClientSecret string
        }
}
	
func main() {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
  		fmt.Println("error:", err)
  		pwd, err2 := os.Getwd()
    	if err2 != nil {
	        fmt.Println(err)
        	os.Exit(1)
    	}
    	fmt.Println("You should place the file config.gcfg in this directory:",pwd)
  		os.Exit(1)
	}
	fmt.Println("Values from config.gcfg: Neurioclientid =", cfg.Neurio.NeurioClientID,"Neurioclientsecret =",cfg.Neurio.NeurioClientSecret)
	fmt.Println("Neurio Webservice API Version reported: ", NeurioAPIVersion() )
	
	token, err1 := NeurioAccessToken(cfg.Neurio.NeurioClientID,cfg.Neurio.NeurioClientSecret)
    if err1 != nil {
        fmt.Println("Error!", err1)
        fmt.Println("Neurio2Go failed to get an access token for the clientID and clientSecret provided.")
        var appuri = "https://my.neur.io/#settings/applications/"+cfg.Neurio.NeurioClientID
        fmt.Println("Check if your App exists by opening:", appuri)
        fmt.Println("or create a new app and update your config file: https://my.neur.io/#settings/applications/register")
        os.Exit(1)
    } else {
        fmt.Println("Access Token: ", token )
    }
   
   var NeurioCurrentUserGlobal UserItem 
   ncuerr := NeurioCurrentUserRaw(token, &NeurioCurrentUserGlobal)
   if ncuerr != nil {
        fmt.Println("Error! ", ncuerr)
        os.Exit(1)
    } 
   
   fmt.Println("Your UserID is",NeurioCurrentUserGlobal.Id)
   
   fmt.Println("Your Sensor ID is",NeurioCurrentUserGlobal.Locations[0].Sensors[0].Channels[3].SensorId)
   fmt.Println("Your last power usage sample is",NeurioLastSample(token,NeurioCurrentUserGlobal.Locations[0].Sensors[0].SensorId),"w")
   
}
