run:
	yarn docs:build
	rm -rf ./docs
	mv ./site/.vuepress/dist docs
	git add .
