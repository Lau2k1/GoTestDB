// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"          // Для работы с BSON
// 	"go.mongodb.org/mongo-driver/mongo"         // Основной MongoDB драйвер
// 	"go.mongodb.org/mongo-driver/mongo/options" // Для настроек подключения
// )

// func main() {
// 	// URI подключения к MongoDB
// 	uri := "mongodb+srv://Lau2k:Dadada70763@gotestbase.t5m4u.mongodb.net/?retryWrites=true&w=majority&appName=GoTestBase"

// 	// Опции клиента
// 	clientOptions := options.Client().ApplyURI(uri)

// 	// Создание контекста с таймаутом
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Подключение к MongoDB
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal("Ошибка подключения к MongoDB:", err)
// 	}

// 	// Проверка подключения
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal("Не удалось подключиться к MongoDB:", err)
// 	}
// 	fmt.Println("Успешное подключение к MongoDB")

// 	// Работа с базой данных и коллекцией
// 	db := client.Database("testdb")         // Название базы данных
// 	collection := db.Collection("testcoll") // Название коллекции

// 	// Пример вставки документа
// 	doc := bson.D{{"name", "Alice"}, {"age", 25}}

// 	insertResult, err := collection.InsertOne(ctx, doc)
// 	if err != nil {
// 		log.Fatal("Ошибка при вставке документа:", err)
// 	}
// 	fmt.Println("ID вставленного документа:", insertResult.InsertedID)

// 	// Завершение работы
// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			log.Fatal("Ошибка при отключении от MongoDB:", err)
// 		}
// 	}()
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// Обновлённый URI
// 	uri := "mongodb+srv://Lau2k:Dadada70763@gotestbase.t5m4u.mongodb.net/gotestbase?retryWrites=true&w=majority&tls=true"

// 	// Опции клиента
// 	clientOptions := options.Client().ApplyURI(uri)

// 	// Создание контекста с увеличенным таймаутом
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	// Подключение к MongoDB
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal("Ошибка подключения к MongoDB:", err)
// 	}

// 	// Проверка подключения
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal("Не удалось подключиться к MongoDB:", err)
// 	}
// 	fmt.Println("Успешное подключение к MongoDB!")
// }

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb://Lau2k:Dadada70763@gotestbase-shard-00-00.t5m4u.mongodb.net:27017,gotestbase-shard-00-01.t5m4u.mongodb.net:27017,gotestbase-shard-00-02.t5m4u.mongodb.net:27017/testdb?authSource=admin&replicaSet=atlas-abcde-shard-0&readPreference=primary&tls=true"

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Не удалось подключиться к MongoDB:", err)
	}
	fmt.Println("Успешное подключение к MongoDB!")
}
