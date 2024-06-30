build:
	@ docker build -t sant0x00/santos0santos0:latest .

run:
	@ docker run --rm -p 8080:8080 santos0santos0/gitops:latest