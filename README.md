Требуется реализовать микросервис согласно требованиям:

1. Сервис должен загрузить данные о странах и телефонных кодах стран в реляционную БД.
    Источник данных:
    Старны: http://country.io/names.json 
    Телефоныне коды: http://country.io/phone.json 
2. Реализовать REST сервис POST /reload, который по требованию обновит данные в БД о странах и телефонных кодах.
3. Выполнять автоматический запуск обновления данных в БД о странах и телефонных кодах, по расписанию, каждую ночь в 2 часа ночи по Москве.
4. Реализовать REST сервис GET /code/${COUNTRYNAME}, который возвращает телефонный код по имени страны, где ${COUNTRYNAME} - это имя страны в любом регистре, например JAMAICA. Возвращать HTTP код 200 в случае успешного ответа, возвращать код 404 в случае, если страна не найдена.
5. Реализовать unit test к сервису GET /code/${COUNTRYNAME}
6. Реализовать логирование ошибок приложения и логирование запросов к реализованным REST сервисам.
7. Написать Dockerfile сборки и деплоя приложения 

-------------------------------------------------

Implement microservice as required:

1. Service must load info about countries and phone codes to the relational DB
    Data sourse: 
        Couuntries: http://country.io/names.json 
        Phone codes: http://country.io/phone.json 
2. Implement RESTful service with the /reload endpoint (POST), wich will refresh info about countries and codes in the DB
3. Implement automatic DB refresh at 2:00AM MSK.
4. Implement RESTful service with the /code/${COUNTRYNAME} endpoint (GET). Service returns phone code, related to country, where ${COUNTRYNAME} is a case insencitive country name, for example JAMAICA. Return HTTP 200 code in case of success. Return HTTP 404 code in case, when country has not been found.
5. Implement unit test for /code/${COUNTRYNAME} (GET)
6. Implement error logging and request logging for implemented RESTful services
7. Make Dockerfile for application build and deploy.
