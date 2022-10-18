# shellcheck shell=bash

task.init() {
	go get ...
}

task.build() {
	# TODO: Remove this eventually
	go build -o './output/build/bin/aggregator' .

	mkdir -p './build' .
	go build -o './build/aggregator' .
}

task.release-nightly() {
	mkdir -p './output/build/bin'

	# Build
	task.build
	tar -C './output' -czf './output/build.tar.gz' './build'

	util.publish './output/build.tar.gz'
}

util.publish() {
	local file="$1"
	bake.assert_not_empty 'file'

	local tag_name='nightly'
	git tag -fa "$tag_name" -m ''
	git push origin ":refs/tags/$tag_name"
	git push --tags
	gh release upload "$tag_name" "$file" --clobber
	gh release edit --draft=false nightly
}
