package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/customer"
	"github.com/stripe/stripe-go/v80/paymentintent"
)

const key = "sk_test_51Q4kQAE2Lje0pxrNEvC0YtsAteE0qVUHUyjtHV4HssJZgRPEAjOWwDqNf3S4SFUWCfiifmWomWsziL1LdQJC0d1Q00bpAmVokP"

func HandlerPaymentSheet(ctx *gin.Context) {
	stripe.Key = key

	// Use an existing Customer ID if this is a returning customer.
	cparams := &stripe.CustomerParams{}
	c, _ := customer.New(cparams)

	// ekparams := &stripe.EphemeralKeyParams{
	// 	Customer:      stripe.String(c.ID),
	// 	StripeVersion: stripe.String("2024-06-20"),
	// }

	piparams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		Customer: stripe.String(c.ID),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter
		// is optional because Stripe enables its functionality by default.
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	pi, _ := paymentintent.New(piparams)

	ctx.JSON(http.StatusOK, struct {
		PaymentIntent  string `json:"paymentIntent"`
		EphemeralKey   string `json:"ephemeralKey"`
		Customer       string `json:"customer"`
		PublishableKey string `json:"publishableKey"`
	}{
		PaymentIntent: pi.ClientSecret,
		// EphemeralKey:   *ekparams.IdempotencyKey,
		Customer:       c.ID,
		PublishableKey: "pk_test_51Q4kQAE2Lje0pxrNTWWudkYsRSVA6KnkOhqLg6s5n1O56lfYA5Z1s30CP1bVcD8NPRbS3Mt8usOEgYNqct0KONRi00iqKMd1Lg",
	})
}
