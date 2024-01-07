package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func browse(id string) browseStruct {
	for {
		data := fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"SA","remoteHost":"95.184.79.221","deviceMake":"","deviceModel":"","visitorData":"CgtWWUJpMmVFSWpoUSjJ7tyrBjIKCgJTQRIEGgAgNg%%3D%%3D","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20231208.01.00","osName":"Windows","osVersion":"10.0"}},"browseId":"%s","params":"EgZ2aWRlb3PyBgQKAjoA"}`, id)
		req, err := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/browse?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false", strings.NewReader(data))
		if err == nil {
			req.Header.Set("Host", "www.youtube.com")
			req.Header.Set("accept", "*/*")
			req.Header.Set("accept-language", "en-US,en;q=0.9")
			req.Header.Set("content-type", "application/json")
			req.Header.Set("origin", "https://www.youtube.com")
			req.Header.Set("referer", "https://www.youtube.com/results?search_query=searchKey")
			req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
			req.Header.Set("sec-ch-ua-arch", "\"x86\"")
			req.Header.Set("sec-ch-ua-bitness", "\"64\"")
			req.Header.Set("sec-ch-ua-full-version", "\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-full-version-list", "\"Not_A Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"120.0.6099.71\", \"Google Chrome\";v=\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-model", "\"\"")
			req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
			req.Header.Set("sec-ch-ua-platform-version", "\"10.0.0\"")
			req.Header.Set("sec-ch-ua-wow64", "?0")
			req.Header.Set("sec-fetch-dest", "empty")
			req.Header.Set("sec-fetch-mode", "same-origin")
			req.Header.Set("sec-fetch-site", "same-origin")
			req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
			req.Header.Set("x-goog-visitor-id", "CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%3D")
			req.Header.Set("x-youtube-bootstrap-logged-in", "false")
			req.Header.Set("x-youtube-client-name", "1")
			req.Header.Set("x-youtube-client-version", "2.20231208.01.00")
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					a := browseStruct{}
					err := json.Unmarshal(body, &a)
					if err == nil {
						return a
					}
					fmt.Println("Error when parsing json.Unmarshal(body, &a): ", err)
				}
			}
		}
	}
}
func browseByCont(id, Cont string) browseContinuationStruct {
	for {
		data := fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"SA","remoteHost":"95.184.79.221","deviceMake":"","deviceModel":"","visitorData":"CgtWWUJpMmVFSWpoUSjJ7tyrBjIKCgJTQRIEGgAgNg%%3D%%3D","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20231208.01.00","osName":"Windows","osVersion":"10.0"}},"browseId":"%s","params":"EgZ2aWRlb3PyBgQKAjoA", "continuation":"%s"}`, id, Cont)
		req, err := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/browse?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false", strings.NewReader(data))
		if err == nil {
			req.Header.Set("Host", "www.youtube.com")
			req.Header.Set("accept", "*/*")
			req.Header.Set("accept-language", "en-US,en;q=0.9")
			req.Header.Set("content-type", "application/json")
			req.Header.Set("origin", "https://www.youtube.com")
			req.Header.Set("referer", "https://www.youtube.com/results?search_query=searchKey")
			req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
			req.Header.Set("sec-ch-ua-arch", "\"x86\"")
			req.Header.Set("sec-ch-ua-bitness", "\"64\"")
			req.Header.Set("sec-ch-ua-full-version", "\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-full-version-list", "\"Not_A Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"120.0.6099.71\", \"Google Chrome\";v=\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-model", "\"\"")
			req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
			req.Header.Set("sec-ch-ua-platform-version", "\"10.0.0\"")
			req.Header.Set("sec-ch-ua-wow64", "?0")
			req.Header.Set("sec-fetch-dest", "empty")
			req.Header.Set("sec-fetch-mode", "same-origin")
			req.Header.Set("sec-fetch-site", "same-origin")
			req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
			req.Header.Set("x-goog-visitor-id", "CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%3D")
			req.Header.Set("x-youtube-bootstrap-logged-in", "false")
			req.Header.Set("x-youtube-client-name", "1")
			req.Header.Set("x-youtube-client-version", "2.20231208.01.00")
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					a := browseContinuationStruct{}
					//fmt.Println(string(body))
					err := json.Unmarshal(body, &a)
					if err == nil {
						return a
					}
					fmt.Println("Error when parsing json.Unmarshal(body, &a): ", err)
				}
			}
		}
	}
}
