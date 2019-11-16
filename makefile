run:
	yarn docs:build
	rm -rf ./docs
	mv .vuepress/dist docs
	git add .

	git commit
	git push
