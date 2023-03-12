package videos

import (
	"errors"
	"strings"

	"github.com/HondaAo/snippet/src/pkg/handle/requests"
	"github.com/HondaAo/snippet/src/pkg/services/idioms"
	"github.com/HondaAo/snippet/src/pkg/services/scriptidioms"
	scriptIdiomsEntity "github.com/HondaAo/snippet/src/pkg/services/scriptidioms/entity"
	"github.com/HondaAo/snippet/src/pkg/services/scripts"

	idiomEntity "github.com/HondaAo/snippet/src/pkg/services/idioms/entity"
	scriptEntity "github.com/HondaAo/snippet/src/pkg/services/scripts/entity"
	"github.com/HondaAo/snippet/src/pkg/services/videos/entity"
)

type VideoUsecaseInterface interface {
	GetVideoByID(videoID string) (*entity.Video, []*scriptEntity.Script, []*scriptIdiomsEntity.ScriptIdioms, []*idiomEntity.Idioms, error)
	GetByIdioms(idioms []string) ([]*entity.Video, error)
	GetVideos(request *requests.VideoSearchType) ([]*entity.Video, error)
	CreateVideo(request *requests.VideoRequest) error
	UpdateVideo(request *requests.VideoUpdateRequest) error
	ChangeDisplayStatus(videoID string) error
	DeleteVideo(videoID string) error
}

type videoUsecase struct {
	videoRepository        VideoRepository
	scriptRepository       scripts.ScriptRepository
	scriptIdiomsRepository scriptidioms.ScriptIdiomsRepository
	idiomsRepository       idioms.IdiomRepository
}

func NewVideoUsecase(
	videoRepository VideoRepository,
	scriptRepository scripts.ScriptRepository,
	scriptIdiomsRepository scriptidioms.ScriptIdiomsRepository,
	idiomsRepository idioms.IdiomRepository,
) *videoUsecase {
	return &videoUsecase{
		videoRepository:        videoRepository,
		scriptRepository:       scriptRepository,
		scriptIdiomsRepository: scriptIdiomsRepository,
		idiomsRepository:       idiomsRepository,
	}
}

func (v *videoUsecase) GetVideoByID(videoID string) (*entity.Video, []*scriptEntity.Script, []*scriptIdiomsEntity.ScriptIdioms, []*idiomEntity.Idioms, error) {
	video, err := v.videoRepository.Find(videoID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	scripts, err := v.scriptRepository.FindByVideoID(videoID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	scriptIdioms, err := v.scriptIdiomsRepository.FindByVideoID(videoID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	idioms := make([]string, 0, len(scriptIdioms))
	for _, i := range scriptIdioms {
		idioms = append(idioms, i.Idiom)
	}

	if len(idioms) == 0 {
		return video, scripts, scriptIdioms, nil, nil
	}

	var trimedIdioms []string
	for _, idiom := range idioms {
		trimedIdioms = append(trimedIdioms, strings.ReplaceAll(idiom, " ", ""))
	}
	idiomEntities, err := v.idiomsRepository.FindIdioms(trimedIdioms)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return video, scripts, scriptIdioms, idiomEntities, nil
}

func (v *videoUsecase) CreateVideo(request *requests.VideoRequest) error {
	video, scripts, scriptIdioms, err := request.NewEntity()
	if err != nil {
		return err
	}
	if err = v.videoRepository.Store(video); err != nil {
		return err
	}

	if err = v.scriptRepository.Store(video.VideoID, scripts); err != nil {
		return err
	}

	idioms := make([]string, 0, len(scriptIdioms))
	for _, i := range scriptIdioms {
		idioms = append(idioms, i.Idiom)
	}

	// idiomsが全てDBにあるか
	_, err = v.idiomsRepository.FindIdioms(idioms)
	if err != nil {
		return err
	}

	if err = v.scriptIdiomsRepository.Store(scriptIdioms); err != nil {
		return err
	}

	return nil
}

func (v *videoUsecase) GetVideoByIdioms(idioms []string) ([]*entity.Video, error) {
	scriptIdioms, err := v.scriptIdiomsRepository.Find(idioms)
	if err != nil {
		return nil, err
	}

	var videoIds []string
	for _, s := range scriptIdioms {
		videoIds = append(videoIds, s.VideoID)
	}
	videos, err := v.videoRepository.FindMany(entity.SearchCondition{
		VideoIDs: videoIds,
		Limit:    entity.DEFAULT_MAX_LIMIT,
		Date:     entity.GetDate(0),
	})
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (v *videoUsecase) GetVideos(request *requests.VideoSearchType) ([]*entity.Video, error) {
	videos, err := v.videoRepository.FindMany(entity.SearchCondition{
		Levels:      request.Levels,
		Limit:       request.Limit,
		CategoryIDs: request.Categories,
		Length:      entity.GetLength(request.Length),
		TypeIDs:     request.Types,
		Date:        entity.GetDate(uint8(request.Date)),
	})

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (v *videoUsecase) ChangeDisplayStatus(videoID string) error {
	if err := v.videoRepository.ChangeStatus(videoID); err != nil {
		return err
	}

	return nil
}

func (v *videoUsecase) UpdateVideo(request *requests.VideoUpdateRequest) error {
	video, scripts, err := request.NewUpdateEntity()
	if err != nil {
		return err
	}

	if err = v.videoRepository.Update(video); err != nil {
		return err
	}

	for _, script := range scripts {
		if err = v.scriptRepository.Update(script); err != nil {
			return err
		}
	}

	return nil
}

func (v *videoUsecase) DeleteVideo(videoID string) error {
	if err := v.videoRepository.Delete(videoID); err != nil {
		return err
	}

	if err := v.scriptRepository.Delete(videoID); err != nil {
		return err
	}

	if err := v.scriptIdiomsRepository.Delete(videoID); err != nil {
		return err
	}

	return nil
}

func (v *videoUsecase) GetByIdioms(idioms []string) ([]*entity.Video, error) {
	idiomEntities, err := v.idiomsRepository.FindIdioms(idioms)
	if err != nil {
		return nil, err
	}

	if len(idiomEntities) == 0 {
		return nil, errors.New("Video Not Found")
	}
	var words []string
	for _, idiom := range idiomEntities {
		words = append(words, idiom.Idiom)
	}
	scriptIdiomsEntities, err := v.scriptIdiomsRepository.Find(words)
	if err != nil {
		return nil, err
	}

	var videoIds []string
	for _, script := range scriptIdiomsEntities {
		videoIds = append(videoIds, script.VideoID)
	}

	videos, err := v.videoRepository.FindMany(entity.SearchCondition{
		VideoIDs: videoIds,
		Limit:    entity.DEFAULT_MAX_LIMIT,
		Date:     entity.GetDate(0),
	})
	if err != nil {
		return nil, err
	}

	return videos, nil
}
