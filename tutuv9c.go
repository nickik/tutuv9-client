package main

import (
    "net/http"
    "time"
    "log"
    "os"

    "path/filepath"

    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"

    "github.com/kardianos/osext"
)

var (
    filePath = determineWorkingDirectory()

    soundFile = filepath.Join(filePath, "sound.mp3")
)

func main() {

    for true {

        iterationStart := time.Now()

        resp, err := http.Get("https://pacific-falls-90446.herokuapp.com/check")
        // Process response
        if err != nil {
            log.Fatal(err)
        }

        if resp.StatusCode == 200 {
            play(soundFile)
	    fmt.Printf("StatusCode == 200 -> play Sound\n")
        } else {
            if resp.StatusCode == 204 {

            } else {

            }
        }

        elapsed := time.Since(iterationStart)

        time.Sleep( (5 * time.Second) - elapsed )
    }
}

func determineWorkingDirectory() string {
    // Get the absolute path this executable is located in.
    executablePath, err := osext.ExecutableFolder()
    if err != nil {
        log.Fatal("Error: Couldn't determine working directory: " + err.Error())
    }
    // Set the working directory to the path the executable is located in.
    os.Chdir(executablePath)
    return ""
}

func play(path string) {
	// Open first sample File
	f, err := os.Open(path)

	// Check for errors when opening the file
	if err != nil {
		log.Fatal(err)
	}

	// Decode the .mp3 File, if you have a .wav file, use wav.Decode(f)
	s, format, _ := mp3.Decode(f)

	// Init the Speaker with the SampleRate of the format and a buffer size of 1/10s
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Channel, which will signal the end of the playback.
	playing := make(chan struct{})

	// Now we Play our Streamer on the Speaker
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		// Callback after the stream Ends
		close(playing)
	})))
	<-playing
}
