package models

import "strconv"

type TerraDriftInputParams struct {
	IsUseTerraformPath  string // Terraform Project Path
	IsUseTerragruntPath string // Terragrunt Project Path
	Concurrency         int    // Concurrency Level

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

func WithIsUseTerraformPath(v string) TerraDriftInputOption {
	return func(o *TerraDriftInputParams) {

		// if len(v) == 0 {
		// 	v = nil
		// }

		o.IsUseTerraformPath = v
	}
}

func WithIsUseTerragruntPath(v string) TerraDriftInputOption {
	return func(o *TerraDriftInputParams) {

		// if len(v) == 0 {
		// 	v = nil
		// }

		o.IsUseTerragruntPath = v
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
