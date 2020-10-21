package sport_trader_case

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	domain "github.com/Shunsuke-ba/goat-server/mod/app/domain/sport_trader_domain"
	"github.com/Shunsuke-ba/goat-server/mod/util/config"
)

type HeadToHeads func(ctx context.Context, team1, team2 string) (headToHead domain.HeadToHeads, err error)

func ProvideHeadToHeadsCase(
	ctx context.Context,
) HeadToHeads {

	return func(ctx context.Context, team1ID, team2ID string) (headToHeads domain.HeadToHeads, err error) {
		url := config.BASEURL + path.Join(
			config.BASETBALL+"-"+config.ACCESSLEVEL_TRIAL+config.VERSION1,
			config.LANGUAGE_CODE,
			// 今の場合
			config.TEAMS,
			team1ID,
			config.VERSUS,
			team2ID,
			config.MATCHESFORMAT+"?api_key="+os.Getenv("BASKETBALL_KEY"),
		)
		/*"sr:competitor:124568",
		"sr:competitor:124582",*/
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		if err = json.Unmarshal(body, &headToHeads); err != nil {
			return
		}

		return
	}
}
