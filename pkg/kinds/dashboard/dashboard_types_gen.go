// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     kinds/gen.go
// Using jennies:
//     GoTypesJenny
//     LatestJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dashboard

// Defines values for SpecStyle.
const (
	SpecStyleDark  SpecStyle = "dark"
	SpecStyleLight SpecStyle = "light"
)

// Defines values for SpecCursorSync.
const (
	SpecDashboardCursorSyncN0 SpecDashboardCursorSync = 0
	SpecDashboardCursorSyncN1 SpecDashboardCursorSync = 1
	SpecDashboardCursorSyncN2 SpecDashboardCursorSync = 2
)

// Defines values for SpecLinkType.
const (
	SpecDashboardLinkTypeDashboards SpecDashboardLinkType = "dashboards"
	SpecDashboardLinkTypeLink       SpecDashboardLinkType = "link"
)

// Defines values for SpecFieldColorModeId.
const (
	SpecFieldColorModeIdContinuousGrYlRd SpecFieldColorModeId = "continuous-GrYlRd"
	SpecFieldColorModeIdFixed            SpecFieldColorModeId = "fixed"
	SpecFieldColorModeIdPaletteClassic   SpecFieldColorModeId = "palette-classic"
	SpecFieldColorModeIdPaletteSaturated SpecFieldColorModeId = "palette-saturated"
	SpecFieldColorModeIdThresholds       SpecFieldColorModeId = "thresholds"
)

// Defines values for SpecFieldColorSeriesByMode.
const (
	SpecFieldColorSeriesByModeLast SpecFieldColorSeriesByMode = "last"
	SpecFieldColorSeriesByModeMax  SpecFieldColorSeriesByMode = "max"
	SpecFieldColorSeriesByModeMin  SpecFieldColorSeriesByMode = "min"
)

// Defines values for SpecGraphPanelType.
const (
	SpecGraphPanelTypeGraph SpecGraphPanelType = "graph"
)

// Defines values for SpecHeatmapPanelType.
const (
	SpecHeatmapPanelTypeHeatmap SpecHeatmapPanelType = "heatmap"
)

// Defines values for SpecLoadingState.
const (
	SpecLoadingStateDone       SpecLoadingState = "Done"
	SpecLoadingStateError      SpecLoadingState = "Error"
	SpecLoadingStateLoading    SpecLoadingState = "Loading"
	SpecLoadingStateNotStarted SpecLoadingState = "NotStarted"
	SpecLoadingStateStreaming  SpecLoadingState = "Streaming"
)

// Defines values for SpecMappingType.
const (
	SpecMappingTypeRange   SpecMappingType = "range"
	SpecMappingTypeRegex   SpecMappingType = "regex"
	SpecMappingTypeSpecial SpecMappingType = "special"
	SpecMappingTypeValue   SpecMappingType = "value"
)

// Defines values for SpecPanelRepeatDirection.
const (
	SpecPanelRepeatDirectionH SpecPanelRepeatDirection = "h"
	SpecPanelRepeatDirectionV SpecPanelRepeatDirection = "v"
)

// Defines values for SpecRangeMapType.
const (
	SpecRangeMapTypeRange   SpecRangeMapType = "range"
	SpecRangeMapTypeRegex   SpecRangeMapType = "regex"
	SpecRangeMapTypeSpecial SpecRangeMapType = "special"
	SpecRangeMapTypeValue   SpecRangeMapType = "value"
)

// Defines values for SpecRegexMapType.
const (
	SpecRegexMapTypeRange   SpecRegexMapType = "range"
	SpecRegexMapTypeRegex   SpecRegexMapType = "regex"
	SpecRegexMapTypeSpecial SpecRegexMapType = "special"
	SpecRegexMapTypeValue   SpecRegexMapType = "value"
)

// Defines values for SpecRowPanelType.
const (
	SpecRowPanelTypeRow SpecRowPanelType = "row"
)

// Defines values for SpecSpecialValueMapOptionsMatch.
const (
	SpecSpecialValueMapOptionsMatchFalse SpecSpecialValueMapOptionsMatch = "false"
	SpecSpecialValueMapOptionsMatchTrue  SpecSpecialValueMapOptionsMatch = "true"
)

// Defines values for SpecSpecialValueMapType.
const (
	SpecSpecialValueMapTypeRange   SpecSpecialValueMapType = "range"
	SpecSpecialValueMapTypeRegex   SpecSpecialValueMapType = "regex"
	SpecSpecialValueMapTypeSpecial SpecSpecialValueMapType = "special"
	SpecSpecialValueMapTypeValue   SpecSpecialValueMapType = "value"
)

// Defines values for SpecSpecialValueMatch.
const (
	SpecSpecialValueMatchEmpty   SpecSpecialValueMatch = "empty"
	SpecSpecialValueMatchFalse   SpecSpecialValueMatch = "false"
	SpecSpecialValueMatchNan     SpecSpecialValueMatch = "nan"
	SpecSpecialValueMatchNull    SpecSpecialValueMatch = "null"
	SpecSpecialValueMatchNullNan SpecSpecialValueMatch = "null+nan"
	SpecSpecialValueMatchTrue    SpecSpecialValueMatch = "true"
)

// Defines values for SpecThresholdsMode.
const (
	SpecThresholdsModeAbsolute   SpecThresholdsMode = "absolute"
	SpecThresholdsModePercentage SpecThresholdsMode = "percentage"
)

// Defines values for SpecValueMapType.
const (
	SpecValueMapTypeRange   SpecValueMapType = "range"
	SpecValueMapTypeRegex   SpecValueMapType = "regex"
	SpecValueMapTypeSpecial SpecValueMapType = "special"
	SpecValueMapTypeValue   SpecValueMapType = "value"
)

// Defines values for SpecVariableHide.
const (
	SpecVariableHideN0 SpecVariableHide = 0
	SpecVariableHideN1 SpecVariableHide = 1
	SpecVariableHideN2 SpecVariableHide = 2
)

// Defines values for SpecVariableType.
const (
	SpecVariableTypeAdhoc      SpecVariableType = "adhoc"
	SpecVariableTypeConstant   SpecVariableType = "constant"
	SpecVariableTypeCustom     SpecVariableType = "custom"
	SpecVariableTypeDatasource SpecVariableType = "datasource"
	SpecVariableTypeInterval   SpecVariableType = "interval"
	SpecVariableTypeQuery      SpecVariableType = "query"
	SpecVariableTypeSystem     SpecVariableType = "system"
	SpecVariableTypeTextbox    SpecVariableType = "textbox"
)

// Dashboard defines model for Dashboard.
type Dashboard struct {
	Spec struct {
		// TODO docs
		Annotations *struct {
			List []SpecAnnotationQuery `json:"list,omitempty"`
		} `json:"annotations,omitempty"`

		// Description of dashboard.
		Description *string `json:"description,omitempty"`

		// Whether a dashboard is editable or not.
		Editable bool `json:"editable"`

		// The month that the fiscal year starts on.  0 = January, 11 = December
		FiscalYearStartMonth *int `json:"fiscalYearStartMonth,omitempty"`

		// For dashboards imported from the https://grafana.com/grafana/dashboards/ portal
		GnetId *string `json:"gnetId,omitempty"`

		// 0 for no shared crosshair or tooltip (default).
		// 1 for shared crosshair.
		// 2 for shared crosshair AND shared tooltip.
		GraphTooltip SpecDashboardCursorSync `json:"graphTooltip"`

		// Unique numeric identifier for the dashboard.
		// TODO must isolate or remove identifiers local to a Grafana instance...?
		Id *int64 `json:"id,omitempty"`

		// TODO docs
		Links []SpecDashboardLink `json:"links,omitempty"`

		// When set to true, the dashboard will redraw panels at an interval matching the pixel width.
		// This will keep data "moving left" regardless of the query refresh rate.  This setting helps
		// avoid dashboards presenting stale live data
		LiveNow *bool         `json:"liveNow,omitempty"`
		Panels  []interface{} `json:"panels,omitempty"`

		// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
		Refresh *interface{} `json:"refresh,omitempty"`

		// This property should only be used in dashboards defined by plugins.  It is a quick check
		// to see if the version has changed since the last time.  Unclear why using the version property
		// is insufficient.
		Revision *int64 `json:"revision,omitempty"`

		// Version of the JSON schema, incremented each time a Grafana update brings
		// changes to said schema.
		// TODO this is the existing schema numbering system. It will be replaced by Thema's themaVersion
		SchemaVersion int `json:"schemaVersion"`

		// TODO docs
		Snapshot *SpecSnapshot `json:"snapshot,omitempty"`

		// Theme of dashboard.
		Style SpecStyle `json:"style"`

		// Tags associated with dashboard.
		Tags []string `json:"tags,omitempty"`

		// TODO docs
		Templating *struct {
			List []SpecVariableModel `json:"list,omitempty"`
		} `json:"templating,omitempty"`

		// Time range for dashboard, e.g. last 6 hours, last 7 days, etc
		Time *struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"time,omitempty"`

		// TODO docs
		// TODO this appears to be spread all over in the frontend. Concepts will likely need tidying in tandem with schema changes
		Timepicker *struct {
			// Whether timepicker is collapsed or not.
			Collapse bool `json:"collapse"`

			// Whether timepicker is enabled or not.
			Enable bool `json:"enable"`

			// Whether timepicker is visible or not.
			Hidden bool `json:"hidden"`

			// Selectable intervals for auto-refresh.
			RefreshIntervals []string `json:"refresh_intervals"`

			// TODO docs
			TimeOptions []string `json:"time_options"`
		} `json:"timepicker,omitempty"`

		// Timezone of dashboard. Accepts IANA TZDB zone ID or "browser" or "utc".
		Timezone *string `json:"timezone,omitempty"`

		// Title of dashboard.
		Title *string `json:"title,omitempty"`

		// Unique dashboard identifier that can be generated by anyone. string (8-40)
		Uid *string `json:"uid,omitempty"`

		// Version of the dashboard, incremented each time the dashboard is updated.
		Version *int `json:"version,omitempty"`

		// TODO docs
		WeekStart *string `json:"weekStart,omitempty"`
	} `json:"spec"`
}

// Theme of dashboard.
type SpecStyle string

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
type SpecAnnotationQuery struct {
	BuiltIn int `json:"builtIn"`

	// Datasource to use for annotation.
	Datasource struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource"`

	// Whether annotation is enabled.
	Enable bool `json:"enable"`

	// Whether to hide annotation.
	Hide *bool `json:"hide,omitempty"`

	// Annotation icon color.
	IconColor *string `json:"iconColor,omitempty"`

	// Name of annotation.
	Name *string `json:"name,omitempty"`

	// Query for annotation data.
	RawQuery *string `json:"rawQuery,omitempty"`
	ShowIn   int     `json:"showIn"`

	// TODO docs
	Target *SpecAnnotationTarget `json:"target,omitempty"`
	Type   string                `json:"type"`
}

// TODO docs
type SpecAnnotationTarget struct {
	Limit    int64    `json:"limit"`
	MatchAny bool     `json:"matchAny"`
	Tags     []string `json:"tags"`
	Type     string   `json:"type"`
}

// 0 for no shared crosshair or tooltip (default).
// 1 for shared crosshair.
// 2 for shared crosshair AND shared tooltip.
type SpecDashboardCursorSync int

// FROM public/app/features/dashboard/state/Models.ts - ish
// TODO docs
type SpecDashboardLink struct {
	AsDropdown  bool     `json:"asDropdown"`
	Icon        string   `json:"icon"`
	IncludeVars bool     `json:"includeVars"`
	KeepTime    bool     `json:"keepTime"`
	Tags        []string `json:"tags"`
	TargetBlank bool     `json:"targetBlank"`
	Title       string   `json:"title"`
	Tooltip     string   `json:"tooltip"`

	// TODO docs
	Type SpecDashboardLinkType `json:"type"`
	Url  string                `json:"url"`
}

// TODO docs
type SpecDashboardLinkType string

// Ref to a DataSource instance
type SpecDataSourceRef struct {
	// The plugin type-id
	Type *string `json:"type,omitempty"`

	// Specific datasource instance
	Uid *string `json:"uid,omitempty"`
}

// TODO docs
type SpecDataTransformerConfig struct {
	// Disabled transformations are skipped
	Disabled *bool              `json:"disabled,omitempty"`
	Filter   *SpecMatcherConfig `json:"filter,omitempty"`

	// Unique identifier of transformer
	Id string `json:"id"`

	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	Options interface{} `json:"options"`
}

// SpecDynamicConfigValue defines model for spec.#DynamicConfigValue.
type SpecDynamicConfigValue struct {
	Id    string       `json:"id"`
	Value *interface{} `json:"value,omitempty"`
}

// TODO docs
type SpecFieldColor struct {
	// Stores the fixed color value if mode is fixed
	FixedColor *string `json:"fixedColor,omitempty"`

	// The main color scheme mode
	Mode string `json:"mode"`

	// TODO docs
	SeriesBy *SpecFieldColorSeriesByMode `json:"seriesBy,omitempty"`
}

// TODO docs
type SpecFieldColorModeId string

// TODO docs
type SpecFieldColorSeriesByMode string

// SpecFieldConfig defines model for spec.#FieldConfig.
type SpecFieldConfig struct {
	// TODO docs
	Color *SpecFieldColor `json:"color,omitempty"`

	// custom is specified by the PanelFieldConfig field
	// in panel plugin schemas.
	Custom map[string]interface{} `json:"custom,omitempty"`

	// Significant digits (for display)
	Decimals *float32 `json:"decimals,omitempty"`

	// Human readable field metadata
	Description *string `json:"description,omitempty"`

	// The display value for this field.  This supports template variables blank is auto
	DisplayName *string `json:"displayName,omitempty"`

	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	DisplayNameFromDS *string `json:"displayNameFromDS,omitempty"`

	// True if data source field supports ad-hoc filters
	Filterable *bool `json:"filterable,omitempty"`

	// The behavior when clicking on a result
	Links []interface{} `json:"links,omitempty"`

	// Convert input values into a display string
	Mappings []interface{} `json:"mappings,omitempty"`
	Max      *float32      `json:"max,omitempty"`
	Min      *float32      `json:"min,omitempty"`

	// Alternative to empty string
	NoValue *string `json:"noValue,omitempty"`

	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	//
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	Path       *string               `json:"path,omitempty"`
	Thresholds *SpecThresholdsConfig `json:"thresholds,omitempty"`

	// Numeric Options
	Unit *string `json:"unit,omitempty"`

	// True if data source can write a value to the path.  Auth/authz are supported separately
	Writeable *bool `json:"writeable,omitempty"`
}

// SpecFieldConfigSource defines model for spec.#FieldConfigSource.
type SpecFieldConfigSource struct {
	Defaults  SpecFieldConfig `json:"defaults"`
	Overrides []struct {
		Matcher    SpecMatcherConfig        `json:"matcher"`
		Properties []SpecDynamicConfigValue `json:"properties"`
	} `json:"overrides"`
}

// Support for legacy graph and heatmap panels.
type SpecGraphPanel struct {
	// @deprecated this is part of deprecated graph panel
	Legend *struct {
		Show     bool    `json:"show"`
		Sort     *string `json:"sort,omitempty"`
		SortDesc *bool   `json:"sortDesc,omitempty"`
	} `json:"legend,omitempty"`
	Type SpecGraphPanelType `json:"type"`
}

// SpecGraphPanelType defines model for SpecGraphPanel.Type.
type SpecGraphPanelType string

// SpecGridPos defines model for spec.#GridPos.
type SpecGridPos struct {
	// H Panel
	H int `json:"h"`

	// Static true if fixed
	Static *bool `json:"static,omitempty"`

	// W Panel
	W int `json:"w"`

	// Panel x
	X int `json:"x"`

	// Panel y
	Y int `json:"y"`
}

// SpecHeatmapPanel defines model for spec.#HeatmapPanel.
type SpecHeatmapPanel struct {
	Type SpecHeatmapPanelType `json:"type"`
}

// SpecHeatmapPanelType defines model for SpecHeatmapPanel.Type.
type SpecHeatmapPanelType string

// SpecLibraryPanelRef defines model for spec.#LibraryPanelRef.
type SpecLibraryPanelRef struct {
	Name string `json:"name"`
	Uid  string `json:"uid"`
}

// SpecLoadingState defines model for spec.#LoadingState.
type SpecLoadingState string

// TODO docs
type SpecMappingType string

// SpecMatcherConfig defines model for spec.#MatcherConfig.
type SpecMatcherConfig struct {
	Id      string       `json:"id"`
	Options *interface{} `json:"options,omitempty"`
}

// Dashboard panels. Panels are canonically defined inline
// because they share a version timeline with the dashboard
// schema; they do not evolve independently.
type SpecPanel struct {
	// The datasource used in all targets.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`

	// Description Description.
	Description *string               `json:"description,omitempty"`
	FieldConfig SpecFieldConfigSource `json:"fieldConfig"`
	GridPos     *SpecGridPos          `json:"gridPos,omitempty"`

	// TODO docs
	Id *int `json:"id,omitempty"`

	// TODO docs
	// TODO tighter constraint
	Interval     *string              `json:"interval,omitempty"`
	LibraryPanel *SpecLibraryPanelRef `json:"libraryPanel,omitempty"`

	// Panel links.
	// TODO fill this out - seems there are a couple variants?
	Links []SpecDashboardLink `json:"links,omitempty"`

	// TODO docs
	MaxDataPoints *float32 `json:"maxDataPoints,omitempty"`

	// options is specified by the PanelOptions field in panel
	// plugin schemas.
	Options map[string]interface{} `json:"options"`

	// FIXME this almost certainly has to be changed in favor of scuemata versions
	PluginVersion *string `json:"pluginVersion,omitempty"`

	// Name of template variable to repeat for.
	Repeat *string `json:"repeat,omitempty"`

	// Direction to repeat in if 'repeat' is set.
	// "h" for horizontal, "v" for vertical.
	// TODO this is probably optional
	RepeatDirection SpecPanelRepeatDirection `json:"repeatDirection"`

	// Id of the repeating panel.
	RepeatPanelId *int64 `json:"repeatPanelId,omitempty"`

	// TODO docs
	Tags []string `json:"tags,omitempty"`

	// TODO docs
	Targets []SpecTarget `json:"targets,omitempty"`

	// TODO docs - seems to be an old field from old dashboard alerts?
	Thresholds []interface{} `json:"thresholds,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeFrom *string `json:"timeFrom,omitempty"`

	// TODO docs
	TimeRegions []interface{} `json:"timeRegions,omitempty"`

	// TODO docs
	// TODO tighter constraint
	TimeShift *string `json:"timeShift,omitempty"`

	// Panel title.
	Title           *string                     `json:"title,omitempty"`
	Transformations []SpecDataTransformerConfig `json:"transformations"`

	// Whether to display the panel without a background.
	Transparent bool `json:"transparent"`

	// The panel plugin type id. May not be empty.
	Type string `json:"type"`
}

// Direction to repeat in if 'repeat' is set.
// "h" for horizontal, "v" for vertical.
// TODO this is probably optional
type SpecPanelRepeatDirection string

// TODO docs
type SpecRangeMap struct {
	Options struct {
		// From to and from are `number | null` in current ts, really not sure what to do
		From float64 `json:"from"`

		// TODO docs
		Result SpecValueMappingResult `json:"result"`
		To     float64                `json:"to"`
	} `json:"options"`
	Type SpecRangeMapType `json:"type"`
}

// SpecRangeMapType defines model for SpecRangeMap.Type.
type SpecRangeMapType string

// TODO docs
type SpecRegexMap struct {
	Options struct {
		Pattern string `json:"pattern"`

		// TODO docs
		Result SpecValueMappingResult `json:"result"`
	} `json:"options"`
	Type SpecRegexMapType `json:"type"`
}

// SpecRegexMapType defines model for SpecRegexMap.Type.
type SpecRegexMapType string

// Row panel
type SpecRowPanel struct {
	Collapsed bool `json:"collapsed"`

	// Name of default datasource.
	Datasource *struct {
		Type *string `json:"type,omitempty"`
		Uid  *string `json:"uid,omitempty"`
	} `json:"datasource,omitempty"`
	GridPos *SpecGridPos  `json:"gridPos,omitempty"`
	Id      int           `json:"id"`
	Panels  []interface{} `json:"panels"`

	// Name of template variable to repeat for.
	Repeat *string          `json:"repeat,omitempty"`
	Title  *string          `json:"title,omitempty"`
	Type   SpecRowPanelType `json:"type"`
}

// SpecRowPanelType defines model for SpecRowPanel.Type.
type SpecRowPanelType string

// TODO docs
type SpecSnapshot struct {
	// TODO docs
	Created string `json:"created"`

	// TODO docs
	Expires string `json:"expires"`

	// TODO docs
	External bool `json:"external"`

	// TODO docs
	ExternalUrl string `json:"externalUrl"`

	// TODO docs
	Id int `json:"id"`

	// TODO docs
	Key string `json:"key"`

	// TODO docs
	Name string `json:"name"`

	// TODO docs
	OrgId int `json:"orgId"`

	// TODO docs
	Updated string `json:"updated"`

	// TODO docs
	Url *string `json:"url,omitempty"`

	// TODO docs
	UserId int `json:"userId"`
}

// TODO docs
type SpecSpecialValueMap struct {
	Options struct {
		Match   SpecSpecialValueMapOptionsMatch `json:"match"`
		Pattern string                          `json:"pattern"`

		// TODO docs
		Result SpecValueMappingResult `json:"result"`
	} `json:"options"`
	Type SpecSpecialValueMapType `json:"type"`
}

// SpecSpecialValueMapOptionsMatch defines model for SpecSpecialValueMap.Options.Match.
type SpecSpecialValueMapOptionsMatch string

// SpecSpecialValueMapType defines model for SpecSpecialValueMap.Type.
type SpecSpecialValueMapType string

// TODO docs
type SpecSpecialValueMatch string

// Schema for panel targets is specified by datasource
// plugins. We use a placeholder definition, which the Go
// schema loader either left open/as-is with the Base
// variant of the Dashboard and Panel families, or filled
// with types derived from plugins in the Instance variant.
// When working directly from CUE, importers can extend this
// type directly to achieve the same effect.
type SpecTarget = map[string]interface{}

// TODO docs
type SpecThreshold struct {
	// TODO docs
	Color string `json:"color"`

	// Threshold index, an old property that is not needed an should only appear in older dashboards
	Index *int32 `json:"index,omitempty"`

	// TODO docs
	// TODO are the values here enumerable into a disjunction?
	// Some seem to be listed in typescript comment
	State *string `json:"state,omitempty"`

	// TODO docs
	// FIXME the corresponding typescript field is required/non-optional, but nulls currently appear here when serializing -Infinity to JSON
	Value *float32 `json:"value,omitempty"`
}

// SpecThresholdsConfig defines model for spec.#ThresholdsConfig.
type SpecThresholdsConfig struct {
	Mode SpecThresholdsMode `json:"mode"`

	// Must be sorted by 'value', first value is always -Infinity
	Steps []SpecThreshold `json:"steps"`
}

// SpecThresholdsMode defines model for spec.#ThresholdsMode.
type SpecThresholdsMode string

// TODO docs
type SpecValueMap struct {
	Options map[string]SpecValueMappingResult `json:"options"`
	Type    SpecValueMapType                  `json:"type"`
}

// SpecValueMapType defines model for SpecValueMap.Type.
type SpecValueMapType string

// TODO docs
type SpecValueMappingResult struct {
	Color *string `json:"color,omitempty"`
	Icon  *string `json:"icon,omitempty"`
	Index *int32  `json:"index,omitempty"`
	Text  *string `json:"text,omitempty"`
}

// SpecVariableHide defines model for spec.#VariableHide.
type SpecVariableHide int

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO what about what's in public/app/features/types.ts?
// TODO there appear to be a lot of different kinds of [template] vars here? if so need a disjunction
type SpecVariableModel struct {
	// Ref to a DataSource instance
	Datasource  *SpecDataSourceRef     `json:"datasource,omitempty"`
	Description *string                `json:"description,omitempty"`
	Error       map[string]interface{} `json:"error,omitempty"`
	Global      bool                   `json:"global"`
	Hide        SpecVariableHide       `json:"hide"`
	Id          string                 `json:"id"`
	Index       int                    `json:"index"`
	Label       *string                `json:"label,omitempty"`
	Name        string                 `json:"name"`

	// TODO: Move this into a separated QueryVariableModel type
	Query        *interface{}     `json:"query,omitempty"`
	RootStateKey *string          `json:"rootStateKey,omitempty"`
	SkipUrlSync  bool             `json:"skipUrlSync"`
	State        SpecLoadingState `json:"state"`

	// FROM: packages/grafana-data/src/types/templateVars.ts
	// TODO docs
	// TODO this implies some wider pattern/discriminated union, probably?
	Type SpecVariableType `json:"type"`
}

// FROM: packages/grafana-data/src/types/templateVars.ts
// TODO docs
// TODO this implies some wider pattern/discriminated union, probably?
type SpecVariableType string
