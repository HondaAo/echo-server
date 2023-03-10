package responses

import (
	idiomEntity "github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
	scriptIdiomsEntity "github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
	scriptEntity "github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
)

type VideoResponse struct {
	Video   Video    `json:"video"`
	Scripts []Script `json:"scripts"`
}

func NewResponse(video *entity.Video, scripts []*scriptEntity.Script, scriptIdioms []*scriptIdiomsEntity.ScriptIdioms, idioms []*idiomEntity.Idioms) VideoResponse {
	videoResponse := NewVideoResponse(video)

	idiomsMap := make(map[string]Idioms)
	for _, i := range idioms {
		idiomsMap[i.Idiom] = Idioms{
			Word:    i.Idiom,
			Meaning: i.Meaning,
			Level:   getLevel(i.Level),
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

	scriptIdiomsMap := make(map[uint64][]ScriptIdioms)
	for _, sim := range scriptIdiomsResponse {
		scriptIdiomsMap[sim.ScriptID] = append(scriptIdiomsMap[sim.ScriptID], sim)
	}

	scriptResponses := make([]Script, 0, len(scripts))
	for _, s := range scripts {
		scriptResponses = append(scriptResponses, Script{
			VideoID:   s.VideoID,
			ScriptID:  s.ScriptID,
			Text:      s.Text,
			Japanese:  s.Japanese,
			TimeStamp: s.TimeStamp,
			Idioms:    scriptIdiomsMap[s.ScriptID],
			CreatedAt: s.CreatedAt.Format("2006-01-02"),
			UpdatedAt: s.UpdatedAt.Format("2006-01-02"),
		})
	}

	return VideoResponse{
		Video:   *videoResponse,
		Scripts: scriptResponses,
	}
}
