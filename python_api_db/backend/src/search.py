from flask_restx import Namespace, Resource
from flask import request, jsonify
from lib.db_connect import MySQLManager

DBManager = MySQLManager("dev")
search_namespace = Namespace("search")
movies_namespace = Namespace("movies")

@search_namespace.route('/')
class GetData(Resource):
    def get(self):
        all_movies = DBManager.get_all_movies()
        return jsonify({"status": "OK", "result": all_movies}) 


@movies_namespace.route('/')
class PostMovies(Resource):
    def post(self):
        movie_info = request.get_json()
        result = DBManager.insert_movie(
            movie_info["name"],
            movie_info["movie_id"],
            movie_info["rating"],
            movie_info["timestamp"]
        )
        return jsonify({"status": "OK", "result": result})