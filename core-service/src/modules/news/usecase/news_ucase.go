package usecase

import (
	"context"
	"time"

	"github.com/jinzhu/copier"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type newsUsecase struct {
	newsRepo       domain.NewsRepository
	categories     domain.CategoryRepository
	userRepo       domain.UserRepository
	tagRepo        domain.TagRepository
	dataTagRepo    domain.DataTagRepository
	areaRepo       domain.AreaRepository
	contextTimeout time.Duration
}

// NewNewsUsecase will create new an newsUsecase object representation of domain.newsUsecase interface
func NewNewsUsecase(n domain.NewsRepository, nc domain.CategoryRepository, u domain.UserRepository, tr domain.TagRepository,
	dtr domain.DataTagRepository, ar domain.AreaRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		newsRepo:       n,
		categories:     nc,
		userRepo:       u,
		tagRepo:        tr,
		dataTagRepo:    dtr,
		areaRepo:       ar,
		contextTimeout: timeout,
	}
}

func (n *newsUsecase) fillDataTags(c context.Context, data []domain.News) ([]domain.News, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the tags
	mapNews := map[int64][]domain.DataTag{}

	for _, news := range data {
		mapNews[news.ID] = []domain.DataTag{}
	}

	// Using goroutine to fetch the list tags
	chanTags := make(chan []domain.DataTag)
	for idx := range mapNews {
		newsID := idx
		g.Go(func() (err error) {
			res, err := n.dataTagRepo.FetchDataTags(ctx, newsID)
			chanTags <- res
			return
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanTags)
	}()

	for listTags := range chanTags {
		newsTags := []domain.DataTag{}
		copier.Copy(&newsTags, &listTags)
		if len(listTags) < 1 {
			continue
		}
		mapNews[listTags[0].DataID] = newsTags
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the tags's data
	for index, element := range data {
		if tags, ok := mapNews[element.ID]; ok {
			data[index].Tags = tags
		}
	}

	return data, nil
}

func (n *newsUsecase) fillDataTagsDetail(c context.Context, data domain.News) (domain.News, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the tags
	mapNews := map[int64][]domain.DataTag{}
	mapNews[data.ID] = []domain.DataTag{}

	// Using goroutine to fetch the list tags
	chanTags := make(chan []domain.DataTag)
	g.Go(func() (err error) {
		res, err := n.dataTagRepo.FetchDataTags(ctx, data.ID)
		chanTags <- res
		return
	})

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanTags)
	}()

	for listTags := range chanTags {
		newsTags := []domain.DataTag{}
		copier.Copy(&newsTags, &listTags)
		if len(listTags) < 1 {
			continue
		}
		mapNews[listTags[0].DataID] = newsTags
	}

	// merge the tags's data
	if tags, ok := mapNews[data.ID]; ok {
		data.Tags = tags
	}

	return data, nil
}

func (n *newsUsecase) fillUserDetails(c context.Context, data []domain.News) ([]domain.News, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the user's id
	mapUsers := map[uuid.UUID]domain.User{}

	for _, news := range data {
		mapUsers[news.CreatedBy.ID] = domain.User{}
	}

	// Using goroutine to fetch the user's detail
	chanUser := make(chan domain.User)
	for authorID := range mapUsers {
		authorID := authorID
		g.Go(func() error {
			res, err := n.userRepo.GetByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanUser <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanUser)
	}()

	for user := range chanUser {
		if user != (domain.User{}) {
			mapUsers[user.ID] = user
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the user's data
	for index, item := range data {
		if a, ok := mapUsers[item.CreatedBy.ID]; ok {
			data[index].CreatedBy = a
		}
	}

	return data, nil
}

func (n *newsUsecase) fillRelatedNews(c context.Context, data []domain.NewsBanner) ([]domain.NewsBanner, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the category
	mapCategories := map[string][]domain.NewsBanner{}

	for _, news := range data {
		mapCategories[news.Category] = []domain.NewsBanner{}
	}

	// Using goroutine to fetch the user's detail
	chanNews := make(chan []domain.News)
	for category := range mapCategories {
		params := domain.Request{PerPage: 4}
		params.Filters = map[string]interface{}{
			"highlight":  "0",
			"is_live":    "1",
			"categories": []string{category},
		}
		g.Go(func() (err error) {
			res, _, err := n.newsRepo.Fetch(ctx, &params)

			chanNews <- res
			return
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanNews)
	}()

	for relatedNews := range chanNews {
		if len(relatedNews) < 1 {
			continue
		}
		relatedNewsBanner := []domain.NewsBanner{}
		copier.Copy(&relatedNewsBanner, &relatedNews)
		mapCategories[relatedNews[0].Category] = relatedNewsBanner
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	// merge the user's data
	for index, item := range data {
		if a, ok := mapCategories[item.Category]; ok {
			data[index].RelatedNews = a
		}
	}

	return data, nil
}

func (n *newsUsecase) getDetail(ctx context.Context, key string, value interface{}) (res domain.News, err error) {
	if key == "slug" {
		res, err = n.newsRepo.GetBySlug(ctx, value.(string))
	} else {
		res, err = n.newsRepo.GetByID(ctx, value.(int64))
	}

	if err != nil {
		return
	}

	resCreatedBy, err := n.userRepo.GetByID(ctx, res.CreatedBy.ID)
	if err != nil {
		return
	}
	res.CreatedBy = resCreatedBy

	res, err = n.fillDataTagsDetail(ctx, res)

	if err != nil {
		return
	}

	detailArea, _ := n.areaRepo.GetByID(ctx, res.Area.ID)
	res.Area = detailArea

	return
}

func (n *newsUsecase) TabStatus(ctx context.Context, au *domain.JwtCustomClaims) (res []domain.TabStatusResponse, err error) {
	var key string
	if au.Role.ID == domain.RoleContributor {
		key = "contributor"
	}
	res, err = n.newsRepo.TabStatus(ctx, au.ID, key)

	if err != nil {
		return
	}

	return
}

func (n *newsUsecase) storeTags(ctx context.Context, newsId int64, tags []string) (err error) {

	for _, tagName := range tags {
		tag := &domain.Tag{
			Name: tagName,
		}

		// check tags already exists
		var checkTag domain.Tag
		checkTag, _ = n.tagRepo.GetTagByName(ctx, tagName)

		if checkTag.Name == "" {
			err = n.tagRepo.StoreTag(ctx, tag)
			if err != nil {
				return
			}
		}

		dataTag := &domain.DataTag{
			DataID:  newsId,
			TagID:   tag.ID,
			TagName: tagName,
			Type:    "news",
		}
		err = n.dataTagRepo.StoreDataTag(ctx, dataTag)

		if err != nil {
			return
		}
	}
	return
}

func (n *newsUsecase) get(c context.Context, params *domain.Request) (res []domain.News, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.newsRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = n.fillUserDetails(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	res, err = n.fillDataTags(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (n *newsUsecase) Fetch(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (
	res []domain.News, total int64, err error) {

	if au.Role.ID == domain.RoleContributor {
		params.Filters["created_by"] = au.ID
	} else if au.Role.ID == domain.RoleGroupAdmin {
		params.Filters["unit_id"] = au.Unit.ID
	}

	return n.get(c, params)
}

func (n *newsUsecase) FetchPublished(c context.Context, params *domain.Request) (res []domain.News, total int64, err error) {
	params.Filters["is_live"] = "1" // only published news that is live
	return n.get(c, params)
}

func (n *newsUsecase) GetByID(c context.Context, id int64) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()
	return n.getDetail(ctx, "id", id)
}

func (n *newsUsecase) GetBySlug(c context.Context, slug string) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.getDetail(ctx, "slug", slug)
	if err != nil {
		return
	}

	// FIXME: prevent abuse page views counter by using cache (redis)
	err = n.newsRepo.AddView(ctx, res.ID)
	if err != nil {
		logrus.Error(err)
	}

	return
}

func (n *newsUsecase) FetchNewsBanner(c context.Context) (res []domain.NewsBanner, err error) {

	news := []domain.News{}
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	news, err = n.newsRepo.FetchNewsBanner(ctx)
	if err != nil {
		return nil, err
	}

	news, err = n.fillUserDetails(ctx, news)
	if err != nil {
		return nil, err
	}

	copier.Copy(&res, &news)

	res, err = n.fillRelatedNews(ctx, res)
	if err != nil {
		return nil, err
	}

	return
}

func (n *newsUsecase) FetchNewsHeadline(c context.Context) (res []domain.News, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.newsRepo.FetchNewsHeadline(ctx)
	if err != nil {
		return nil, err
	}

	res, err = n.fillUserDetails(ctx, res)
	if err != nil {
		return nil, err
	}

	return
}

func (n *newsUsecase) AddShare(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()
	return n.newsRepo.AddShare(ctx, id)
}

func (n *newsUsecase) Store(c context.Context, dt *domain.StoreNewsRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	dt.CreatedAt = time.Now()
	dt.UpdatedAt = time.Now()

	if err = n.newsRepo.Store(ctx, dt); err != nil {
		return
	}

	if dt.Status == "PUBLISHED" {
		dt.Slug = helpers.MakeSlug(dt.Title, dt.ID)
		helpers.SetPropLiveNews(dt)
		n.newsRepo.Update(ctx, dt.ID, dt)
	}

	if err = n.storeTags(ctx, dt.ID, dt.Tags); err != nil {
		return
	}

	return
}

func (n *newsUsecase) Update(c context.Context, id int64, dt *domain.StoreNewsRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	news, err := n.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	dt.Slug = news.Slug

	if err := n.dataTagRepo.DeleteDataTag(ctx, id, "news"); err != nil {
		logrus.Error(err)
	}

	if err = n.storeTags(ctx, id, dt.Tags); err != nil {
		logrus.Error(err)
	}

	return n.newsRepo.Update(ctx, id, dt)
}

func (n *newsUsecase) UpdateStatus(c context.Context, id int64, status string) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	news, err := n.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	news.Status = status

	newsRequest := domain.StoreNewsRequest{
		StartDate: helpers.ConvertTimeToString(news.StartDate.Time),
		EndDate:   helpers.ConvertTimeToString(news.EndDate.Time),
		AreaID:    news.Area.ID,
	}
	copier.Copy(&newsRequest, &news)

	if status == "PUBLISHED" {
		newsRequest.Slug = helpers.MakeSlug(newsRequest.Title, newsRequest.ID)
		helpers.SetPropLiveNews(&newsRequest)
	}

	return n.newsRepo.Update(ctx, id, &newsRequest)
}

func (u *newsUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.newsRepo.Delete(ctx, id)
}
