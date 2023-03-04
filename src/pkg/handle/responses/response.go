package responses

import (
	idiomEntity "github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
	scriptIdiomsEntity "github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
	scriptEntity "github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
)

type VideoResponse struct {
	Video        Video          `json:"video"`
	Scripts      []Script       `json:"scripts"`
	ScriptIdioms []ScriptIdioms `json:"script_idioms"`
}

func NewResponse(video *entity.Video, scripts []*scriptEntity.Script, scriptIdioms []*scriptIdiomsEntity.ScriptIdioms, idioms []*idiomEntity.Idioms) VideoResponse {
	level := getVideoLevel(video.Level)
	videoResponse := Video{
		VideoID:   video.VideoID,
		URL:       video.URL,
		Start:     video.Start,
		End:       video.End,
		Level:     level,
		Display:   video.Display,
		CreatedAt: video.CreatedAt.Format("2006-01-02"),
		UpdatedAt: video.UpdatedAt.Format("2006-01-02"),
	}

	scriptResponses := make([]Script, 0, len(scripts))
	for _, s := range scripts {
		scriptResponses = append(scriptResponses, Script{
			VideoID:   s.VideoID,
			ScriptID:  s.ScriptID,
			Text:      s.Text,
			Japanese:  s.Japanese,
			TimeStamp: s.TimeStamp,
			CreatedAt: s.CreatedAt.Format("2006-01-02"),
			UpdatedAt: s.UpdatedAt.Format("2006-01-02"),
		})
	}

	idiomsMap := make(map[string]Idioms)
	for _, i := range idioms {
		idiomsMap[i.Idiom] = Idioms{
			Word:    i.Idiom,
			Meaning: i.Meaning,
			Level:   uint64(i.Level),
		}
	}

	scriptIdiomsResponse := make([]ScriptIdioms, 0, len(scriptIdioms))
	for _, si := range scriptIdioms {
		scriptIdiomsResponse = append(scriptIdiomsResponse, ScriptIdioms{
			Idiom:    idiomsMap[si.Idiom],
			VideoID:  si.VideoID,
			ScriptID: si.ScriptID,
		})
	}

	return VideoResponse{
		Video:        videoResponse,
		Scripts:      scriptResponses,
		ScriptIdioms: scriptIdiomsResponse,
	}
}

func getVideoLevel(intLevel entity.VideoLevel) string {
	switch intLevel {
	case 1:
		return "A1"
	case 2:
		return "A2"
	case 3:
		return "B1"
	case 4:
		return "B2"
	case 5:
		return "C1"
	default:
		return "F"
	}
}
