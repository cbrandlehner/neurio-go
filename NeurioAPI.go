package main

import (
	"fmt"
	"errors"
	"net/url"
	"github.com/franela/goreq"
)

// returns the version of the Neurio API
func NeurioAPIVersion() string {
	var Statusitem StatusItem	
	res, err := goreq.Request{
	    Method:      "GET",
	    Uri:         "https://api.neur.io/v1/status",
	    Accept: "application/json",
	    ContentType: "application/json",
	    Compression: goreq.Gzip(),
	}.Do()
	if err != nil {
		fmt.Println("Error: ",  err)
	}
	res.Body.FromJsonTo(&Statusitem)
	// fmt.Println("API Version: ", Statusitem.ApiVersion)
	return Statusitem.ApiVersion
}
	
  // returns the oAuth Access Token
	func NeurioAccessToken(clientID string,clientSecret string) (string, error) {
		var Tokenitem TokenItem
		form := url.Values{}
		form.Add("grant_type", "client_credentials")
		form.Add("client_id", clientID)
		form.Add("client_secret", clientSecret)
		res, err := goreq.Request{
		    Method:      "POST",
	    	Uri:         "https://api.neur.io/v1/oauth2/token",
	    	Accept: "application/json",
	    	ContentType: "application/x-www-form-urlencoded",
	    	Body:        form.Encode(),
		}.Do()
		if err != nil {
			fmt.Println("Error: ",  err)
		} 
		res.Body.FromJsonTo(&Tokenitem)
		if Tokenitem.Access_token == "" {
			err = errors.New("Failed to get an access token")
		} else {
			// fmt.Println("Access Token: ", Tokenitem.Access_token)
		}
		return Tokenitem.Access_token , err
	}
	
	// returns the user information
	func NeurioCurrentUserRaw(accessToken string, NCU *UserItem) (error) {
		res, err := goreq.Request{
		    Method:      "GET",
	    	Uri:         "https://api.neur.io/v1/users/current",
	    	Accept: "application/json",
	    	ContentType: "application/json",
	    	Compression: goreq.Gzip(),
		}.WithHeader("Authorization", "Bearer "+accessToken).Do()
		if err != nil {
			fmt.Println("func NeurioCurrentUserRaw, Error: ",  err)
		} else {
			res.Body.FromJsonTo(&NCU)
		}
		return err
	}
	
	// returns the last power sample
	func NeurioLastSample(accessToken string, sensorid string) int {
		var Powersample PowerSample
		res, err := goreq.Request{
		    Method:      "GET",
	    	Uri:       "https://api.neur.io/v1/samples/live/last?sensorId="+sensorid,
	    	Accept: "application/json",
	    	ContentType: "application/json",
	    	Compression: goreq.Gzip(),
		}.WithHeader("Authorization", "Bearer "+accessToken).Do()
		if err != nil {
			fmt.Println("func NeurioLastSample, Error: ",  err)
		}
		xerr := res.Body.FromJsonTo(&Powersample)
		if xerr != nil {
			fmt.Println("func NeurioLastSample FromJsonTo, Error: ",  xerr)
		} else {
			// fmt.Println(Powersample.ConsumptionPower)
		}
		return Powersample.ConsumptionPower
	}

	func myNeurioAppliances(accessToken string, location string, NAL *ApplianceList) (error) {
		res, err := goreq.Request{
		    Method:      "GET",
	    	Uri:       "https://api.neur.io/v1/appliances?locationId="+location,
	    	Accept: "application/json",
	    	ContentType: "application/json",
	    	Compression: goreq.Gzip(),
		}.WithHeader("Authorization", "Bearer "+accessToken).Do()
		if err != nil {
			fmt.Println("func NeurioAppliances, Error: ",  err)
		}
		xerr := res.Body.FromJsonTo(&NAL)
		if xerr != nil {
			fmt.Println("func NeurioLastSample FromJsonTo, Error: ",  xerr)
		} else {
			// TODO add some optional debug code here
		}
		return err
	}
