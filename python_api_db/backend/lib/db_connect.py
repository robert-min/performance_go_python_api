from . import DB_CONNECTION
from sqlalchemy import create_engine, select
from sqlalchemy.orm import Session
from flask_restx import abort
from model import Movies

class MySQLManager:
    """
    Database Manager
    """

    def __init__(self, env) -> None:
        user = DB_CONNECTION[env]['user']
        passwd = DB_CONNECTION[env]['passwd']
        host = DB_CONNECTION[env]['host']
        port = DB_CONNECTION[env]['port']
        db = DB_CONNECTION[env]['db']
        charset = DB_CONNECTION[env]['charset']
        engine = create_engine(f"mysql+pymysql://{user}:{passwd}@{host}:{port}/{db}?{charset}",
                                echo=False)

        self.session = Session(engine)

    
    def insert_movie(self, name, movie_id, rating, timestamp):
        """
        Insert movie data to movies table on robertmin db.
        """
        try:
            with self.session as session:
                content = Movies(
                    name=name,
                    movie_id=movie_id,
                    rating=rating,
                    timestamp=timestamp
                )
                session.add(content)
                session.commit()

        except Exception as e:
            return abort(404, status="Fail", message="Problems to DB connection", result=e)
        
        return name

    def get_all_movies(self):
        """
        Get all movies from movies table on robertmin db
        """
        try:
            all_movies = list()
            with self.session as session:
                sql = select(Movies)
                for row in session.execute(sql):
                    all_movies.append({
                        "name": row.Movies.name,
                        "movie_id": row.Movies.movie_id,
                        "rating": row.Movies.rating,
                        "timestamp": row.Movies.timestamp
                    })

        except Exception as e:
            return abort(404, status="Fail", message="Problems to DB connection", result=e)
        
        return all_movies


