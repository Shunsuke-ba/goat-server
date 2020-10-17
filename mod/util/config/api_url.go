package config

import "time"

const (
	//URL = "https://api.sportradar.com/basketball-{access_level}{version}/{language_code}/schedules/{year}-{month}-{day}/results.{format}?api_key={your_api_key}"
	//URL = "https://api.sportradar.com/basketball-t1/ja/schedules/2020-10-17/results.json?api_key=yzsq7fh3jpzpamrf6yy5jm8r"
	BASEURL = "https://api.sportradar.com/"
	BASETBALL = "basketball"
	ACCESSLEVEL_TRIAL = "t"
	VERSION1 = "1"
	LANGUAGE_CODE = "ja"
	RESULTFORMAT = "results.json"
)

func TimeNow() string {
	return "schedules/" + time.Now().Format("2006-01-02")
}