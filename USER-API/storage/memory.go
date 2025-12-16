package storage
import(
	"fmt"
	"user-api/models"
)
var users = make(map{string}models.User)

func GetAllUsers()[]models.User{
	return []models.User{}
}
func CreateUser(user models.User)error{
	fmt.Printf("Создан новый пользователь: %+v\n",users)
	return nil
}