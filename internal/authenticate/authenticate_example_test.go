package authenticate_test

import (
	"fmt"
	"go.uber.org/zap"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/logger"
)

func ExampleGetUserIDFromCookie() {

	cookieValue := "30316566376539342d346130332d363633312d616163312d3132366664393030643461313d236637f1eb14e449f64a77d044f63a120fd5ee06d57c4f06684ae32c36e344"

	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	authService := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)
	_, err := authService.GetUserIDFromCookie(cookieValue)
	if err == nil {
		fmt.Println("Cookie is valid")
	} else {
		fmt.Println("Cookie is not valid")
	}

	// Output:
	// Cookie is not valid
}
