package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

type UploadFile struct {
	DocumentAttachmentID int     `json:"document_attachment_id"`
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	Description          string  `json:"description"`
}

type Payload struct {
	CIFCode     string       `json:"cifcode"`
	Type        string       `json:"type"`
	Frequency   int          `json:"frequency"`
	Notes       string       `json:"notes"`
	UploadFiles []UploadFile `json:"upload_files"`
}

const (
	apiURL     = "http://localhost:5432/api/v1/hex"
	authToken  = "Bearer ey..." // Replace with valid token
	concurrent = 20
	total      = 20000
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Printf("Running with %d CPU core(s)\n", runtime.NumCPU())

	startTime := time.Now()

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrent)
	var mu sync.Mutex
	visitFreq, callFreq := 1, 1

	for i := 0; i < total; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			var payload Payload
			var typ string
			var freq int

			mu.Lock()
			if i%2 == 0 {
				typ = "visit"
				payload = buildVisitPayload(visitFreq)
				freq = visitFreq
				visitFreq++
			} else {
				typ = "call"
				payload = buildCallPayload(callFreq)
				freq = callFreq
				callFreq++
			}
			mu.Unlock()

			timestamp := time.Now().Format("15:04:05.000")
			fmt.Printf("[%s] Sending %s #%d\n", timestamp, typ, freq)

			sendRequest(payload)
		}(i)
	}

	wg.Wait()

	totalDuration := time.Since(startTime).Seconds()
	fmt.Printf("All requests done in %.2f seconds\n", totalDuration)
}

/*
Code ini hasil referensi dari cara kerja Burp Suite : Turbo Intruder

Only few time got this kind of error
 1. Server kamu overload atau lambat merespons : terjadi saat banyak request paralel dikirim, tapi server tidak cukup cepat membalas (bottleneck di backend).
 2. Terlalu banyak concurrent (goroutine aktif) : misalnya set concurrent = 20, tapi server hanya kuat 8 paralel request.
 3. Timeout terlalu pendek

 Solusi :
 Error: Post "http://localhost:5432/api/v1/hex": due to [context deadline exceeded (Client.Timeout exceeded while awaiting headers)]
 add to 30 * time.Second atau 60 * time.Second

 Jadi analisis dari struktur code yang terdiri dari 3 layer (Controller → Usecase → Repository) menunjukkan bahwa race condition yang menyebabkan duplikasi _id_generator sangat mungkin terjadi, terlepas dari penggunaan pointer

Penggunaan pointer seperti:

func Create (ctx context.Context, request *model.TODOCore, ...)

...tidak menyebabkan race condition selama objek activityRequest tidak diubah secara parallel oleh goroutine lain, yang dalam konteks Turbo Intruder memang tidak terjadi karena setiap HTTP request punya objek sendiri.

Jadi: Pointer bukan sumber race, tapi mekanisme ID generator + insert paralel (use : atomic transaction database) + tidak ada unique constraint yang menyebabkan duplikasi.
*/

func sendRequest(data Payload) {
	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	if resp.StatusCode == 401 {
		fmt.Println("Unauthorized detected — exiting...")
		os.Exit(0)
	}
}

func buildVisitPayload(freq int) Payload {
	return Payload{
		CIFCode:   "CIF001",
		Type:      "visit",
		Frequency: freq,
		Notes:     fmt.Sprintf("Automated Test - visit #%d", freq),
		UploadFiles: []UploadFile{
			{106, 120.12345678, 8.12345678, "file A"},
			{107, 120.12345678, 8.12345678, "file B"},
		},
	}
}

func buildCallPayload(freq int) Payload {
	return Payload{
		CIFCode:   "CIF001",
		Type:      "call",
		Frequency: freq,
		Notes:     fmt.Sprintf("Automated Test - call #%d", freq),
		UploadFiles: []UploadFile{
			{106, 120.12345678, 8.12345678, "file A"},
			{107, 120.12345678, 8.12345678, "file B"},
		},
	}
}
