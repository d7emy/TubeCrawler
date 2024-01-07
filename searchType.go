package main

type searchContent struct {
	ChannelRenderer struct {
		ChannelID string `json:"channelId"`
		Title     struct {
			SimpleText string `json:"simpleText"`
		} `json:"title"`
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
				BrowseID         string `json:"browseId"`
				CanonicalBaseURL string `json:"canonicalBaseUrl"`
			} `json:"browseEndpoint"`
		} `json:"navigationEndpoint"`
		Thumbnail struct {
			Thumbnails []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		DescriptionSnippet struct {
			Runs []struct {
				Text string `json:"text"`
				Bold bool   `json:"bold,omitempty"`
			} `json:"runs"`
		} `json:"descriptionSnippet"`
		ShortBylineText struct {
			Runs []struct {
				Text               string `json:"text"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
			} `json:"runs"`
		} `json:"shortBylineText"`
		VideoCountText struct {
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
			SimpleText string `json:"simpleText"`
		} `json:"videoCountText"`
		SubscriptionButton struct {
			Subscribed bool `json:"subscribed"`
		} `json:"subscriptionButton"`
		OwnerBadges []struct {
			MetadataBadgeRenderer struct {
				Icon struct {
					IconType string `json:"iconType"`
				} `json:"icon"`
				Style             string `json:"style"`
				Tooltip           string `json:"tooltip"`
				TrackingParams    string `json:"trackingParams"`
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"metadataBadgeRenderer"`
		} `json:"ownerBadges"`
		SubscriberCountText struct {
			SimpleText string `json:"simpleText"`
		} `json:"subscriberCountText"`
		SubscribeButton struct {
			ButtonRenderer struct {
				Style      string `json:"style"`
				Size       string `json:"size"`
				IsDisabled bool   `json:"isDisabled"`
				Text       struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
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
					SignInEndpoint struct {
						NextEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SearchEndpoint struct {
								Query string `json:"query"`
							} `json:"searchEndpoint"`
						} `json:"nextEndpoint"`
						ContinueAction string `json:"continueAction"`
					} `json:"signInEndpoint"`
				} `json:"navigationEndpoint"`
				TrackingParams string `json:"trackingParams"`
			} `json:"buttonRenderer"`
		} `json:"subscribeButton"`
		TrackingParams string `json:"trackingParams"`
		LongBylineText struct {
			Runs []struct {
				Text               string `json:"text"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
			} `json:"runs"`
		} `json:"longBylineText"`
	} `json:"channelRenderer,omitempty"`
	VideoRenderer struct {
		VideoID   string `json:"videoId"`
		Thumbnail struct {
			Thumbnails []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		Title struct {
			Runs []struct {
				Text string `json:"text"`
			} `json:"runs"`
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
		} `json:"title"`
		LongBylineText struct {
			Runs []struct {
				Text               string `json:"text"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
			} `json:"runs"`
		} `json:"longBylineText"`
		PublishedTimeText struct {
			SimpleText string `json:"simpleText"`
		} `json:"publishedTimeText"`
		LengthText struct {
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
			SimpleText string `json:"simpleText"`
		} `json:"lengthText"`
		ViewCountText struct {
			SimpleText string `json:"simpleText"`
		} `json:"viewCountText"`
		NavigationEndpoint struct {
			ClickTrackingParams string `json:"clickTrackingParams"`
			CommandMetadata     struct {
				WebCommandMetadata struct {
					URL         string `json:"url"`
					WebPageType string `json:"webPageType"`
					RootVe      int    `json:"rootVe"`
				} `json:"webCommandMetadata"`
			} `json:"commandMetadata"`
			WatchEndpoint struct {
				VideoID                            string `json:"videoId"`
				Params                             string `json:"params"`
				PlayerParams                       string `json:"playerParams"`
				WatchEndpointSupportedOnesieConfig struct {
					HTML5PlaybackOnesieConfig struct {
						CommonConfig struct {
							URL string `json:"url"`
						} `json:"commonConfig"`
					} `json:"html5PlaybackOnesieConfig"`
				} `json:"watchEndpointSupportedOnesieConfig"`
			} `json:"watchEndpoint"`
		} `json:"navigationEndpoint"`
		OwnerBadges []struct {
			MetadataBadgeRenderer struct {
				Icon struct {
					IconType string `json:"iconType"`
				} `json:"icon"`
				Style             string `json:"style"`
				Tooltip           string `json:"tooltip"`
				TrackingParams    string `json:"trackingParams"`
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"metadataBadgeRenderer"`
		} `json:"ownerBadges"`
		OwnerText struct {
			Runs []struct {
				Text               string `json:"text"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
			} `json:"runs"`
		} `json:"ownerText"`
		ShortBylineText struct {
			Runs []struct {
				Text               string `json:"text"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
			} `json:"runs"`
		} `json:"shortBylineText"`
		TrackingParams     string `json:"trackingParams"`
		ShowActionMenu     bool   `json:"showActionMenu"`
		ShortViewCountText struct {
			Accessibility struct {
				AccessibilityData struct {
					Label string `json:"label"`
				} `json:"accessibilityData"`
			} `json:"accessibility"`
			SimpleText string `json:"simpleText"`
		} `json:"shortViewCountText"`
		Menu struct {
			MenuRenderer struct {
				Items []struct {
					MenuServiceItemRenderer struct {
						Text struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"text"`
						Icon struct {
							IconType string `json:"iconType"`
						} `json:"icon"`
						ServiceEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									SendPost bool `json:"sendPost"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SignalServiceEndpoint struct {
								Signal  string `json:"signal"`
								Actions []struct {
									ClickTrackingParams  string `json:"clickTrackingParams"`
									AddToPlaylistCommand struct {
										OpenMiniplayer      bool   `json:"openMiniplayer"`
										VideoID             string `json:"videoId"`
										ListType            string `json:"listType"`
										OnCreateListCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													SendPost bool   `json:"sendPost"`
													APIURL   string `json:"apiUrl"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											CreatePlaylistServiceEndpoint struct {
												VideoIds []string `json:"videoIds"`
												Params   string   `json:"params"`
											} `json:"createPlaylistServiceEndpoint"`
										} `json:"onCreateListCommand"`
										VideoIds []string `json:"videoIds"`
									} `json:"addToPlaylistCommand"`
								} `json:"actions"`
							} `json:"signalServiceEndpoint"`
						} `json:"serviceEndpoint"`
						TrackingParams string `json:"trackingParams"`
					} `json:"menuServiceItemRenderer,omitempty"`
					MenuServiceItemDownloadRenderer struct {
						ServiceEndpoint struct {
							ClickTrackingParams  string `json:"clickTrackingParams"`
							OfflineVideoEndpoint struct {
								VideoID      string `json:"videoId"`
								OnAddCommand struct {
									ClickTrackingParams      string `json:"clickTrackingParams"`
									GetDownloadActionCommand struct {
										VideoID string `json:"videoId"`
										Params  string `json:"params"`
									} `json:"getDownloadActionCommand"`
								} `json:"onAddCommand"`
							} `json:"offlineVideoEndpoint"`
						} `json:"serviceEndpoint"`
						TrackingParams string `json:"trackingParams"`
					} `json:"menuServiceItemDownloadRenderer,omitempty"`
				} `json:"items"`
				TrackingParams string `json:"trackingParams"`
				Accessibility  struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
			} `json:"menuRenderer"`
		} `json:"menu"`
		ChannelThumbnailSupportedRenderers struct {
			ChannelThumbnailWithLinkRenderer struct {
				Thumbnail struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"thumbnail"`
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
						BrowseID         string `json:"browseId"`
						CanonicalBaseURL string `json:"canonicalBaseUrl"`
					} `json:"browseEndpoint"`
				} `json:"navigationEndpoint"`
				Accessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
			} `json:"channelThumbnailWithLinkRenderer"`
		} `json:"channelThumbnailSupportedRenderers"`
		ThumbnailOverlays []struct {
			ThumbnailOverlayTimeStatusRenderer struct {
				Text struct {
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					SimpleText string `json:"simpleText"`
				} `json:"text"`
				Style string `json:"style"`
			} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
			ThumbnailOverlayToggleButtonRenderer struct {
				IsToggled     bool `json:"isToggled"`
				UntoggledIcon struct {
					IconType string `json:"iconType"`
				} `json:"untoggledIcon"`
				ToggledIcon struct {
					IconType string `json:"iconType"`
				} `json:"toggledIcon"`
				UntoggledTooltip         string `json:"untoggledTooltip"`
				ToggledTooltip           string `json:"toggledTooltip"`
				UntoggledServiceEndpoint struct {
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
				} `json:"untoggledServiceEndpoint"`
				ToggledServiceEndpoint struct {
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
				} `json:"toggledServiceEndpoint"`
				UntoggledAccessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"untoggledAccessibility"`
				ToggledAccessibility struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"toggledAccessibility"`
				TrackingParams string `json:"trackingParams"`
			} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
			ThumbnailOverlayNowPlayingRenderer struct {
				Text struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"text"`
			} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
			ThumbnailOverlayLoadingPreviewRenderer struct {
				Text struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"text"`
			} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
		} `json:"thumbnailOverlays"`
		RichThumbnail struct {
			MovingThumbnailRenderer struct {
				MovingThumbnailDetails struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
					LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
				} `json:"movingThumbnailDetails"`
				EnableHoveredLogging bool `json:"enableHoveredLogging"`
				EnableOverlay        bool `json:"enableOverlay"`
			} `json:"movingThumbnailRenderer"`
		} `json:"richThumbnail"`
		DetailedMetadataSnippets []struct {
			SnippetText struct {
				Runs []struct {
					Text string `json:"text"`
					Bold bool   `json:"bold,omitempty"`
				} `json:"runs"`
			} `json:"snippetText"`
			SnippetHoverText struct {
				Runs []struct {
					Text string `json:"text"`
				} `json:"runs"`
			} `json:"snippetHoverText"`
			MaxOneLine bool `json:"maxOneLine"`
		} `json:"detailedMetadataSnippets"`
		InlinePlaybackEndpoint struct {
			ClickTrackingParams string `json:"clickTrackingParams"`
			CommandMetadata     struct {
				WebCommandMetadata struct {
					URL         string `json:"url"`
					WebPageType string `json:"webPageType"`
					RootVe      int    `json:"rootVe"`
				} `json:"webCommandMetadata"`
			} `json:"commandMetadata"`
			WatchEndpoint struct {
				VideoID              string `json:"videoId"`
				Params               string `json:"params"`
				PlayerParams         string `json:"playerParams"`
				PlayerExtraURLParams []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"playerExtraUrlParams"`
				WatchEndpointSupportedOnesieConfig struct {
					HTML5PlaybackOnesieConfig struct {
						CommonConfig struct {
							URL string `json:"url"`
						} `json:"commonConfig"`
					} `json:"html5PlaybackOnesieConfig"`
				} `json:"watchEndpointSupportedOnesieConfig"`
			} `json:"watchEndpoint"`
		} `json:"inlinePlaybackEndpoint"`
		SearchVideoResultEntityKey string `json:"searchVideoResultEntityKey"`
	} `json:"videoRenderer,omitempty"`
	ReelShelfRenderer struct {
		Title struct {
			SimpleText string `json:"simpleText"`
		} `json:"title"`
		Button struct {
			MenuRenderer struct {
				Items []struct {
					MenuNavigationItemRenderer struct {
						Text struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"text"`
						Icon struct {
							IconType string `json:"iconType"`
						} `json:"icon"`
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									IgnoreNavigation bool `json:"ignoreNavigation"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							UserFeedbackEndpoint struct {
								AdditionalDatas []struct {
									UserFeedbackEndpointProductSpecificValueData struct {
										Key   string `json:"key"`
										Value string `json:"value"`
									} `json:"userFeedbackEndpointProductSpecificValueData"`
								} `json:"additionalDatas"`
							} `json:"userFeedbackEndpoint"`
						} `json:"navigationEndpoint"`
						TrackingParams string `json:"trackingParams"`
						Accessibility  struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"menuNavigationItemRenderer"`
				} `json:"items"`
				TrackingParams string `json:"trackingParams"`
				Accessibility  struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
			} `json:"menuRenderer"`
		} `json:"button"`
		Items []struct {
			ReelItemRenderer struct {
				VideoID  string `json:"videoId"`
				Headline struct {
					SimpleText string `json:"simpleText"`
				} `json:"headline"`
				Thumbnail struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
					IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
				} `json:"thumbnail"`
				ViewCountText struct {
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					SimpleText string `json:"simpleText"`
				} `json:"viewCountText"`
				NavigationEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					ReelWatchEndpoint struct {
						VideoID      string `json:"videoId"`
						PlayerParams string `json:"playerParams"`
						Thumbnail    struct {
							Thumbnails []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"thumbnails"`
							IsOriginalAspectRatio bool `json:"isOriginalAspectRatio"`
						} `json:"thumbnail"`
						Overlay struct {
							ReelPlayerOverlayRenderer struct {
								Style                     string `json:"style"`
								TrackingParams            string `json:"trackingParams"`
								ReelPlayerNavigationModel string `json:"reelPlayerNavigationModel"`
							} `json:"reelPlayerOverlayRenderer"`
						} `json:"overlay"`
						Params           string `json:"params"`
						SequenceProvider string `json:"sequenceProvider"`
						SequenceParams   string `json:"sequenceParams"`
						LoggingContext   struct {
							VssLoggingContext struct {
								SerializedContextData string `json:"serializedContextData"`
							} `json:"vssLoggingContext"`
							QoeLoggingContext struct {
								SerializedContextData string `json:"serializedContextData"`
							} `json:"qoeLoggingContext"`
						} `json:"loggingContext"`
						UstreamerConfig string `json:"ustreamerConfig"`
					} `json:"reelWatchEndpoint"`
				} `json:"navigationEndpoint"`
				Menu struct {
					MenuRenderer struct {
						Items []struct {
							MenuNavigationItemRenderer struct {
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
								Icon struct {
									IconType string `json:"iconType"`
								} `json:"icon"`
								NavigationEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									CommandMetadata     struct {
										WebCommandMetadata struct {
											IgnoreNavigation bool `json:"ignoreNavigation"`
										} `json:"webCommandMetadata"`
									} `json:"commandMetadata"`
									UserFeedbackEndpoint struct {
										AdditionalDatas []struct {
											UserFeedbackEndpointProductSpecificValueData struct {
												Key   string `json:"key"`
												Value string `json:"value"`
											} `json:"userFeedbackEndpointProductSpecificValueData"`
										} `json:"additionalDatas"`
									} `json:"userFeedbackEndpoint"`
								} `json:"navigationEndpoint"`
								TrackingParams string `json:"trackingParams"`
								Accessibility  struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"menuNavigationItemRenderer"`
						} `json:"items"`
						TrackingParams string `json:"trackingParams"`
						Accessibility  struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"menuRenderer"`
				} `json:"menu"`
				TrackingParams string `json:"trackingParams"`
				Accessibility  struct {
					AccessibilityData struct {
						Label string `json:"label"`
					} `json:"accessibilityData"`
				} `json:"accessibility"`
				Style                  string `json:"style"`
				VideoType              string `json:"videoType"`
				InlinePlaybackEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WatchEndpoint struct {
						VideoID              string `json:"videoId"`
						PlayerParams         string `json:"playerParams"`
						PlayerExtraURLParams []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
						} `json:"playerExtraUrlParams"`
						WatchEndpointSupportedOnesieConfig struct {
							HTML5PlaybackOnesieConfig struct {
								CommonConfig struct {
									URL string `json:"url"`
								} `json:"commonConfig"`
							} `json:"html5PlaybackOnesieConfig"`
						} `json:"watchEndpointSupportedOnesieConfig"`
					} `json:"watchEndpoint"`
				} `json:"inlinePlaybackEndpoint"`
				LoggingDirectives struct {
					TrackingParams string `json:"trackingParams"`
					Visibility     struct {
						Types string `json:"types"`
					} `json:"visibility"`
					EnableDisplayloggerExperiment bool `json:"enableDisplayloggerExperiment"`
				} `json:"loggingDirectives"`
			} `json:"reelItemRenderer"`
		} `json:"items"`
		TrackingParams string `json:"trackingParams"`
		Icon           struct {
			IconType string `json:"iconType"`
		} `json:"icon"`
	} `json:"reelShelfRenderer,omitempty"`
	ShelfRenderer struct {
		Title struct {
			SimpleText string `json:"simpleText"`
		} `json:"title"`
		Content struct {
			VerticalListRenderer struct {
				Items []struct {
					VideoRenderer struct {
						VideoID   string `json:"videoId"`
						Thumbnail struct {
							Thumbnails []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"thumbnails"`
						} `json:"thumbnail"`
						Title struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
						} `json:"title"`
						LongBylineText struct {
							Runs []struct {
								Text               string `json:"text"`
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
										BrowseID         string `json:"browseId"`
										CanonicalBaseURL string `json:"canonicalBaseUrl"`
									} `json:"browseEndpoint"`
								} `json:"navigationEndpoint"`
							} `json:"runs"`
						} `json:"longBylineText"`
						PublishedTimeText struct {
							SimpleText string `json:"simpleText"`
						} `json:"publishedTimeText"`
						LengthText struct {
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							SimpleText string `json:"simpleText"`
						} `json:"lengthText"`
						ViewCountText struct {
							SimpleText string `json:"simpleText"`
						} `json:"viewCountText"`
						NavigationEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							WatchEndpoint struct {
								VideoID                            string `json:"videoId"`
								PlayerParams                       string `json:"playerParams"`
								WatchEndpointSupportedOnesieConfig struct {
									HTML5PlaybackOnesieConfig struct {
										CommonConfig struct {
											URL string `json:"url"`
										} `json:"commonConfig"`
									} `json:"html5PlaybackOnesieConfig"`
								} `json:"watchEndpointSupportedOnesieConfig"`
							} `json:"watchEndpoint"`
						} `json:"navigationEndpoint"`
						OwnerText struct {
							Runs []struct {
								Text               string `json:"text"`
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
										BrowseID         string `json:"browseId"`
										CanonicalBaseURL string `json:"canonicalBaseUrl"`
									} `json:"browseEndpoint"`
								} `json:"navigationEndpoint"`
							} `json:"runs"`
						} `json:"ownerText"`
						ShortBylineText struct {
							Runs []struct {
								Text               string `json:"text"`
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
										BrowseID         string `json:"browseId"`
										CanonicalBaseURL string `json:"canonicalBaseUrl"`
									} `json:"browseEndpoint"`
								} `json:"navigationEndpoint"`
							} `json:"runs"`
						} `json:"shortBylineText"`
						TrackingParams     string `json:"trackingParams"`
						ShowActionMenu     bool   `json:"showActionMenu"`
						ShortViewCountText struct {
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							SimpleText string `json:"simpleText"`
						} `json:"shortViewCountText"`
						Menu struct {
							MenuRenderer struct {
								Items []struct {
									MenuServiceItemRenderer struct {
										Text struct {
											Runs []struct {
												Text string `json:"text"`
											} `json:"runs"`
										} `json:"text"`
										Icon struct {
											IconType string `json:"iconType"`
										} `json:"icon"`
										ServiceEndpoint struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													SendPost bool `json:"sendPost"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											SignalServiceEndpoint struct {
												Signal  string `json:"signal"`
												Actions []struct {
													ClickTrackingParams  string `json:"clickTrackingParams"`
													AddToPlaylistCommand struct {
														OpenMiniplayer      bool   `json:"openMiniplayer"`
														VideoID             string `json:"videoId"`
														ListType            string `json:"listType"`
														OnCreateListCommand struct {
															ClickTrackingParams string `json:"clickTrackingParams"`
															CommandMetadata     struct {
																WebCommandMetadata struct {
																	SendPost bool   `json:"sendPost"`
																	APIURL   string `json:"apiUrl"`
																} `json:"webCommandMetadata"`
															} `json:"commandMetadata"`
															CreatePlaylistServiceEndpoint struct {
																VideoIds []string `json:"videoIds"`
																Params   string   `json:"params"`
															} `json:"createPlaylistServiceEndpoint"`
														} `json:"onCreateListCommand"`
														VideoIds []string `json:"videoIds"`
													} `json:"addToPlaylistCommand"`
												} `json:"actions"`
											} `json:"signalServiceEndpoint"`
										} `json:"serviceEndpoint"`
										TrackingParams string `json:"trackingParams"`
									} `json:"menuServiceItemRenderer,omitempty"`
									MenuServiceItemDownloadRenderer struct {
										ServiceEndpoint struct {
											ClickTrackingParams  string `json:"clickTrackingParams"`
											OfflineVideoEndpoint struct {
												VideoID      string `json:"videoId"`
												OnAddCommand struct {
													ClickTrackingParams      string `json:"clickTrackingParams"`
													GetDownloadActionCommand struct {
														VideoID string `json:"videoId"`
														Params  string `json:"params"`
													} `json:"getDownloadActionCommand"`
												} `json:"onAddCommand"`
											} `json:"offlineVideoEndpoint"`
										} `json:"serviceEndpoint"`
										TrackingParams string `json:"trackingParams"`
									} `json:"menuServiceItemDownloadRenderer,omitempty"`
								} `json:"items"`
								TrackingParams string `json:"trackingParams"`
								Accessibility  struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"menuRenderer"`
						} `json:"menu"`
						ChannelThumbnailSupportedRenderers struct {
							ChannelThumbnailWithLinkRenderer struct {
								Thumbnail struct {
									Thumbnails []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"thumbnails"`
								} `json:"thumbnail"`
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
										BrowseID         string `json:"browseId"`
										CanonicalBaseURL string `json:"canonicalBaseUrl"`
									} `json:"browseEndpoint"`
								} `json:"navigationEndpoint"`
								Accessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"accessibility"`
							} `json:"channelThumbnailWithLinkRenderer"`
						} `json:"channelThumbnailSupportedRenderers"`
						ThumbnailOverlays []struct {
							ThumbnailOverlayTimeStatusRenderer struct {
								Text struct {
									Accessibility struct {
										AccessibilityData struct {
											Label string `json:"label"`
										} `json:"accessibilityData"`
									} `json:"accessibility"`
									SimpleText string `json:"simpleText"`
								} `json:"text"`
								Style string `json:"style"`
							} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`
							ThumbnailOverlayToggleButtonRenderer struct {
								IsToggled     bool `json:"isToggled"`
								UntoggledIcon struct {
									IconType string `json:"iconType"`
								} `json:"untoggledIcon"`
								ToggledIcon struct {
									IconType string `json:"iconType"`
								} `json:"toggledIcon"`
								UntoggledTooltip         string `json:"untoggledTooltip"`
								ToggledTooltip           string `json:"toggledTooltip"`
								UntoggledServiceEndpoint struct {
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
								} `json:"untoggledServiceEndpoint"`
								ToggledServiceEndpoint struct {
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
								} `json:"toggledServiceEndpoint"`
								UntoggledAccessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"untoggledAccessibility"`
								ToggledAccessibility struct {
									AccessibilityData struct {
										Label string `json:"label"`
									} `json:"accessibilityData"`
								} `json:"toggledAccessibility"`
								TrackingParams string `json:"trackingParams"`
							} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
							ThumbnailOverlayNowPlayingRenderer struct {
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
							} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
							ThumbnailOverlayLoadingPreviewRenderer struct {
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
							} `json:"thumbnailOverlayLoadingPreviewRenderer,omitempty"`
						} `json:"thumbnailOverlays"`
						RichThumbnail struct {
							MovingThumbnailRenderer struct {
								MovingThumbnailDetails struct {
									Thumbnails []struct {
										URL    string `json:"url"`
										Width  int    `json:"width"`
										Height int    `json:"height"`
									} `json:"thumbnails"`
									LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
								} `json:"movingThumbnailDetails"`
								EnableHoveredLogging bool `json:"enableHoveredLogging"`
								EnableOverlay        bool `json:"enableOverlay"`
							} `json:"movingThumbnailRenderer"`
						} `json:"richThumbnail"`
						InlinePlaybackEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									URL         string `json:"url"`
									WebPageType string `json:"webPageType"`
									RootVe      int    `json:"rootVe"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							WatchEndpoint struct {
								VideoID              string `json:"videoId"`
								PlayerParams         string `json:"playerParams"`
								PlayerExtraURLParams []struct {
									Key   string `json:"key"`
									Value string `json:"value"`
								} `json:"playerExtraUrlParams"`
								WatchEndpointSupportedOnesieConfig struct {
									HTML5PlaybackOnesieConfig struct {
										CommonConfig struct {
											URL string `json:"url"`
										} `json:"commonConfig"`
									} `json:"html5PlaybackOnesieConfig"`
								} `json:"watchEndpointSupportedOnesieConfig"`
							} `json:"watchEndpoint"`
						} `json:"inlinePlaybackEndpoint"`
						SearchVideoResultEntityKey string `json:"searchVideoResultEntityKey"`
					} `json:"videoRenderer"`
				} `json:"items"`
				CollapsedItemCount       int `json:"collapsedItemCount"`
				CollapsedStateButtonText struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
				} `json:"collapsedStateButtonText"`
				TrackingParams string `json:"trackingParams"`
			} `json:"verticalListRenderer"`
		} `json:"content"`
		TrackingParams string `json:"trackingParams"`
	} `json:"shelfRenderer,omitempty"`
}
type SearchStruct struct {
	EstimatedResults string `json:"estimatedResults"`
	Contents         struct {
		TwoColumnSearchResultsRenderer struct {
			PrimaryContents struct {
				SectionListRenderer struct {
					Contents []struct {
						ItemSectionRenderer struct {
							Contents       []searchContent `json:"contents"`
							TrackingParams string          `json:"trackingParams"`
						} `json:"itemSectionRenderer,omitempty"`
						ContinuationItemRenderer struct {
							Trigger              string `json:"trigger"`
							ContinuationEndpoint struct {
								ClickTrackingParams string `json:"clickTrackingParams"`
								CommandMetadata     struct {
									WebCommandMetadata struct {
										SendPost bool   `json:"sendPost"`
										APIURL   string `json:"apiUrl"`
									} `json:"webCommandMetadata"`
								} `json:"commandMetadata"`
								ContinuationCommand struct {
									Token   string `json:"token"`
									Request string `json:"request"`
								} `json:"continuationCommand"`
							} `json:"continuationEndpoint"`
							LoggingDirectives struct {
								TrackingParams string `json:"trackingParams"`
							} `json:"loggingDirectives"`
						} `json:"continuationItemRenderer,omitempty"`
					} `json:"contents"`
				} `json:"sectionListRenderer"`
			} `json:"primaryContents"`
		} `json:"twoColumnSearchResultsRenderer"`
	} `json:"contents"`
}
type SearchContinuationStruct struct {
	OnResponseReceivedCommands []struct {
		ClickTrackingParams           string `json:"clickTrackingParams"`
		AppendContinuationItemsAction struct {
			ContinuationItems []struct {
				ItemSectionRenderer struct {
					Contents       []searchContent `json:"contents"`
					TrackingParams string          `json:"trackingParams"`
				} `json:"itemSectionRenderer,omitempty"`
				ContinuationItemRenderer struct {
					Trigger              string `json:"trigger"`
					ContinuationEndpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								SendPost bool   `json:"sendPost"`
								APIURL   string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						ContinuationCommand struct {
							Token   string `json:"token"`
							Request string `json:"request"`
						} `json:"continuationCommand"`
					} `json:"continuationEndpoint"`
					LoggingDirectives struct {
						TrackingParams string `json:"trackingParams"`
					} `json:"loggingDirectives"`
				} `json:"continuationItemRenderer,omitempty"`
			} `json:"continuationItems"`
			TargetID string `json:"targetId"`
		} `json:"appendContinuationItemsAction"`
	} `json:"onResponseReceivedCommands"`
}
