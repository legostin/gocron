
**Gocron -- минималистичное решение для запуска задач, чаще, чем раз в минуту**

[![Go Report Card](https://goreportcard.com/badge/github.com/legostin/gocron)](https://goreportcard.com/report/github.com/legostin/gocron)
[![Go Build Card](https://travis-ci.org/legostin/gocron.svg?branch=master)](https://travis-ci.org/legostin/gocron)

**Запуск:**
1. У Вас должен быть установлен Go актуальной версии, и прописан $GOPATH
2. Входим в директорию, куда мы хотим установить скрипт `cd /home;`
3. Клонируем репозиторий `git clone https://github.com/legostin/gocron;`
4. `cd gocron;`
5. Задачи описываются в файле cronjob.json
6. после описания нужных нам задач `go build main.go`
7. ./main &


**Описание задач:**

Простейшая задача

```
{
    "Tasks":[
        {
            "Period":15,
            "Command":"ping -c 4 google.com",
            "Output":"./logs.txt"
        }
    ]
}   
```

В данном случае мы будем вызывать команду `ping -c 4 google.com >> ./logs.txt`  каждые 15 секунд

Усложним пример. Добавим периоды, когда скрипт не будет вызываться

```
{
       "Period": 84,
       "SleepTime": [
         "19:00:00-23:59:59",
         "00:00:00-13:00:00"
       ],
       "SleepDays": [
         "Mon"
       ],
       "Command": "ping -c 4 google.com",
       "Output": "./logs2.txt"
}
```
     
Данный скрипт будет вызывать команду каждые 84 секунды, в промежутке времени между 13 часами дня и 19 часами вечера, каждый день, кроме понедельника

Также имеется возможность вызывать команду в строго заданное время
```
{
       "Time": [
         "14:00:00",
         "14:30:00",
       ],
       "Command": "ping -c 4 google.com",
       "Output": "./logs3.txt"
}
  ```
     
Данный скрипт будет срабатывать дважды в день, в строго заданное время

А еще можно задать конкретную дату и время для запуска скрипта 
```
{
       "DateTime": [
         "2018-05-13 14:07:22",
         "2018-05-13 13:55:10",
         "2018-05-14 15:00:00",
         "2018-05-15 17:00:00"
       ],
       "Command": "ping -c 4 google.com",
       "Output": "./logs4.txt"
}
 ```
 
За примером заполнения файла с задачами можете заглянуть в файл [cronjob.json](https://github.com/legostin/gocron/blob/master/cronjob.json)