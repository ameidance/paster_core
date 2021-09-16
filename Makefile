build: build.sh
	sh build.sh

run: output/bootstrap.sh
	sh output/bootstrap.sh

clean: output
	rm -rf output
