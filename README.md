# MEGO

A meeting organizer for Exchange server.

## How it works:
* It connects to Exchage server using [ews](http://github.com/mhewedy/ews)
* The configuration is defined in [app.conf](https://github.com/mhewedy/mego/blob/master/mego-api/app.conf), so you need to place it in the same directory as the binary and configure it properly
* Room list is defined in [rooms.csv](https://github.com/mhewedy/mego/blob/master/mego-api/rooms.csv), so you need to place it in the same directory as the binary and fill the values according to your environment.
* The user credentials entered at login screen is encryped and saved in memory with an encryption key itself is defined only when the server starts. (see [enc.go](https://github.com/mhewedy/mego/blob/master/mego-api/user/enc.go))    
So when a request made to the server with the valid user token, then the token got verified first then the `username` is being exctracted and used to get the encryped password from the in-memory user map and then the password got decrypted by the key defined at the server startup.

## Screenshots

1. Search is done using some basic indexing and searching in-memory scheme.
<img src="https://github.com/mhewedy/mego/raw/master/screenshots/1.png" width="70%">

2. Availabity is returned based on the user input
<img src="https://github.com/mhewedy/mego/raw/master/screenshots/2.png" width="70%">

3. User can send meeting request and add optional attendees as well, where thier availabity time will not checked.
<img src="https://github.com/mhewedy/mego/raw/master/screenshots/3.png" width="70%">

## Download
Download the latest release from [the releases section](https://github.com/mhewedy/mego/releases/latest) (Linux, mac and win releases available) 

