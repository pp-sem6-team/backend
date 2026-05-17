package ml

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/pp-sem6-team/backend/internal/domain"
	"github.com/pp-sem6-team/backend/internal/logger"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

type HTTPClient struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

func NewHTTPClient(
	baseURL string,
	apiKey string,
	timeout time.Duration,
) Client {
	return &HTTPClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		http: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPClient) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		strings.TrimRight(c.baseURL, "/")+"/health",
		nil,
	)
	if err != nil {
		return ErrRequestFailed
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return ErrRequestFailed
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrRequestFailed
	}

	return nil
}

func (c *HTTPClient) Analyze(
	ctx context.Context,
	filename string,
	data []byte,
) (*Result, error) {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
	h.Set("Content-Type", http.DetectContentType(data))

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, ErrRequestFailed
	}

	_, err = part.Write(data)
	if err != nil {
		return nil, ErrRequestFailed
	}

	if err := writer.Close(); err != nil {
		return nil, ErrRequestFailed
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		strings.TrimRight(c.baseURL, "/")+"/predict",
		&body,
	)
	if err != nil {
		return nil, ErrRequestFailed
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-API-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		logger.Log.Error("ml request failed (network error)", zap.Error(err))
		return nil, ErrRequestFailed
	}
	defer resp.Body.Close()

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		logger.Log.Info("ml content type check",
			zap.String("detect", http.DetectContentType(data)),
		)

		logger.Log.Error("ml non-200 response",
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(bodyBytes)),
		)

		return nil, ErrRequestFailed
	}

	rawJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrInvalidResponse
	}

	var parsed map[string]any

	if err := json.Unmarshal(rawJSON, &parsed); err != nil {
		return nil, ErrInvalidResponse
	}

	skinTypeValue, ok := parsed["skin_type"].(string)
	if !ok {
		return nil, ErrInvalidResponse
	}

	skinType := domain.SkinType(skinTypeValue)

	switch skinType {
	case domain.Oily, domain.Dry, domain.Normal, domain.Combination:
	default:
		return nil, ErrInvalidResponse
	}

	delete(parsed, "skin_type")

	cleanJSON, err := json.Marshal(parsed)
	if err != nil {
		return nil, ErrInvalidResponse
	}

	return &Result{
		SkinType: skinType,
		Data:     datatypes.JSON(cleanJSON),
	}, nil
}
