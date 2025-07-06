# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /v1/version">client.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#ProviderLookupResponse">ProviderLookupResponse</a>

Methods:

- <code title="get /v1/providers/lookup">client.Providers.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#ProviderService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#ProviderLookupParams">ProviderLookupParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#ProviderLookupResponse">ProviderLookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Database

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#DatabaseCheckHealthResponse">DatabaseCheckHealthResponse</a>

Methods:

- <code title="get /v1/database/health">client.Database.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#DatabaseService.CheckHealth">CheckHealth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#DatabaseCheckHealthResponse">DatabaseCheckHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Fax

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxListProvidersResponse">FaxListProvidersResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxGetStatusResponse">FaxGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxSendResponse">FaxSendResponse</a>

Methods:

- <code title="get /v1/fax/providers">client.Fax.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxService.ListProviders">ListProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxListProvidersResponse">FaxListProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fax/status/{id}">client.Fax.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxGetStatusResponse">FaxGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/fax/send">client.Fax.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxSendParams">FaxSendParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go">bluehive</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bluehive-go#FaxSendResponse">FaxSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
