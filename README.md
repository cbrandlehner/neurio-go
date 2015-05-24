# neurio-go
Neurio energy sensor client in golang

This is the unofficial Go(lang) library for the [Neurio](http://www.neur.io) sensor [API](https://api.neur.io/docs/)

The library currently supports:

    OAuth 2 authentication (including token request) – /v1/oauth2
    reading Neurio API version and status - /v1/status
    Reading user information - /v1/users/current
    Consumption samples (last recent sample) – /v1/samples
    Printing list of appliances - /v1/appliances
    
<h1>Request API Access Key</h1>
Create a new APP on the [My Neurio cloud](https://my.neur.io/#settings/applications/register) to get you Client ID and Client Secret.

<h1>Configure</h1>
Edit the file config.gcfg and replace the placeholder with your Client ID and Client Secret.

