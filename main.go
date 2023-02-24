package main

import (
	"fmt"
	"net/http"

	"context"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

func main() {
    http.HandleFunc("/", speechToTextHandler)
    http.ListenAndServe(":8080", nil)
}

func speechToTextHandler(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

        // Creates a client.
        client, err := speech.NewClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }
        defer client.Close()

        // The path to the remote audio file to transcribe.
        fileURI := "gs://cloud-samples-data/speech/brooklyn_bridge.raw"

        // Detects speech in the audio file.
        resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
                Config: &speechpb.RecognitionConfig{
                        Encoding:        speechpb.RecognitionConfig_LINEAR16,
                        SampleRateHertz: 16000,
                        LanguageCode:    "en-US",
                },
                Audio: &speechpb.RecognitionAudio{
                        AudioSource: &speechpb.RecognitionAudio_Uri{Uri: fileURI},
                },
        })
        if err != nil {
                log.Fatalf("failed to recognize: %v", err)
        }

        // Prints the results.
        for _, result := range resp.Results {
                for _, alt := range result.Alternatives {
                        fmt.Printf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
                }
        }

}
