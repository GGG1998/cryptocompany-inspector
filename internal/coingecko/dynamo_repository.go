package coingecko

type (
	DetailPlatform struct {
		Decimal         int    `dynamodbc:"Decimal"`
		ContractAddress string `dynamodbc:"ContractAddress"`
	}

	Links struct {
		Name string `dynamodbc:"Name"`
		Url  string `dynamodbc:"Url"`
	}

	Scores struct {
		Name  string
		Value float64
	}

	Rank struct {
		PlatformName  string `dynamodbv:"PlatformName"`
		MarketcapRank string `dynamodbv:"MarketcapRank"`
		PlatformRank  int    `dynamodbv:"PlatformRank"`
	}

	CommunityData struct {
		TwitterFollowers         int `dynamodbv:"TwitterFollowers"`
		RedditAveragePosts48H    int `dynamodbv:"RedditAveragePosts48H"`
		RedditAverageComments48H int `dynamodbv:"RedditAverageComments48H"`
		RedditAverageActive48H   int `dynamodbv:"RedditAverageActive48H"`
		TelegramChannelUserCount int `dynamodbv:"TelegramChannelUserCount"`
		WatchlistPortfolioUsers  int `dynamodbv:"WatchlistPortfolioUsers"`
	}

	CryptoCompanies struct {
		// LocalSecondary Index
		GlobalRank    string `dynamodbv:"GlobalRank"`
		CommunityRank string `dynamodbv:"CommunityRank"`

		// Main Index
		SortKey string `dynamodbc:"SortKey"`
		Id      string `dynamodbv:"Id"`

		CreatedAt       string                    `dynamodbv:"CreatedAt"`
		UpdateAt        string                    `dynamodbv:"UpdateAt"`
		AssetPlatformId string                    `dynamodbv:"AssetPlatformId"`
		Symbol          string                    `dynamodbv:"Symbol"`
		Name            string                    `dynamodbv:"Name"`
		Notice          string                    `dynamodbv:"Notice"`
		Description     string                    `dynamodbv:"Description"`
		Platforms       map[string]string         `dynamodbv:"Platforms"`
		DetailPlatform  map[string]DetailPlatform `dynamodbv:"DetailPlatform"`
		PlatformRank    map[string]Rank           `dynamodbv:"PlatformRank"`
		Links           []Links                   `dynamodbv:"Links"`
		ForumUrls       []Links                   `dynamodbv:"ForumUrls"`
		ChatUrls        []Links                   `dynamodbv:"ChatUrls"`
		RepoUrls        []Links                   `dynamodbv:"RepoUrls"`
		Socialmedia     []string                  `dynamodbv:"Socialmedia"`
		Categories      []string                  `dynamodbv:"Categories"`
	}
)
