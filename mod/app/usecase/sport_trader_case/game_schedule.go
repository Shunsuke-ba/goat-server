package sport_trader_case

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	domain "github.com/Shunsuke-ba/goat-server/mod/app/domain/sport_trader_domain"
	"github.com/Shunsuke-ba/goat-server/mod/util/config"
)

type GameSchedules func(ctx context.Context) (gameSchedules domain.GameSchedules, err error)

func ProvideGameShedulesCase(
	ctx context.Context,
) GameSchedules {

	return func(ctx context.Context) (gameSchedule domain.GameSchedules, err error) {
		url := config.BASEURL + path.Join(
			config.BASETBALL+"-"+config.ACCESSLEVEL_TRIAL+config.VERSION1,
			config.LANGUAGE_CODE,
			// 今の場合
			config.TimeNow(),
			config.SCHEDULEFORMAT+"?api_key="+os.Getenv("BASKETBALL_KEY"),
		)
		resp, err := http.Get(url)
		if err != nil {
			err = errors.New("ProvideGameScheduleCase http.Get fail")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.New("ProvideGameScheduleCase ioutil.ReadAll fail")
			return
		}

		if err = json.Unmarshal(body, &gameSchedule); err != nil {
			return
		}

		gameSchedule = gameSchedule.JapanValidation()

		return
	}
}
