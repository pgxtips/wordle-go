package globals

import(
    "app/internal/models"
)

var serverData map[string]models.AppData

func Init(){
    serverData = make(map[string]models.AppData, 1)
}

func GetServerData() map[string]models.AppData {
    return serverData
}

func AddServerData(uuid string){
    serverData[uuid] = models.AppData{
        Id: uuid,
        CurrentData: "",
    }
}

func RemoveServerData(uuid string){
    delete(serverData, uuid)
}





