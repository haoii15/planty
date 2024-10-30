package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
    bun.BaseModel `bun:"table:users,alias:u"`

    UserID    string `bun:"userid,pk,type:uuid,default:gen_random_uuid()" json:"userId"`
    FirstName string `bun:"firstname,notnull" json:"firstName"`
    LastName  string `bun:"lastname,notnull" json:"lastName"`
    Username  string `bun:"username,unique,notnull" json:"username"`
    Password  string `bun:"password,notnull" json:"-"`
}

type Plant struct {
    bun.BaseModel `bun:"table:plants,alias:p"`

    PlantID string `bun:"plantid,pk,type:uuid,default:gen_random_uuid()" json:"plantId"`
    RoomID  string `bun:"roomid,type:uuid,notnull" json:"roomId"`
    Name    string `bun:"name,notnull" json:"name"`
}

type Home struct {
    bun.BaseModel `bun:"table:homes,alias:h"`

    HomeID string `bun:"homeid,pk,type:uuid,default:gen_random_uuid()" json:"homeId"`
    UserID string `bun:"userid,type:uuid,notnull" json:"userId"`
		Name  string `bun:"name,notnull" json:"name"`
	}
	
	type Room struct {
		bun.BaseModel `bun:"table:rooms,alias:r"`
		
    RoomID string `bun:"roomid,pk,type:uuid,default:gen_random_uuid()" json:"roomId"`
    HomeID string `bun:"homeid,type:uuid,notnull" json:"homeId"`
		Name  string `bun:"name,notnull" json:"name"`
}

type SensorData struct {
    bun.BaseModel `bun:"table:sensor_data,alias:s"`

    PlantID   string    `bun:"plantid,type:uuid,notnull" json:"plantId"`
    Timestamp time.Time `bun:"timestamp,notnull" json:"timestamp"`
    Data      int       `bun:"data,notnull" json:"data"`
}
