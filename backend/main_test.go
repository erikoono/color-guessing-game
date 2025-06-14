package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// テスト実行前の設定
	rand.Seed(time.Now().UnixNano())
	
	// テスト実行
	os.Exit(m.Run())
}

func TestColorHandler(t *testing.T) {
	// デフォルトのテスト（count=3）
	req := httptest.NewRequest(http.MethodGet, "/api/color", nil)
	w := httptest.NewRecorder()

	colorHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var body ColorResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// デフォルトでは3つのオプションを返すように変更
	if len(body.Options) != 3 {
		t.Errorf("expected 3 options, got %d", len(body.Options))
	}

	found := false
	for _, c := range body.Options {
		if c == body.Correct {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("correct color %s not in options %v", body.Correct, body.Options)
	}
}

func TestColorHandlerWithCount4(t *testing.T) {
	// count=4のテスト
	req := httptest.NewRequest(http.MethodGet, "/api/color?count=4", nil)
	w := httptest.NewRecorder()

	colorHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var body ColorResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// count=4の場合は4つのオプションを返す
	if len(body.Options) != 4 {
		t.Errorf("expected 4 options, got %d", len(body.Options))
	}

	found := false
	for _, c := range body.Options {
		if c == body.Correct {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("correct color %s not in options %v", body.Correct, body.Options)
	}
}

func TestColorHandlerWithInvalidCount(t *testing.T) {
	// 無効なcountパラメータのテスト
	req := httptest.NewRequest(http.MethodGet, "/api/color?count=invalid", nil)
	w := httptest.NewRecorder()

	colorHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var body ColorResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	// 無効なcountの場合はデフォルトの3つのオプションを返す
	if len(body.Options) != 3 {
		t.Errorf("expected 3 options, got %d", len(body.Options))
	}
}