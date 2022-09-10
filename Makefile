SHELL := /bin/bash

## help: Show this help message
help:
	@ printf "\033[33m%s:\033[0m\n" 'Available commands'
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## clean: Clean build & benchmark files
clean:
	@ rm -f Barmak
	@ rm -f graphic/operations.png
	@ rm -f graphic/time_operations.png

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
benchmark: clean
	@ cd kafka ; \
	go test -bench=. -run=^# | tee ../graphic/out.dat ; \
	awk '/Benchmark/{count ++; gsub(/BenchmarkProduce/,""); gsub("-[^-]",""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3)}' ../graphic/out.dat > ../graphic/final.dat ; \
	MAX=$$(cat ../graphic/final.dat | sort -k3 -t',' -nr | head -1 | awk -F',' '{print $$3}'); \
	MAX=$$(echo $$MAX + 200 | bc); \
	gnuplot -e "file_path='../graphic/final.dat'" -e "graphic_file_name='../graphic/operations.png'" -e "y_label='Num of Ops'" -e "y_range_min='000000000''" -e "y_range_max='$$MAX'" -e "column_1=1" -e "column_2=3" ../graphic/performance.gp ; \
	gnuplot -e "file_path='../graphic/final.dat'" -e "graphic_file_name='../graphic/time_operations.png'" -e "y_label='Each Op in NS'" -e "y_range_min='000''" -e "y_range_max='110000000'" -e "column_1=1" -e "column_2=4" ../graphic/performance.gp ; \
	rm -f ../graphic/out.dat ../graphic/final.dat ; \
	echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."
