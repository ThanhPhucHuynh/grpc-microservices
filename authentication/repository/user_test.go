package repository

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ThanhPhucHuynh/authentication/models"
	"github.com/ThanhPhucHuynh/db"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)


func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUserRepositorySave(t *testing.T)  {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)

	assert.NoError(t, err)
	defer conn.Close()


	id := bson.NewObjectId()

	user := &models.User{
		Id: id,
		Name: "TEST",
		Email: fmt.Sprintf("%s@email.com",id.Hex()),
		Password: "123546",
		Created: time.Now(),
		Updated: time.Now(),
	}

	r:= NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)


	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)

	// assert.Equal(t,user., found.Id)
	// assert.NotNil(t,found)


}