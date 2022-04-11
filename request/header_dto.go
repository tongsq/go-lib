package request

type HeaderDto struct {
	Referer                 string
	UserAgent               string
	Host                    string
	UpgradeInsecureRequests string
	Accept                  string
	AcceptEncoding          string
	AcceptLanguage          string
	SecFetchDest            string
	SecFetchMode            string
	XRequestedWith          string
	ContentType             string
	Other                   map[string]string
}
