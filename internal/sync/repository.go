package sync

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"kssyncservice_go/config"
	"kssyncservice_go/pkg/db"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type SyncRepository struct {
	Database *db.Db
}

func NewSyncRepository(database *db.Db) *SyncRepository {
	return &SyncRepository{
		Database: database,
	}
}

func calculateDataHash(data []Tmp_Ksservice) []Tmp_Ksservice {
	ns := []Tmp_Ksservice{}
	for _, value := range data {
		values := reflect.ValueOf(value)
		var builder strings.Builder
		for i := 0; i < values.NumField(); i++ {
			builder.WriteString(fmt.Sprintf("%v", values.Field(i)))
		}
		hasher := sha1.New()
		hasher.Write([]byte(builder.String()))
		hash := hasher.Sum(nil)
		value.Line_Hash = hex.EncodeToString(hash)
		ns = append(ns, value);
	}

	return ns;
}

func filterHashedData(hashedData []Tmp_Ksservice, insertHashes []string) []Ksservice {
	insertMap := make(map[string]bool)
	for i := 0; i < len(insertHashes); i++ {
		insertMap[insertHashes[i]] = true
	}

	ns := []Ksservice{}

	for _, value := range hashedData {
		if insertMap[value.Line_Hash] {
			ns = append(ns, Ksservice(value))
		}
	}

	return ns;
}

func (repo *SyncRepository) Synchronize(conf *config.Config) {

	response, err := http.Get(fmt.Sprintf("http://%s:%s%s", conf.ExternalServiceConfig.KS_SERVICE_HOST, conf.ExternalServiceConfig.KS_SERVICE_PORT, conf.ExternalServiceConfig.KS_SERVICE_METHOD))
	if err != nil {
		fmt.Println(time.Now(), " - error fetch data!")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(time.Now(), " - error read data!")
	}

	var data ServicesResponse
	json.Unmarshal(body, &data)

	hashedData := calculateDataHash(data.Data)

	repo.Database.DB.Migrator().CreateTable(Tmp_Ksservice{})
	repo.Database.DB.Omit("id").Create(&hashedData)
	defer repo.Database.DB.Migrator().DropTable(Tmp_Ksservice{})


	var insertHashes []string
	var deleteHashes []string 
	
	repo.Database.DB.Model(&Tmp_Ksservice{}).Select("tmp_services.line_hash").Joins("full outer join services on services.line_hash = tmp_services.line_hash").Where("services.line_hash is NULL").Scan(&insertHashes)
	repo.Database.DB.Model(&Tmp_Ksservice{}).Select("services.line_hash").Joins("full outer join services on services.line_hash = tmp_services.line_hash").Where("tmp_services.line_hash is NULL").Scan(&deleteHashes)
	
	if len(deleteHashes) > 0 {
		repo.Database.DB.Where("line_hash IN ?", deleteHashes).Unscoped().Delete(&Ksservice{})
	}
	
	insertData := filterHashedData(hashedData, insertHashes)
	if len(insertData) > 0 {
		repo.Database.DB.Omit("id").Create(&insertData)
	}
}
