package swagger

import (
	"html/template"
)

// Config stores SwaggerUI configuration variables
type Config struct {
	// This parameter can be used to name different swagger document instances.
	// default: ""
	InstanceName string `json:"-"`

	// Title pointing to title of HTML page.
	// default: "Swagger UI"
	Title string `json:"-"`

	// URL to fetch external configuration document from.
	// default: ""
	ConfigURL string `json:"configUrl,omitempty"`

	// The URL pointing to API definition (normally swagger.json or swagger.yaml).
	// default: "doc.json"
	URL string `json:"url,omitempty"`

	// Enables overriding configuration parameters via URL search params.
	// default: false
	QueryConfigEnabled bool `json:"queryConfigEnabled,omitempty"`

	// The name of a component available via the plugin system to use as the top-level layout for Swagger UI.
	// default: "StandaloneLayout"
	Layout string `json:"layout,omitempty"`

	// An array of plugin functions to use in Swagger UI.
	// default: [SwaggerUIBundle.plugins.DownloadUrl]
	Plugins []template.JS `json:"-"`

	// An array of presets to use in Swagger UI. Usually, you'll want to include ApisPreset if you use this option.
	// default: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset]
	Presets []template.JS `json:"-"`

	// If set to true, enables deep linking for tags and operations.
	// default: true
	DeepLinking bool `json:"deepLinking"`

	// Controls the display of operationId in operations list.
	// default: false
	DisplayOperationId bool `json:"displayOperationId,omitempty"`

	// The default expansion depth for models (set to -1 completely hide the models).
	// default: 1
	DefaultModelsExpandDepth int `json:"defaultModelsExpandDepth,omitempty"`

	// The default expansion depth for the model on the model-example section.
	// default: 1
	DefaultModelExpandDepth int `json:"defaultModelExpandDepth,omitempty"`

	// Controls how the model is shown when the API is first rendered.
	// The user can always switch the rendering for a given model by clicking the 'Model' and 'Example Value' links.
	// default: "example"
	DefaultModelRendering string `json:"defaultModelRendering,omitempty"`

	// Controls the display of the request duration (in milliseconds) for "Try it out" requests.
	// default: false
	DisplayRequestDuration bool `json:"displayRequestDuration,omitempty"`

	// Controls the default expansion setting for the operations and tags.
	// 'list' (default, expands only the tags),
	// 'full' (expands the tags and operations),
	// 'none' (expands nothing)
	DocExpansion string `json:"docExpansion,omitempty"`

	// If set, enables filtering. The top bar will show an edit box that you can use to filter the tagged operations that are shown.
	// Can be Boolean to enable or disable, or a string, in which case filtering will be enabled using that string as the filter expression.
	// Filtering is case sensitive matching the filter expression anywhere inside the tag.
	// default: false
	Filter FilterConfig `json:"-"`

	// If set, limits the number of tagged operations displayed to at most this many. The default is to show all operations.
	// default: 0
	MaxDisplayedTags int `json:"maxDisplayedTags,omitempty"`

	// Controls the display of vendor extension (x-) fields and values for Operations, Parameters, Responses, and Schema.
	// default: false
	ShowExtensions bool `json:"showExtensions,omitempty"`

	// Controls the display of extensions (pattern, maxLength, minLength, maximum, minimum) fields and values for Parameters.
	// default: false
	ShowCommonExtensions bool `json:"showCommonExtensions,omitempty"`

	// Apply a sort to the tag list of each API. It can be 'alpha' (sort by paths alphanumerically) or a function (see Array.prototype.sort().
	// to learn how to write a sort function). Two tag name strings are passed to the sorter for each pass.
	// default: "" -> Default is the order determined by Swagger UI.
	TagsSorter template.JS `json:"-"`

	// Provides a mechanism to be notified when Swagger UI has finished rendering a newly provided definition.
	// default: "" -> Function=NOOP
	OnComplete template.JS `json:"-"`

	// An object with the activate and theme properties.
	SyntaxHighlight *SyntaxHighlightConfig `json:"-"`

	// Controls whether the "Try it out" section should be enabled by default.
	// default: false
	TryItOutEnabled bool `json:"tryItOutEnabled,omitempty"`

	// Enables the request snippet section. When disabled, the legacy curl snippet will be used.
	// default: false
	RequestSnippetsEnabled bool `json:"requestSnippetsEnabled,omitempty"`

	// OAuth redirect URL.
	// default: ""
	OAuth2RedirectUrl string `json:"oauth2RedirectUrl,omitempty"`

	// MUST be a function. Function to intercept remote definition, "Try it out", and OAuth 2.0 requests.
	// Accepts one argument requestInterceptor(request) and must return the modified request, or a Promise that resolves to the modified request.
	// default: ""
	RequestInterceptor template.JS `json:"-"`

	// If set, MUST be an array of command line options available to the curl command. This can be set on the mutated request in the requestInterceptor function.
	// For example request.curlOptions = ["-g", "--limit-rate 20k"]
	// default: nil
	RequestCurlOptions []string `json:"request.curlOptions,omitempty"`

	// MUST be a function. Function to intercept remote definition, "Try it out", and OAuth 2.0 responses.
	// Accepts one argument responseInterceptor(response) and must return the modified response, or a Promise that resolves to the modified response.
	// default: ""
	ResponseInterceptor template.JS `json:"-"`

	// If set to true, uses the mutated request returned from a requestInterceptor to produce the curl command in the UI,
	// otherwise the request before the requestInterceptor was applied is used.
	// default: true
	ShowMutatedRequest bool `json:"showMutatedRequest"`

	// List of HTTP methods that have the "Try it out" feature enabled. An empty array disables "Try it out" for all operations.
	// This does not filter the operations from the display.
	// Possible values are ["get", "put", "post", "delete", "options", "head", "patch", "trace"]
	// default: nil
	SupportedSubmitMethods []string `json:"supportedSubmitMethods,omitempty"`

	// By default, Swagger UI attempts to validate specs against swagger.io's online validator. You can use this parameter to set a different validator URL.
	// For example for locally deployed validators (https://github.com/swagger-api/validator-badge).
	// Setting it to either none, 127.0.0.1 or localhost will disable validation.
	// default: ""
	ValidatorUrl string `json:"validatorUrl,omitempty"`

	// If set to true, enables passing credentials, as defined in the Fetch standard, in CORS requests that are sent by the browser.
	// Note that Swagger UI cannot currently set cookies cross-domain (see https://github.com/swagger-api/swagger-js/issues/1163).
	// as a result, you will have to rely on browser-supplied cookies (which this setting enables sending) that Swagger UI cannot control.
	// default: false
	WithCredentials bool `json:"withCredentials,omitempty"`

	// Function to set default values to each property in model. Accepts one argument modelPropertyMacro(property), property is immutable.
	// default: ""
	ModelPropertyMacro template.JS `json:"-"`

	// Function to set default value to parameters. Accepts two arguments parameterMacro(operation, parameter).
	// Operation and parameter are objects passed for context, both remain immutable.
	// default: ""
	ParameterMacro template.JS `json:"-"`

	// If set to true, it persists authorization data and it would not be lost on browser close/refresh.
	// default: false
	PersistAuthorization bool `json:"persistAuthorization,omitempty"`

	// Configuration information for OAuth2, optional if using OAuth2
	OAuth *OAuthConfig `json:"-"`

	// (authDefinitionKey, username, password) => action
	// Programmatically set values for a Basic authorization scheme.
	// default: ""
	PreauthorizeBasic template.JS `json:"-"`

	// (authDefinitionKey, apiKeyValue) => action
	// Programmatically set values for an API key or Bearer authorization scheme.
	// In case of OpenAPI 3.0 Bearer scheme, apiKeyValue must contain just the token itself without the Bearer prefix.
	// default: ""
	PreauthorizeApiKey template.JS `json:"-"`

	// Applies custom CSS styles.
	// default: ""
	CustomStyle template.CSS `json:"-"`

	// Applies custom JavaScript scripts.
	// default ""
	CustomScript template.JS `json:"-"`
}

type FilterConfig struct {
	Enabled    bool
	Expression string
}

func (fc FilterConfig) Value() interface{} {
	if fc.Expression != "" {
		return fc.Expression
	}
	return fc.Enabled
}

type SyntaxHighlightConfig struct {
	// Whether syntax highlighting should be activated or not.
	// default: true
	Activate bool `json:"activate"`
	// Highlight.js syntax coloring theme to use.
	// Possible values are ["agate", "arta", "monokai", "nord", "obsidian", "tomorrow-night"]
	// default: "agate"
	Theme string `json:"theme,omitempty"`
}

func (shc SyntaxHighlightConfig) Value() interface{} {
	if shc.Activate {
		return shc
	}
	return false
}

type OAuthConfig struct {
	// ID of the client sent to the OAuth2 provider.
	// default: ""
	ClientId string `json:"clientId,omitempty"`

	// Never use this parameter in your production environment.
	// It exposes cruicial security information. This feature is intended for dev/test environments only.
	// Secret of the client sent to the OAuth2 provider.
	// default: ""
	ClientSecret string `json:"clientSecret,omitempty"`

	// Application name, displayed in authorization popup.
	// default: ""
	AppName string `json:"appName,omitempty"`

	// Realm query parameter (for oauth1) added to authorizationUrl and tokenUrl.
	// default: ""
	Realm string `json:"realm,omitempty"`

	// String array of initially selected oauth scopes
	// default: nil
	Scopes []string `json:"scopes,omitempty"`

	// Additional query parameters added to authorizationUrl and tokenUrl.
	// default: nil
	AdditionalQueryStringParams map[string]string `json:"additionalQueryStringParams,omitempty"`

	// Unavailable	Only activated for the accessCode flow.
	// During the authorization_code request to the tokenUrl, pass the Client Password using the HTTP Basic Authentication scheme
	// (Authorization header with Basic base64encode(client_id + client_secret)).
	// default: false
	UseBasicAuthenticationWithAccessCodeGrant bool `json:"useBasicAuthenticationWithAccessCodeGrant,omitempty"`

	// Only applies to authorizatonCode flows.
	// Proof Key for Code Exchange brings enhanced security for OAuth public clients.
	// default: false
	UsePkceWithAuthorizationCodeGrant bool `json:"usePkceWithAuthorizationCodeGrant,omitempty"`
}

var (
	ConfigDefault = Config{
		Title:  "Swagger UI",
		Layout: "StandaloneLayout",
		Plugins: []template.JS{
			template.JS("SwaggerUIBundle.plugins.DownloadUrl"),
		},
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
		DeepLinking:              true,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DefaultModelRendering:    "example",
		DocExpansion:             "list",
		SyntaxHighlight: &SyntaxHighlightConfig{
			Activate: true,
			Theme:    "agate",
		},
		ShowMutatedRequest: true,
	}
)

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.Title == "" {
		cfg.Title = ConfigDefault.Title
	}

	if cfg.Layout == "" {
		cfg.Layout = ConfigDefault.Layout
	}

	if cfg.DefaultModelRendering == "" {
		cfg.DefaultModelRendering = ConfigDefault.DefaultModelRendering
	}

	if cfg.DocExpansion == "" {
		cfg.DocExpansion = ConfigDefault.DocExpansion
	}

	if cfg.Plugins == nil {
		cfg.Plugins = ConfigDefault.Plugins
	}

	if cfg.Presets == nil {
		cfg.Presets = ConfigDefault.Presets
	}

	if cfg.SyntaxHighlight == nil {
		cfg.SyntaxHighlight = ConfigDefault.SyntaxHighlight
	}

	return cfg
}
