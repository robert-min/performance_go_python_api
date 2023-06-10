import os
import sys
from flask import Flask
from flask_restx import Api
from flask_cors import CORS
from . import search


src_path = os.path.abspath(os.path.join(__file__, 'src'))
backend_path = os.path.abspath(os.path.join(src_path, os.path.pardir))
lib_path = os.path.abspath(os.path.join(backend_path, 'lib'))

if src_path not in sys.path:
    sys.path.append(src_path)

if backend_path not in sys.path:
    sys.path.append(backend_path)

if lib_path not in sys.path:
    sys.path.append(lib_path)

app = Flask(__name__)

def create_app():
    # flask_restx
    api = Api(app, title="python_api_db", doc="/doc")
    api.add_namespace(search.search_namespace, path="/search")
    api.add_namespace(search.movies_namespace, path="/movies")

    # CORS
    CORS(app, resources={r"/search/*": {"origins": "*"}})

    return app

