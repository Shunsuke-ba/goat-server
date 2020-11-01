package basketball

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	domain "github.com/Shunsuke-ba/goat-server/mod/app/domain/basketball"
	"github.com/Shunsuke-ba/goat-server/mod/util/config"
)

type Schedules func(ctx context.Context) (schedules domain.BasketballSchedules, err error)

func ProvideShedulesCase(
	ctx context.Context,
) Schedules {

	return func(ctx context.Context) (schedules domain.BasketballSchedules, err error) {
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

		if err = json.Unmarshal(body, &schedules); err != nil {
			return
		}

		schedules = schedules.JapanValidation()

		return
	}
}
