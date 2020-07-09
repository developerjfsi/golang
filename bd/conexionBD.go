package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root_user_admin:VA5Zz8F5xxmmNqYL@cluster.bilou.mongodb.net/test?authSource=admin&replicaSet=atlas-ofgfov-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/*ConectarBD es la función que me permite conectar la base de datos */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")

	return client
}

/*ChequeoConnection es para saber si existe una conexión a la base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
