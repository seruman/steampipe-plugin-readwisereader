package readwisereader

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	"github.com/seruman/go-readwisereader"
)

func tableDocuments() *plugin.Table {
	return &plugin.Table{
		Name:        "readwisereader_documents",
		Description: "Docs and stuff",
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the document."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "The URL of the document."},
			{Name: "source_url", Type: proto.ColumnType_STRING, Description: "The URL of the source of the document."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The title of the document."},
			{Name: "author", Type: proto.ColumnType_STRING, Description: "The author of the document."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "The source of the document."},
			{Name: "category", Type: proto.ColumnType_STRING, Description: "The category of the document."},
			{Name: "location", Type: proto.ColumnType_STRING, Description: "The location of the document."},
			{Name: "tags", Type: proto.ColumnType_JSON, Description: "The tags of the document."},
			{Name: "site_name", Type: proto.ColumnType_STRING, Description: "The site name of the document."},
			{Name: "word_count", Type: proto.ColumnType_INT, Description: "The word count of the document."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The creation time of the document."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The update time of the document."},
			{Name: "notes", Type: proto.ColumnType_STRING, Description: "The notes of the document."},
			{Name: "published_date", Type: proto.ColumnType_TIMESTAMP, Description: "The published date of the document."},
			{Name: "summary", Type: proto.ColumnType_STRING, Description: "The summary of the document."},
			{Name: "image_url", Type: proto.ColumnType_STRING, Description: "The image URL of the document."},
			{Name: "parent_id", Type: proto.ColumnType_STRING, Description: "The parent ID of the document."},
			{Name: "reading_progress", Type: proto.ColumnType_DOUBLE, Description: "The reading progress of the document."},
			{Name: "first_opened_at", Type: proto.ColumnType_TIMESTAMP, Description: "The first opened time of the document."},
			{Name: "last_opened_at", Type: proto.ColumnType_TIMESTAMP, Description: "The last opened time of the document."},
			{Name: "saved_at", Type: proto.ColumnType_TIMESTAMP, Description: "The save time of the document."},
			{Name: "last_moved_at", Type: proto.ColumnType_TIMESTAMP, Description: "The last moved time of the document."},
		},
		List: &plugin.ListConfig{
			Hydrate: tableDocumentsList,
		},
	}
}

func tableDocumentsList(
	ctx context.Context,
	d *plugin.QueryData,
	h *plugin.HydrateData,
) (interface{}, error) {
	var listparams readwisereader.ListParams
	for col, qual := range d.EqualsQuals {
		switch col {
		case "id":
			listparams.ID = qual.GetStringValue()
		case "location":
			listparams.Location = readwisereader.Location(qual.GetStringValue())
		case "category":
			listparams.Category = readwisereader.Category(qual.GetStringValue())
		}
	}

	client := connect(ctx, d)

	for page, err := range client.ListPaginate(ctx, listparams) {
		if err != nil {
			return nil, err
		}

		for _, doc := range page.Results {
			dd := docy{
				ID:              doc.ID,
				URL:             doc.URL,
				SourceURL:       doc.SourceURL,
				Title:           doc.Title,
				Author:          doc.Author,
				Source:          doc.Source,
				Category:        string(doc.Category),
				Location:        string(doc.Location),
				Tags:            doc.Tags,
				SiteName:        doc.SiteName,
				WordCount:       doc.WordCount,
				CreatedAt:       doc.CreatedAt,
				UpdatedAt:       doc.UpdatedAt,
				Notes:           doc.Notes,
				PublishedDate:   doc.PublishedDate,
				Summary:         doc.Summary,
				ImageURL:        doc.ImageURL,
				ParentID:        doc.ParentID,
				ReadingProgress: doc.ReadingProgress,
				FirstOpenedAt:   doc.FirstOpenedAt,
				LastOpenedAt:    doc.LastOpenedAt,
				SavedAt:         doc.SavedAt,
				LastMovedAt:     doc.LastMovedAt,
			}
			d.StreamListItem(ctx, dd)

			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func connect(ctx context.Context, d *plugin.QueryData) *readwisereader.Client {
	cacheKey := "readwisereader"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*readwisereader.Client)
	}

	cfg := GetConfig(d.Connection)

	if cfg.Token == nil {
		panic("Token must be set")
	}

	client := readwisereader.NewClient(*cfg.Token)
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client
}

type docy struct {
	ID              string         `json:"id"`
	URL             string         `json:"url"`
	SourceURL       string         `json:"source_url"`
	Title           string         `json:"title"`
	Author          string         `json:"author"`
	Source          string         `json:"source"`
	Category        string         `json:"category"`
	Location        string         `json:"location"`
	Tags            map[string]any `json:"tags"`
	SiteName        string         `json:"site_name"`
	WordCount       int            `json:"word_count"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Notes           string         `json:"notes"`
	PublishedDate   time.Time      `json:"published_date"`
	Summary         string         `json:"summary"`
	ImageURL        string         `json:"image_url"`
	ParentID        string         `json:"parent_id"`
	ReadingProgress float64        `json:"reading_progress"`
	FirstOpenedAt   time.Time      `json:"first_opened_at"`
	LastOpenedAt    time.Time      `json:"last_opened_at"`
	SavedAt         time.Time      `json:"saved_at"`
	LastMovedAt     time.Time      `json:"last_moved_at"`
}
