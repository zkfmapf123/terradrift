package models

import "strconv"

type TerraDriftInputParams struct {
	AccessKey       string
	AccessSecretKey string

	Concurrency int              // Concurrency Level
	SlackParams SlackInputParams // Slack Channel
}

type SlackInputParams struct {
	Token   string
	Channel string
}

type TerraDriftInputOption func(options *TerraDriftInputParams)

func TerraDriftInput(opts ...TerraDriftInputOption) TerraDriftInputParams {

	o := &TerraDriftInputParams{}

	for _, opt := range opts {
		opt(o)
	}

	return *o
}

func WithAccessKey(v string) TerraDriftInputOption {
	return func(options *TerraDriftInputParams) {
		options.AccessKey = v
	}
}

func WithAccessSecretKey(v string) TerraDriftInputOption {
	return func(options *TerraDriftInputParams) {
		options.AccessSecretKey = v
	}
}

func WithConcurreny(v string) TerraDriftInputOption {
	return func(options *TerraDriftInputParams) {

		if v == "" {
			options.Concurrency = 0
		} else {
			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			options.Concurrency = n
		}

	}
}

func WithSlackToken(token string) TerraDriftInputOption {
	return func(options *TerraDriftInputParams) {
		if token == "" {
			token = ""
		}

		options.SlackParams.Token = token
	}
}

func WithSlackChannel(channel string) TerraDriftInputOption {

	return func(options *TerraDriftInputParams) {
		if channel == "" {
			channel = ""
		}

		options.SlackParams.Channel = channel
	}
}
