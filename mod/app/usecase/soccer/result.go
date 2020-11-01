package soccer

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	domain "github.com/Shunsuke-ba/goat-server/mod/app/domain/soccer"
	"github.com/Shunsuke-ba/goat-server/mod/util/config"
)

type Results func(ctx context.Context) (gameResults domain.SoccerResults, err error)

func ProvideResultsCase(
	ctx context.Context,
) Results {

	return func(ctx context.Context) (results domain.SoccerResults, err error) {
		url := config.BASEURL + path.Join(
			config.SOCCER+"-"+config.ACCESSLEVEL_TRIAL+config.VERSION3,
			config.ASIA,
			config.LANGUAGE_CODE,
			// 今の場合
			config.TimeNow(),
			config.RESULTFORMAT+"?api_key="+os.Getenv("SOCCER_KEY"),
		)
		resp, err := http.Get(url)
		if err != nil {
			err = errors.New("ProvideGameResultsCase http.Get fail")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.New("ProvideGameResultsCase ioutil.ReadAll fail")
			return
		}

		if err = json.Unmarshal(body, &results); err != nil {
			err = errors.New("ProvideGameResultsCase json.Unmarshal failed")
			return
		}

		//results = results.JapanValidation()

		return
	}
}
