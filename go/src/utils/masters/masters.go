package masters

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/bluele/gcache"
)

//go:generate go-bindata -pkg masters -out leaderboard.go ./files/leaderboard.tmpl.html
//go:generate go-bindata -pkg masters -out styles.go ./files/styles.css
//go:generate go-bindata -pkg masters -out logo.go ./files/logo.png

var leaderBoard = template.Must(template.New("leaderboard").Funcs(map[string]interface{}{
	"add": func(a, b int) int { return a + b },
}).Parse(string(files_leaderboard_tmpl_html())))

// Cache for a 30 seconds here to avoid any sort of detection by pga (if it exists).
var tmplCache = gcache.New(10).Expiration(time.Minute).Build()

func RenderTmpl(leaderboard *LeaderBoardResp) ([]byte, error) {
	v, err := tmplCache.Get("leaderboard")
	if err == gcache.KeyNotFoundError {
		// Think this cache guards contestants from race when updating...
		// Wow this is shitty programming.
		if err := OrderContestants(leaderboard.Leaderboard.Players, Contestants); err != nil {
			return nil, err
		}

		mst, err := time.LoadLocation("America/Denver")
		if err != nil {
			return nil, err
		}

		args := &LeaderBoardArgs{
			Contestants: Contestants,
			LastUpdated: time.Now().In(mst).Format(time.RFC822),
		}

		b := &bytes.Buffer{}
		if err := leaderBoard.Execute(b, args); err != nil {
			return nil, err
		}

		bytes := b.Bytes()
		tmplCache.Set("leaderboard", bytes)

		return bytes, nil
	}

	return v.([]byte), nil

}

func GetStyles() io.ReadSeeker {
	return bytes.NewReader(files_styles_css())
}

func GetLogo() io.ReadSeeker {
	return bytes.NewReader(files_logo_png())
}

func GetLeaderboard() (*LeaderBoardResp, error) {
	resp, err := http.Get("https://statdata.pgatour.com/r/014/leaderboard-v2mini.json")
	if err != nil {
		return nil, err
	}

	//f, err := os.Create("~/Desktop/masters.json")
	//if err != nil {
	//	panic(err)
	//}
	//tee := io.TeeReader(resp.Body, f) // Remove this

	leaderBoard := LeaderBoardResp{}
	if err := json.NewDecoder(resp.Body).Decode(&leaderBoard); err != nil {
		return nil, err
	}

	//leaderBoard := LeaderBoardResp{}
	//if err := fsio.ReadAndLockFileJSON("/Users/hdh/Desktop/masters.json", &leaderBoard); err != nil {
	//	return nil, err
	//}

	return &leaderBoard, nil
}
