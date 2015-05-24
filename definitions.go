package main

import (

)

type StatusItem struct {
	    RequestsTimedOut string
		MaxOpenConnections string
		MaxOpenRequests string
		OpenConnections string
		TotalConnections string
		OpenRequests string
		TotalRequests string
		Uptime string
		ApiVersion string
	}

type TokenItem struct {
		Access_token string
		Token_type string
		Expires_in string
		Created_at string
	}

type UserItem struct {
		Email string
		Name string
		Password string
		Id string
		Status string
		CreatedAt string
		//locations string
		Locations []struct {
        	Name	string
        	Timezone   string
        	Sensors [] struct {
        		SensorId string
        		Email string
        		InstallCode string
        		SensorType string
        		LocationId string
        		Channels [] struct {
        			SensorId string
        			Channel int
        			ChannelType string
        			Name string
        			Start string
        			Id string
        		}
        		StartTime string
        		Id string
        	}
			EnergyRate float64
        	CreatedAt string
        	Id      string    	
    	}
	}

 type PowerSample struct {
	ConsumptionPower int
	ConsumptionEnergy int
	GenerationPower int
	GenerationEnergy int
	Timestamp string
}

 type ApplianceList []struct { 
 		Label string `json:"label"` 
 		Name string `json:"name"` 
 		Locationid string `json:"locationId"` 
 		Tags []string `json:"tags"` 
 		Createdat string `json:"createdAt"` 
 		ID string `json:"id"` 
 		}
