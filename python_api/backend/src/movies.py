from flask_restx import Namespace, Resource
import requests

movies_namespace = Namespace("movies")

@movies_namespace.route('/')
class GetMovies(Resource):
    def get(self):
        # docker network connect
        url = "http://python_api_db:8080/search"
        resp = requests.get(url)
        movies_data = resp.json()["result"]

        return {"status": "OK", "result": movies_data}