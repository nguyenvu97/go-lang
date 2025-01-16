package middleweres

import (
	"awesomeProject/global"
	"awesomeProject/pkg/response"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"

	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)
type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter
	publicAPIRateLimiter      *limiter.Limiter
	userPrivateAPIRateLimiter *limiter.Limiter
}

func NewRateLimiter () *RateLimiter {
	rateLimit := &RateLimiter{
		globalRateLimiter:  rateLimiter("100-S"),
		publicAPIRateLimiter: rateLimiter("80-s"),
		userPrivateAPIRateLimiter : rateLimiter("60-s"),
	}
	return rateLimit
}

func rateLimiter(interval string) *limiter.Limiter {
	store,err := redisStore.NewStoreWithOptions(global.Rdb,limiter.StoreOptions{
		Prefix: "rate-limiter",
		MaxRetry: 3,
		CleanUpInterval: time.Hour,
	})
	if err != nil{
		fmt.Printf("err %e", err)
	}
	rate ,err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		fmt.Printf("err %e", err)
	}
	instance := limiter.New(store,rate)

	return instance

}

// global 
func(rl * RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		key := "global" // unit
		log.Println("global--->")
		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}
		if limitContext.Reached {
			log.Printf("Rate limit breached GLOBAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try later"})
			return
		}
		c.Next()

	}
}

// public
func(rl * RateLimiter)PublicRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path // urlPath: /ping/80 80 >
		rateLimitPath := rl.filterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			log.Println("Client Ip--->", c.ClientIP())
			key := fmt.Sprintf("%s-%s", "111-222-333-44", urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

	}

}

// private
func (rl * RateLimiter) PrivateRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)
		if rateLimitPath != nil {
			userId := c.Request.Context().Value("subjectUUID")
			if userId == nil {
				 response.ErrorResponse(c,5001)
				 return
			}
			fmt.Printf("data %d",userId)
			key := fmt.Sprintf("%d-%s", userId, urlPath)
			limitContext, err := rateLimitPath.Get(c, key)
			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}
			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}
		c.Next()
	}
}

func(rl * RateLimiter)filterLimitUrlPath(urlPath string) * limiter.Limiter{

	if urlPath == "/api/v1/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/api/v1/admin/add-product" || urlPath == "/ping/50" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}

}


