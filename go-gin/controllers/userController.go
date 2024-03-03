package controllers


import ( 
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "go-gin/helpers"
	"go-gin/models"
	"go-gin/database"
	"golang.org/x/crypto/bcrypt"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	)

)

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

