package main

//go:generate go install github.com/1340691923/ElasticView/cmd/ev_plugin_builder

//go:generate go install github.com/1340691923/ElasticView/cmd/ev_plugin_zip

//go:generate swag init -g main.go -o docs -exclude cmd,dist,frontend

//go:generate gowatch

//go:generate ev_plugin_builder

//go:generate ev_plugin_zip
