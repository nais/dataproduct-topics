package collect

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type TopicsResponse struct {
	Topics []string
}

func GetKafkaTopicsFromKafkaAdminRest(ctx context.Context, url string) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/topics", url), nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Errorf("Close response body: %s", err)
		}
	}()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 status code")
	}

	topicList := &TopicsResponse{}
	err = json.NewDecoder(response.Body).Decode(topicList)
	if err != nil {
		return nil, err
	}

	return topicList.Topics, nil
}
