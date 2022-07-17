SHELL := /bin/bash

## help: Show this help message
help:
	@ echo -e "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: Build app's binary
build:
	@ go build -a -installsuffix cgo -o Barmak .

## run: Run the app
run: build
	@ ./Barmak

## test: Run unit tests
test:
	@ go test ./kafka -count=1

## benchmark: Run benchmark tests
benchmark:
	@ cd kafka ; \
	go test -bench=. -run=^# | tee ../graphic/out.dat ; \
	awk '/Benchmark/{count ++; gsub(/BenchmarkProduce/,""); gsub("-[^-]",""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' ../graphic/out.dat > ../graphic/final.dat ; \
	gnuplot -e "file_path='../graphic/final.dat'" -e "graphic_file_name='../graphic/operations.png'" -e "y_label='Num of Ops'" -e "y_range_min='000000000''" -e "y_range_max='1300'" -e "column_1=1" -e "column_2=3" ../graphic/performance.gp ; \
	gnuplot -e "file_path='../graphic/final.dat'" -e "graphic_file_name='../graphic/time_operations.png'" -e "y_label='Each Op in NS'" -e "y_range_min='000''" -e "y_range_max='110000000'" -e "column_1=1" -e "column_2=4" ../graphic/performance.gp ; \
	echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."
