import csv
import requests
file = open("./origin/ratings.csv", "r")
all_ratings = csv.reader(file)
next(all_ratings)

for idx, row in enumerate(all_ratings):
    # limit 100
    if idx == 100:
        break
    requests.post("http://localhost:8080/movies", 
                    headers={"Content-Type": "application/json"},
                    data={
                        "name": row[0],
                        "movie_id": row[1],
                        "rating": row[2],
                        "timestamp": row[3]
                    })
    