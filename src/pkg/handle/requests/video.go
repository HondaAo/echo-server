package requests

import (
	scriptIdiomsEntity "github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
	scriptEntity "github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
	"github.com/google/uuid"
)

type VideoRequest struct {
	Title   string          `json:"title"`
	URL     string          `json:"url"`
	Start   float64         `json:"start"`
	End     float64         `json:"end"`
	Level   uint64          `json:"level"`
	Scripts []ScriptRequest `json:"scripts"`
}

func (v *VideoRequest) NewEntity() (*entity.Video, []scriptEntity.Script, []scriptIdiomsEntity.ScriptIdioms, error) {
	videoID := NewUUID()

	video := &entity.Video{
		VideoID: videoID.String(),
		Title:   v.Title,
		URL:     v.URL,
		Start:   v.Start,
		End:     v.End,
		Level:   entity.VideoLevel(v.Level),
	}

	scriptEntities := make([]scriptEntity.Script, 0, len(v.Scripts))
	scriptIdioms := make([]scriptIdiomsEntity.ScriptIdioms, 0)
	for _, s := range v.Scripts {
		scriptEntities = append(scriptEntities, scriptEntity.Script{
			VideoID:   videoID.String(),
			ScriptID:  s.ScriptID,
			Text:      s.Text,
			Japanese:  s.Japanese,
			TimeStamp: s.TimeStamp,
		})

		for _, i := range s.ScriptIdioms {
			scriptIdioms = append(scriptIdioms, scriptIdiomsEntity.ScriptIdioms{
				Idiom:    i,
				VideoID:  video.VideoID,
				ScriptID: s.ScriptID,
			})
		}
	}

	return video, scriptEntities, scriptIdioms, nil
}

type VideoUpdateRequest struct {
	VideoID string          `json:"video_id"`
	Title   string          `json:"title"`
	URL     string          `json:"url"`
	Start   float64         `json:"start"`
	End     float64         `json:"end"`
	Level   uint64          `json:"level"`
	Scripts []ScriptRequest `json:"scripts"`
}

func (v *VideoUpdateRequest) NewUpdateEntity() (*entity.Video, []*scriptEntity.Script, error) {
	video := &entity.Video{
		VideoID: v.VideoID,
		Title:   v.Title,
		URL:     v.URL,
		Start:   v.Start,
		End:     v.End,
		Level:   entity.VideoLevel(v.Level),
	}

	scriptEntities := make([]*scriptEntity.Script, 0, len(v.Scripts))
	for _, s := range v.Scripts {
		scriptEntities = append(scriptEntities, &scriptEntity.Script{
			VideoID:   v.VideoID,
			ScriptID:  s.ScriptID,
			Text:      s.Text,
			Japanese:  s.Japanese,
			TimeStamp: s.TimeStamp,
		})
	}

	return video, scriptEntities, nil
}

func NewUUID() uuid.UUID {
	return uuid.New()
}
