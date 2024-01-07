package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func player(VideoID string) playerStruct {
	for {
		data := fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"SA","remoteHost":"95.184.79.221","deviceMake":"","deviceModel":"","visitorData":"CgtWWUJpMmVFSWpoUSjJ7tyrBjIKCgJTQRIEGgAgNg%%3D%%3D","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20231208.01.00","osName":"Windows","osVersion":"10.0"}},"videoId":"%s","params":"","playbackContext":{"contentPlaybackContext":{"currentUrl":"/watch?v=2btwjqpX-bg&pp=YAHIAQE%%3D","vis":5,"splay":false,"autoCaptionsDefaultOn":false,"autonavState":"STATE_OFF","html5Preference":"HTML5_PREF_WANTS","signatureTimestamp":19697,"autoplay":true,"autonav":true,"referer":"https://www.youtube.com/","lactMilliseconds":"-1","watchAmbientModeContext":{"hasShownAmbientMode":true,"hasToggledOffAmbientMode":true}}},"racyCheckOk":false,"contentCheckOk":false}`, VideoID)
		req, err := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/player?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false", strings.NewReader(data))
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
					a := playerStruct{}
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

type playerStruct struct {
	ResponseContext struct {
		ServiceTrackingParams []struct {
			Service string `json:"service"`
			Params  []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"params"`
		} `json:"serviceTrackingParams"`
		MaxAgeSeconds             int `json:"maxAgeSeconds"`
		MainAppWebResponseContext struct {
			LoggedOut     bool   `json:"loggedOut"`
			TrackingParam string `json:"trackingParam"`
		} `json:"mainAppWebResponseContext"`
		WebResponseContextExtensionData struct {
			HasDecorated bool `json:"hasDecorated"`
		} `json:"webResponseContextExtensionData"`
	} `json:"responseContext"`
	PlayabilityStatus struct {
		Status          string `json:"status"`
		PlayableInEmbed bool   `json:"playableInEmbed"`
		Miniplayer      struct {
			MiniplayerRenderer struct {
				PlaybackMode string `json:"playbackMode"`
			} `json:"miniplayerRenderer"`
		} `json:"miniplayer"`
		ContextParams string `json:"contextParams"`
	} `json:"playabilityStatus"`
	StreamingData struct {
		ExpiresInSeconds string `json:"expiresInSeconds"`
		Formats          []struct {
			Itag             int    `json:"itag"`
			MimeType         string `json:"mimeType"`
			Bitrate          int    `json:"bitrate"`
			Width            int    `json:"width"`
			Height           int    `json:"height"`
			LastModified     string `json:"lastModified"`
			Quality          string `json:"quality"`
			Fps              int    `json:"fps"`
			QualityLabel     string `json:"qualityLabel"`
			ProjectionType   string `json:"projectionType"`
			AudioQuality     string `json:"audioQuality"`
			ApproxDurationMs string `json:"approxDurationMs"`
			AudioSampleRate  string `json:"audioSampleRate"`
			AudioChannels    int    `json:"audioChannels"`
			SignatureCipher  string `json:"signatureCipher"`
		} `json:"formats"`
		AdaptiveFormats []struct {
			Itag      int    `json:"itag"`
			MimeType  string `json:"mimeType"`
			Bitrate   int    `json:"bitrate"`
			Width     int    `json:"width,omitempty"`
			Height    int    `json:"height,omitempty"`
			InitRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"initRange"`
			IndexRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"indexRange"`
			LastModified     string `json:"lastModified"`
			ContentLength    string `json:"contentLength"`
			Quality          string `json:"quality"`
			Fps              int    `json:"fps,omitempty"`
			QualityLabel     string `json:"qualityLabel,omitempty"`
			ProjectionType   string `json:"projectionType"`
			AverageBitrate   int    `json:"averageBitrate"`
			ApproxDurationMs string `json:"approxDurationMs"`
			SignatureCipher  string `json:"signatureCipher"`
			ColorInfo        struct {
				Primaries               string `json:"primaries"`
				TransferCharacteristics string `json:"transferCharacteristics"`
				MatrixCoefficients      string `json:"matrixCoefficients"`
			} `json:"colorInfo,omitempty"`
			HighReplication bool    `json:"highReplication,omitempty"`
			AudioQuality    string  `json:"audioQuality,omitempty"`
			AudioSampleRate string  `json:"audioSampleRate,omitempty"`
			AudioChannels   int     `json:"audioChannels,omitempty"`
			LoudnessDb      float64 `json:"loudnessDb,omitempty"`
		} `json:"adaptiveFormats"`
	} `json:"streamingData"`
	PlayerAds []struct {
		PlayerLegacyDesktopWatchAdsRenderer struct {
			PlayerAdParams struct {
				ShowContentThumbnail bool   `json:"showContentThumbnail"`
				EnabledEngageTypes   string `json:"enabledEngageTypes"`
			} `json:"playerAdParams"`
			GutParams struct {
				Tag string `json:"tag"`
			} `json:"gutParams"`
			ShowCompanion bool `json:"showCompanion"`
			ShowInstream  bool `json:"showInstream"`
			UseGut        bool `json:"useGut"`
		} `json:"playerLegacyDesktopWatchAdsRenderer"`
	} `json:"playerAds"`
	PlaybackTracking struct {
		VideostatsPlaybackURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsPlaybackUrl"`
		VideostatsDelayplayURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"videostatsDelayplayUrl"`
		VideostatsWatchtimeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"videostatsWatchtimeUrl"`
		PtrackingURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"ptrackingUrl"`
		QoeURL struct {
			BaseURL string `json:"baseUrl"`
		} `json:"qoeUrl"`
		AtrURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"atrUrl"`
		VideostatsScheduledFlushWalltimeSeconds []int `json:"videostatsScheduledFlushWalltimeSeconds"`
		VideostatsDefaultFlushIntervalSeconds   int   `json:"videostatsDefaultFlushIntervalSeconds"`
		YoutubeRemarketingURL                   struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"youtubeRemarketingUrl"`
		GoogleRemarketingURL struct {
			BaseURL                 string `json:"baseUrl"`
			ElapsedMediaTimeSeconds int    `json:"elapsedMediaTimeSeconds"`
		} `json:"googleRemarketingUrl"`
	} `json:"playbackTracking"`
	VideoDetails struct {
		VideoID          string   `json:"videoId"`
		Title            string   `json:"title"`
		LengthSeconds    string   `json:"lengthSeconds"`
		Keywords         []string `json:"keywords"`
		ChannelID        string   `json:"channelId"`
		IsOwnerViewing   bool     `json:"isOwnerViewing"`
		ShortDescription string   `json:"shortDescription"`
		IsCrawlable      bool     `json:"isCrawlable"`
		Thumbnail        struct {
			Thumbnails []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		AllowRatings      bool   `json:"allowRatings"`
		ViewCount         string `json:"viewCount"`
		Author            string `json:"author"`
		IsPrivate         bool   `json:"isPrivate"`
		IsUnpluggedCorpus bool   `json:"isUnpluggedCorpus"`
		IsLiveContent     bool   `json:"isLiveContent"`
	} `json:"videoDetails"`
	Annotations []struct {
		PlayerAnnotationsExpandedRenderer struct {
			FeaturedChannel struct {
				StartTimeMs string `json:"startTimeMs"`
				EndTimeMs   string `json:"endTimeMs"`
				Watermark   struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"watermark"`
				TrackingParams     string `json:"trackingParams"`
				NavigationEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
							APIURL      string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					BrowseEndpoint struct {
						BrowseID string `json:"browseId"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
				ChannelName     string `json:"channelName"`
				SubscribeButton struct {
					SubscribeButtonRenderer struct {
						ButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"buttonText"`
						Subscribed           bool   `json:"subscribed"`
						Enabled              bool   `json:"enabled"`
						Type                 string `json:"type"`
						ChannelID            string `json:"channelId"`
						ShowPreferences      bool   `json:"showPreferences"`
						SubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"subscribedButtonText"`
						UnsubscribedButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribedButtonText"`
						TrackingParams        string `json:"trackingParams"`
						UnsubscribeButtonText struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"unsubscribeButtonText"`
						ServiceEndpoints []struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									SendPost bool   `json:"sendPost"`
									APIURL   string `json:"apiUrl"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SubscribeEndpoint struct {
								ChannelIds []string `json:"channelIds"`
								Params     string   `json:"params"`
							} `json:"subscribeEndpoint,omitempty"`
							SignalServiceEndpoint struct {
								Signal  string `json:"signal"`
								Actions []struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									OpenPopupAction     struct {
										Popup struct {
											ConfirmDialogRenderer struct {
												TrackingParams string `json:"trackingParams"`
												DialogMessages []struct {
													Runs []struct {
														Text string `json:"text"`
													} `json:"runs"`
												} `json:"dialogMessages"`
												ConfirmButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														ServiceEndpoint struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	SendPost bool   `json:"sendPost"`
																	APIURL   string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															UnsubscribeEndpoint struct {
																ChannelIds []string `json:"channelIds"`
																Params     string   `json:"params"`
															} `json:"unsubscribeEndpoint"`
														} `json:"serviceEndpoint"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"confirmButton"`
												CancelButton struct {
													ButtonRenderer struct {
														Style      string `json:"style"`
														Size       string `json:"size"`
														IsDisabled bool   `json:"isDisabled"`
														Text       struct {
															Runs []struct {
																Text string `json:"text"`
															} `json:"runs"`
														} `json:"text"`
														Accessibility struct {
															Label string `json:"label"`
														} `json:"accessibility"`
														TrackingParams string `json:"trackingParams"`
													} `json:"buttonRenderer"`
												} `json:"cancelButton"`
												PrimaryIsCancel bool `json:"primaryIsCancel"`
											} `json:"confirmDialogRenderer"`
										} `json:"popup"`
										PopupType string `json:"popupType"`
									} `json:"openPopupAction"`
								} `json:"actions"`
							} `json:"signalServiceEndpoint,omitempty"`
						} `json:"serviceEndpoints"`
						SubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"subscribeAccessibility"`
						UnsubscribeAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"unsubscribeAccessibility"`
						SignInEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL string `json:"url"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
						} `json:"signInEndpoint"`
					} `json:"subscribeButtonRenderer"`
				} `json:"subscribeButton"`
			} `json:"featuredChannel"`
			AllowSwipeDismiss bool   `json:"allowSwipeDismiss"`
			AnnotationID      string `json:"annotationId"`
		} `json:"playerAnnotationsExpandedRenderer"`
	} `json:"annotations"`
	PlayerConfig struct {
		AudioConfig struct {
			LoudnessDb              float64 `json:"loudnessDb"`
			PerceptualLoudnessDb    float64 `json:"perceptualLoudnessDb"`
			EnablePerFormatLoudness bool    `json:"enablePerFormatLoudness"`
		} `json:"audioConfig"`
		StreamSelectionConfig struct {
			MaxBitrate string `json:"maxBitrate"`
		} `json:"streamSelectionConfig"`
		MediaCommonConfig struct {
			DynamicReadaheadConfig struct {
				MaxReadAheadMediaTimeMs int `json:"maxReadAheadMediaTimeMs"`
				MinReadAheadMediaTimeMs int `json:"minReadAheadMediaTimeMs"`
				ReadAheadGrowthRateMs   int `json:"readAheadGrowthRateMs"`
			} `json:"dynamicReadaheadConfig"`
		} `json:"mediaCommonConfig"`
		WebPlayerConfig struct {
			UseCobaltTvosDash       bool `json:"useCobaltTvosDash"`
			WebPlayerActionsPorting struct {
				GetSharePanelCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WebPlayerShareEntityServiceEndpoint struct {
						SerializedShareEntity string `json:"serializedShareEntity"`
					} `json:"webPlayerShareEntityServiceEndpoint"`
				} `json:"getSharePanelCommand"`
				SubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					SubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"subscribeEndpoint"`
				} `json:"subscribeCommand"`
				UnsubscribeCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					UnsubscribeEndpoint struct {
						ChannelIds []string `json:"channelIds"`
						Params     string   `json:"params"`
					} `json:"unsubscribeEndpoint"`
				} `json:"unsubscribeCommand"`
				AddToWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							AddedVideoID string `json:"addedVideoId"`
							Action       string `json:"action"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"addToWatchLaterCommand"`
				RemoveFromWatchLaterCommand struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							SendPost bool   `json:"sendPost"`
							APIURL   string `json:"apiUrl"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					PlaylistEditEndpoint struct {
						PlaylistID string `json:"playlistId"`
						Actions    []struct {
							Action         string `json:"action"`
							RemovedVideoID string `json:"removedVideoId"`
						} `json:"actions"`
					} `json:"playlistEditEndpoint"`
				} `json:"removeFromWatchLaterCommand"`
			} `json:"webPlayerActionsPorting"`
		} `json:"webPlayerConfig"`
	} `json:"playerConfig"`
	Storyboards struct {
		PlayerStoryboardSpecRenderer struct {
			Spec                           string `json:"spec"`
			RecommendedLevel               int    `json:"recommendedLevel"`
			HighResolutionRecommendedLevel int    `json:"highResolutionRecommendedLevel"`
		} `json:"playerStoryboardSpecRenderer"`
	} `json:"storyboards"`
	Microformat struct {
		PlayerMicroformatRenderer struct {
			Thumbnail struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"thumbnail"`
			Embed struct {
				IframeURL string `json:"iframeUrl"`
				Width     int    `json:"width"`
				Height    int    `json:"height"`
			} `json:"embed"`
			Title struct {
				SimpleText string `json:"simpleText"`
			} `json:"title"`
			Description struct {
				SimpleText string `json:"simpleText"`
			} `json:"description"`
			LengthSeconds      string   `json:"lengthSeconds"`
			OwnerProfileURL    string   `json:"ownerProfileUrl"`
			ExternalChannelID  string   `json:"externalChannelId"`
			IsFamilySafe       bool     `json:"isFamilySafe"`
			AvailableCountries []string `json:"availableCountries"`
			IsUnlisted         bool     `json:"isUnlisted"`
			HasYpcMetadata     bool     `json:"hasYpcMetadata"`
			ViewCount          string   `json:"viewCount"`
			Category           string   `json:"category"`
			PublishDate        string   `json:"publishDate"`
			OwnerChannelName   string   `json:"ownerChannelName"`
			UploadDate         string   `json:"uploadDate"`
		} `json:"playerMicroformatRenderer"`
	} `json:"microformat"`
	Cards struct {
		CardCollectionRenderer struct {
			Cards []struct {
				CardRenderer struct {
					Teaser struct {
						SimpleCardTeaserRenderer struct {
							Message struct {
								SimpleText string `json:"simpleText"`
							} `json:"message"`
							TrackingParams       string `json:"trackingParams"`
							Prominent            bool   `json:"prominent"`
							LogVisibilityUpdates bool   `json:"logVisibilityUpdates"`
							OnTapCommand         struct {
								ClickTrackingParams                   string `json:"clickTrackingParams"`
								ChangeEngagementPanelVisibilityAction struct {
									TargetID   string `json:"targetId"`
									Visibility string `json:"visibility"`
								} `json:"changeEngagementPanelVisibilityAction"`
							} `json:"onTapCommand"`
						} `json:"simpleCardTeaserRenderer"`
					} `json:"teaser"`
					CueRanges []struct {
						StartCardActiveMs string `json:"startCardActiveMs"`
						EndCardActiveMs   string `json:"endCardActiveMs"`
						TeaserDurationMs  string `json:"teaserDurationMs"`
						IconAfterTeaserMs string `json:"iconAfterTeaserMs"`
					} `json:"cueRanges"`
					TrackingParams string `json:"trackingParams"`
				} `json:"cardRenderer"`
			} `json:"cards"`
			HeaderText struct {
				SimpleText string `json:"simpleText"`
			} `json:"headerText"`
			Icon struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"icon"`
			CloseButton struct {
				InfoCardIconRenderer struct {
					TrackingParams string `json:"trackingParams"`
				} `json:"infoCardIconRenderer"`
			} `json:"closeButton"`
			TrackingParams           string `json:"trackingParams"`
			AllowTeaserDismiss       bool   `json:"allowTeaserDismiss"`
			LogIconVisibilityUpdates bool   `json:"logIconVisibilityUpdates"`
		} `json:"cardCollectionRenderer"`
	} `json:"cards"`
	TrackingParams string `json:"trackingParams"`
	Attestation    struct {
		PlayerAttestationRenderer struct {
			Challenge    string `json:"challenge"`
			BotguardData struct {
				Program            string `json:"program"`
				InterpreterSafeURL struct {
					PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
				} `json:"interpreterSafeUrl"`
				ServerEnvironment int `json:"serverEnvironment"`
			} `json:"botguardData"`
		} `json:"playerAttestationRenderer"`
	} `json:"attestation"`
	Messages []struct {
		MealbarPromoRenderer struct {
			Icon struct {
				Thumbnails []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"thumbnails"`
			} `json:"icon"`
			MessageTexts []struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTexts"`
			ActionButton struct {
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams    string `json:"clickTrackingParams"`
						CommandExecutorCommand struct {
							Commands []struct {
								ClickTrackingParams string `json:"clickTrackingParams,omitempty"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL         string `json:"url"`
										WebPageType string `json:"webPageType"`
										RootVe      int    `json:"rootVe"`
										APIURL      string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								BrowseEndpoint struct {
									BrowseID string `json:"browseId"`
									Params   string `json:"params"`
								} `json:"browseEndpoint,omitempty"`
								FeedbackEndpoint struct {
									FeedbackToken string `json:"feedbackToken"`
									UIActions     struct {
										HideEnclosingContainer bool `json:"hideEnclosingContainer"`
									} `json:"uiActions"`
								} `json:"feedbackEndpoint,omitempty"`
							} `json:"commands"`
						} `json:"commandExecutorCommand"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"actionButton"`
			DismissButton struct {
				ButtonRenderer struct {
					Style string `json:"style"`
					Size  string `json:"size"`
					Text  struct {
						Runs []struct {
							Text string `json:"text"`
						} `json:"runs"`
					} `json:"text"`
					TrackingParams string `json:"trackingParams"`
					Command        struct {
						ClickTrackingParams    string `json:"clickTrackingParams"`
						CommandExecutorCommand struct {
							Commands []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								FeedbackEndpoint struct {
									FeedbackToken string `json:"feedbackToken"`
									UIActions     struct {
										HideEnclosingContainer bool `json:"hideEnclosingContainer"`
									} `json:"uiActions"`
								} `json:"feedbackEndpoint"`
							} `json:"commands"`
						} `json:"commandExecutorCommand"`
					} `json:"command"`
				} `json:"buttonRenderer"`
			} `json:"dismissButton"`
			TriggerCondition    string `json:"triggerCondition"`
			Style               string `json:"style"`
			TrackingParams      string `json:"trackingParams"`
			ImpressionEndpoints []struct {
				ClickTrackingParams string `json:"clickTrackingParams"`
				CommandMetadata     struct {
					WebCommandMetadata struct {
						SendPost bool   `json:"sendPost"`
						APIURL   string `json:"apiUrl"`
					} `json:"webCommandMetadata"`
				} `json:"commandMetadata"`
				FeedbackEndpoint struct {
					FeedbackToken string `json:"feedbackToken"`
					UIActions     struct {
						HideEnclosingContainer bool `json:"hideEnclosingContainer"`
					} `json:"uiActions"`
				} `json:"feedbackEndpoint"`
			} `json:"impressionEndpoints"`
			IsVisible    bool `json:"isVisible"`
			MessageTitle struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"messageTitle"`
			EnableSharedFeatureForImpressionHandling bool `json:"enableSharedFeatureForImpressionHandling"`
		} `json:"mealbarPromoRenderer"`
	} `json:"messages"`
	Endscreen struct {
		EndscreenRenderer struct {
			Elements []struct {
				EndscreenElementRenderer struct {
					Style string `json:"style"`
					Image struct {
						Thumbnails []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"thumbnails"`
					} `json:"image"`
					Icon struct {
						Thumbnails []struct {
							URL string `json:"url"`
						} `json:"thumbnails"`
					} `json:"icon"`
					Left        float64     `json:"left"`
					Width       float64     `json:"width"`
					Top         float64     `json:"top"`
					AspectRatio interface{} `json:"aspectRatio"`
					StartMs     string      `json:"startMs"`
					EndMs       string      `json:"endMs"`
					Title       struct {
						Accessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
						SimpleText string `json:"simpleText"`
					} `json:"title"`
					Metadata struct {
						SimpleText string `json:"simpleText"`
					} `json:"metadata"`
					CallToAction struct {
						SimpleText string `json:"simpleText"`
					} `json:"callToAction"`
					Dismiss struct {
						SimpleText string `json:"simpleText"`
					} `json:"dismiss"`
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID string `json:"browseId"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					HovercardButton struct {
						SubscribeButtonRenderer struct {
							ButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"buttonText"`
							Subscribed           bool   `json:"subscribed"`
							Enabled              bool   `json:"enabled"`
							Type                 string `json:"type"`
							ChannelID            string `json:"channelId"`
							ShowPreferences      bool   `json:"showPreferences"`
							SubscribedButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"subscribedButtonText"`
							UnsubscribedButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"unsubscribedButtonText"`
							TrackingParams        string `json:"trackingParams"`
							UnsubscribeButtonText struct {
								Runs []struct {
									Text string `json:"text"`
								} `json:"runs"`
							} `json:"unsubscribeButtonText"`
							ServiceEndpoints []struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								SubscribeEndpoint struct {
									ChannelIds []string `json:"channelIds"`
									Params     string   `json:"params"`
								} `json:"subscribeEndpoint,omitempty"`
								SignalServiceEndpoint struct {
									Signal  string `json:"signal"`
									Actions []struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										OpenPopupAction     struct {
											Popup struct {
												ConfirmDialogRenderer struct {
													TrackingParams string `json:"trackingParams"`
													DialogMessages []struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"dialogMessages"`
													ConfirmButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															ServiceEndpoint struct {
																ClickTrackingParams string `json:"clickTrackingParams"`
																CommandMetadata     struct {
																	WebCommandMetadata struct {
																		SendPost bool   `json:"sendPost"`
																		APIURL   string `json:"apiUrl"`
																	} `json:"webCommandMetadata"`
																} `json:"commandMetadata"`
																UnsubscribeEndpoint struct {
																	ChannelIds []string `json:"channelIds"`
																	Params     string   `json:"params"`
																} `json:"unsubscribeEndpoint"`
															} `json:"serviceEndpoint"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"confirmButton"`
													CancelButton struct {
														ButtonRenderer struct {
															Style      string `json:"style"`
															Size       string `json:"size"`
															IsDisabled bool   `json:"isDisabled"`
															Text       struct {
																Runs []struct {
																	Text string `json:"text"`
																} `json:"runs"`
															} `json:"text"`
															Accessibility struct {
																Label string `json:"label"`
															} `json:"accessibility"`
															TrackingParams string `json:"trackingParams"`
														} `json:"buttonRenderer"`
													} `json:"cancelButton"`
													PrimaryIsCancel bool `json:"primaryIsCancel"`
												} `json:"confirmDialogRenderer"`
											} `json:"popup"`
											PopupType string `json:"popupType"`
										} `json:"openPopupAction"`
									} `json:"actions"`
								} `json:"signalServiceEndpoint,omitempty"`
							} `json:"serviceEndpoints"`
							SubscribeAccessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"subscribeAccessibility"`
							UnsubscribeAccessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"unsubscribeAccessibility"`
							SignInEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										URL string `json:"url"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
							} `json:"signInEndpoint"`
						} `json:"subscribeButtonRenderer"`
					} `json:"hovercardButton"`
					TrackingParams string `json:"trackingParams"`
					IsSubscribe    bool   `json:"isSubscribe"`
					ID             string `json:"id"`
				} `json:"endscreenElementRenderer"`
			} `json:"elements"`
			StartMs        string `json:"startMs"`
			TrackingParams string `json:"trackingParams"`
		} `json:"endscreenRenderer"`
	} `json:"endscreen"`
	AdPlacements []struct {
		AdPlacementRenderer struct {
			Config struct {
				AdPlacementConfig struct {
					Kind         string `json:"kind"`
					AdTimeOffset struct {
						OffsetStartMilliseconds string `json:"offsetStartMilliseconds"`
						OffsetEndMilliseconds   string `json:"offsetEndMilliseconds"`
					} `json:"adTimeOffset"`
					HideCueRangeMarker bool `json:"hideCueRangeMarker"`
				} `json:"adPlacementConfig"`
			} `json:"config"`
			Renderer struct {
				InstreamVideoAdRenderer struct {
					PlayerOverlay struct {
						InstreamAdPlayerOverlayRenderer struct {
							SkipOrPreviewRenderer struct {
								SkipAdRenderer struct {
									PreskipRenderer struct {
										AdPreviewRenderer struct {
											Thumbnail struct {
												Thumbnail struct {
													Thumbnails []struct {
														URL    string `json:"url"`
														Width  int    `json:"width"`
														Height int    `json:"height"`
													} `json:"thumbnails"`
												} `json:"thumbnail"`
												TrackingParams string `json:"trackingParams"`
											} `json:"thumbnail"`
											TrackingParams     string `json:"trackingParams"`
											TemplatedCountdown struct {
												TemplatedAdText struct {
													Text           string `json:"text"`
													IsTemplated    bool   `json:"isTemplated"`
													TrackingParams string `json:"trackingParams"`
												} `json:"templatedAdText"`
											} `json:"templatedCountdown"`
											DurationMilliseconds int `json:"durationMilliseconds"`
										} `json:"adPreviewRenderer"`
									} `json:"preskipRenderer"`
									SkippableRenderer struct {
										SkipButtonRenderer struct {
											Message struct {
												Text           string `json:"text"`
												IsTemplated    bool   `json:"isTemplated"`
												TrackingParams string `json:"trackingParams"`
											} `json:"message"`
											TrackingParams string `json:"trackingParams"`
										} `json:"skipButtonRenderer"`
									} `json:"skippableRenderer"`
									TrackingParams         string `json:"trackingParams"`
									SkipOffsetMilliseconds int    `json:"skipOffsetMilliseconds"`
								} `json:"skipAdRenderer"`
							} `json:"skipOrPreviewRenderer"`
							TrackingParams          string `json:"trackingParams"`
							VisitAdvertiserRenderer struct {
								ButtonRenderer struct {
									Style string `json:"style"`
									Text  struct {
										SimpleText string `json:"simpleText"`
									} `json:"text"`
									Icon struct {
										IconType string `json:"iconType"`
									} `json:"icon"`
									NavigationEndpoint struct {
										ClickTrackingParams string `json:"clickTrackingParams"`
										CommandMetadata     struct {
											WebCommandMetadata struct {
												URL         string `json:"url"`
												WebPageType string `json:"webPageType"`
												RootVe      int    `json:"rootVe"`
											} `json:"webCommandMetadata"`
										} `json:"commandMetadata"`
										URLEndpoint struct {
											URL                string `json:"url"`
											Target             string `json:"target"`
											AttributionSrcMode string `json:"attributionSrcMode"`
										} `json:"urlEndpoint"`
									} `json:"navigationEndpoint"`
									TrackingParams string `json:"trackingParams"`
								} `json:"buttonRenderer"`
							} `json:"visitAdvertiserRenderer"`
							AdBadgeRenderer struct {
								SimpleAdBadgeRenderer struct {
									Text struct {
										Text           string `json:"text"`
										IsTemplated    bool   `json:"isTemplated"`
										TrackingParams string `json:"trackingParams"`
									} `json:"text"`
									TrackingParams string `json:"trackingParams"`
								} `json:"simpleAdBadgeRenderer"`
							} `json:"adBadgeRenderer"`
							AdDurationRemaining struct {
								AdDurationRemainingRenderer struct {
									TemplatedCountdown struct {
										TemplatedAdText struct {
											Text           string `json:"text"`
											IsTemplated    bool   `json:"isTemplated"`
											TrackingParams string `json:"trackingParams"`
										} `json:"templatedAdText"`
									} `json:"templatedCountdown"`
									TrackingParams string `json:"trackingParams"`
								} `json:"adDurationRemainingRenderer"`
							} `json:"adDurationRemaining"`
							AdInfoRenderer struct {
								AdHoverTextButtonRenderer struct {
									Button struct {
										ButtonRenderer struct {
											Style           string `json:"style"`
											Size            string `json:"size"`
											IsDisabled      bool   `json:"isDisabled"`
											ServiceEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												OpenPopupAction     struct {
													Popup struct {
														AboutThisAdRenderer struct {
															URL struct {
																PrivateDoNotAccessOrElseTrustedResourceURLWrappedValue string `json:"privateDoNotAccessOrElseTrustedResourceUrlWrappedValue"`
															} `json:"url"`
															TrackingParams string `json:"trackingParams"`
														} `json:"aboutThisAdRenderer"`
													} `json:"popup"`
													PopupType string `json:"popupType"`
												} `json:"openPopupAction"`
											} `json:"serviceEndpoint"`
											Icon struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer"`
									} `json:"button"`
									HoverText struct {
										SimpleText string `json:"simpleText"`
									} `json:"hoverText"`
									TrackingParams string `json:"trackingParams"`
								} `json:"adHoverTextButtonRenderer"`
							} `json:"adInfoRenderer"`
							FlyoutCtaRenderer struct {
								FlyoutCtaRenderer struct {
									Image struct {
										Thumbnail struct {
											Thumbnails []struct {
												URL string `json:"url"`
											} `json:"thumbnails"`
										} `json:"thumbnail"`
										TrackingParams string `json:"trackingParams"`
									} `json:"image"`
									Headline struct {
										Text           string `json:"text"`
										IsTemplated    bool   `json:"isTemplated"`
										TrackingParams string `json:"trackingParams"`
									} `json:"headline"`
									Description struct {
										Text           string `json:"text"`
										TrackingParams string `json:"trackingParams"`
									} `json:"description"`
									ActionButton struct {
										ButtonRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														URL         string `json:"url"`
														WebPageType string `json:"webPageType"`
														RootVe      int    `json:"rootVe"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												URLEndpoint struct {
													URL                string `json:"url"`
													Target             string `json:"target"`
													AttributionSrcMode string `json:"attributionSrcMode"`
												} `json:"urlEndpoint"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
										} `json:"buttonRenderer"`
									} `json:"actionButton"`
									StartMs        int    `json:"startMs"`
									TrackingParams string `json:"trackingParams"`
								} `json:"flyoutCtaRenderer"`
							} `json:"flyoutCtaRenderer"`
							AdLayoutLoggingData struct {
								SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
							} `json:"adLayoutLoggingData"`
							ElementID        string `json:"elementId"`
							InPlayerSlotID   string `json:"inPlayerSlotId"`
							InPlayerLayoutID string `json:"inPlayerLayoutId"`
						} `json:"instreamAdPlayerOverlayRenderer"`
					} `json:"playerOverlay"`
					TrackingParams                string `json:"trackingParams"`
					LayoutID                      string `json:"layoutId"`
					AssociatedPlayerBytesLayoutID string `json:"associatedPlayerBytesLayoutId"`
				} `json:"instreamVideoAdRenderer"`
			} `json:"renderer"`
			AdSlotLoggingData struct {
				SerializedSlotAdServingDataEntry string `json:"serializedSlotAdServingDataEntry"`
			} `json:"adSlotLoggingData"`
		} `json:"adPlacementRenderer"`
	} `json:"adPlacements"`
	AdSlots []struct {
		AdSlotRenderer struct {
			AdSlotMetadata struct {
				SlotID            string `json:"slotId"`
				SlotType          string `json:"slotType"`
				AdSlotLoggingData struct {
					SerializedSlotAdServingDataEntry string `json:"serializedSlotAdServingDataEntry"`
				} `json:"adSlotLoggingData"`
				TriggerEvent string `json:"triggerEvent"`
			} `json:"adSlotMetadata"`
			FulfillmentContent struct {
				FulfilledLayout struct {
					PlayerBytesAdLayoutRenderer struct {
						AdLayoutMetadata struct {
							LayoutID            string `json:"layoutId"`
							LayoutType          string `json:"layoutType"`
							AdLayoutLoggingData struct {
								SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
							} `json:"adLayoutLoggingData"`
						} `json:"adLayoutMetadata"`
						RenderingContent struct {
							InstreamVideoAdRenderer struct {
								SkipOffsetMilliseconds int `json:"skipOffsetMilliseconds"`
								Pings                  struct {
									ImpressionPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"impressionPings"`
									ErrorPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"errorPings"`
									MutePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"mutePings"`
									UnmutePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"unmutePings"`
									PausePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"pausePings"`
									RewindPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"rewindPings"`
									ResumePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"resumePings"`
									SkipPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"skipPings"`
									ClosePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"closePings"`
									ProgressPings []struct {
										BaseURL            string `json:"baseUrl"`
										OffsetMilliseconds int    `json:"offsetMilliseconds"`
										AttributionSrcMode string `json:"attributionSrcMode"`
									} `json:"progressPings"`
									FullscreenPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"fullscreenPings"`
									ActiveViewViewablePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"activeViewViewablePings"`
									EndFullscreenPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"endFullscreenPings"`
									ActiveViewMeasurablePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"activeViewMeasurablePings"`
									AbandonPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"abandonPings"`
									ActiveViewFullyViewableAudibleHalfDurationPings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"activeViewFullyViewableAudibleHalfDurationPings"`
									CompletePings []struct {
										BaseURL string `json:"baseUrl"`
									} `json:"completePings"`
									ActiveViewTracking struct {
										TrafficType string `json:"trafficType"`
									} `json:"activeViewTracking"`
								} `json:"pings"`
								ClickthroughEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									CommandMetadata     struct {
										WebCommandMetadata struct {
											URL         string `json:"url"`
											WebPageType string `json:"webPageType"`
											RootVe      int    `json:"rootVe"`
										} `json:"webCommandMetadata"`
									} `json:"commandMetadata"`
									URLEndpoint struct {
										URL                string `json:"url"`
										Target             string `json:"target"`
										AttributionSrcMode string `json:"attributionSrcMode"`
									} `json:"urlEndpoint"`
								} `json:"clickthroughEndpoint"`
								CsiParameters []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"csiParameters"`
								PlayerVars                  string `json:"playerVars"`
								ElementID                   string `json:"elementId"`
								TrackingParams              string `json:"trackingParams"`
								LegacyInfoCardVastExtension string `json:"legacyInfoCardVastExtension"`
								SodarExtensionData          struct {
									Siub string `json:"siub"`
									Bgub string `json:"bgub"`
									Scs  string `json:"scs"`
									Bgp  string `json:"bgp"`
								} `json:"sodarExtensionData"`
								ExternalVideoID     string `json:"externalVideoId"`
								AdLayoutLoggingData struct {
									SerializedAdServingDataEntry string `json:"serializedAdServingDataEntry"`
								} `json:"adLayoutLoggingData"`
								LayoutID string `json:"layoutId"`
							} `json:"instreamVideoAdRenderer"`
						} `json:"renderingContent"`
						LayoutExitNormalTriggers []struct {
							ID                               string `json:"id"`
							OnLayoutSelfExitRequestedTrigger struct {
								TriggeringLayoutID string `json:"triggeringLayoutId"`
							} `json:"onLayoutSelfExitRequestedTrigger"`
						} `json:"layoutExitNormalTriggers"`
						LayoutExitSkipTriggers []struct {
							ID                   string `json:"id"`
							SkipRequestedTrigger struct {
								TriggeringLayoutID string `json:"triggeringLayoutId"`
							} `json:"skipRequestedTrigger"`
						} `json:"layoutExitSkipTriggers"`
						LayoutExitMuteTriggers []struct {
							ID                   string `json:"id"`
							SkipRequestedTrigger struct {
								TriggeringLayoutID string `json:"triggeringLayoutId"`
							} `json:"skipRequestedTrigger"`
						} `json:"layoutExitMuteTriggers"`
					} `json:"playerBytesAdLayoutRenderer"`
				} `json:"fulfilledLayout"`
			} `json:"fulfillmentContent"`
			SlotEntryTrigger struct {
				ID                                 string `json:"id"`
				BeforeContentVideoIDStartedTrigger struct {
				} `json:"beforeContentVideoIdStartedTrigger"`
			} `json:"slotEntryTrigger"`
			SlotFulfillmentTriggers []struct {
				ID                   string `json:"id"`
				SlotIDEnteredTrigger struct {
					TriggeringSlotID string `json:"triggeringSlotId"`
				} `json:"slotIdEnteredTrigger"`
			} `json:"slotFulfillmentTriggers"`
			SlotExpirationTriggers []struct {
				ID                  string `json:"id"`
				SlotIDExitedTrigger struct {
					TriggeringSlotID string `json:"triggeringSlotId"`
				} `json:"slotIdExitedTrigger,omitempty"`
				OnNewPlaybackAfterContentVideoIDTrigger struct {
				} `json:"onNewPlaybackAfterContentVideoIdTrigger,omitempty"`
			} `json:"slotExpirationTriggers"`
		} `json:"adSlotRenderer"`
	} `json:"adSlots"`
	AdBreakHeartbeatParams string `json:"adBreakHeartbeatParams"`
	FrameworkUpdates       struct {
		EntityBatchUpdate struct {
			Mutations []struct {
				EntityKey string `json:"entityKey"`
				Type      string `json:"type"`
				Payload   struct {
					OfflineabilityEntity struct {
						Key                     string `json:"key"`
						AddToOfflineButtonState string `json:"addToOfflineButtonState"`
					} `json:"offlineabilityEntity"`
				} `json:"payload"`
			} `json:"mutations"`
			Timestamp struct {
				Seconds string `json:"seconds"`
				Nanos   int    `json:"nanos"`
			} `json:"timestamp"`
		} `json:"entityBatchUpdate"`
	} `json:"frameworkUpdates"`
}
