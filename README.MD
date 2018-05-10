**Gocron -- минималистичное решение для запуска задач, чаще, чем раз в минуту**

Запуск:
1. У Вас должен быть установлен Go актуальной версии, и прописан $GOPATH
2. Сами задачи записываются в файл cronjob.json
3. Запуск `go build main.go`

        
         {
             "Tasks":[
                 {
                     "Period":25,
                     "Command":"wget -O- https://google.com",
                     "Output":"./logs.txt"
                 },
                 {
                     "Period":10,
                     "Command":"ping google.com",
                     "Output":"./logs2.txt"
                 }
            ]
        }

Период записывается в секундах

Колличество задач неограничено
