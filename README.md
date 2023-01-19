# Job Scraper

<img width="382" alt="Captura de pantalla 2023-01-18 a la(s) 5 57 18 p m" src="https://user-images.githubusercontent.com/63562180/213313500-506f0053-6b21-4ccd-8010-85f770ac9442.png">

1. ``` git clone git@github.com:DICAPISAR/job_scraper.git ```
2. ``` cd job_scraper ```
3. ``` docker-compose -p "job_scraper" build ```
4. ``` docker-compose up -d```

End points

1. Linkedin

```
curl --location --request POST 'http://localhost:3000/linkedin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Developer Java", // title job
    "countToFind": 10, // results count that you want to save on database
    "location": "Colombia" // name location: Country or city (Bogotá, Colombia)
}'
```
