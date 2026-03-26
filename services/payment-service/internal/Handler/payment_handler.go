package Handler

//import (
//	"net/http"
//	"os"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stripe/stripe-go/v76"
//	"github.com/stripe/stripe-go/v76/checkout/session"
//
//	"backend/payment-service/internal/Models"
//	"backend/payment-service/internal/db"
//)
//
//func CreateCheckoutSession(c *gin.Context) {
//	var req struct {
//		UserID uint `json:"user_id"`
//		PlanID uint `json:"plan_id"`
//	}
//
//	if err := c.BindJSON(&req); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	var plan Models.Plan
//	if err := db.InitDB.First(&plan, req.PlanID).Error; err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
//		return
//	}
//
//	params := &stripe.CheckoutSessionParams{
//		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
//		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
//		LineItems: []*stripe.CheckoutSessionLineItemParams{
//			{
//				Price:    stripe.String(plan.StripePrice),
//				Quantity: stripe.Int64(1),
//			},
//		},
//		SuccessURL: stripe.String(os.Getenv("FRONTEND_URL") + "/success"),
//		CancelURL:  stripe.String(os.Getenv("FRONTEND_URL") + "/cancel"),
//		Metadata: map[string]string{
//			"user_id": strconv.Itoa(int(req.UserID)),
//			"plan_id": strconv.Itoa(int(req.PlanID)),
//		},
//	}
//
//	s, err := session.New(params)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"url": s.URL})
//}
