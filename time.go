package lemmy

import (
	"encoding/json"
	"time"
)

type LemmyTime struct {
	time.Time
}

func (lt LemmyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(lt.Time.Format("2006-01-02T15:04:05"))
}

func (lt *LemmyTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return err
	}

	if timeStr == "" {
		lt.Time = time.Unix(0, 0)
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", timeStr)
	if err != nil {
		t, err = time.Parse(time.RFC3339, timeStr)
		if err != nil {
			return err
		}
	}

	lt.Time = t
	return nil
}
