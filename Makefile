start:
	go run main.go

build_image:
	docker build -t simple_web_app .

run_server:
	docker run -p 5000:5000 simple_web_app