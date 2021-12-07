include META

.PHONY: generate-docs
generate-docs:
	@echo "[GENERATE DOCS] Generating API documents"
	@echo " - Updating document version"
	@echo " - Initializing swag"
	@swag init --parseDependency --parseInternal --generatedTime --parseDepth 3

.PHONY: git
git:
	@echo "[BUILD] Committing and pushing to remote repository"
	@echo " - Committing"
	@git add META
	@git commit -am "v$(VERSION)"
	@echo " - Tagging"
	@git tag v${VERSION}
	@echo " - Pushing"
	@git push --tags origin ${BRANCH}