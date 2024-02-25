# DistroGo

This project is an unofficial API wrapper for DistroKid, developed based on the reversal of the DistroKid iOS application. It aims to provide a programmatic way to interact with some of the functionalities offered by DistroKid. As this is an unofficial tool, please use it responsibly and with consideration of DistroKid's policies and terms of service.

## Current Features

As of now, the wrapper supports the following features:

- **Releases Getters**: Retrieve information about music releases, including title, artist, release date, and more.
- **Tracks Getters**: Fetch details and statistics for individual tracks, such as duration, ISRC, play counts, and other relevant data.

## Planned Features

The project is in its early stages, and there are plans to expand its capabilities. Future updates may include:

- Editing release and track information.
- Uploading new tracks and releases.
- Managing artist profiles and accounts.

## Contributions

Contributions are highly welcome! If you're interested in enhancing this API wrapper, consider contributing in the following areas:

- Expanding the feature set to cover more DistroKid functionalities.
- Improving the existing codebase for better efficiency and reliability.
- Writing documentation and examples to help other developers.

## Usage
To use DistroGo, you need to retrieved your DistroKid API bearer token. You can find it by inspecting the DistroKid iOS application's network requests or by using a proxy to intercept the requests made by the app. Once you have your bearer token, you can use it to create a new DistroKid client and start interacting with the API.
```go
package main

import (
  "fmt"

  "github.com/szerookii/distrogo"
)

func main() {
  distrokid := distrogo.NewDistroKid("your-bearer-token")
  
  releases, err := distrokid.GetReleases()
  
  if err != nil {
    panic(err)
  }

  fmt.Printf("Found %d releases\n", len(releases))
}
```

## Disclaimer

This API wrapper is not officially affiliated with, authorized, maintained, sponsored, or endorsed by DistroKid or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
