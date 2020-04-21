package instacart

type delivery struct {
	Container struct {
		Title       string      `json:"title"`
		Path        string      `json:"path"`
		InitialStep interface{} `json:"initial_step"`
		Modules     []struct {
			ID   string `json:"id"`
			Data struct {
				Image struct {
					URL string `json:"url"`
					Alt string `json:"alt"`
				} `json:"image"`
				ResponsiveImage struct {
					URL        string `json:"url"`
					Alt        string `json:"alt"`
					Responsive struct {
						Template string `json:"template"`
						Defaults struct {
							Width int `json:"width"`
						} `json:"defaults"`
					} `json:"responsive"`
					Sizes []interface{} `json:"sizes"`
				} `json:"responsive_image"`
				Title                     string      `json:"title"`
				FormattedTitle            interface{} `json:"formatted_title"`
				DescriptionLines          []string    `json:"description_lines"`
				FormattedDescriptionLines interface{} `json:"formatted_description_lines"`
				PrimaryAction             interface{} `json:"primary_action"`
				TrackingEventNames        struct {
				} `json:"tracking_event_names"`
			} `json:"data"`
			Types           []string      `json:"types"`
			Layouts         []string      `json:"layouts"`
			Steps           []interface{} `json:"steps"`
			StepTransitions struct {
			} `json:"step_transitions"`
			AsyncDataPath         interface{}   `json:"async_data_path"`
			AsyncDataDependencies []interface{} `json:"async_data_dependencies"`
			PublicRollbarToken    string        `json:"public_rollbar_token"`
		} `json:"modules"`
		Attributes     []interface{} `json:"attributes"`
		AsyncDataPath  interface{}   `json:"async_data_path"`
		TrackingParams struct {
			ProductFlow         string `json:"product_flow"`
			SourceType          string `json:"source_type"`
			SourceValue         string `json:"source_value"`
			RetailerInfoVersion int    `json:"retailer_info_version"`
			PageViewID          string `json:"page_view_id"`
		} `json:"tracking_params"`
		TrackingEventNames struct {
			View  string `json:"view"`
			Close string `json:"close"`
		} `json:"tracking_event_names"`
		PublicRollbarToken string        `json:"public_rollbar_token"`
		DataDependencies   []interface{} `json:"data_dependencies"`
		Image              interface{}   `json:"image"`
		Images             []interface{} `json:"images"`
	} `json:"container"`
	Meta struct {
		TriggeredAction  interface{}   `json:"triggered_action"`
		AnalyticsActions []interface{} `json:"analytics_actions"`
		CacheTTL         int           `json:"cache_ttl"`
	} `json:"meta"`
}
