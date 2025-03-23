# tank ai backend
MVP для имплементации возможностей движка `github.com/Eight-Stones/ecs-tank-engine`
## Disclaimer
Код проекта жуткая солянка из `срезанных углов`, компромиссов и отсутствующей обработки ошибок, собран как демонстрация работы движка 
и попытка попробовать его возможности, для дальнейшей модернизации, скорее всего линтер, архитектурный линтер и отсутствие тестов
могут свести с ума, но именно MVP нужно для сбора обратной связи
## Как запустить
Команда запуска `go run cmd/game/game.go` 
## Конфигурация 
Файл с настройками по умолчанию расположен `config/config.toml`
## API для работы с игрой
Путь `docs/proto_source/game`.  
Содержит следующие методы:
- rpc Games(GamesReq) returns (GamesResp) {} - Отвечает за список игр, которые сейчас созданы
- rpc Join(JoinReq) returns (JoinResp) {} - Отправляет запрос на присоединение
- rpc Ready(ReadyReq) returns (ReadyResp) {} - Позволяет узнать запущена ли игра
## API для игрока
Путь `docs/proto_source/action`.  
Содержит следующие методы
- rpc Info(InfoReq) returns (InfoResp) {} - Возвращает данные по игроку, не имеет отката
- rpc Rotate(RotateReq) returns (RotateResp) {} - Поворачивает танк в направлении
- rpc Move(MoveReq) returns (MoveResp) {} - Двигает танк в направлении
- rpc Shoot(ShootReq) returns (ShootResp) {} - Осуществляет выстрел в текущем направлении танка
- rpc Vision(VisionReq) returns (VisionResp) {} - Запрашивает небольшой обзор вокруг игрока, не имеет отката
- rpc Radar(RadarReq) returns (RadarResp) {} - Запрашивает большой обзор вокруг игрока

Каждый метод возвращает код, с толкованием можно ознакомится в `docs/proto_source/action`