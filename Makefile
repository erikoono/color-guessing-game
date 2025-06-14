.PHONY: test test-frontend test-backend coverage coverage-frontend

# テスト実行
test: test-frontend test-backend

test-frontend:
	docker-compose up --build frontend-test

test-backend:
	docker-compose up --build backend-test

# カバレッジレポート生成
coverage: coverage-frontend coverage-backend

coverage-frontend:
	docker-compose up --build frontend-coverage
	
coverage-backend:
	docker-compose up --build backend-coverage

# ローカル環境でのテスト実行
test-local:
	cd frontend && npm test
	cd backend && go test ./...

# ローカル環境でのカバレッジレポート生成
coverage-local:
	cd frontend && npm run test:coverage
	cd backend && go test -coverprofile=coverage.out ./...
	cd backend && go tool cover -html=coverage.out -o coverage.html

# クリーンアップ
clean:
	rm -rf coverage
	rm -rf frontend/coverage
