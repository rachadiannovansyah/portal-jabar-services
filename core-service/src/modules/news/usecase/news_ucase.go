package usecase

import (
	"context"
	"time"

	"github.com/jinzhu/copier"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
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
	searchRepo     domain.SearchRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewNewsUsecase will create new an newsUsecase object representation of domain.newsUsecase interface
func NewNewsUsecase(n domain.NewsRepository, nc domain.CategoryRepository, u domain.UserRepository, tr domain.TagRepository,
	dtr domain.DataTagRepository, ar domain.AreaRepository, sr domain.SearchRepository, cfg *config.Config, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		newsRepo:       n,
		categories:     nc,
		userRepo:       u,
		tagRepo:        tr,
		dataTagRepo:    dtr,
		areaRepo:       ar,
		searchRepo:     sr,
		cfg:            cfg,
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
			res, err := n.dataTagRepo.FetchDataTags(ctx, newsID, domain.ConstNews)
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
		res, err := n.dataTagRepo.FetchDataTags(ctx, data.ID, domain.ConstNews)
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

// fill data area news
func (n *newsUsecase) fillDataArea(c context.Context, data []domain.News) ([]domain.News, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the area's id
	mapAreas := map[int64]domain.Area{}

	for _, news := range data {
		mapAreas[news.Area.ID] = domain.Area{}
	}

	// Using goroutine to fetch the area's detail
	chanArea := make(chan domain.Area)
	for areaID := range mapAreas {
		areaID := areaID
		g.Go(func() error {
			res, err := n.areaRepo.GetByID(ctx, areaID)
			if err != nil {
				return err
			}
			chanArea <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanArea)
	}()

	for area := range chanArea {
		if area != (domain.Area{}) {
			mapAreas[area.ID] = area
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the area's data
	for index, item := range data {
		if a, ok := mapAreas[item.Area.ID]; ok {
			data[index].Area = a
		}
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
		params := domain.Request{PerPage: 4, SortBy: "n.views", SortOrder: "ASC"}
		params.Filters = map[string]interface{}{
			"highlight":  "0",
			"is_live":    "1",
			"categories": []string{category},
		}
		g.Go(func() (err error) {
			res, _, err := n.newsRepo.Fetch(ctx, &params)
			res, _ = n.fillUserDetails(ctx, res)

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
	return n.newsRepo.TabStatus(ctx, filterByRoleAcces(au, &domain.Request{}))
}

func (n *newsUsecase) storeTags(ctx context.Context, newsId int64, tags []string) (err error) {

	for _, tagName := range tags {
		// 20 is max length of tags name
		tagName = helpers.Substr(tagName, 20)

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
			Type:    domain.ConstNews,
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

	res, err = n.fillDataArea(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	return
}

func filterByRoleAcces(au *domain.JwtCustomClaims, params *domain.Request) *domain.Request {

	if params.Filters == nil {
		params.Filters = map[string]interface{}{}
	}

	if au.Role.ID == domain.RoleContributor {
		params.Filters["created_by"] = au.ID
	} else if helpers.IsAdminOPD(au) {
		params.Filters["unit_id"] = au.Unit.ID
	}

	return params
}

func (n *newsUsecase) Fetch(c context.Context, au *domain.JwtCustomClaims, params *domain.Request) (
	res []domain.News, total int64, err error) {
	return n.get(c, filterByRoleAcces(au, params))
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

func (n *newsUsecase) GetViewsBySlug(c context.Context, slug string) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	err = n.newsRepo.AddView(ctx, slug)
	res, err = n.newsRepo.GetBySlug(ctx, slug)
	if err != nil {
		logrus.Error(err)
	}

	return
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
		helpers.SetPropLiveNews(dt)
	}

	// set slug for the news
	dt.Slug = helpers.MakeSlug(dt.Title, dt.ID)
	n.newsRepo.Update(ctx, dt.ID, dt)

	if err = n.storeTags(ctx, dt.ID, dt.Tags); err != nil {
		return
	}

	if dt.Status != "PUBLISHED" {
		dt.PublishedAt = &time.Time{}
	}

	// FIXME: make a function to prepare data for search index
	err = n.searchRepo.Store(ctx, n.cfg.ELastic.IndexContent, &domain.Search{
		ID:          int(dt.ID),
		Domain:      "news",
		Title:       dt.Title,
		Excerpt:     dt.Excerpt,
		Content:     dt.Content,
		Slug:        dt.Slug,
		Category:    dt.Category,
		Thumbnail:   *dt.Image,
		PublishedAt: dt.PublishedAt,
		CreatedAt:   dt.CreatedAt,
		UpdatedAt:   dt.UpdatedAt,
		IsActive:    dt.IsLive == 1,
	})

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

	if err := n.dataTagRepo.DeleteDataTag(ctx, id, domain.ConstNews); err != nil {
		logrus.Error(err)
	}

	if err = n.storeTags(ctx, id, dt.Tags); err != nil {
		logrus.Error(err)
	}
	err = n.newsRepo.Update(ctx, id, dt)
	if err != nil {
		return
	}

	if news.Status != "PUBLISHED" {
		news.PublishedAt = &time.Time{}
	}

	if esErr := n.searchRepo.Update(ctx, n.cfg.ELastic.IndexContent, int(id), &domain.Search{
		Domain:      "news",
		Title:       dt.Title,
		Excerpt:     dt.Excerpt,
		Content:     dt.Content,
		Slug:        dt.Slug,
		Category:    dt.Category,
		Thumbnail:   *dt.Image,
		PublishedAt: news.PublishedAt,
		UpdatedAt:   time.Now(),
		IsActive:    news.IsLive == 1,
	}); esErr != nil {
		logrus.Error(esErr)
	}

	return
}

func (n *newsUsecase) UpdateStatus(c context.Context, id int64, status string) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	news, err := n.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	// check if end_date is null, otherwise set it to default date
	if news.EndDate == nil {
		news.EndDate = &time.Time{}
	}

	// set status
	news.Status = status
	publishedAt := time.Now()
	newsRequest := domain.StoreNewsRequest{
		StartDate: helpers.ConvertTimeToString(*news.StartDate),
		EndDate:   helpers.ConvertTimeToString(*news.EndDate),
		AreaID:    news.Area.ID,
	}
	copier.Copy(&newsRequest, &news)

	if status == "PUBLISHED" {
		newsRequest.Slug = helpers.MakeSlug(newsRequest.Title, newsRequest.ID)
		helpers.SetPropLiveNews(&newsRequest)
	} else if status == "ARCHIVED" {
		// archived news will set is_live to 0
		newsRequest.IsLive = 0
	}

	err = n.newsRepo.Update(ctx, id, &newsRequest)
	if err != nil {
		return
	}

	esErr := n.searchRepo.Update(ctx, n.cfg.ELastic.IndexContent, int(id), &domain.Search{
		Domain:      "news",
		Title:       news.Title,
		Excerpt:     news.Excerpt,
		Content:     news.Content,
		Slug:        news.Slug,
		Category:    news.Category,
		Thumbnail:   *news.Image,
		PublishedAt: &publishedAt,
		UpdatedAt:   time.Now(),
		IsActive:    newsRequest.IsLive == 1,
	})
	if esErr != nil {
		logrus.Error(esErr)
	}

	return
}

func (u *newsUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err = u.newsRepo.Delete(ctx, id); err != nil {
		return
	}

	esErr := u.searchRepo.Delete(ctx, u.cfg.ELastic.IndexContent, int(id), "news")
	if err != nil {
		logrus.Error(esErr)
	}

	return
}
