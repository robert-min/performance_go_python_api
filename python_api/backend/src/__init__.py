import os
import sys
from flask import Flask
from flask_restx import Api
from flask_cors import CORS
from . import movies

src_path = os.path.abspath(os.path.join(__file__, 'src'))

if src_path not in sys.path:
    sys.path.append(src_path)

app = Flask(__name__)

def create_app():
    # flask_restx
    api = Api(app, title="python_api", doc="/doc")
    api.add_namespace(movies.movies_namespace, path="/movies")

    # CORS
    CORS(app, resources={r"/movies/*": {"origins": "*"}})

    return app

