run:
	vuepress build
	rm -rf ./docs
	mv .vuepress/dist docs
	git add .

	git commit
	git push
