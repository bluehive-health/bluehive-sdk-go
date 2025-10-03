# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /v1/version">client.Version.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#ProviderLookupResponse">ProviderLookupResponse</a>

Methods:

- <code title="get /v1/providers/lookup">client.Providers.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#ProviderService.Lookup">Lookup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#ProviderLookupParams">ProviderLookupParams</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#ProviderLookupResponse">ProviderLookupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Database

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#DatabaseCheckHealthResponse">DatabaseCheckHealthResponse</a>

Methods:

- <code title="get /v1/database/health">client.Database.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#DatabaseService.CheckHealth">CheckHealth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#DatabaseCheckHealthResponse">DatabaseCheckHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Fax

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxListProvidersResponse">FaxListProvidersResponse</a>
- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxGetStatusResponse">FaxGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxSendResponse">FaxSendResponse</a>

Methods:

- <code title="get /v1/fax/providers">client.Fax.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxService.ListProviders">ListProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxListProvidersResponse">FaxListProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fax/status/{id}">client.Fax.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxGetStatusResponse">FaxGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/fax/send">client.Fax.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxSendParams">FaxSendParams</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#FaxSendResponse">FaxSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Employers

Response Types:

- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerNewResponse">EmployerNewResponse</a>
- <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerGetResponse">EmployerGetResponse</a>

Methods:

- <code title="post /v1/employers">client.Employers.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerNewParams">EmployerNewParams</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerNewResponse">EmployerNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/employers/{employerId}">client.Employers.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, employerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#EmployerGetResponse">EmployerGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Hl7

Methods:

- <code title="post /v1/hl7/">client.Hl7.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#Hl7Service.Process">Process</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go">githubcombluehivehealthbluehivesdkgo</a>.<a href="https://pkg.go.dev/github.com/bluehive-health/bluehive-sdk-go#Hl7ProcessParams">Hl7ProcessParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
