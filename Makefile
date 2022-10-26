define build_binaries
	@echo "- Building binaries..."
	@GOOS=linux GOARCH=amd64 go build -o bin/findAll cmd/findAll/main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/post cmd/post/main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/findOne cmd/findOne/main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/deleteOne cmd/deleteOne/main.go
	@echo "Finished building binaries"
endef

define zip_files
	@echo "- Zipping files..."
	@for file in bin/*; do \
		zip -j $$file.zip $$file; \
		rm $$file; \
	done
	@echo "Finished zipping files"
endef

define clean_up
	@echo "- Cleaning up..."
	@rm -rf bin
endef

define deploy_to_aws
	@echo "- Deploying to AWS..."
	@serverless deploy --stage dev
	@echo "Finished deploying to AWS"
endef


deploy:
	@rm -rf bin/
	${build_binaries}
	${zip_files}
	${deploy_to_aws}
	${clean_up}

